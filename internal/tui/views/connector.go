package views

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/connector"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// ConnectorView handles connector configuration UI
type ConnectorView struct {
	app           *tview.Application
	pages         *tview.Pages
	configService *server.ConfigService
	onBack        func()
	statusBar     *tview.TextView
}

// NewConnectorView creates a new connector configuration view
func NewConnectorView(app *tview.Application, pages *tview.Pages, configService *server.ConfigService, statusBar *tview.TextView, onBack func()) *ConnectorView {
	return &ConnectorView{
		app:           app,
		pages:         pages,
		configService: configService,
		onBack:        onBack,
		statusBar:     statusBar,
	}
}

// Show displays the connector configuration menu
func (v *ConnectorView) Show() {
	list := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	// Get all connectors from all services
	services := v.configService.GetServices()

	// HTTP Connectors
	httpCount := v.countConnectorsByType(services, connector.ConnectorTypeHTTP)
	list.AddItem(
		fmt.Sprintf("[::b]"+i18n.T("connector.http")+"[::-] [yellow](%d)[-]", httpCount),
		i18n.T("connector.http.desc"),
		'h',
		func() { v.showHTTPConnectors() },
	)

	// AJP Connectors
	ajpCount := v.countConnectorsByType(services, connector.ConnectorTypeAJP)
	list.AddItem(
		fmt.Sprintf("[::b]"+i18n.T("connector.ajp")+"[::-] [yellow](%d)[-]", ajpCount),
		i18n.T("connector.ajp.desc"),
		'a',
		func() { v.showAJPConnectors() },
	)

	// SSL/TLS Configuration
	sslCount := v.countSSLConnectors(services)
	list.AddItem(
		fmt.Sprintf("[::b]"+i18n.T("connector.ssl")+"[::-] [yellow](%d)[-]", sslCount),
		i18n.T("connector.ssl.desc"),
		's',
		func() { v.showSSLConnectors() },
	)

	// Executors
	executorCount := 0
	for _, svc := range services {
		executorCount += len(svc.Executors)
	}
	list.AddItem(
		fmt.Sprintf("[::b]"+i18n.T("connector.executor")+"[::-] [yellow](%d)[-]", executorCount),
		i18n.T("connector.executor.desc"),
		'e',
		func() { v.showExecutors() },
	)

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 'b', v.onBack)

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.connector.http")
		case 1:
			helpPanel.SetHelpKey("help.connector.ajp")
		case 2:
			helpPanel.SetHelpKey("help.connector.https")
		case 3:
			helpPanel.SetHelpKey("help.server.executor")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.connector.http")

	list.SetBorder(true).SetTitle(" " + i18n.T("connector.title") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Create flex layout with list and help panel
	flex := tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("connector-config", flex, true)
	v.app.SetFocus(list)
}

func (v *ConnectorView) countConnectorsByType(services []server.Service, connType connector.ConnectorType) int {
	count := 0
	for _, svc := range services {
		for _, conn := range svc.Connectors {
			if connector.GetConnectorType(conn.Protocol) == connType && !conn.SSLEnabled {
				count++
			}
		}
	}
	return count
}

func (v *ConnectorView) countSSLConnectors(services []server.Service) int {
	count := 0
	for _, svc := range services {
		for _, conn := range svc.Connectors {
			if conn.SSLEnabled {
				count++
			}
		}
	}
	return count
}

