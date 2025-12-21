package views

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// ServerView handles server configuration UI
type ServerView struct {
	app           *tview.Application
	pages         *tview.Pages
	configService *server.ConfigService
	onBack        func()
	statusBar     *tview.TextView
}

// NewServerView creates a new server configuration view
func NewServerView(app *tview.Application, pages *tview.Pages, configService *server.ConfigService, statusBar *tview.TextView, onBack func()) *ServerView {
	return &ServerView{
		app:           app,
		pages:         pages,
		configService: configService,
		onBack:        onBack,
		statusBar:     statusBar,
	}
}

// Show displays the server settings menu
func (v *ServerView) Show() {
	srv := v.configService.GetServer()
	if srv == nil {
		v.showError("Failed to load server configuration")
		return
	}

	list := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	// Server Settings
	list.AddItem(
		fmt.Sprintf(i18n.T("server.port")+": [yellow]%d[-]", srv.Port),
		i18n.T("server.port.desc"),
		's',
		func() { v.showServerSettings() },
	)

	// Listeners
	list.AddItem(
		fmt.Sprintf(i18n.T("server.listeners")+": [yellow]%d[-]", len(srv.Listeners)),
		i18n.T("server.listeners.desc"),
		'l',
		func() { v.showListeners() },
	)

	// Services
	serviceCount := len(srv.Services)
	list.AddItem(
		fmt.Sprintf(i18n.T("server.services")+": [yellow]%d[-]", serviceCount),
		i18n.T("server.services.desc"),
		'v',
		func() { v.showServices() },
	)

	// Global Resources
	resourceCount := 0
	if srv.Resources != nil {
		resourceCount = len(srv.Resources.Resources)
	}
	list.AddItem(
		fmt.Sprintf(i18n.T("server.globalresources")+": [yellow]%d[-]", resourceCount),
		i18n.T("server.globalresources.desc"),
		'g',
		func() { v.showGlobalResources() },
	)

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 'b', v.onBack)

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.server.port")
		case 1:
			helpPanel.SetHelpKey("help.server.listener")
		case 2:
			helpPanel.SetHelpKey("help.server.service")
		case 3:
			helpPanel.SetHelpKey("help.server.executor")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.server.port")

	list.SetBorder(true).SetTitle(" " + i18n.T("server.title") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Create flex layout with list and help panel
	flex := tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("server-config", flex, true)
	v.app.SetFocus(list)
}

// showServerSettings shows the server settings form
func (v *ServerView) showServerSettings() {
	srv := v.configService.GetServer()

	form := tview.NewForm()
	preview := NewPreviewPanel()

	// Function to update preview
	updatePreview := func() {
		tempSrv := server.Server{
			Port:     srv.Port,
			Shutdown: srv.Shutdown,
		}

		if port, err := strconv.Atoi(form.GetFormItem(0).(*tview.InputField).GetText()); err == nil {
			tempSrv.Port = port
		}
		tempSrv.Shutdown = form.GetFormItem(1).(*tview.InputField).GetText()

		preview.SetXMLPreview(GenerateServerXML(&tempSrv))
	}

	form.AddInputField(i18n.T("server.port"), strconv.Itoa(srv.Port), 10, func(text string, lastChar rune) bool {
		// Only allow digits
		return lastChar >= '0' && lastChar <= '9'
	}, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.shutdown"), srv.Shutdown, 20, nil, func(text string) {
		updatePreview()
	})

	form.AddButton(i18n.T("common.save"), func() {
		portStr := form.GetFormItem(0).(*tview.InputField).GetText()
		shutdown := form.GetFormItem(1).(*tview.InputField).GetText()

		port, err := strconv.Atoi(portStr)
		if err != nil {
			v.showError(i18n.T("server.settings.invalidport"))
			return
		}

		v.configService.UpdateServerPort(port)
		v.configService.UpdateShutdownCommand(shutdown)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("server.settings.saved") + "[-]")
		v.Show()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.Show()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.settings") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("server-settings", layout, true)
	v.app.SetFocus(form)
}

