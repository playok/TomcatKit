package views

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// QuickTemplatesView provides TUI for applying quick configuration templates
type QuickTemplatesView struct {
	app           *tview.Application
	pages         *tview.Pages
	mainPages     *tview.Pages
	statusBar     *tview.TextView
	onReturn      func()
	configService *server.ConfigService
	catalinaBase  string
}

// NewQuickTemplatesView creates a new quick templates view
func NewQuickTemplatesView(app *tview.Application, mainPages *tview.Pages, statusBar *tview.TextView, catalinaBase string, onReturn func()) *QuickTemplatesView {
	return &QuickTemplatesView{
		app:           app,
		mainPages:     mainPages,
		statusBar:     statusBar,
		onReturn:      onReturn,
		configService: server.NewConfigService(catalinaBase),
		catalinaBase:  catalinaBase,
	}
}

// Load initializes the view
func (v *QuickTemplatesView) Load() error {
	if err := v.configService.Load(); err != nil {
		return fmt.Errorf("failed to load server.xml: %w", err)
	}

	v.pages = tview.NewPages()
	v.showMainMenu()

	v.mainPages.AddAndSwitchToPage("quicktemplates", v.pages, true)
	return nil
}

// showMainMenu displays the quick templates menu
func (v *QuickTemplatesView) showMainMenu() {
	list := tview.NewList().
		AddItem("[::b]"+i18n.T("qt.virtualthread")+"[::-]", i18n.T("qt.virtualthread.desc"), 'v', func() {
			v.showVirtualThreadTemplate()
		}).
		AddItem("[::b]"+i18n.T("qt.https")+"[::-]", i18n.T("qt.https.desc"), 's', func() {
			v.showHTTPSTemplate()
		}).
		AddItem("[::b]"+i18n.T("qt.connpool")+"[::-]", i18n.T("qt.connpool.desc"), 'p', func() {
			v.showConnectionPoolTemplate()
		}).
		AddItem("[::b]"+i18n.T("qt.gzip")+"[::-]", i18n.T("qt.gzip.desc"), 'g', func() {
			v.showGzipTemplate()
		}).
		AddItem("[::b]"+i18n.T("qt.accesslog")+"[::-]", i18n.T("qt.accesslog.desc"), 'a', func() {
			v.showAccessLogTemplate()
		}).
		AddItem("[::b]"+i18n.T("qt.security")+"[::-]", i18n.T("qt.security.desc"), 'h', func() {
			v.showSecurityHardeningTemplate()
		}).
		AddItem("[::b]"+i18n.T("qt.apache")+"[::-]", i18n.T("qt.apache.desc"), 'j', func() {
			v.showApacheHttpdTemplate()
		}).
		AddItem("[::b]"+i18n.T("qt.nginx")+"[::-]", i18n.T("qt.nginx.desc"), 'n', func() {
			v.showNginxProxyTemplate()
		}).
		AddItem("[::b]"+i18n.T("qt.haproxy")+"[::-]", i18n.T("qt.haproxy.desc"), 'l', func() {
			v.showHAProxyTemplate()
		}).
		AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
			v.mainPages.RemovePage("quicktemplates")
			v.onReturn()
		})

	list.SetBorder(true).SetTitle(" " + i18n.T("qt.title") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.mainPages.RemovePage("quicktemplates")
			v.onReturn()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("menu", list, true)
	v.setStatus(i18n.T("qt.select"))
}