// showHTTPConnectors shows HTTP connector list
func (v *ConnectorView) showHTTPConnectors() {
	list := tview.NewList().ShowSecondaryText(true)

	services := v.configService.GetServices()
	for svcIdx, svc := range services {
		for connIdx, conn := range svc.Connectors {
			connType := connector.GetConnectorType(conn.Protocol)
			if connType == connector.ConnectorTypeHTTP && !conn.SSLEnabled {
				si, ci := svcIdx, connIdx
				protocol := connector.GetProtocolDescription(conn.Protocol)
				list.AddItem(
					fmt.Sprintf("Port [yellow]%d[-] - %s", conn.Port, protocol),
					fmt.Sprintf("Service: %s, Threads: %d-%d", svc.Name, conn.MinSpareThreads, conn.MaxThreads),
					0,
					func() { v.showConnectorDetail(si, ci) },
				)
			}
		}
	}

	list.AddItem("[green]+ "+i18n.T("connector.http.add")+"[-]", i18n.T("connector.http.add.desc"), 'n', func() {
		v.showAddConnector(connector.ConnectorTypeHTTP)
	})

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("connector.returnmenu"), 'b', func() {
		v.Show()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("connector.http") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("http-connectors", list, true)
	v.app.SetFocus(list)
}

// showAJPConnectors shows AJP connector list
func (v *ConnectorView) showAJPConnectors() {
	list := tview.NewList().ShowSecondaryText(true)

	services := v.configService.GetServices()
	for svcIdx, svc := range services {
		for connIdx, conn := range svc.Connectors {
			if connector.GetConnectorType(conn.Protocol) == connector.ConnectorTypeAJP {
				si, ci := svcIdx, connIdx
				protocol := connector.GetProtocolDescription(conn.Protocol)
				secretStatus := "No Secret"
				if conn.Secret != "" {
					secretStatus = "Secret Set"
				}
				list.AddItem(
					fmt.Sprintf("Port [yellow]%d[-] - %s", conn.Port, protocol),
					fmt.Sprintf("Service: %s, %s", svc.Name, secretStatus),
					0,
					func() { v.showAJPConnectorDetail(si, ci) },
				)
			}
		}
	}

	list.AddItem("[green]+ "+i18n.T("connector.ajp.add")+"[-]", i18n.T("connector.ajp.add.desc"), 'n', func() {
		v.showAddConnector(connector.ConnectorTypeAJP)
	})

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("connector.returnmenu"), 'b', func() {
		v.Show()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("connector.ajp") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("ajp-connectors", list, true)
	v.app.SetFocus(list)
}

// showSSLConnectors shows SSL/TLS connector list
func (v *ConnectorView) showSSLConnectors() {
	list := tview.NewList().ShowSecondaryText(true)

	services := v.configService.GetServices()
	for svcIdx, svc := range services {
		for connIdx, conn := range svc.Connectors {
			if conn.SSLEnabled {
				si, ci := svcIdx, connIdx
				keystoreInfo := "Keystore not configured"
				if conn.KeystoreFile != "" {
					keystoreInfo = fmt.Sprintf("Keystore: %s", conn.KeystoreFile)
				}
				list.AddItem(
					fmt.Sprintf("Port [yellow]%d[-] - HTTPS", conn.Port),
					fmt.Sprintf("Service: %s, %s", svc.Name, keystoreInfo),
					0,
					func() { v.showSSLConnectorDetail(si, ci) },
				)
			}
		}
	}

	list.AddItem("[green]+ "+i18n.T("connector.ssl.add")+"[-]", i18n.T("connector.ssl.add.desc"), 'n', func() {
		v.showAddSSLConnector()
	})

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("connector.returnmenu"), 'b', func() {
		v.Show()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("connector.ssl") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("ssl-connectors", list, true)
	v.app.SetFocus(list)
}

// showConnectorDetail shows HTTP connector detail form
func (v *ConnectorView) showConnectorDetail(serviceIndex, connectorIndex int) {
	svc := v.configService.GetService(serviceIndex)
	if svc == nil || connectorIndex >= len(svc.Connectors) {
		return
	}

	conn := &svc.Connectors[connectorIndex]

	form := tview.NewForm()
	preview := NewPreviewPanel()
	formReady := false

	// Function to update preview based on current form values
	updatePreview := func() {
		if !formReady {
			return
		}
		tempConn := server.Connector{
			Port:              conn.Port,
			Protocol:          conn.Protocol,
			ConnectionTimeout: conn.ConnectionTimeout,
			RedirectPort:      conn.RedirectPort,
			MaxThreads:        conn.MaxThreads,
			MinSpareThreads:   conn.MinSpareThreads,
			AcceptCount:       conn.AcceptCount,
			Executor:          conn.Executor,
		}

		// Get current form values
		if port, err := strconv.Atoi(form.GetFormItem(0).(*tview.InputField).GetText()); err == nil {
			tempConn.Port = port
		}
		_, tempConn.Protocol = form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
		if timeout, err := strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText()); err == nil {
			tempConn.ConnectionTimeout = timeout
		}
		if redirect, err := strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText()); err == nil {
			tempConn.RedirectPort = redirect
		}
		if maxThreads, err := strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText()); err == nil {
			tempConn.MaxThreads = maxThreads
		}
		if minThreads, err := strconv.Atoi(form.GetFormItem(5).(*tview.InputField).GetText()); err == nil {
			tempConn.MinSpareThreads = minThreads
		}
		if acceptCount, err := strconv.Atoi(form.GetFormItem(6).(*tview.InputField).GetText()); err == nil {
			tempConn.AcceptCount = acceptCount
		}
		tempConn.Executor = form.GetFormItem(7).(*tview.InputField).GetText()

		preview.SetXMLPreview(GenerateConnectorXML(&tempConn))
	}

	// Basic settings with change handlers
	form.AddInputField(i18n.T("connector.port"), strconv.Itoa(conn.Port), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddDropDown(i18n.T("connector.protocol"), connector.AvailableHTTPProtocols(), v.getProtocolIndex(conn.Protocol, connector.AvailableHTTPProtocols()), func(text string, index int) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.timeout"), strconv.Itoa(conn.ConnectionTimeout), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.redirect"), strconv.Itoa(conn.RedirectPort), 10, acceptDigits, func(text string) {
		updatePreview()
	})

	// Thread settings
	form.AddInputField(i18n.T("connector.maxthreads"), strconv.Itoa(conn.MaxThreads), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.minthreads"), strconv.Itoa(conn.MinSpareThreads), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.acceptcount"), strconv.Itoa(conn.AcceptCount), 10, acceptDigits, func(text string) {
		updatePreview()
	})

	// Executor reference
	form.AddInputField(i18n.T("connector.executor.optional"), conn.Executor, 30, nil, func(text string) {
		updatePreview()
	})

	formReady = true
	updatePreview()

	form.AddButton(i18n.T("common.save.short"), func() {
		conn.Port, _ = strconv.Atoi(form.GetFormItem(0).(*tview.InputField).GetText())
		_, protocol := form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
		conn.Protocol = protocol
		conn.ConnectionTimeout, _ = strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText())
		conn.RedirectPort, _ = strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())
		conn.MaxThreads, _ = strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText())
		conn.MinSpareThreads, _ = strconv.Atoi(form.GetFormItem(5).(*tview.InputField).GetText())
		conn.AcceptCount, _ = strconv.Atoi(form.GetFormItem(6).(*tview.InputField).GetText())
		conn.Executor = form.GetFormItem(7).(*tview.InputField).GetText()

		v.configService.UpdateService(serviceIndex, *svc)
		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("connector.updated.http") + "[-]")
		v.showHTTPConnectors()
	})

	form.AddButton(i18n.T("common.delete"), func() {
		v.showConfirm(i18n.T("connector.delete.title"), fmt.Sprintf(i18n.T("connector.delete.confirm"), conn.Port), func(confirmed bool) {
			if confirmed {
				svc.Connectors = append(svc.Connectors[:connectorIndex], svc.Connectors[connectorIndex+1:]...)
				v.configService.UpdateService(serviceIndex, *svc)
				if err := v.configService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]" + i18n.T("connector.deleted") + "[-]")
			}
			v.showHTTPConnectors()
		})
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showHTTPConnectors()
	})

	form.SetBorder(true).SetTitle(fmt.Sprintf(" %s - %s %d ", i18n.T("connector.http"), i18n.T("connector.port"), conn.Port)).SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("http-connector-detail", layout, true)
	v.app.SetFocus(form)
}