// showListeners shows the listeners list
func (v *ServerView) showListeners() {
	listeners := v.configService.GetListeners()

	list := tview.NewList().ShowSecondaryText(true)

	for i, listener := range listeners {
		idx := i // Capture for closure
		desc := server.GetListenerDescription(listener.ClassName)
		shortName := getShortClassName(listener.ClassName)
		list.AddItem(
			fmt.Sprintf("[%d] %s", idx+1, shortName),
			desc,
			0,
			func() { v.showListenerDetail(idx) },
		)
	}

	list.AddItem("[green]"+i18n.T("server.listener.add")+"[-]", i18n.T("server.globalresource.add.desc"), 'a', func() {
		v.showAddListener()
	})

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 'b', func() {
		v.Show()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("server.listeners") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("listeners", list, true)
	v.app.SetFocus(list)
}

// showListenerDetail shows details of a specific listener
func (v *ServerView) showListenerDetail(index int) {
	listeners := v.configService.GetListeners()
	if index < 0 || index >= len(listeners) {
		return
	}
	listener := listeners[index]

	form := tview.NewForm()
	preview := NewPreviewPanel()

	hasSSLEngine := listener.SSLEngine != ""
	hasSSLSeed := listener.SSLRandomSeed != ""

	// Function to update preview
	updatePreview := func() {
		tempListener := server.Listener{
			ClassName:     listener.ClassName,
			SSLEngine:     listener.SSLEngine,
			SSLRandomSeed: listener.SSLRandomSeed,
		}

		tempListener.ClassName = form.GetFormItem(0).(*tview.InputField).GetText()
		fieldIdx := 1
		if hasSSLEngine {
			tempListener.SSLEngine = form.GetFormItem(fieldIdx).(*tview.InputField).GetText()
			fieldIdx++
		}
		if hasSSLSeed {
			tempListener.SSLRandomSeed = form.GetFormItem(fieldIdx).(*tview.InputField).GetText()
		}

		preview.SetXMLPreview(GenerateListenerXML(&tempListener))
	}

	form.AddInputField(i18n.T("server.listener.classname"), listener.ClassName, 60, nil, func(text string) {
		updatePreview()
	})

	// Add listener-specific fields
	if hasSSLEngine {
		form.AddInputField(i18n.T("server.listener.sslengine"), listener.SSLEngine, 20, nil, func(text string) {
			updatePreview()
		})
	}
	if hasSSLSeed {
		form.AddInputField(i18n.T("server.listener.sslseed"), listener.SSLRandomSeed, 20, nil, func(text string) {
			updatePreview()
		})
	}

	form.AddButton(i18n.T("common.save"), func() {
		className := form.GetFormItem(0).(*tview.InputField).GetText()
		listeners[index].ClassName = className

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("server.listener.updated") + "[-]")
		v.showListeners()
	})

	form.AddButton(i18n.T("server.listener.delete"), func() {
		v.showConfirm(i18n.T("server.listener.delete"), i18n.T("server.confirm.delete"), func(confirmed bool) {
			if confirmed {
				v.configService.RemoveListener(index)
				if err := v.configService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]" + i18n.T("server.listener.deleted") + "[-]")
			}
			v.showListeners()
		})
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showListeners()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.listener.detail") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("listener-detail", layout, true)
	v.app.SetFocus(form)
}

// showAddListener shows the add listener form
func (v *ServerView) showAddListener() {
	// Common listener options
	listenerOptions := []string{
		server.ListenerVersionLogger,
		server.ListenerAprLifecycle,
		server.ListenerJreMemoryLeak,
		server.ListenerGlobalResources,
		server.ListenerThreadLocalLeak,
	}

	list := tview.NewList().ShowSecondaryText(true)

	for _, className := range listenerOptions {
		cn := className // Capture
		desc := server.GetListenerDescription(className)
		list.AddItem(getShortClassName(className), desc, 0, func() {
			v.configService.AddListener(server.Listener{ClassName: cn})
			if err := v.configService.Save(); err != nil {
				v.showError(fmt.Sprintf("Failed to save: %v", err))
				return
			}
			v.setStatus("[green]Listener added[-]")
			v.showListeners()
		})
	}

	list.AddItem("[yellow]"+i18n.T("server.listener.custom")+"[-]", i18n.T("server.listener.custom.desc"), 'c', func() {
		v.showCustomListenerForm()
	})

	list.AddItem("[red]"+i18n.T("common.cancel")+"[-]", i18n.T("common.return"), 0, func() {
		v.showListeners()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("server.listener.add") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("add-listener", list, true)
	v.app.SetFocus(list)
}

// showCustomListenerForm shows a form for custom listener
func (v *ServerView) showCustomListenerForm() {
	form := tview.NewForm()
	form.AddInputField(i18n.T("server.listener.classname"), "", 60, nil, nil)

	form.AddButton(i18n.T("common.add"), func() {
		className := form.GetFormItem(0).(*tview.InputField).GetText()
		if className == "" {
			v.showError(i18n.T("server.listener.classrequired"))
			return
		}

		v.configService.AddListener(server.Listener{ClassName: className})
		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]" + i18n.T("server.listener.added") + "[-]")
		v.showListeners()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showAddListener()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.listener.custom.title") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("custom-listener", form, true)
	v.app.SetFocus(form)
}

// showServices shows the services list
func (v *ServerView) showServices() {
	services := v.configService.GetServices()

	list := tview.NewList().ShowSecondaryText(true)

	for i, svc := range services {
		idx := i
		connectorCount := len(svc.Connectors)
		list.AddItem(
			fmt.Sprintf("Service: [yellow]%s[-]", svc.Name),
			fmt.Sprintf("Engine: %s, Connectors: %d", svc.Engine.Name, connectorCount),
			0,
			func() { v.showServiceDetail(idx) },
		)
	}

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 'b', func() {
		v.Show()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("server.services") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("services", list, true)
	v.app.SetFocus(list)
}

// showServiceDetail shows details of a specific service
func (v *ServerView) showServiceDetail(index int) {
	svc := v.configService.GetService(index)
	if svc == nil {
		return
	}

	list := tview.NewList().ShowSecondaryText(true)

	// Service name
	list.AddItem(
		fmt.Sprintf(i18n.T("server.service.name")+": [yellow]%s[-]", svc.Name),
		i18n.T("server.service.name"),
		'n',
		func() { v.showEditServiceName(index) },
	)

	// Engine settings
	list.AddItem(
		fmt.Sprintf(i18n.T("server.engine")+": [yellow]%s[-]", svc.Engine.Name),
		fmt.Sprintf(i18n.T("server.service.engine.desc"), svc.Engine.DefaultHost, svc.Engine.JvmRoute),
		'e',
		func() { v.showEngineSettings(index) },
	)

	// Executors
	list.AddItem(
		fmt.Sprintf(i18n.T("server.executor")+": [yellow]%d[-]", len(svc.Executors)),
		i18n.T("server.executor"),
		'x',
		func() { v.showExecutors(index) },
	)

	// Connectors (preview)
	list.AddItem(
		fmt.Sprintf(i18n.T("server.service.connectors")+": [yellow]%d[-]", len(svc.Connectors)),
		i18n.T("server.service.connectors.desc"),
		'c',
		nil,
	)

	// Hosts
	list.AddItem(
		fmt.Sprintf(i18n.T("server.service.hosts")+": [yellow]%d[-]", len(svc.Engine.Hosts)),
		i18n.T("server.service.hosts.desc"),
		'h',
		nil,
	)

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 'b', func() {
		v.showServices()
	})

	list.SetBorder(true).SetTitle(fmt.Sprintf(" Service: %s ", svc.Name)).SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("service-detail", list, true)
	v.app.SetFocus(list)
}

// showEditServiceName shows form to edit service name
func (v *ServerView) showEditServiceName(index int) {
	svc := v.configService.GetService(index)
	if svc == nil {
		return
	}

	form := tview.NewForm()
	preview := NewPreviewPanel()

	// Function to update preview
	updatePreview := func() {
		tempSvc := server.Service{
			Name: svc.Name,
		}
		tempSvc.Name = form.GetFormItem(0).(*tview.InputField).GetText()
		preview.SetXMLPreview(GenerateServiceXML(&tempSvc))
	}

	form.AddInputField(i18n.T("server.service.name"), svc.Name, 30, nil, func(text string) {
		updatePreview()
	})

	form.AddButton(i18n.T("common.save"), func() {
		name := form.GetFormItem(0).(*tview.InputField).GetText()
		svc.Name = name
		v.configService.UpdateService(index, *svc)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("server.service.updated") + "[-]")
		v.showServiceDetail(index)
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showServiceDetail(index)
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.service.edit") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("edit-service-name", layout, true)
	v.app.SetFocus(form)
}

// showEngineSettings shows the engine settings form
func (v *ServerView) showEngineSettings(serviceIndex int) {
	svc := v.configService.GetService(serviceIndex)
	if svc == nil {
		return
	}

	form := tview.NewForm()
	preview := NewPreviewPanel()

	// Function to update preview
	updatePreview := func() {
		tempEngine := server.Engine{
			Name:        svc.Engine.Name,
			DefaultHost: svc.Engine.DefaultHost,
			JvmRoute:    svc.Engine.JvmRoute,
		}
		tempEngine.Name = form.GetFormItem(0).(*tview.InputField).GetText()
		tempEngine.DefaultHost = form.GetFormItem(1).(*tview.InputField).GetText()
		tempEngine.JvmRoute = form.GetFormItem(2).(*tview.InputField).GetText()

		preview.SetXMLPreview(GenerateEngineXML(&tempEngine))
	}

	form.AddInputField(i18n.T("server.engine.name"), svc.Engine.Name, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.engine.defaulthost"), svc.Engine.DefaultHost, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.engine.jvmroute"), svc.Engine.JvmRoute, 30, nil, func(text string) {
		updatePreview()
	})

	form.AddButton(i18n.T("common.save"), func() {
		svc.Engine.Name = form.GetFormItem(0).(*tview.InputField).GetText()
		svc.Engine.DefaultHost = form.GetFormItem(1).(*tview.InputField).GetText()
		svc.Engine.JvmRoute = form.GetFormItem(2).(*tview.InputField).GetText()

		v.configService.UpdateService(serviceIndex, *svc)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("server.engine.saved") + "[-]")
		v.showServiceDetail(serviceIndex)
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showServiceDetail(serviceIndex)
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.engine.settings") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("engine-settings", layout, true)
	v.app.SetFocus(form)
}

// showExecutors shows the executors list
func (v *ServerView) showExecutors(serviceIndex int) {
	svc := v.configService.GetService(serviceIndex)
	if svc == nil {
		return
	}

	list := tview.NewList().ShowSecondaryText(true)

	for i, exec := range svc.Executors {
		idx := i
		list.AddItem(
			fmt.Sprintf(i18n.T("server.executor")+": [yellow]%s[-]", exec.Name),
			fmt.Sprintf(i18n.T("server.executor.threads"), exec.MinSpareThreads, exec.MaxThreads, exec.MaxIdleTime),
			0,
			func() { v.showExecutorDetail(serviceIndex, idx) },
		)
	}

	list.AddItem("[green]"+i18n.T("server.executor.add")+"[-]", i18n.T("server.executor.add"), 'a', func() {
		v.showAddExecutor(serviceIndex)
	})

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 'b', func() {
		v.showServiceDetail(serviceIndex)
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("server.executor") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("executors", list, true)
	v.app.SetFocus(list)
}

// showExecutorDetail shows executor detail form
func (v *ServerView) showExecutorDetail(serviceIndex, executorIndex int) {
	svc := v.configService.GetService(serviceIndex)
	if svc == nil || executorIndex >= len(svc.Executors) {
		return
	}

	exec := &svc.Executors[executorIndex]

	form := tview.NewForm()
	preview := NewPreviewPanel()

	// Function to update preview
	updatePreview := func() {
		tempExec := server.Executor{
			ClassName:       exec.ClassName,
			Name:            exec.Name,
			NamePrefix:      exec.NamePrefix,
			MaxThreads:      exec.MaxThreads,
			MinSpareThreads: exec.MinSpareThreads,
			MaxIdleTime:     exec.MaxIdleTime,
			MaxQueueSize:    exec.MaxQueueSize,
		}

		tempExec.Name = form.GetFormItem(0).(*tview.InputField).GetText()
		tempExec.NamePrefix = form.GetFormItem(1).(*tview.InputField).GetText()
		if maxThreads, err := strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText()); err == nil {
			tempExec.MaxThreads = maxThreads
		}
		if minThreads, err := strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText()); err == nil {
			tempExec.MinSpareThreads = minThreads
		}
		if maxIdle, err := strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText()); err == nil {
			tempExec.MaxIdleTime = maxIdle
		}

		preview.SetXMLPreview(GenerateExecutorXML(&tempExec))
	}

	form.AddInputField(i18n.T("server.executor.name"), exec.Name, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.executor.prefix"), exec.NamePrefix, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.executor.maxthreads"), strconv.Itoa(exec.MaxThreads), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.executor.minthreads"), strconv.Itoa(exec.MinSpareThreads), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.executor.maxidle"), strconv.Itoa(exec.MaxIdleTime), 10, acceptDigits, func(text string) {
		updatePreview()
	})

	form.AddButton(i18n.T("common.save"), func() {
		exec.Name = form.GetFormItem(0).(*tview.InputField).GetText()
		exec.NamePrefix = form.GetFormItem(1).(*tview.InputField).GetText()
		exec.MaxThreads, _ = strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText())
		exec.MinSpareThreads, _ = strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())
		exec.MaxIdleTime, _ = strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText())

		v.configService.UpdateService(serviceIndex, *svc)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("server.executor.updated") + "[-]")
		v.showExecutors(serviceIndex)
	})

	form.AddButton(i18n.T("common.delete"), func() {
		v.showConfirm(i18n.T("common.delete"), i18n.T("server.confirm.delete"), func(confirmed bool) {
			if confirmed {
				svc.Executors = append(svc.Executors[:executorIndex], svc.Executors[executorIndex+1:]...)
				v.configService.UpdateService(serviceIndex, *svc)
				if err := v.configService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]" + i18n.T("server.executor.deleted") + "[-]")
			}
			v.showExecutors(serviceIndex)
		})
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showExecutors(serviceIndex)
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.executor.edit") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("executor-detail", layout, true)
	v.app.SetFocus(form)
}

