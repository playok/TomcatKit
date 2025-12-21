package views

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// HostView handles virtual host and context configuration
type HostView struct {
	app           *tview.Application
	pages         *tview.Pages
	mainPages     *tview.Pages
	statusBar     *tview.TextView
	onReturn      func()
	configService *server.ConfigService
}

// NewHostView creates a new host view
func NewHostView(app *tview.Application, mainPages *tview.Pages, configService *server.ConfigService, statusBar *tview.TextView, onReturn func()) *HostView {
	return &HostView{
		app:           app,
		pages:         tview.NewPages(),
		mainPages:     mainPages,
		statusBar:     statusBar,
		onReturn:      onReturn,
		configService: configService,
	}
}

// Show displays the host view
func (v *HostView) Show() {
	v.showMainMenu()
	v.mainPages.AddAndSwitchToPage("host", v.pages, true)
}

// showMainMenu shows the main host/context menu
func (v *HostView) showMainMenu() {
	menu := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	cfg := v.configService.GetConfig()
	hostCount := 0
	contextCount := 0
	if cfg != nil && len(cfg.Services) > 0 {
		hostCount = len(cfg.Services[0].Engine.Hosts)
		for _, host := range cfg.Services[0].Engine.Hosts {
			contextCount += len(host.Contexts)
		}
	}

	menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("host.virtualhost")+"[::-] [yellow](%d)[-]", hostCount),
		i18n.T("host.virtualhost.desc"), 'h', func() {
			v.showHostList()
		})

	menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("host.context")+"[::-] [yellow](%d)[-]", contextCount),
		i18n.T("host.context.desc"), 'c', func() {
			v.showContextSelector()
		})

	menu.AddItem("[::b]"+i18n.T("host.engine")+"[::-]",
		i18n.T("host.engine.desc"), 'e', func() {
			v.showEngineForm()
		})

	menu.AddItem("", "", 0, nil)
	menu.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 0, func() {
		v.onReturn()
	})

	// Update help panel when selection changes
	menu.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.host.virtualhost")
		case 1:
			helpPanel.SetHelpKey("help.host.context")
		case 2:
			helpPanel.SetHelpKey("help.host.engine")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.host.virtualhost")

	menu.SetBorder(true).SetTitle(" " + i18n.T("host.title") + " ").SetBorderColor(tcell.ColorDarkCyan)
	menu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.onReturn()
			return nil
		}
		return event
	})

	// Create flex layout with menu and help panel
	flex := tview.NewFlex().
		AddItem(menu, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("menu", flex, true)
	v.app.SetFocus(menu)
}

// showEngineForm shows the engine configuration form
func (v *HostView) showEngineForm() {
	cfg := v.configService.GetConfig()
	if cfg == nil || len(cfg.Services) == 0 {
		v.setStatus("[red]No service configuration found[-]")
		return
	}

	engine := &cfg.Services[0].Engine
	form := tview.NewForm()
	preview := NewPreviewPanel()

	updatePreview := func() {
		tempEngine := server.Engine{
			Name:        form.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
			DefaultHost: form.GetFormItemByLabel("Default Host").(*tview.InputField).GetText(),
			JvmRoute:    form.GetFormItemByLabel("JVM Route").(*tview.InputField).GetText(),
		}
		preview.SetXMLPreview(GenerateEngineXML(&tempEngine))
	}

	form.AddInputField("Name", engine.Name, 30, nil, func(text string) { updatePreview() })
	form.AddInputField("Default Host", engine.DefaultHost, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("JVM Route", engine.JvmRoute, 30, nil, func(text string) { updatePreview() })

	form.AddButton(i18n.T("common.save.short"), func() {
		engine.Name = form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
		engine.DefaultHost = form.GetFormItemByLabel("Default Host").(*tview.InputField).GetText()
		engine.JvmRoute = form.GetFormItemByLabel("JVM Route").(*tview.InputField).GetText()

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Engine settings saved successfully[-]")
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMainMenu()
	})

	form.SetBorder(true).SetTitle(" Engine Settings ").SetBorderColor(tcell.ColorBlue)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	// Initial preview
	updatePreview()

	// Layout with form and preview
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("engine-form", layout, true)
	v.app.SetFocus(form)
}

// showHostList shows the list of virtual hosts
func (v *HostView) showHostList() {
	list := tview.NewList().ShowSecondaryText(true)

	cfg := v.configService.GetConfig()
	if cfg != nil && len(cfg.Services) > 0 {
		engine := &cfg.Services[0].Engine
		for i := range engine.Hosts {
			host := &engine.Hosts[i]
			info := fmt.Sprintf("appBase: %s, contexts: %d", host.AppBase, len(host.Contexts))
			isDefault := ""
			if host.Name == engine.DefaultHost {
				isDefault = " [green](default)[-]"
			}
			list.AddItem(fmt.Sprintf("%s%s", host.Name, isDefault), info, 0, func() {
				v.showHostForm(host, false)
			})
		}
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add Virtual Host[-]", "Create new virtual host", 'a', func() {
		newHost := &server.Host{
			Name:            "newhost.example.com",
			AppBase:         "webapps",
			UnpackWARs:      true,
			AutoDeploy:      true,
			DeployOnStartup: true,
		}
		v.showHostForm(newHost, true)
	})
	list.AddItem("[red]Back[-]", "Return to menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Virtual Hosts ").SetBorderColor(tcell.ColorGreen)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("host-list", list, true)
	v.app.SetFocus(list)
}