// showAJPConnectorDetail shows AJP connector detail form
func (v *ConnectorView) showAJPConnectorDetail(serviceIndex, connectorIndex int) {
	svc := v.configService.GetService(serviceIndex)
	if svc == nil || connectorIndex >= len(svc.Connectors) {
		return
	}

	conn := &svc.Connectors[connectorIndex]

	form := tview.NewForm()
	preview := NewPreviewPanel()
	formReady := false

	// Function to update preview
	updatePreview := func() {
		if !formReady {
			return
		}
		tempConn := server.Connector{
			Port:           conn.Port,
			Protocol:       conn.Protocol,
			RedirectPort:   conn.RedirectPort,
			SecretRequired: conn.SecretRequired,
			Secret:         conn.Secret,
			Executor:       conn.Executor,
		}

		if port, err := strconv.Atoi(form.GetFormItem(0).(*tview.InputField).GetText()); err == nil {
			tempConn.Port = port
		}
		_, tempConn.Protocol = form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
		if redirect, err := strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText()); err == nil {
			tempConn.RedirectPort = redirect
		}
		secretReqIdx, _ := form.GetFormItem(3).(*tview.DropDown).GetCurrentOption()
		tempConn.SecretRequired = secretReqIdx == 0
		tempConn.Secret = form.GetFormItem(4).(*tview.InputField).GetText()
		tempConn.Executor = form.GetFormItem(5).(*tview.InputField).GetText()

		preview.SetXMLPreview(GenerateConnectorXML(&tempConn))
	}

	// Basic settings
	form.AddInputField(i18n.T("connector.port"), strconv.Itoa(conn.Port), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddDropDown(i18n.T("connector.protocol"), connector.AvailableAJPProtocols(), v.getProtocolIndex(conn.Protocol, connector.AvailableAJPProtocols()), func(text string, index int) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.redirect"), strconv.Itoa(conn.RedirectPort), 10, acceptDigits, func(text string) {
		updatePreview()
	})

	// Security settings
	secretRequired := i18n.T("common.no")
	if conn.SecretRequired {
		secretRequired = i18n.T("common.yes")
	}
	form.AddDropDown(i18n.T("connector.secretrequired"), []string{i18n.T("common.yes"), i18n.T("common.no")}, indexOf(secretRequired, []string{i18n.T("common.yes"), i18n.T("common.no")}), func(text string, index int) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.secret"), conn.Secret, 40, nil, func(text string) {
		updatePreview()
	})

	// Executor reference
	form.AddInputField(i18n.T("connector.executor.optional"), conn.Executor, 30, nil, func(text string) {
		updatePreview()
	})

	formReady = true
	updatePreview()

	form.AddButton(i18n.T("common.save.short"), func() {
		conn.Port, _ = strconv.Atoi(form.GetFormItem(0).(*tview.InputField).GetText())
		_, protocol := form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
		conn.Protocol = protocol
		conn.RedirectPort, _ = strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText())
		secretReqIdx, _ := form.GetFormItem(3).(*tview.DropDown).GetCurrentOption()
		conn.SecretRequired = secretReqIdx == 0
		conn.Secret = form.GetFormItem(4).(*tview.InputField).GetText()
		conn.Executor = form.GetFormItem(5).(*tview.InputField).GetText()

		v.configService.UpdateService(serviceIndex, *svc)
		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("connector.updated.ajp") + "[-]")
		v.showAJPConnectors()
	})

	form.AddButton(i18n.T("common.delete"), func() {
		v.showConfirm(i18n.T("connector.delete.title"), fmt.Sprintf(i18n.T("connector.delete.ajp.confirm"), conn.Port), func(confirmed bool) {
			if confirmed {
				svc.Connectors = append(svc.Connectors[:connectorIndex], svc.Connectors[connectorIndex+1:]...)
				v.configService.UpdateService(serviceIndex, *svc)
				if err := v.configService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]" + i18n.T("connector.deleted") + "[-]")
			}
			v.showAJPConnectors()
		})
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showAJPConnectors()
	})

	form.SetBorder(true).SetTitle(fmt.Sprintf(" %s - %s %d ", i18n.T("connector.ajp"), i18n.T("connector.port"), conn.Port)).SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("ajp-connector-detail", layout, true)
	v.app.SetFocus(form)
}