// showVirtualThreadTemplate shows the Virtual Thread configuration template
func (v *QuickTemplatesView) showVirtualThreadTemplate() {
	cfg := v.configService.GetServer()

	// Help panel
	helpPanel := HelpPanel("help.qt.virtualthread")

	// Check if virtual thread executor already exists
	hasVirtualExecutor := false
	for _, svc := range cfg.Services {
		for _, exec := range svc.Executors {
			if exec.IsVirtualThread() {
				hasVirtualExecutor = true
				break
			}
		}
	}

	// Get available connectors for selection
	var connectorOptions []string
	var connectorPorts []int
	if len(cfg.Services) > 0 {
		for _, conn := range cfg.Services[0].Connectors {
			if conn.Protocol == "" || strings.Contains(conn.Protocol, "HTTP") {
				connectorOptions = append(connectorOptions, fmt.Sprintf("Port %d (%s)", conn.Port, conn.Protocol))
				connectorPorts = append(connectorPorts, conn.Port)
			}
		}
	}

	form := tview.NewForm()

	// Warning if executor exists
	if hasVirtualExecutor {
		form.AddTextView("Warning", "[red]Virtual Thread executor already exists![white]", 50, 1, false, false)
	}

	// Executor name
	executorName := "virtualThreadExecutor"
	form.AddInputField("Executor Name", executorName, 30, nil, func(text string) {
		executorName = text
	})

	// Name prefix
	namePrefix := "virt-exec-"
	form.AddInputField("Thread Name Prefix", namePrefix, 20, nil, func(text string) {
		namePrefix = text
	})

	// Max queue size
	maxQueueSize := "100"
	form.AddInputField("Max Queue Size", maxQueueSize, 10, acceptNumber, func(text string) {
		maxQueueSize = text
	})

	// Connector selection
	selectedConnector := 0
	if len(connectorOptions) > 0 {
		form.AddDropDown("Apply to Connector", connectorOptions, 0, func(option string, index int) {
			selectedConnector = index
		})
	}

	form.AddButton("[white:green]Apply Template[-:-]", func() {
		// Create virtual thread executor
		executor := server.NewVirtualThreadExecutor(executorName)
		executor.NamePrefix = namePrefix

		// Add executor to first service
		if len(cfg.Services) > 0 {
			cfg.Services[0].Executors = append(cfg.Services[0].Executors, *executor)

			// Update connector to use the executor
			if len(connectorPorts) > selectedConnector {
				port := connectorPorts[selectedConnector]
				for i := range cfg.Services[0].Connectors {
					if cfg.Services[0].Connectors[i].Port == port {
						cfg.Services[0].Connectors[i].Executor = executorName
						break
					}
				}
			}
		}

		// Save configuration
		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		v.setStatus("Virtual Thread executor applied successfully!")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Virtual Thread Template ")

	// Create layout with form and help panel
	flex := CreateFormWithHelp(form, "help.qt.virtualthread", "")

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	// Remove unused helpPanel reference
	_ = helpPanel

	v.pages.AddAndSwitchToPage("virtual-thread", flex, true)
}

// showHTTPSTemplate shows the HTTPS configuration template
func (v *QuickTemplatesView) showHTTPSTemplate() {
	form := tview.NewForm()

	port := "8443"
	form.AddInputField("HTTPS Port", port, 10, acceptNumber, func(text string) {
		port = text
	})

	keystoreFile := "${catalina.base}/conf/localhost-rsa.jks"
	form.AddInputField("Keystore File", keystoreFile, 50, nil, func(text string) {
		keystoreFile = text
	})

	keystorePass := ""
	form.AddPasswordField("Keystore Password", keystorePass, 30, '*', func(text string) {
		keystorePass = text
	})

	keystoreTypes := []string{"JKS", "PKCS12"}
	keystoreTypeIdx := 0
	form.AddDropDown("Keystore Type", keystoreTypes, keystoreTypeIdx, func(option string, index int) {
		keystoreTypeIdx = index
	})

	form.AddButton("[white:green]Apply Template[-:-]", func() {
		cfg := v.configService.GetServer()

		// Create HTTPS connector
		connector := &server.Connector{
			Port:              8443,
			Protocol:          "org.apache.coyote.http11.Http11NioProtocol",
			SSLEnabled:        true,
			Scheme:            "https",
			Secure:            true,
			KeystoreFile:      keystoreFile,
			KeystorePass:      keystorePass,
			KeystoreType:      keystoreTypes[keystoreTypeIdx],
			MaxThreads:        150,
			ConnectionTimeout: 20000,
		}

		// Parse port
		if p := parsePort(port); p > 0 {
			connector.Port = p
		}

		// Add to first service
		if len(cfg.Services) > 0 {
			cfg.Services[0].Connectors = append(cfg.Services[0].Connectors, *connector)
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		v.setStatus("HTTPS connector added successfully!")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" HTTPS Template ")

	// Create layout with form and help panel
	flex := CreateFormWithHelp(form, "help.qt.https", "")

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("https-template", flex, true)
}

// showConnectionPoolTemplate shows the connection pool tuning template
func (v *QuickTemplatesView) showConnectionPoolTemplate() {
	cfg := v.configService.GetServer()

	form := tview.NewForm()

	profiles := []string{"Development (25-100)", "Production (150-400)", "High Traffic (400-800)", "Custom"}
	form.AddDropDown("Profile", profiles, 1, nil)

	maxThreads := "200"
	form.AddInputField("Max Threads", maxThreads, 10, acceptNumber, func(text string) {
		maxThreads = text
	})

	minSpareThreads := "25"
	form.AddInputField("Min Spare Threads", minSpareThreads, 10, acceptNumber, func(text string) {
		minSpareThreads = text
	})

	acceptCount := "100"
	form.AddInputField("Accept Count", acceptCount, 10, acceptNumber, func(text string) {
		acceptCount = text
	})

	connectionTimeout := "20000"
	form.AddInputField("Connection Timeout (ms)", connectionTimeout, 10, acceptNumber, func(text string) {
		connectionTimeout = text
	})

	form.AddButton("[white:green]Apply to All HTTP Connectors[-:-]", func() {
		maxT := parsePort(maxThreads)
		minS := parsePort(minSpareThreads)
		accC := parsePort(acceptCount)
		connT := parsePort(connectionTimeout)

		// Update all HTTP connectors
		if len(cfg.Services) > 0 {
			for i := range cfg.Services[0].Connectors {
				conn := &cfg.Services[0].Connectors[i]
				if conn.Protocol == "" || strings.Contains(conn.Protocol, "HTTP") {
					conn.MaxThreads = maxT
					conn.MinSpareThreads = minS
					conn.AcceptCount = accC
					conn.ConnectionTimeout = connT
				}
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		v.setStatus("Connection pool settings applied!")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Connection Pool Tuning ")

	// Create layout with form and help panel
	flex := CreateFormWithHelp(form, "help.qt.connpool", "")

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("pool-template", flex, true)
}

// showGzipTemplate shows the Gzip compression template
func (v *QuickTemplatesView) showGzipTemplate() {
	cfg := v.configService.GetServer()

	form := tview.NewForm()

	minSize := "2048"
	form.AddInputField("Min Compression Size (bytes)", minSize, 10, acceptNumber, func(text string) {
		minSize = text
	})

	form.AddButton("[white:green]Apply Template[-:-]", func() {
		minS := parsePort(minSize)

		// Update all HTTP connectors with compression settings
		if len(cfg.Services) > 0 {
			for i := range cfg.Services[0].Connectors {
				conn := &cfg.Services[0].Connectors[i]
				if conn.Protocol == "" || strings.Contains(conn.Protocol, "HTTP") {
					conn.Compression = "on"
					conn.CompressionMinSize = minS
					conn.CompressibleMimeType = "text/html,text/xml,text/plain,text/css,text/javascript,application/javascript,application/json,application/xml"
				}
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		v.setStatus("Gzip compression enabled!")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Gzip Compression ")

	// Create layout with form and help panel
	flex := CreateFormWithHelp(form, "help.qt.gzip", "")

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("gzip-template", flex, true)
}

// showAccessLogTemplate shows the access log configuration template
func (v *QuickTemplatesView) showAccessLogTemplate() {
	form := tview.NewForm()

	// Help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey("help.qt.accesslog")

	// Preview panel
	previewPanel := NewPreviewPanel()

	// Variables for form fields
	directory := "logs"
	prefix := "localhost_access_log"
	suffix := ".txt"
	selectedPattern := 0 // 0=combined, 1=common, 2=custom
	customPattern := "%h %l %u %t \"%r\" %s %b"

	// Function to update XML preview
	updatePreview := func() {
		pattern := "combined"
		switch selectedPattern {
		case 0:
			pattern = "combined"
		case 1:
			pattern = "common"
		case 2:
			pattern = customPattern
		}

		valve := server.Valve{
			ClassName: "org.apache.catalina.valves.AccessLogValve",
			Directory: directory,
			Prefix:    prefix,
			Suffix:    suffix,
			Pattern:   pattern,
			Rotatable: true,
		}
		previewPanel.SetXMLPreview(GenerateValveXML(&valve))
	}

	patterns := []string{"combined (Recommended)", "common", "custom"}
	form.AddDropDown("Log Pattern", patterns, 0, func(option string, index int) {
		selectedPattern = index
		// Show/hide custom pattern field
		if index == 2 {
			// Custom pattern selected - show input field
			if form.GetFormItemByLabel("Custom Pattern") == nil {
				form.AddInputField("Custom Pattern", customPattern, 50, nil, func(text string) {
					customPattern = text
					updatePreview()
				})
			}
		}
		updatePreview()
	})

	form.AddInputField("Directory", directory, 30, nil, func(text string) {
		directory = text
		updatePreview()
	})

	form.AddInputField("File Prefix", prefix, 30, nil, func(text string) {
		prefix = text
		updatePreview()
	})

	form.AddInputField("File Suffix", suffix, 10, nil, func(text string) {
		suffix = text
		updatePreview()
	})

	form.AddButton("[white:green]Apply Template[-:-]", func() {
		cfg := v.configService.GetServer()

		// Determine pattern based on selection
		pattern := "combined"
		switch selectedPattern {
		case 0:
			pattern = "combined"
		case 1:
			pattern = "common"
		case 2:
			pattern = customPattern
		}

		// Create access log valve
		valve := server.Valve{
			ClassName: "org.apache.catalina.valves.AccessLogValve",
			Directory: directory,
			Prefix:    prefix,
			Suffix:    suffix,
			Pattern:   pattern,
			Rotatable: true,
		}

		// Add to default host
		if len(cfg.Services) > 0 {
			for i := range cfg.Services[0].Engine.Hosts {
				cfg.Services[0].Engine.Hosts[i].Valves = append(cfg.Services[0].Engine.Hosts[i].Valves, valve)
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		v.setStatus("Access log configured successfully!")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Access Log Template ")

	// Initial preview
	updatePreview()

	// Create layout with form, help panel, and preview
	flex := CreateFormWithHelpAndPreview(form, helpPanel, previewPanel)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("accesslog-template", flex, true)
}

// showSecurityHardeningTemplate shows security hardening options
func (v *QuickTemplatesView) showSecurityHardeningTemplate() {
	form := tview.NewForm()

	disableShutdown := true
	form.AddCheckbox("Disable Shutdown Port", disableShutdown, func(checked bool) {
		disableShutdown = checked
	})

	removeServerInfo := true
	form.AddCheckbox("Remove Server Info from Errors", removeServerInfo, func(checked bool) {
		removeServerInfo = checked
	})

	addSecurityListener := true
	form.AddCheckbox("Add Security Listener", addSecurityListener, func(checked bool) {
		addSecurityListener = checked
	})

	form.AddButton("[white:green]Apply Template[-:-]", func() {
		cfg := v.configService.GetServer()

		// Disable shutdown port
		if disableShutdown {
			cfg.Port = -1
			cfg.Shutdown = "DISABLED_" + generateRandomString(8)
		}

		// Add security listener
		if addSecurityListener {
			hasSecurityListener := false
			for _, l := range cfg.Listeners {
				if l.ClassName == "org.apache.catalina.security.SecurityListener" {
					hasSecurityListener = true
					break
				}
			}
			if !hasSecurityListener {
				cfg.Listeners = append(cfg.Listeners, server.Listener{
					ClassName: "org.apache.catalina.security.SecurityListener",
				})
			}
		}

		// Add error report valve with showServerInfo=false
		if removeServerInfo && len(cfg.Services) > 0 {
			for i := range cfg.Services[0].Engine.Hosts {
				hasErrorValve := false
				for j := range cfg.Services[0].Engine.Hosts[i].Valves {
					if cfg.Services[0].Engine.Hosts[i].Valves[j].ClassName == "org.apache.catalina.valves.ErrorReportValve" {
						cfg.Services[0].Engine.Hosts[i].Valves[j].ShowServerInfo = false
						hasErrorValve = true
					}
				}
				if !hasErrorValve {
					cfg.Services[0].Engine.Hosts[i].Valves = append(cfg.Services[0].Engine.Hosts[i].Valves, server.Valve{
						ClassName:      "org.apache.catalina.valves.ErrorReportValve",
						ShowServerInfo: false,
					})
				}
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		v.setStatus("Security hardening applied!")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Security Hardening ")

	// Create layout with form and help panel
	flex := CreateFormWithHelp(form, "help.qt.security", "")

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("security-template", flex, true)
}

// setStatus updates the status bar
func (v *QuickTemplatesView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(" " + message)
	}
}

// parsePort parses a string to int, returns 0 on error
func parsePort(s string) int {
	var port int
	fmt.Sscanf(s, "%d", &port)
	return port
}

// generateRandomString generates a random string of given length
func generateRandomString(length int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[i%len(chars)]
	}
	return string(result)
}

// showApacheHttpdTemplate shows the Apache httpd (mod_jk/AJP) configuration template
func (v *QuickTemplatesView) showApacheHttpdTemplate() {
	form := tview.NewForm()

	ajpPort := "8009"
	form.AddInputField("AJP Port", ajpPort, 10, acceptNumber, func(text string) {
		ajpPort = text
	})

	address := "127.0.0.1"
	form.AddInputField("Bind Address", address, 20, nil, func(text string) {
		address = text
	})

	secret := generateRandomString(16)
	form.AddInputField("AJP Secret", secret, 30, nil, func(text string) {
		secret = text
	})

	addRemoteIpValve := true
	form.AddCheckbox("Add RemoteIpValve", addRemoteIpValve, func(checked bool) {
		addRemoteIpValve = checked
	})

	form.AddButton("[white:green]Apply Template[-:-]", func() {
		cfg := v.configService.GetServer()

		// Create AJP connector
		connector := &server.Connector{
			Port:              parsePort(ajpPort),
			Protocol:          "AJP/1.3",
			Address:           address,
			SecretRequired:    true,
			Secret:            secret,
			RedirectPort:      8443,
			ConnectionTimeout: 20000,
		}

		// Add to first service
		if len(cfg.Services) > 0 {
			// Check if AJP connector already exists on the port
			exists := false
			for _, c := range cfg.Services[0].Connectors {
				if c.Port == connector.Port && strings.Contains(c.Protocol, "AJP") {
					exists = true
					break
				}
			}
			if !exists {
				cfg.Services[0].Connectors = append(cfg.Services[0].Connectors, *connector)
			}

			// Add RemoteIpValve to Engine if requested
			if addRemoteIpValve {
				hasRemoteIpValve := false
				for _, valve := range cfg.Services[0].Engine.Valves {
					if valve.ClassName == server.ValveRemoteIp {
						hasRemoteIpValve = true
						break
					}
				}
				if !hasRemoteIpValve {
					cfg.Services[0].Engine.Valves = append(cfg.Services[0].Engine.Valves, server.Valve{
						ClassName:       server.ValveRemoteIp,
						RemoteIpHeader:  "X-Forwarded-For",
						ProtocolHeader:  "X-Forwarded-Proto",
						InternalProxies: "127\\.0\\.0\\.1|::1",
					})
				}
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		// Show Apache configuration example
		v.showApacheConfigExample(ajpPort, secret)
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Apache httpd (AJP) Template ")

	// Create layout with form and help panel
	flex := CreateFormWithHelp(form, "help.qt.apache", "")

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("apache-template", flex, true)
}

// showApacheConfigExample shows Apache httpd configuration example
func (v *QuickTemplatesView) showApacheConfigExample(ajpPort, secret string) {
	configText := fmt.Sprintf(`[yellow]Apache httpd Configuration Applied![white]

[green]AJP Connector created on port %s[white]

Copy the following configuration to your Apache httpd:

[cyan]Option 1: mod_proxy_ajp (Recommended)[white]
─────────────────────────────────────
<VirtualHost *:80>
    ServerName example.com

    ProxyPreserveHost On
    ProxyPass / ajp://127.0.0.1:%s/ secret=%s
    ProxyPassReverse / ajp://127.0.0.1:%s/

    # Optional: Forward client IP
    RequestHeader set X-Forwarded-Proto "http"
</VirtualHost>

[cyan]Option 2: mod_jk (workers.properties)[white]
─────────────────────────────────────
# workers.properties
worker.list=tomcat1
worker.tomcat1.type=ajp13
worker.tomcat1.host=127.0.0.1
worker.tomcat1.port=%s
worker.tomcat1.secret=%s

# httpd.conf
JkMount /* tomcat1

[yellow]Note:[white] Restart Apache httpd after applying configuration.
`, ajpPort, ajpPort, secret, ajpPort, ajpPort, secret)

	textView := tview.NewTextView().
		SetText(configText).
		SetDynamicColors(true).
		SetScrollable(true)

	textView.SetBorder(true).SetTitle(" Apache httpd Configuration ")

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyEnter {
			v.setStatus("Apache httpd AJP connector configured successfully!")
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("apache-config", textView, true)
	v.setStatus("Press Enter or Escape to continue")
}

// showNginxProxyTemplate shows the nginx reverse proxy configuration template
func (v *QuickTemplatesView) showNginxProxyTemplate() {
	form := tview.NewForm()

	// Get HTTP connectors
	cfg := v.configService.GetServer()
	var httpPorts []string
	var httpPortValues []int
	if len(cfg.Services) > 0 {
		for _, conn := range cfg.Services[0].Connectors {
			if conn.Protocol == "" || strings.Contains(conn.Protocol, "HTTP") {
				httpPorts = append(httpPorts, fmt.Sprintf("Port %d", conn.Port))
				httpPortValues = append(httpPortValues, conn.Port)
			}
		}
	}

	selectedPort := 0
	if len(httpPorts) > 0 {
		form.AddDropDown("HTTP Connector", httpPorts, 0, func(option string, index int) {
			selectedPort = index
		})
	} else {
		form.AddTextView("Warning", "[red]No HTTP connector found![white]", 40, 1, false, false)
	}

	internalProxies := "127\\.0\\.0\\.1|10\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}|192\\.168\\.\\d{1,3}\\.\\d{1,3}"
	form.AddInputField("Internal Proxies (regex)", internalProxies, 60, nil, func(text string) {
		internalProxies = text
	})

	enableProxyProtocol := false
	form.AddCheckbox("Handle X-Forwarded-Proto", true, func(checked bool) {
		enableProxyProtocol = checked
	})

	form.AddButton("[white:green]Apply Template[-:-]", func() {
		cfg := v.configService.GetServer()

		if len(cfg.Services) > 0 {
			// Add RemoteIpValve to Engine
			hasRemoteIpValve := false
			for i, valve := range cfg.Services[0].Engine.Valves {
				if valve.ClassName == server.ValveRemoteIp {
					// Update existing valve
					cfg.Services[0].Engine.Valves[i].InternalProxies = internalProxies
					cfg.Services[0].Engine.Valves[i].RemoteIpHeader = "X-Forwarded-For"
					if enableProxyProtocol {
						cfg.Services[0].Engine.Valves[i].ProtocolHeader = "X-Forwarded-Proto"
					}
					hasRemoteIpValve = true
					break
				}
			}
			if !hasRemoteIpValve {
				valve := server.Valve{
					ClassName:       server.ValveRemoteIp,
					RemoteIpHeader:  "X-Forwarded-For",
					InternalProxies: internalProxies,
				}
				if enableProxyProtocol {
					valve.ProtocolHeader = "X-Forwarded-Proto"
				}
				cfg.Services[0].Engine.Valves = append(cfg.Services[0].Engine.Valves, valve)
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		// Get the selected port for config example
		port := 8080
		if len(httpPortValues) > selectedPort {
			port = httpPortValues[selectedPort]
		}

		// Show nginx configuration example
		v.showNginxConfigExample(port)
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" nginx Reverse Proxy Template ")

	// Create layout with form and help panel
	flex := CreateFormWithHelp(form, "help.qt.nginx", "")

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("nginx-template", flex, true)
}

// showNginxConfigExample shows nginx configuration example
func (v *QuickTemplatesView) showNginxConfigExample(tomcatPort int) {
	configText := fmt.Sprintf(`[yellow]nginx Reverse Proxy Configuration Applied![white]

[green]RemoteIpValve configured for nginx proxy[white]

Copy the following configuration to your nginx:

[cyan]Basic Configuration[white]
─────────────────────────────────────
upstream tomcat {
    server 127.0.0.1:%d;
    keepalive 32;
}

server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://tomcat;
        proxy_http_version 1.1;

        # Forward headers
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host $host;
        proxy_set_header X-Forwarded-Port $server_port;

        # Connection settings
        proxy_set_header Connection "";
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }
}

[cyan]HTTPS Configuration (with SSL termination)[white]
─────────────────────────────────────
server {
    listen 443 ssl http2;
    server_name example.com;

    ssl_certificate /etc/nginx/ssl/cert.pem;
    ssl_certificate_key /etc/nginx/ssl/key.pem;

    location / {
        proxy_pass http://tomcat;
        proxy_http_version 1.1;

        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto https;
        proxy_set_header X-Forwarded-Host $host;
        proxy_set_header X-Forwarded-Port 443;

        proxy_set_header Connection "";
    }
}

[yellow]Note:[white] Restart nginx after applying configuration.
`, tomcatPort)

	textView := tview.NewTextView().
		SetText(configText).
		SetDynamicColors(true).
		SetScrollable(true)

	textView.SetBorder(true).SetTitle(" nginx Configuration ")

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyEnter {
			v.setStatus("nginx reverse proxy configured successfully!")
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("nginx-config", textView, true)
	v.setStatus("Press Enter or Escape to continue")
}

// showHAProxyTemplate shows the HAProxy load balancer configuration template
func (v *QuickTemplatesView) showHAProxyTemplate() {
	form := tview.NewForm()

	// Get HTTP connectors
	cfg := v.configService.GetServer()
	var httpPorts []string
	var httpPortValues []int
	if len(cfg.Services) > 0 {
		for _, conn := range cfg.Services[0].Connectors {
			if conn.Protocol == "" || strings.Contains(conn.Protocol, "HTTP") {
				httpPorts = append(httpPorts, fmt.Sprintf("Port %d", conn.Port))
				httpPortValues = append(httpPortValues, conn.Port)
			}
		}
	}

	selectedPort := 0
	if len(httpPorts) > 0 {
		form.AddDropDown("HTTP Connector", httpPorts, 0, func(option string, index int) {
			selectedPort = index
		})
	}

	// Sticky session configuration
	enableStickySession := true
	form.AddCheckbox("Enable Sticky Sessions", enableStickySession, func(checked bool) {
		enableStickySession = checked
	})

	jvmRoute := "tomcat1"
	form.AddInputField("JVM Route (node ID)", jvmRoute, 20, nil, func(text string) {
		jvmRoute = text
	})

	internalProxies := "127\\.0\\.0\\.1|10\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}|192\\.168\\.\\d{1,3}\\.\\d{1,3}"
	form.AddInputField("Internal Proxies (regex)", internalProxies, 60, nil, func(text string) {
		internalProxies = text
	})

	form.AddButton("[white:green]Apply Template[-:-]", func() {
		cfg := v.configService.GetServer()

		if len(cfg.Services) > 0 {
			// Set jvmRoute for sticky sessions
			if enableStickySession && jvmRoute != "" {
				cfg.Services[0].Engine.JvmRoute = jvmRoute
			}

			// Add RemoteIpValve to Engine
			hasRemoteIpValve := false
			for i, valve := range cfg.Services[0].Engine.Valves {
				if valve.ClassName == server.ValveRemoteIp {
					// Update existing valve
					cfg.Services[0].Engine.Valves[i].InternalProxies = internalProxies
					cfg.Services[0].Engine.Valves[i].RemoteIpHeader = "X-Forwarded-For"
					cfg.Services[0].Engine.Valves[i].ProtocolHeader = "X-Forwarded-Proto"
					hasRemoteIpValve = true
					break
				}
			}
			if !hasRemoteIpValve {
				cfg.Services[0].Engine.Valves = append(cfg.Services[0].Engine.Valves, server.Valve{
					ClassName:       server.ValveRemoteIp,
					RemoteIpHeader:  "X-Forwarded-For",
					ProtocolHeader:  "X-Forwarded-Proto",
					InternalProxies: internalProxies,
				})
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus("Error saving: " + err.Error())
			return
		}

		// Get the selected port for config example
		port := 8080
		if len(httpPortValues) > selectedPort {
			port = httpPortValues[selectedPort]
		}

		// Show HAProxy configuration example
		v.showHAProxyConfigExample(port, jvmRoute, enableStickySession)
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" HAProxy Load Balancer Template ")

	// Create layout with form and help panel
	flex := CreateFormWithHelp(form, "help.qt.haproxy", "")

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("haproxy-template", flex, true)
}

// showHAProxyConfigExample shows HAProxy configuration example
func (v *QuickTemplatesView) showHAProxyConfigExample(tomcatPort int, jvmRoute string, stickySession bool) {
	stickyConfig := ""
	cookieConfig := ""
	serverSuffix := ""

	if stickySession {
		cookieConfig = "\n    cookie JSESSIONID prefix nocache"
		serverSuffix = fmt.Sprintf(" cookie %s", jvmRoute)
	}

	if stickySession {
		stickyConfig = fmt.Sprintf(`
[cyan]Sticky Session Notes:[white]
  - jvmRoute "%s" configured in Tomcat Engine
  - HAProxy uses JSESSIONID cookie for session affinity
  - Session ID format: <session-id>.<jvmRoute>
`, jvmRoute)
	}

	configText := fmt.Sprintf(`[yellow]HAProxy Load Balancer Configuration Applied![white]

[green]RemoteIpValve configured for HAProxy[white]%s

Copy the following configuration to your HAProxy:

[cyan]HTTP Mode (Layer 7) - Recommended[white]
─────────────────────────────────────
global
    log /dev/log local0
    maxconn 4096
    daemon

defaults
    log     global
    mode    http
    option  httplog
    option  dontlognull
    option  forwardfor
    option  http-server-close
    timeout connect 5s
    timeout client  30s
    timeout server  30s

frontend http_front
    bind *:80
    default_backend tomcat_backend

    # Optional: HTTPS redirect
    # redirect scheme https if !{ ssl_fc }

frontend https_front
    bind *:443 ssl crt /etc/haproxy/certs/cert.pem
    default_backend tomcat_backend

backend tomcat_backend
    balance roundrobin
    option httpchk GET /
    http-check expect status 200%s

    # Add more servers for load balancing
    server tomcat1 127.0.0.1:%d check%s
    # server tomcat2 192.168.1.11:%d check cookie tomcat2
    # server tomcat3 192.168.1.12:%d check cookie tomcat3

[cyan]TCP Mode (Layer 4) - SSL Passthrough[white]
─────────────────────────────────────
frontend tcp_front
    bind *:443
    mode tcp
    default_backend tomcat_tcp_backend

backend tomcat_tcp_backend
    mode tcp
    balance roundrobin
    server tomcat1 127.0.0.1:8443 check

[cyan]Health Check Endpoint (Optional)[white]
─────────────────────────────────────
# Add to your Tomcat application
# Create: /health or /status endpoint returning HTTP 200

# Alternative: Use Tomcat Manager status
# option httpchk GET /manager/status
# http-check expect status 200

[cyan]HAProxy Stats (Optional)[white]
─────────────────────────────────────
listen stats
    bind *:8404
    stats enable
    stats uri /stats
    stats refresh 10s
    stats admin if LOCALHOST

[yellow]Note:[white] Restart HAProxy after applying configuration.
`, stickyConfig, cookieConfig, tomcatPort, serverSuffix, tomcatPort, tomcatPort)

	textView := tview.NewTextView().
		SetText(configText).
		SetDynamicColors(true).
		SetScrollable(true)

	textView.SetBorder(true).SetTitle(" HAProxy Configuration ")

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyEnter {
			v.setStatus("HAProxy load balancer configured successfully!")
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("haproxy-config", textView, true)
	v.setStatus("Press Enter or Escape to continue")
}