// showAddExecutor shows form to add new executor
func (v *ServerView) showAddExecutor(serviceIndex int) {
	svc := v.configService.GetService(serviceIndex)
	if svc == nil {
		return
	}

	form := tview.NewForm()
	form.AddInputField(i18n.T("server.executor.name"), "tomcatThreadPool", 30, nil, nil)
	form.AddInputField(i18n.T("server.executor.prefix"), "catalina-exec-", 30, nil, nil)
	form.AddInputField(i18n.T("server.executor.maxthreads"), "200", 10, acceptDigits, nil)
	form.AddInputField(i18n.T("server.executor.minthreads"), "25", 10, acceptDigits, nil)
	form.AddInputField(i18n.T("server.executor.maxidle"), "60000", 10, acceptDigits, nil)

	form.AddButton(i18n.T("common.add"), func() {
		maxThreads, _ := strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText())
		minSpare, _ := strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())
		maxIdle, _ := strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText())

		exec := server.Executor{
			Name:            form.GetFormItem(0).(*tview.InputField).GetText(),
			NamePrefix:      form.GetFormItem(1).(*tview.InputField).GetText(),
			MaxThreads:      maxThreads,
			MinSpareThreads: minSpare,
			MaxIdleTime:     maxIdle,
		}

		svc.Executors = append(svc.Executors, exec)
		v.configService.UpdateService(serviceIndex, *svc)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("server.executor.added") + "[-]")
		v.showExecutors(serviceIndex)
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showExecutors(serviceIndex)
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.executor.add") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("add-executor", form, true)
	v.app.SetFocus(form)
}