// showSSLConnectorDetail shows SSL connector detail form
func (v *ConnectorView) showSSLConnectorDetail(serviceIndex, connectorIndex int) {
	svc := v.configService.GetService(serviceIndex)
	if svc == nil || connectorIndex >= len(svc.Connectors) {
		return
	}

	conn := &svc.Connectors[connectorIndex]

	form := tview.NewForm()
	preview := NewPreviewPanel()
	formReady := false

	// Function to update preview
	updatePreview := func() {
		if !formReady {
			return
		}
		tempConn := server.Connector{
			Port:              conn.Port,
			Protocol:          conn.Protocol,
			ConnectionTimeout: conn.ConnectionTimeout,
			MaxThreads:        conn.MaxThreads,
			MinSpareThreads:   conn.MinSpareThreads,
			SSLEnabled:        true,
			Scheme:            "https",
			Secure:            true,
			SSLProtocol:       conn.SSLProtocol,
			KeystoreFile:      conn.KeystoreFile,
			KeystoreType:      conn.KeystoreType,
			ClientAuth:        conn.ClientAuth,
		}

		if port, err := strconv.Atoi(form.GetFormItem(0).(*tview.InputField).GetText()); err == nil {
			tempConn.Port = port
		}
		_, tempConn.Protocol = form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
		if timeout, err := strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText()); err == nil {
			tempConn.ConnectionTimeout = timeout
		}
		if maxThreads, err := strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText()); err == nil {
			tempConn.MaxThreads = maxThreads
		}
		if minThreads, err := strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText()); err == nil {
			tempConn.MinSpareThreads = minThreads
		}
		_, tempConn.SSLProtocol = form.GetFormItem(5).(*tview.DropDown).GetCurrentOption()
		tempConn.KeystoreFile = form.GetFormItem(6).(*tview.InputField).GetText()
		// Password not shown in preview for security
		_, tempConn.KeystoreType = form.GetFormItem(8).(*tview.DropDown).GetCurrentOption()
		_, tempConn.ClientAuth = form.GetFormItem(9).(*tview.DropDown).GetCurrentOption()

		preview.SetXMLPreview(GenerateConnectorXML(&tempConn))
	}

	// Basic settings
	form.AddInputField(i18n.T("connector.port"), strconv.Itoa(conn.Port), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddDropDown(i18n.T("connector.protocol"), connector.AvailableHTTPProtocols(), v.getProtocolIndex(conn.Protocol, connector.AvailableHTTPProtocols()), func(text string, index int) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.timeout"), strconv.Itoa(conn.ConnectionTimeout), 10, acceptDigits, func(text string) {
		updatePreview()
	})

	// Thread settings
	form.AddInputField(i18n.T("connector.maxthreads"), strconv.Itoa(conn.MaxThreads), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.minthreads"), strconv.Itoa(conn.MinSpareThreads), 10, acceptDigits, func(text string) {
		updatePreview()
	})

	// SSL settings
	form.AddDropDown(i18n.T("connector.sslprotocol"), connector.SSLProtocols(), indexOf(conn.SSLProtocol, connector.SSLProtocols()), func(text string, index int) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.keystorefile"), conn.KeystoreFile, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddPasswordField(i18n.T("connector.keystorepass"), conn.KeystorePass, 30, '*', nil)
	form.AddDropDown(i18n.T("connector.keystoretype"), connector.KeystoreTypes(), indexOf(conn.KeystoreType, connector.KeystoreTypes()), func(text string, index int) {
		updatePreview()
	})
	form.AddDropDown(i18n.T("connector.clientauth"), connector.ClientAuthOptions(), indexOf(conn.ClientAuth, connector.ClientAuthOptions()), func(text string, index int) {
		updatePreview()
	})

	formReady = true
	updatePreview()

	form.AddButton(i18n.T("common.save.short"), func() {
		conn.Port, _ = strconv.Atoi(form.GetFormItem(0).(*tview.InputField).GetText())
		_, protocol := form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
		conn.Protocol = protocol
		conn.ConnectionTimeout, _ = strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText())
		conn.MaxThreads, _ = strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())
		conn.MinSpareThreads, _ = strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText())
		_, conn.SSLProtocol = form.GetFormItem(5).(*tview.DropDown).GetCurrentOption()
		conn.KeystoreFile = form.GetFormItem(6).(*tview.InputField).GetText()
		conn.KeystorePass = form.GetFormItem(7).(*tview.InputField).GetText()
		_, conn.KeystoreType = form.GetFormItem(8).(*tview.DropDown).GetCurrentOption()
		_, conn.ClientAuth = form.GetFormItem(9).(*tview.DropDown).GetCurrentOption()

		v.configService.UpdateService(serviceIndex, *svc)
		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("connector.updated.ssl") + "[-]")
		v.showSSLConnectors()
	})

	form.AddButton(i18n.T("common.delete"), func() {
		v.showConfirm(i18n.T("connector.delete.title"), fmt.Sprintf(i18n.T("connector.delete.ssl.confirm"), conn.Port), func(confirmed bool) {
			if confirmed {
				svc.Connectors = append(svc.Connectors[:connectorIndex], svc.Connectors[connectorIndex+1:]...)
				v.configService.UpdateService(serviceIndex, *svc)
				if err := v.configService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]" + i18n.T("connector.deleted") + "[-]")
			}
			v.showSSLConnectors()
		})
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showSSLConnectors()
	})

	form.SetBorder(true).SetTitle(fmt.Sprintf(" %s - %s %d ", i18n.T("connector.ssl"), i18n.T("connector.port"), conn.Port)).SetBorderColor(tcell.ColorDarkCyan)

	// Initial preview
	updatePreview()

	// Create layout with form on top and preview on bottom
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	v.pages.AddAndSwitchToPage("ssl-connector-detail", layout, true)
	v.app.SetFocus(form)
}