// showHostForm shows the host edit form
func (v *HostView) showHostForm(host *server.Host, isNew bool) {
	form := tview.NewForm()
	preview := NewPreviewPanel()

	// Aliases
	aliasStr := ""
	for _, alias := range host.Aliases {
		if aliasStr != "" {
			aliasStr += ", "
		}
		aliasStr += alias.Name
	}

	updatePreview := func() {
		tempHost := server.Host{
			Name:             GetFormText(form, "Name (hostname)"),
			AppBase:          GetFormText(form, "App Base"),
			WorkDir:          GetFormText(form, "Work Dir"),
			UnpackWARs:       GetFormBool(form, "Unpack WARs"),
			AutoDeploy:       GetFormBool(form, "Auto Deploy"),
			DeployOnStartup:  GetFormBool(form, "Deploy On Startup"),
			CreateDirs:       GetFormBool(form, "Create Dirs"),
			DeployXML:        GetFormBool(form, "Deploy XML"),
			CopyXML:          GetFormBool(form, "Copy XML"),
			DeployIgnore:     GetFormText(form, "Deploy Ignore (regex)"),
			StartStopThreads: GetFormInt(form, "Start/Stop Threads"),
		}
		preview.SetXMLPreview(GenerateHostXML(&tempHost))
	}

	// Basic settings
	form.AddInputField("Name (hostname)", host.Name, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("App Base", host.AppBase, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("Work Dir", host.WorkDir, 40, nil, func(text string) { updatePreview() })

	// Deployment settings
	form.AddCheckbox("Unpack WARs", host.UnpackWARs, func(checked bool) { updatePreview() })
	form.AddCheckbox("Auto Deploy", host.AutoDeploy, func(checked bool) { updatePreview() })
	form.AddCheckbox("Deploy On Startup", host.DeployOnStartup, func(checked bool) { updatePreview() })
	form.AddCheckbox("Create Dirs", host.CreateDirs, func(checked bool) { updatePreview() })
	form.AddCheckbox("Deploy XML", host.DeployXML, func(checked bool) { updatePreview() })
	form.AddCheckbox("Copy XML", host.CopyXML, func(checked bool) { updatePreview() })

	// Advanced settings
	form.AddInputField("Deploy Ignore (regex)", host.DeployIgnore, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("Start/Stop Threads", strconv.Itoa(host.StartStopThreads), 10, acceptNumber, func(text string) { updatePreview() })

	form.AddInputField("Aliases (comma-separated)", aliasStr, 50, nil, nil)

	originalName := host.Name

	form.AddButton(i18n.T("common.save.short"), func() {
		host.Name = GetFormText(form, "Name (hostname)")
		host.AppBase = GetFormText(form, "App Base")
		host.WorkDir = GetFormText(form, "Work Dir")
		host.UnpackWARs = GetFormBool(form, "Unpack WARs")
		host.AutoDeploy = GetFormBool(form, "Auto Deploy")
		host.DeployOnStartup = GetFormBool(form, "Deploy On Startup")
		host.CreateDirs = GetFormBool(form, "Create Dirs")
		host.DeployXML = GetFormBool(form, "Deploy XML")
		host.CopyXML = GetFormBool(form, "Copy XML")
		host.DeployIgnore = GetFormText(form, "Deploy Ignore (regex)")
		host.StartStopThreads = GetFormInt(form, "Start/Stop Threads")

		// Parse aliases
		host.Aliases = nil
		for _, a := range ParseCommaSeparated(GetFormText(form, "Aliases (comma-separated)")) {
			host.Aliases = append(host.Aliases, server.Alias{Name: a})
		}

		cfg := v.configService.GetConfig()
		if cfg != nil && len(cfg.Services) > 0 {
			engine := &cfg.Services[0].Engine
			if isNew {
				engine.Hosts = append(engine.Hosts, *host)
			} else {
				// Find and update
				for i := range engine.Hosts {
					if engine.Hosts[i].Name == originalName {
						engine.Hosts[i] = *host
						break
					}
				}
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Virtual host saved successfully[-]")
		v.showHostList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.contexts"), func() {
			v.showContextList(host)
		})

		form.AddButton(i18n.T("common.delete"), func() {
			v.confirmDelete("Virtual Host", host.Name, func() {
				cfg := v.configService.GetConfig()
				if cfg != nil && len(cfg.Services) > 0 {
					engine := &cfg.Services[0].Engine
					for i := range engine.Hosts {
						if engine.Hosts[i].Name == host.Name {
							engine.Hosts = append(engine.Hosts[:i], engine.Hosts[i+1:]...)
							break
						}
					}
				}
				if err := v.configService.Save(); err != nil {
					v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					return
				}
				v.setStatus("[green]Virtual host deleted[-]")
				v.showHostList()
			})
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showHostList()
	})

	title := " Edit Virtual Host "
	if isNew {
		title = " Add Virtual Host "
	}
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorGreen)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showHostList()
			return nil
		}
		return event
	})

	// Initial preview
	updatePreview()

	// Layout with form and preview
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("host-form", layout, true)
	v.app.SetFocus(form)
}

// showContextSelector shows the host selector for context management
func (v *HostView) showContextSelector() {
	list := tview.NewList().ShowSecondaryText(true)

	cfg := v.configService.GetConfig()
	if cfg != nil && len(cfg.Services) > 0 {
		for i := range cfg.Services[0].Engine.Hosts {
			host := &cfg.Services[0].Engine.Hosts[i]
			info := fmt.Sprintf("%d contexts", len(host.Contexts))
			list.AddItem(host.Name, info, 0, func() {
				v.showContextList(host)
			})
		}
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[red]Back[-]", "Return to menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Select Host for Context Management ").SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("context-selector", list, true)
	v.app.SetFocus(list)
}

// showContextList shows the list of contexts for a host
func (v *HostView) showContextList(host *server.Host) {
	list := tview.NewList().ShowSecondaryText(true)

	for i := range host.Contexts {
		ctx := &host.Contexts[i]
		path := ctx.Path
		if path == "" {
			path = "/"
		}
		info := fmt.Sprintf("docBase: %s", ctx.DocBase)
		list.AddItem(path, info, 0, func() {
			v.showContextForm(host, ctx, false)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add Context[-]", "Create new web application context", 'a', func() {
		newCtx := &server.Context{
			Path:       "/newapp",
			DocBase:    "newapp",
			Reloadable: false,
			Cookies:    true,
		}
		v.showContextForm(host, newCtx, true)
	})
	list.AddItem("[red]Back[-]", "Return to host list", 0, func() {
		v.showHostList()
	})

	list.SetBorder(true).SetTitle(fmt.Sprintf(" Contexts for %s ", host.Name)).SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showHostList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("context-list", list, true)
	v.app.SetFocus(list)
}

// showContextForm shows the context edit form
func (v *HostView) showContextForm(host *server.Host, ctx *server.Context, isNew bool) {
	form := tview.NewForm()
	preview := NewPreviewPanel()

	updatePreview := func() {
		tempCtx := server.Context{
			Path:                form.GetFormItemByLabel("Path").(*tview.InputField).GetText(),
			DocBase:             form.GetFormItemByLabel("DocBase").(*tview.InputField).GetText(),
			Reloadable:          form.GetFormItemByLabel("Reloadable").(*tview.Checkbox).IsChecked(),
			CrossContext:        form.GetFormItemByLabel("Cross Context").(*tview.Checkbox).IsChecked(),
			Privileged:          form.GetFormItemByLabel("Privileged").(*tview.Checkbox).IsChecked(),
			Cookies:             form.GetFormItemByLabel("Cookies").(*tview.Checkbox).IsChecked(),
			SessionCookieName:   form.GetFormItemByLabel("Session Cookie Name").(*tview.InputField).GetText(),
			SessionCookiePath:   form.GetFormItemByLabel("Session Cookie Path").(*tview.InputField).GetText(),
			SessionCookieDomain: form.GetFormItemByLabel("Session Cookie Domain").(*tview.InputField).GetText(),
			UseHttpOnly:         form.GetFormItemByLabel("Use HttpOnly").(*tview.Checkbox).IsChecked(),
			AntiResourceLocking: form.GetFormItemByLabel("Anti Resource Locking").(*tview.Checkbox).IsChecked(),
			SwallowOutput:       form.GetFormItemByLabel("Swallow Output").(*tview.Checkbox).IsChecked(),
			CachingAllowed:      form.GetFormItemByLabel("Caching Allowed").(*tview.Checkbox).IsChecked(),
		}
		maxSize := form.GetFormItemByLabel("Cache Max Size (KB)").(*tview.InputField).GetText()
		tempCtx.CacheMaxSize, _ = strconv.Atoi(maxSize)
		ttl := form.GetFormItemByLabel("Cache TTL (ms)").(*tview.InputField).GetText()
		tempCtx.CacheTTL, _ = strconv.Atoi(ttl)
		preview.SetXMLPreview(GenerateContextXML(&tempCtx))
	}

	// Basic settings
	form.AddInputField("Path", ctx.Path, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("DocBase", ctx.DocBase, 40, nil, func(text string) { updatePreview() })

	// Behavior settings
	form.AddCheckbox("Reloadable", ctx.Reloadable, func(checked bool) { updatePreview() })
	form.AddCheckbox("Cross Context", ctx.CrossContext, func(checked bool) { updatePreview() })
	form.AddCheckbox("Privileged", ctx.Privileged, func(checked bool) { updatePreview() })

	// Session settings
	form.AddCheckbox("Cookies", ctx.Cookies, func(checked bool) { updatePreview() })
	form.AddInputField("Session Cookie Name", ctx.SessionCookieName, 30, nil, func(text string) { updatePreview() })
	form.AddInputField("Session Cookie Path", ctx.SessionCookiePath, 30, nil, func(text string) { updatePreview() })
	form.AddInputField("Session Cookie Domain", ctx.SessionCookieDomain, 30, nil, func(text string) { updatePreview() })
	form.AddCheckbox("Use HttpOnly", ctx.UseHttpOnly, func(checked bool) { updatePreview() })

	// Resource handling
	form.AddCheckbox("Anti Resource Locking", ctx.AntiResourceLocking, func(checked bool) { updatePreview() })
	form.AddCheckbox("Swallow Output", ctx.SwallowOutput, func(checked bool) { updatePreview() })

	// Cache settings
	form.AddCheckbox("Caching Allowed", ctx.CachingAllowed, func(checked bool) { updatePreview() })
	form.AddInputField("Cache Max Size (KB)", strconv.Itoa(ctx.CacheMaxSize), 10, acceptNumber, func(text string) { updatePreview() })
	form.AddInputField("Cache TTL (ms)", strconv.Itoa(ctx.CacheTTL), 10, acceptNumber, func(text string) { updatePreview() })

	originalPath := ctx.Path

	form.AddButton(i18n.T("common.save.short"), func() {
		ctx.Path = form.GetFormItemByLabel("Path").(*tview.InputField).GetText()
		ctx.DocBase = form.GetFormItemByLabel("DocBase").(*tview.InputField).GetText()
		ctx.Reloadable = form.GetFormItemByLabel("Reloadable").(*tview.Checkbox).IsChecked()
		ctx.CrossContext = form.GetFormItemByLabel("Cross Context").(*tview.Checkbox).IsChecked()
		ctx.Privileged = form.GetFormItemByLabel("Privileged").(*tview.Checkbox).IsChecked()
		ctx.Cookies = form.GetFormItemByLabel("Cookies").(*tview.Checkbox).IsChecked()
		ctx.SessionCookieName = form.GetFormItemByLabel("Session Cookie Name").(*tview.InputField).GetText()
		ctx.SessionCookiePath = form.GetFormItemByLabel("Session Cookie Path").(*tview.InputField).GetText()
		ctx.SessionCookieDomain = form.GetFormItemByLabel("Session Cookie Domain").(*tview.InputField).GetText()
		ctx.UseHttpOnly = form.GetFormItemByLabel("Use HttpOnly").(*tview.Checkbox).IsChecked()
		ctx.AntiResourceLocking = form.GetFormItemByLabel("Anti Resource Locking").(*tview.Checkbox).IsChecked()
		ctx.SwallowOutput = form.GetFormItemByLabel("Swallow Output").(*tview.Checkbox).IsChecked()
		ctx.CachingAllowed = form.GetFormItemByLabel("Caching Allowed").(*tview.Checkbox).IsChecked()
		ctx.CacheMaxSize, _ = strconv.Atoi(form.GetFormItemByLabel("Cache Max Size (KB)").(*tview.InputField).GetText())
		ctx.CacheTTL, _ = strconv.Atoi(form.GetFormItemByLabel("Cache TTL (ms)").(*tview.InputField).GetText())

		if isNew {
			host.Contexts = append(host.Contexts, *ctx)
		} else {
			// Find and update
			for i := range host.Contexts {
				if host.Contexts[i].Path == originalPath {
					host.Contexts[i] = *ctx
					break
				}
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Context saved successfully[-]")
		v.showContextList(host)
	})

	if !isNew {
		form.AddButton(i18n.T("common.parameters"), func() {
			v.showParameterList(host, ctx)
		})

		form.AddButton(i18n.T("host.sessionmanager"), func() {
			v.showManagerForm(host, ctx)
		})

		form.AddButton(i18n.T("common.delete"), func() {
			v.confirmDelete("Context", ctx.Path, func() {
				for i := range host.Contexts {
					if host.Contexts[i].Path == ctx.Path {
						host.Contexts = append(host.Contexts[:i], host.Contexts[i+1:]...)
						break
					}
				}
				if err := v.configService.Save(); err != nil {
					v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					return
				}
				v.setStatus("[green]Context deleted[-]")
				v.showContextList(host)
			})
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showContextList(host)
	})

	title := " Edit Context "
	if isNew {
		title = " Add Context "
	}
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorYellow)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showContextList(host)
			return nil
		}
		return event
	})

	// Initial preview
	updatePreview()

	// Layout with form and preview
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("context-form", layout, true)
	v.app.SetFocus(form)
}

// showParameterList shows the list of context parameters
func (v *HostView) showParameterList(host *server.Host, ctx *server.Context) {
	list := tview.NewList().ShowSecondaryText(true)

	for i := range ctx.Parameters {
		param := &ctx.Parameters[i]
		info := fmt.Sprintf("= %s", param.Value)
		if len(info) > 40 {
			info = info[:37] + "..."
		}
		list.AddItem(param.Name, info, 0, func() {
			v.showParameterForm(host, ctx, param, false)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add Parameter[-]", "Create new context parameter", 'a', func() {
		newParam := &server.Parameter{
			Name:     "newParam",
			Value:    "",
			Override: true,
		}
		v.showParameterForm(host, ctx, newParam, true)
	})
	list.AddItem("[red]Back[-]", "Return to context", 0, func() {
		v.showContextForm(host, ctx, false)
	})

	list.SetBorder(true).SetTitle(fmt.Sprintf(" Parameters for %s ", ctx.Path)).SetBorderColor(tcell.ColorPurple)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showContextForm(host, ctx, false)
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("parameter-list", list, true)
	v.app.SetFocus(list)
}

// showParameterForm shows the parameter edit form
func (v *HostView) showParameterForm(host *server.Host, ctx *server.Context, param *server.Parameter, isNew bool) {
	form := tview.NewForm()
	preview := NewPreviewPanel()

	updatePreview := func() {
		tempParam := server.Parameter{
			Name:        form.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
			Value:       form.GetFormItemByLabel("Value").(*tview.InputField).GetText(),
			Override:    form.GetFormItemByLabel("Override").(*tview.Checkbox).IsChecked(),
			Description: form.GetFormItemByLabel("Description").(*tview.InputField).GetText(),
		}
		preview.SetXMLPreview(GenerateParameterXML(&tempParam))
	}

	form.AddInputField("Name", param.Name, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("Value", param.Value, 50, nil, func(text string) { updatePreview() })
	form.AddCheckbox("Override", param.Override, func(checked bool) { updatePreview() })
	form.AddInputField("Description", param.Description, 50, nil, func(text string) { updatePreview() })

	originalName := param.Name

	form.AddButton(i18n.T("common.save.short"), func() {
		param.Name = form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
		param.Value = form.GetFormItemByLabel("Value").(*tview.InputField).GetText()
		param.Override = form.GetFormItemByLabel("Override").(*tview.Checkbox).IsChecked()
		param.Description = form.GetFormItemByLabel("Description").(*tview.InputField).GetText()

		if isNew {
			ctx.Parameters = append(ctx.Parameters, *param)
		} else {
			for i := range ctx.Parameters {
				if ctx.Parameters[i].Name == originalName {
					ctx.Parameters[i] = *param
					break
				}
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Parameter saved successfully[-]")
		v.showParameterList(host, ctx)
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			v.confirmDelete("Parameter", param.Name, func() {
				for i := range ctx.Parameters {
					if ctx.Parameters[i].Name == param.Name {
						ctx.Parameters = append(ctx.Parameters[:i], ctx.Parameters[i+1:]...)
						break
					}
				}
				if err := v.configService.Save(); err != nil {
					v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					return
				}
				v.setStatus("[green]Parameter deleted[-]")
				v.showParameterList(host, ctx)
			})
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showParameterList(host, ctx)
	})

	title := " Edit Parameter "
	if isNew {
		title = " Add Parameter "
	}
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorPurple)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showParameterList(host, ctx)
			return nil
		}
		return event
	})

	// Initial preview
	updatePreview()

	// Layout with form and preview
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("parameter-form", layout, true)
	v.app.SetFocus(form)
}

// showManagerForm shows the session manager configuration form
func (v *HostView) showManagerForm(host *server.Host, ctx *server.Context) {
	if ctx.Manager == nil {
		ctx.Manager = &server.Manager{}
	}

	form := tview.NewForm()
	preview := NewPreviewPanel()

	// Manager class options
	managerClasses := []string{
		"",
		"org.apache.catalina.session.StandardManager",
		"org.apache.catalina.session.PersistentManager",
	}
	classIdx := 0
	for i, c := range managerClasses {
		if c == ctx.Manager.ClassName {
			classIdx = i
			break
		}
	}

	updatePreview := func() {
		tempMgr := server.Manager{}
		_, tempMgr.ClassName = form.GetFormItemByLabel("Manager Class").(*tview.DropDown).GetCurrentOption()
		maxActive := form.GetFormItemByLabel("Max Active Sessions").(*tview.InputField).GetText()
		tempMgr.MaxActiveSessions, _ = strconv.Atoi(maxActive)
		sidLen := form.GetFormItemByLabel("Session ID Length").(*tview.InputField).GetText()
		tempMgr.SessionIdLength, _ = strconv.Atoi(sidLen)
		maxInactive := form.GetFormItemByLabel("Max Inactive Interval (sec)").(*tview.InputField).GetText()
		tempMgr.MaxInactiveInterval, _ = strconv.Atoi(maxInactive)
		tempMgr.Pathname = form.GetFormItemByLabel("Session File Path").(*tview.InputField).GetText()
		procExp := form.GetFormItemByLabel("Process Expires Frequency").(*tview.InputField).GetText()
		tempMgr.ProcessExpiresFrequency, _ = strconv.Atoi(procExp)
		preview.SetXMLPreview(GenerateManagerXML(&tempMgr))
	}

	form.AddDropDown("Manager Class", managerClasses, classIdx, func(option string, optionIndex int) { updatePreview() })

	form.AddInputField("Max Active Sessions", strconv.Itoa(ctx.Manager.MaxActiveSessions), 10, acceptNumber, func(text string) { updatePreview() })
	form.AddInputField("Session ID Length", strconv.Itoa(ctx.Manager.SessionIdLength), 10, acceptNumber, func(text string) { updatePreview() })
	form.AddInputField("Max Inactive Interval (sec)", strconv.Itoa(ctx.Manager.MaxInactiveInterval), 10, acceptNumber, func(text string) { updatePreview() })
	form.AddInputField("Session File Path", ctx.Manager.Pathname, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("Process Expires Frequency", strconv.Itoa(ctx.Manager.ProcessExpiresFrequency), 10, acceptNumber, func(text string) { updatePreview() })

	form.AddButton(i18n.T("common.save.short"), func() {
		_, ctx.Manager.ClassName = form.GetFormItemByLabel("Manager Class").(*tview.DropDown).GetCurrentOption()
		ctx.Manager.MaxActiveSessions, _ = strconv.Atoi(form.GetFormItemByLabel("Max Active Sessions").(*tview.InputField).GetText())
		ctx.Manager.SessionIdLength, _ = strconv.Atoi(form.GetFormItemByLabel("Session ID Length").(*tview.InputField).GetText())
		ctx.Manager.MaxInactiveInterval, _ = strconv.Atoi(form.GetFormItemByLabel("Max Inactive Interval (sec)").(*tview.InputField).GetText())
		ctx.Manager.Pathname = form.GetFormItemByLabel("Session File Path").(*tview.InputField).GetText()
		ctx.Manager.ProcessExpiresFrequency, _ = strconv.Atoi(form.GetFormItemByLabel("Process Expires Frequency").(*tview.InputField).GetText())

		// Remove manager if all fields are empty/default
		if ctx.Manager.ClassName == "" && ctx.Manager.MaxActiveSessions == 0 &&
			ctx.Manager.SessionIdLength == 0 && ctx.Manager.MaxInactiveInterval == 0 {
			ctx.Manager = nil
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Session manager saved successfully[-]")
		v.showContextForm(host, ctx, false)
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showContextForm(host, ctx, false)
	})

	form.SetBorder(true).SetTitle(" Session Manager ").SetBorderColor(tcell.ColorBlue)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showContextForm(host, ctx, false)
			return nil
		}
		return event
	})

	// Initial preview
	updatePreview()

	// Layout with form and preview
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("manager-form", layout, true)
	v.app.SetFocus(form)
}

// confirmDelete shows a confirmation dialog
func (v *HostView) confirmDelete(itemType, name string, onConfirm func()) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Delete %s '%s'?", itemType, name)).
		AddButtons([]string{"Delete", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Delete" {
				onConfirm()
			} else {
				v.showHostList()
			}
		})
	v.pages.AddAndSwitchToPage("confirm-delete", modal, true)
}

// setStatus updates the status bar
func (v *HostView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(fmt.Sprintf(" %s", message))
	}
}

// acceptNumber is an input field accept function for numbers
func acceptNumber(text string, lastChar rune) bool {
	if lastChar == '-' && len(text) == 1 {
		return true
	}
	return lastChar >= '0' && lastChar <= '9'
}