// showGlobalResources shows global JNDI resources
func (v *ServerView) showGlobalResources() {
	resources := v.configService.GetGlobalResources()

	list := tview.NewList().ShowSecondaryText(true)

	if resources != nil {
		for i, res := range resources.Resources {
			idx := i
			list.AddItem(
				fmt.Sprintf("[yellow]%s[-]", res.Name),
				fmt.Sprintf("Type: %s", res.Type),
				0,
				func() { v.showGlobalResourceDetail(idx) },
			)
		}
	}

	list.AddItem("[green]"+i18n.T("server.globalresource.add")+"[-]", i18n.T("server.globalresource.add.desc"), 'a', func() {
		v.showAddGlobalResource()
	})

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 'b', func() {
		v.Show()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("server.globalresources") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("global-resources", list, true)
	v.app.SetFocus(list)
}

// showGlobalResourceDetail shows a global resource detail
func (v *ServerView) showGlobalResourceDetail(index int) {
	resources := v.configService.GetGlobalResources()
	if resources == nil || index >= len(resources.Resources) {
		return
	}

	res := &resources.Resources[index]

	form := tview.NewForm()
	preview := NewPreviewPanel()

	// Function to update preview
	updatePreview := func() {
		tempRes := server.Resource{
			Name:        res.Name,
			Auth:        res.Auth,
			Type:        res.Type,
			Description: res.Description,
			Factory:     res.Factory,
			Pathname:    res.Pathname,
		}

		tempRes.Name = form.GetFormItem(0).(*tview.InputField).GetText()
		tempRes.Auth = form.GetFormItem(1).(*tview.InputField).GetText()
		tempRes.Type = form.GetFormItem(2).(*tview.InputField).GetText()
		tempRes.Description = form.GetFormItem(3).(*tview.InputField).GetText()
		tempRes.Factory = form.GetFormItem(4).(*tview.InputField).GetText()
		tempRes.Pathname = form.GetFormItem(5).(*tview.InputField).GetText()

		preview.SetXMLPreview(GenerateResourceXML(&tempRes))
	}

	form.AddInputField(i18n.T("jndi.resource.name"), res.Name, 40, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.globalresource.auth"), res.Auth, 20, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.globalresource.type"), res.Type, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.globalresource.description"), res.Description, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.globalresource.factory"), res.Factory, 60, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("server.globalresource.pathname"), res.Pathname, 40, nil, func(text string) {
		updatePreview()
	})

	form.AddButton(i18n.T("common.save"), func() {
		res.Name = form.GetFormItem(0).(*tview.InputField).GetText()
		res.Auth = form.GetFormItem(1).(*tview.InputField).GetText()
		res.Type = form.GetFormItem(2).(*tview.InputField).GetText()
		res.Description = form.GetFormItem(3).(*tview.InputField).GetText()
		res.Factory = form.GetFormItem(4).(*tview.InputField).GetText()
		res.Pathname = form.GetFormItem(5).(*tview.InputField).GetText()

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("server.globalresource.updated") + "[-]")
		v.showGlobalResources()
	})

	form.AddButton(i18n.T("common.delete"), func() {
		v.showConfirm(i18n.T("common.delete"), i18n.T("server.confirm.delete"), func(confirmed bool) {
			if confirmed {
				v.configService.RemoveGlobalResource(index)
				if err := v.configService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]" + i18n.T("server.globalresource.deleted") + "[-]")
			}
			v.showGlobalResources()
		})
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showGlobalResources()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.globalresource.edit") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("global-resource-detail", layout, true)
	v.app.SetFocus(form)
}