// showAddConnector shows form to add a new connector
func (v *ConnectorView) showAddConnector(connType connector.ConnectorType) {
	var defaultConn server.Connector
	var protocols []string
	var title string

	switch connType {
	case connector.ConnectorTypeHTTP:
		defaultConn = connector.DefaultHTTPConnector()
		protocols = connector.AvailableHTTPProtocols()
		title = i18n.T("connector.http.add.title")
	case connector.ConnectorTypeAJP:
		defaultConn = connector.DefaultAJPConnector()
		protocols = connector.AvailableAJPProtocols()
		title = i18n.T("connector.ajp.add.title")
	default:
		return
	}

	// Select service first
	services := v.configService.GetServices()
	if len(services) == 0 {
		v.showError(i18n.T("connector.error.noservices"))
		return
	}

	serviceNames := make([]string, len(services))
	for i, svc := range services {
		serviceNames[i] = svc.Name
	}

	form := tview.NewForm()
	form.AddDropDown(i18n.T("server.service"), serviceNames, 0, nil)
	form.AddInputField(i18n.T("connector.port"), strconv.Itoa(defaultConn.Port), 10, acceptDigits, nil)
	form.AddDropDown(i18n.T("connector.protocol"), protocols, 0, nil)

	if connType == connector.ConnectorTypeHTTP {
		form.AddInputField(i18n.T("connector.timeout"), strconv.Itoa(defaultConn.ConnectionTimeout), 10, acceptDigits, nil)
		form.AddInputField(i18n.T("connector.maxthreads"), strconv.Itoa(defaultConn.MaxThreads), 10, acceptDigits, nil)
		form.AddInputField(i18n.T("connector.minthreads"), strconv.Itoa(defaultConn.MinSpareThreads), 10, acceptDigits, nil)
	} else if connType == connector.ConnectorTypeAJP {
		form.AddDropDown(i18n.T("connector.secretrequired"), []string{i18n.T("common.yes"), i18n.T("common.no")}, 0, nil)
		form.AddInputField(i18n.T("connector.secret"), "", 40, nil, nil)
	}

	form.AddInputField(i18n.T("connector.redirect"), strconv.Itoa(defaultConn.RedirectPort), 10, acceptDigits, nil)

	form.AddButton(i18n.T("common.add"), func() {
		svcIdx, _ := form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()
		svc := v.configService.GetService(svcIdx)
		if svc == nil {
			return
		}

		newConn := defaultConn
		newConn.Port, _ = strconv.Atoi(form.GetFormItem(1).(*tview.InputField).GetText())
		_, newConn.Protocol = form.GetFormItem(2).(*tview.DropDown).GetCurrentOption()

		if connType == connector.ConnectorTypeHTTP {
			newConn.ConnectionTimeout, _ = strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())
			newConn.MaxThreads, _ = strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText())
			newConn.MinSpareThreads, _ = strconv.Atoi(form.GetFormItem(5).(*tview.InputField).GetText())
			newConn.RedirectPort, _ = strconv.Atoi(form.GetFormItem(6).(*tview.InputField).GetText())
		} else if connType == connector.ConnectorTypeAJP {
			secretReqIdx, _ := form.GetFormItem(3).(*tview.DropDown).GetCurrentOption()
			newConn.SecretRequired = secretReqIdx == 0
			newConn.Secret = form.GetFormItem(4).(*tview.InputField).GetText()
			newConn.RedirectPort, _ = strconv.Atoi(form.GetFormItem(5).(*tview.InputField).GetText())
		}

		svc.Connectors = append(svc.Connectors, newConn)
		v.configService.UpdateService(svcIdx, *svc)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("connector.added") + "[-]")
		if connType == connector.ConnectorTypeHTTP {
			v.showHTTPConnectors()
		} else {
			v.showAJPConnectors()
		}
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		if connType == connector.ConnectorTypeHTTP {
			v.showHTTPConnectors()
		} else {
			v.showAJPConnectors()
		}
	})

	form.SetBorder(true).SetTitle(fmt.Sprintf(" %s ", title)).SetBorderColor(tcell.ColorGreen)
	v.pages.AddAndSwitchToPage("add-connector", form, true)
	v.app.SetFocus(form)
}

// showAddSSLConnector shows form to add a new SSL connector
func (v *ConnectorView) showAddSSLConnector() {
	defaultConn := connector.DefaultHTTPSConnector()

	services := v.configService.GetServices()
	if len(services) == 0 {
		v.showError(i18n.T("connector.error.noservices"))
		return
	}

	serviceNames := make([]string, len(services))
	for i, svc := range services {
		serviceNames[i] = svc.Name
	}

	form := tview.NewForm()
	form.AddDropDown(i18n.T("server.service"), serviceNames, 0, nil)
	form.AddInputField(i18n.T("connector.port"), strconv.Itoa(defaultConn.Port), 10, acceptDigits, nil)
	form.AddDropDown(i18n.T("connector.protocol"), connector.AvailableHTTPProtocols(), 0, nil)
	form.AddInputField(i18n.T("connector.timeout"), strconv.Itoa(defaultConn.ConnectionTimeout), 10, acceptDigits, nil)
	form.AddInputField(i18n.T("connector.maxthreads"), strconv.Itoa(defaultConn.MaxThreads), 10, acceptDigits, nil)

	// SSL settings
	form.AddDropDown(i18n.T("connector.sslprotocol"), connector.SSLProtocols(), 0, nil)
	form.AddInputField(i18n.T("connector.keystorefile"), defaultConn.KeystoreFile, 50, nil, nil)
	form.AddPasswordField(i18n.T("connector.keystorepass"), defaultConn.KeystorePass, 30, '*', nil)
	form.AddDropDown(i18n.T("connector.keystoretype"), connector.KeystoreTypes(), 0, nil)
	form.AddDropDown(i18n.T("connector.clientauth"), connector.ClientAuthOptions(), 0, nil)

	form.AddButton(i18n.T("common.add"), func() {
		svcIdx, _ := form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()
		svc := v.configService.GetService(svcIdx)
		if svc == nil {
			return
		}

		newConn := defaultConn
		newConn.Port, _ = strconv.Atoi(form.GetFormItem(1).(*tview.InputField).GetText())
		_, newConn.Protocol = form.GetFormItem(2).(*tview.DropDown).GetCurrentOption()
		newConn.ConnectionTimeout, _ = strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())
		newConn.MaxThreads, _ = strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText())
		_, newConn.SSLProtocol = form.GetFormItem(5).(*tview.DropDown).GetCurrentOption()
		newConn.KeystoreFile = form.GetFormItem(6).(*tview.InputField).GetText()
		newConn.KeystorePass = form.GetFormItem(7).(*tview.InputField).GetText()
		_, newConn.KeystoreType = form.GetFormItem(8).(*tview.DropDown).GetCurrentOption()
		_, newConn.ClientAuth = form.GetFormItem(9).(*tview.DropDown).GetCurrentOption()

		svc.Connectors = append(svc.Connectors, newConn)
		v.configService.UpdateService(svcIdx, *svc)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("connector.added.ssl") + "[-]")
		v.showSSLConnectors()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showSSLConnectors()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("connector.ssl.add.title") + " ").SetBorderColor(tcell.ColorGreen)
	v.pages.AddAndSwitchToPage("add-ssl-connector", form, true)
	v.app.SetFocus(form)
}