// showAddGlobalResource shows form to add global resource
func (v *ServerView) showAddGlobalResource() {
	form := tview.NewForm()
	form.AddInputField(i18n.T("jndi.resource.name"), "UserDatabase", 40, nil, nil)
	form.AddInputField(i18n.T("server.globalresource.auth"), "Container", 20, nil, nil)
	form.AddInputField(i18n.T("server.globalresource.type"), "org.apache.catalina.UserDatabase", 50, nil, nil)
	form.AddInputField(i18n.T("server.globalresource.description"), "User database that can be updated and saved", 50, nil, nil)
	form.AddInputField(i18n.T("server.globalresource.factory"), "org.apache.catalina.users.MemoryUserDatabaseFactory", 60, nil, nil)
	form.AddInputField(i18n.T("server.globalresource.pathname"), "conf/tomcat-users.xml", 40, nil, nil)

	form.AddButton(i18n.T("common.add"), func() {
		res := server.Resource{
			Name:        form.GetFormItem(0).(*tview.InputField).GetText(),
			Auth:        form.GetFormItem(1).(*tview.InputField).GetText(),
			Type:        form.GetFormItem(2).(*tview.InputField).GetText(),
			Description: form.GetFormItem(3).(*tview.InputField).GetText(),
			Factory:     form.GetFormItem(4).(*tview.InputField).GetText(),
			Pathname:    form.GetFormItem(5).(*tview.InputField).GetText(),
		}

		v.configService.AddGlobalResource(res)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("server.globalresource.added") + "[-]")
		v.showGlobalResources()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showGlobalResources()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("server.globalresource.add") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("add-global-resource", form, true)
	v.app.SetFocus(form)
}

// Helper functions

func (v *ServerView) showError(message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			v.pages.SwitchToPage("server-config")
		})
	modal.SetBackgroundColor(tcell.ColorDarkRed)
	v.pages.AddAndSwitchToPage("error", modal, true)
}

func (v *ServerView) showConfirm(title, message string, callback func(bool)) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			callback(buttonIndex == 0)
		})
	modal.SetBorder(true).SetTitle(fmt.Sprintf(" %s ", title))
	v.pages.AddAndSwitchToPage("confirm", modal, true)
}

func (v *ServerView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(fmt.Sprintf(" %s", message))
	}
}

// getShortClassName returns the short class name from full qualified name
func getShortClassName(fullName string) string {
	for i := len(fullName) - 1; i >= 0; i-- {
		if fullName[i] == '.' {
			return fullName[i+1:]
		}
	}
	return fullName
}

// acceptDigits is a validation function that only accepts digits
func acceptDigits(text string, lastChar rune) bool {
	return lastChar >= '0' && lastChar <= '9'
}