// showExecutors shows executor list
func (v *ConnectorView) showExecutors() {
	list := tview.NewList().ShowSecondaryText(true)

	services := v.configService.GetServices()
	for svcIdx, svc := range services {
		for execIdx, exec := range svc.Executors {
			si, ei := svcIdx, execIdx
			list.AddItem(
				fmt.Sprintf("[yellow]%s[-]", exec.Name),
				fmt.Sprintf("Service: %s, Threads: %d-%d, MaxIdle: %dms", svc.Name, exec.MinSpareThreads, exec.MaxThreads, exec.MaxIdleTime),
				0,
				func() { v.showExecutorDetail(si, ei) },
			)
		}
	}

	list.AddItem("[green]+ "+i18n.T("connector.executor.add")+"[-]", i18n.T("connector.executor.add.desc"), 'n', func() {
		v.showAddExecutor()
	})

	list.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("connector.returnmenu"), 'b', func() {
		v.Show()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("connector.executor.title") + " ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("executors", list, true)
	v.app.SetFocus(list)
}

// showExecutorDetail shows executor detail form
func (v *ConnectorView) showExecutorDetail(serviceIndex, executorIndex int) {
	svc := v.configService.GetService(serviceIndex)
	if svc == nil || executorIndex >= len(svc.Executors) {
		return
	}

	exec := &svc.Executors[executorIndex]

	form := tview.NewForm()
	preview := NewPreviewPanel()

	// Function to update preview based on current form values
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

	form.AddInputField(i18n.T("connector.executor.name"), exec.Name, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.executor.nameprefix"), exec.NamePrefix, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.maxthreads"), strconv.Itoa(exec.MaxThreads), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.minthreads"), strconv.Itoa(exec.MinSpareThreads), 10, acceptDigits, func(text string) {
		updatePreview()
	})
	form.AddInputField(i18n.T("connector.executor.maxidle"), strconv.Itoa(exec.MaxIdleTime), 10, acceptDigits, func(text string) {
		updatePreview()
	})

	form.AddButton(i18n.T("common.save.short"), func() {
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

		v.setStatus("[green]" + i18n.T("connector.executor.updated") + "[-]")
		v.showExecutors()
	})

	form.AddButton(i18n.T("common.delete"), func() {
		v.showConfirm(i18n.T("connector.executor.delete.title"), fmt.Sprintf(i18n.T("connector.executor.delete.confirm"), exec.Name), func(confirmed bool) {
			if confirmed {
				svc.Executors = append(svc.Executors[:executorIndex], svc.Executors[executorIndex+1:]...)
				v.configService.UpdateService(serviceIndex, *svc)
				if err := v.configService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]" + i18n.T("connector.executor.deleted") + "[-]")
			}
			v.showExecutors()
		})
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showExecutors()
	})

	form.SetBorder(true).SetTitle(fmt.Sprintf(" %s: %s ", i18n.T("connector.executor"), exec.Name)).SetBorderColor(tcell.ColorDarkCyan)

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

// showAddExecutor shows form to add a new executor
func (v *ConnectorView) showAddExecutor() {
	services := v.configService.GetServices()
	if len(services) == 0 {
		v.showError(i18n.T("connector.error.noservices"))
		return
	}

	serviceNames := make([]string, len(services))
	for i, svc := range services {
		serviceNames[i] = svc.Name
	}

	form := tview.NewForm()
	form.AddDropDown(i18n.T("server.service"), serviceNames, 0, nil)
	form.AddInputField(i18n.T("connector.executor.name"), "tomcatThreadPool", 30, nil, nil)
	form.AddInputField(i18n.T("connector.executor.nameprefix"), "catalina-exec-", 30, nil, nil)
	form.AddInputField(i18n.T("connector.maxthreads"), "200", 10, acceptDigits, nil)
	form.AddInputField(i18n.T("connector.minthreads"), "25", 10, acceptDigits, nil)
	form.AddInputField(i18n.T("connector.executor.maxidle"), "60000", 10, acceptDigits, nil)

	form.AddButton(i18n.T("common.add"), func() {
		svcIdx, _ := form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()
		svc := v.configService.GetService(svcIdx)
		if svc == nil {
			return
		}

		maxThreads, _ := strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())
		minSpare, _ := strconv.Atoi(form.GetFormItem(4).(*tview.InputField).GetText())
		maxIdle, _ := strconv.Atoi(form.GetFormItem(5).(*tview.InputField).GetText())

		exec := server.Executor{
			Name:            form.GetFormItem(1).(*tview.InputField).GetText(),
			NamePrefix:      form.GetFormItem(2).(*tview.InputField).GetText(),
			MaxThreads:      maxThreads,
			MinSpareThreads: minSpare,
			MaxIdleTime:     maxIdle,
		}

		svc.Executors = append(svc.Executors, exec)
		v.configService.UpdateService(svcIdx, *svc)

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}

		v.setStatus("[green]" + i18n.T("connector.executor.added") + "[-]")
		v.showExecutors()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showExecutors()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("connector.executor.add.title") + " ").SetBorderColor(tcell.ColorGreen)
	v.pages.AddAndSwitchToPage("add-executor", form, true)
	v.app.SetFocus(form)
}

// Helper functions
func (v *ConnectorView) getProtocolIndex(protocol string, protocols []string) int {
	for i, p := range protocols {
		if p == protocol {
			return i
		}
	}
	return 0
}

func indexOf(value string, list []string) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return 0
}

func (v *ConnectorView) showError(message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			v.Show()
		})
	modal.SetBackgroundColor(tcell.ColorDarkRed)
	v.pages.AddAndSwitchToPage("error", modal, true)
}

func (v *ConnectorView) showConfirm(title, message string, callback func(bool)) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			callback(buttonIndex == 0)
		})
	modal.SetBorder(true).SetTitle(fmt.Sprintf(" %s ", title))
	v.pages.AddAndSwitchToPage("confirm", modal, true)
}

func (v *ConnectorView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(fmt.Sprintf(" %s", message))
	}
}
