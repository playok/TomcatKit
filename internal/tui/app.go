package tui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/detector"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/playok/tomcatkit/internal/tui/views"
	"github.com/rivo/tview"
)

// AppOptions contains initialization options for the app
type AppOptions struct {
	CatalinaHome    string
	CatalinaBase    string
	SettingsManager *config.SettingsManager
}

// App represents the main TUI application
type App struct {
	app             *tview.Application
	pages           *tview.Pages
	mainMenu        *tview.List
	statusBar       *tview.TextView
	infoPanel       *tview.TextView
	instance        *config.TomcatInstance
	settingsManager *config.SettingsManager
}

// NewApp creates a new TUI application instance with default options
func NewApp() *App {
	return NewAppWithOptions(&AppOptions{})
}

// NewAppWithOptions creates a new TUI application instance with options
func NewAppWithOptions(opts *AppOptions) *App {
	a := &App{
		app:             tview.NewApplication(),
		pages:           tview.NewPages(),
		settingsManager: opts.SettingsManager,
	}

	// Load language setting from settings
	if a.settingsManager != nil {
		if savedLang := a.settingsManager.GetLanguage(); savedLang != "" {
			i18n.SetLanguage(i18n.Language(savedLang))
		}
	}

	// Set instance from options if provided
	if opts.CatalinaHome != "" {
		a.instance = &config.TomcatInstance{
			CatalinaHome: opts.CatalinaHome,
			CatalinaBase: opts.CatalinaBase,
		}
		// Detect version
		d := detector.NewDetector()
		if instances, _ := d.DetectAll(); instances != nil {
			for _, inst := range instances {
				if inst.CatalinaHome == opts.CatalinaHome {
					a.instance = inst
					break
				}
			}
		}
	}

	a.setupUI()

	// If no instance selected, try to load from settings or show selector
	if a.instance == nil && a.settingsManager != nil {
		if lastInstance := a.settingsManager.GetLastInstance(); lastInstance != nil {
			// Verify the path still exists
			serverXml := filepath.Join(lastInstance.CatalinaBase, "conf", "server.xml")
			if _, err := os.Stat(serverXml); err == nil {
				a.instance = lastInstance
				a.updateInstanceInfo()
			}
		}
	}

	return a
}

// setupUI initializes the user interface components
func (a *App) setupUI() {
	// Create header
	header := tview.NewTextView().
		SetText("[::b]TomcatKit[-:-:-] - Apache Tomcat 9.0 Configuration Helper").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	header.SetBorder(true).SetBorderColor(tcell.ColorDarkCyan)

	// Create info panel
	a.infoPanel = tview.NewTextView().
		SetDynamicColors(true)
	noInstanceText := fmt.Sprintf("[yellow::b]%s[-::-]\n\n[white]%s[-]\n\n  [green]1.[-] %s\n  [green]2.[-] %s\n\n[gray]%s[-]",
		i18n.T("instance.info.noselected"),
		i18n.T("instance.info.getstarted"),
		i18n.T("instance.info.step1"),
		i18n.T("instance.info.step2"),
		i18n.T("instance.info.autodetect"))
	a.infoPanel.SetText(noInstanceText)
	a.infoPanel.SetBorder(true).SetTitle(" " + i18n.T("instance.info") + " ").SetBorderColor(tcell.ColorYellow)

	// Create main menu with all configuration categories
	a.mainMenu = tview.NewList().
		ShowSecondaryText(true)

	// Select Tomcat Instance (moved to top for visibility)
	a.mainMenu.AddItem("[green::b]► "+i18n.T("instance.title")+"[-::-]", i18n.T("instance.manual.desc"), 't', func() {
		a.showInstanceSelector()
	})

	// Separator
	a.mainMenu.AddItem("─────────────────────────", "", 0, nil)

	// Server Configuration
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.server")+"[::-]", i18n.T("menu.server.desc"), 's', func() {
		a.showServerMenu()
	})

	// Connector Configuration
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.connector")+"[::-]", i18n.T("menu.connector.desc"), 'c', func() {
		a.showConnectorMenu()
	})

	// Security & Authentication
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.security")+"[::-]", i18n.T("menu.security.desc"), 'a', func() {
		a.showSecurityMenu()
	})

	// JNDI Resources
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.jndi")+"[::-]", i18n.T("menu.jndi.desc"), 'j', func() {
		a.showJNDIMenu()
	})

	// Virtual Hosts & Contexts
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.host")+"[::-]", i18n.T("menu.host.desc"), 'h', func() {
		a.showHostContextMenu()
	})

	// Valves
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.valve")+"[::-]", i18n.T("menu.valve.desc"), 'v', func() {
		a.showValveMenu()
	})

	// Clustering
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.cluster")+"[::-]", i18n.T("menu.cluster.desc"), 'l', func() {
		a.showClusterMenu()
	})

	// Logging
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.logging")+"[::-]", i18n.T("menu.logging.desc"), 'g', func() {
		a.showLoggingMenu()
	})

	// Context (context.xml)
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.context")+"[::-]", i18n.T("menu.context.desc"), 'x', func() {
		a.showContextMenu()
	})

	// Web (web.xml)
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.web")+"[::-]", i18n.T("menu.web.desc"), 'w', func() {
		a.showWebMenu()
	})

	// Separator
	a.mainMenu.AddItem("─────────────────────────", "", 0, nil)

	// Quick Templates
	a.mainMenu.AddItem("[green::b]"+i18n.T("menu.quicktemplates")+"[::-]", i18n.T("menu.quicktemplates.desc"), 'Q', func() {
		a.showQuickTemplatesMenu()
	})

	// Separator
	a.mainMenu.AddItem("─────────────────────────", "", 0, nil)

	// Quit
	a.mainMenu.AddItem("[red]"+i18n.T("menu.exit")+"[-]", i18n.T("menu.exit.desc"), 'q', func() {
		a.app.Stop()
	})

	a.mainMenu.SetBorder(true).SetTitle(" " + i18n.T("menu.title") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Create status bar
	a.statusBar = tview.NewTextView().
		SetText(" [yellow]↑↓[-] Navigate | [yellow]Enter[-] Select | [yellow]Esc[-] Back | [yellow]F2[-] Lang | [yellow]q[-] Quit").
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)

	// Create main layout with left menu and right info panel
	contentArea := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(a.mainMenu, 0, 2, true).
		AddItem(a.infoPanel, 0, 1, false)

	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 3, 0, false).
		AddItem(contentArea, 0, 1, true).
		AddItem(a.statusBar, 1, 0, false)

	a.pages.AddPage("main", mainLayout, true, true)

	// Setup global key handlers (F2 for language selection only)
	// ESC key is handled by individual views for hierarchical navigation
	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// F2 key for language selection
		if event.Key() == tcell.KeyF2 {
			a.showLanguageSelector()
			return nil
		}
		return event
	})

	a.app.SetRoot(a.pages, true)
}

// updateInstanceInfo updates the info panel with current instance details
func (a *App) updateInstanceInfo() {
	if a.instance == nil {
		noInstanceText := fmt.Sprintf("[yellow::b]%s[-::-]\n\n[white]%s[-]\n\n  [green]1.[-] %s\n  [green]2.[-] %s\n\n[gray]%s[-]",
			i18n.T("instance.info.noselected"),
			i18n.T("instance.info.getstarted"),
			i18n.T("instance.info.step1"),
			i18n.T("instance.info.step2"),
			i18n.T("instance.info.autodetect"))
		a.infoPanel.SetText(noInstanceText)
		a.infoPanel.SetBorder(true).SetTitle(" " + i18n.T("instance.info") + " ").SetBorderColor(tcell.ColorYellow)
		return
	}

	status := "[red]" + i18n.T("instance.stopped") + "[-]"
	if a.instance.IsRunning {
		status = fmt.Sprintf("[green]%s[-] (PID: %d)", i18n.T("instance.running"), a.instance.PID)
	}

	info := fmt.Sprintf("[::b]%s[::-]\n\n[yellow]%s:[-]       %s\n[yellow]CATALINA_HOME:[-] %s\n[yellow]CATALINA_BASE:[-] %s\n[yellow]%s:[-]        %s\n\n[green]%s[-]",
		i18n.T("instance.info"),
		i18n.T("instance.version"),
		a.instance.Version,
		a.instance.CatalinaHome,
		a.instance.CatalinaBase,
		i18n.T("instance.status"),
		status,
		i18n.T("instance.ready"))

	a.infoPanel.SetText(info)
	a.infoPanel.SetBorder(true).SetTitle(" " + i18n.T("instance.info") + " ").SetBorderColor(tcell.ColorYellow)

	// Save to settings
	if a.settingsManager != nil {
		a.settingsManager.SetLastInstance(a.instance)
		a.settingsManager.Save()
	}
}

// rebuildMainMenu rebuilds the main menu with the current language
func (a *App) rebuildMainMenu() {
	a.mainMenu.Clear()

	// Select Tomcat Instance (moved to top for visibility)
	a.mainMenu.AddItem("[green::b]► "+i18n.T("instance.title")+"[-::-]", i18n.T("instance.manual.desc"), 't', func() {
		a.showInstanceSelector()
	})

	// Separator
	a.mainMenu.AddItem("─────────────────────────", "", 0, nil)

	// Server Configuration
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.server")+"[::-]", i18n.T("menu.server.desc"), 's', func() {
		a.showServerMenu()
	})

	// Connector Configuration
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.connector")+"[::-]", i18n.T("menu.connector.desc"), 'c', func() {
		a.showConnectorMenu()
	})

	// Security & Authentication
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.security")+"[::-]", i18n.T("menu.security.desc"), 'a', func() {
		a.showSecurityMenu()
	})

	// JNDI Resources
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.jndi")+"[::-]", i18n.T("menu.jndi.desc"), 'j', func() {
		a.showJNDIMenu()
	})

	// Virtual Hosts & Contexts
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.host")+"[::-]", i18n.T("menu.host.desc"), 'h', func() {
		a.showHostContextMenu()
	})

	// Valves
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.valve")+"[::-]", i18n.T("menu.valve.desc"), 'v', func() {
		a.showValveMenu()
	})

	// Clustering
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.cluster")+"[::-]", i18n.T("menu.cluster.desc"), 'l', func() {
		a.showClusterMenu()
	})

	// Logging
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.logging")+"[::-]", i18n.T("menu.logging.desc"), 'g', func() {
		a.showLoggingMenu()
	})

	// Context (context.xml)
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.context")+"[::-]", i18n.T("menu.context.desc"), 'x', func() {
		a.showContextMenu()
	})

	// Web (web.xml)
	a.mainMenu.AddItem("[::b]"+i18n.T("menu.web")+"[::-]", i18n.T("menu.web.desc"), 'w', func() {
		a.showWebMenu()
	})

	// Separator
	a.mainMenu.AddItem("─────────────────────────", "", 0, nil)

	// Quick Templates
	a.mainMenu.AddItem("[green::b]"+i18n.T("menu.quicktemplates")+"[::-]", i18n.T("menu.quicktemplates.desc"), 'Q', func() {
		a.showQuickTemplatesMenu()
	})

	// Separator
	a.mainMenu.AddItem("─────────────────────────", "", 0, nil)

	// Quit
	a.mainMenu.AddItem("[red]"+i18n.T("menu.exit")+"[-]", i18n.T("menu.exit.desc"), 'q', func() {
		a.app.Stop()
	})

	a.mainMenu.SetBorder(true).SetTitle(" " + i18n.T("menu.title") + " ").SetBorderColor(tcell.ColorDarkCyan)
}

// showInstanceSelector shows the Tomcat instance selection dialog
func (a *App) showInstanceSelector() {
	d := detector.NewDetector()
	instances, _ := d.DetectAll()

	list := tview.NewList().ShowSecondaryText(true)

	// Add recent instances from settings first
	if a.settingsManager != nil {
		recentInstances := a.settingsManager.GetRecentInstances()
		if len(recentInstances) > 0 {
			list.AddItem("[::b]"+i18n.T("instance.recent")+"[::-]", "─────────────────────", 0, nil)
			for _, recent := range recentInstances {
				inst := recent // capture for closure
				// Check if path still exists
				serverXml := filepath.Join(inst.CatalinaBase, "conf", "server.xml")
				if _, err := os.Stat(serverXml); err == nil {
					list.AddItem(
						fmt.Sprintf("[yellow]%s[-]", inst.CatalinaHome),
						fmt.Sprintf("%s: %s", i18n.T("instance.version"), inst.Version),
						0,
						func() {
							a.instance = &inst
							a.updateInstanceInfo()
							a.pages.SwitchToPage("main")
							a.app.SetFocus(a.mainMenu)
						},
					)
				}
			}
			list.AddItem("", "", 0, nil) // Spacer
		}
	}

	// Add detected instances
	if len(instances) > 0 {
		list.AddItem("[::b]"+i18n.T("instance.detected")+"[::-]", "─────────────────────", 0, nil)
		for _, inst := range instances {
			instance := inst // capture for closure
			status := ""
			if instance.IsRunning {
				status = " [green](" + i18n.T("instance.running") + ")[-]"
			}
			list.AddItem(
				fmt.Sprintf("Tomcat %s%s", instance.Version, status),
				instance.CatalinaHome,
				0,
				func() {
					a.instance = instance
					a.updateInstanceInfo()
					a.pages.SwitchToPage("main")
					a.app.SetFocus(a.mainMenu)
				},
			)
		}
	} else {
		list.AddItem("[gray]"+i18n.T("instance.none")+"[-]", i18n.T("instance.manual.desc"), 0, nil)
	}

	list.AddItem("", "", 0, nil) // Spacer

	list.AddItem("[green]► "+i18n.T("instance.manual")+"[-]", i18n.T("instance.manual.desc"), 'm', func() {
		a.showManualPathInput()
	})

	list.AddItem("[red]"+i18n.T("common.cancel")+"[-]", i18n.T("common.return"), 0, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("instance.title") + " ").SetBorderColor(tcell.ColorGreen)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			a.pages.SwitchToPage("main")
			a.app.SetFocus(a.mainMenu)
			return nil
		}
		return event
	})

	a.pages.AddAndSwitchToPage("instance-selector", list, true)
	a.app.SetFocus(list)
}

// showManualPathInput shows a form for manual path entry
func (a *App) showManualPathInput() {
	form := tview.NewForm()

	// Get default values
	defaultHome := os.Getenv("CATALINA_HOME")
	defaultBase := os.Getenv("CATALINA_BASE")

	form.AddInputField(i18n.T("instance.path.home"), defaultHome, 50, nil, nil)
	form.AddInputField(i18n.T("instance.path.base"), defaultBase, 50, nil, nil)

	form.AddButton(i18n.T("instance.path.validate"), func() {
		catalinaHome := form.GetFormItem(0).(*tview.InputField).GetText()
		catalinaBase := form.GetFormItem(1).(*tview.InputField).GetText()

		if catalinaHome == "" {
			a.setStatus("[red]" + i18n.T("instance.path.required") + "[-]")
			return
		}

		if catalinaBase == "" {
			catalinaBase = catalinaHome
		}

		// Validate path
		serverXml := filepath.Join(catalinaBase, "conf", "server.xml")
		if _, err := os.Stat(serverXml); os.IsNotExist(err) {
			a.setStatus("[red]" + i18n.T("instance.path.invalid") + "[-]")
			return
		}

		// Detect version
		d := detector.NewDetector()
		version := "unknown"
		if instances, _ := d.DetectAll(); instances != nil {
			for _, inst := range instances {
				if inst.CatalinaHome == catalinaHome {
					version = inst.Version
					break
				}
			}
		}

		a.instance = &config.TomcatInstance{
			CatalinaHome: catalinaHome,
			CatalinaBase: catalinaBase,
			Version:      version,
		}
		a.updateInstanceInfo()
		a.setStatus("[green]" + i18n.T("instance.selected") + "[-]")
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		a.showInstanceSelector()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("instance.path.title") + " ").SetBorderColor(tcell.ColorGreen)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			a.showInstanceSelector()
			return nil
		}
		return event
	})

	// Add help text
	helpText := tview.NewTextView().
		SetDynamicColors(true).
		SetText(fmt.Sprintf("[yellow]%s[-]\n[yellow]%s[-]\n\n%s",
			i18n.T("instance.path.help.home"),
			i18n.T("instance.path.help.base"),
			i18n.T("instance.path.help.xml")))

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(helpText, 5, 0, false)

	layout.SetBorder(true).SetTitle(" " + i18n.T("instance.path.title") + " ")

	a.pages.AddAndSwitchToPage("manual-path", layout, true)
	a.app.SetFocus(form)
}

// setStatus updates the status bar
func (a *App) setStatus(message string) {
	a.statusBar.SetText(fmt.Sprintf(" %s", message))
}

// showServerMenu shows the server configuration menu
func (a *App) showServerMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and load server configuration service
	configService := server.NewConfigService(a.instance.CatalinaBase)
	if err := configService.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load server.xml:\n%v", err))
		return
	}

	// Create and show server view
	serverView := views.NewServerView(a.app, a.pages, configService, a.statusBar, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	serverView.Show()
}

func (a *App) showConnectorMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and load server configuration service
	configService := server.NewConfigService(a.instance.CatalinaBase)
	if err := configService.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load server.xml:\n%v", err))
		return
	}

	// Create and show connector view
	connectorView := views.NewConnectorView(a.app, a.pages, configService, a.statusBar, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	connectorView.Show()
}

func (a *App) showSecurityMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and load server configuration service
	configService := server.NewConfigService(a.instance.CatalinaBase)
	if err := configService.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load server.xml:\n%v", err))
		return
	}

	// Create and show security view
	securityView := views.NewSecurityView(a.app, a.pages, configService, a.instance.CatalinaBase, a.statusBar, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	securityView.Show()
}

func (a *App) showJNDIMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and show JNDI view
	jndiView := views.NewJNDIView(a.app, a.pages, a.instance.CatalinaBase, a.statusBar, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	jndiView.Show()
}

func (a *App) showHostContextMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and load server configuration service
	configService := server.NewConfigService(a.instance.CatalinaBase)
	if err := configService.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load server.xml:\n%v", err))
		return
	}

	// Create and show host view
	hostView := views.NewHostView(a.app, a.pages, configService, a.statusBar, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	hostView.Show()
}

func (a *App) showValveMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and load server configuration service
	configService := server.NewConfigService(a.instance.CatalinaBase)
	if err := configService.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load server.xml:\n%v", err))
		return
	}

	// Create and show valve view
	valveView := views.NewValveView(a.app, a.pages, configService, a.statusBar, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	valveView.Show()
}

func (a *App) showClusterMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and load server configuration service
	configService := server.NewConfigService(a.instance.CatalinaBase)
	if err := configService.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load server.xml:\n%v", err))
		return
	}

	// Create and show cluster view
	clusterView := views.NewClusterView(a.app, a.pages, configService, a.statusBar, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	clusterView.Show()
}

func (a *App) showLoggingMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and show logging view
	loggingView := views.NewLoggingView(a.app, a.pages, a.statusBar, a.instance.CatalinaBase, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	if err := loggingView.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load logging configuration:\n%v", err))
		return
	}
}

func (a *App) showContextMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and show context view
	contextView := views.NewContextView(a.app, a.pages, a.statusBar, a.instance.CatalinaBase, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	if err := contextView.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load context.xml:\n%v", err))
		return
	}
}

func (a *App) showWebMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and show web view
	webView := views.NewWebView(a.app, a.pages, a.statusBar, a.instance.CatalinaBase, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	if err := webView.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load web.xml:\n%v", err))
		return
	}
}

func (a *App) showQuickTemplatesMenu() {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	// Create and show quick templates view
	quickTemplatesView := views.NewQuickTemplatesView(a.app, a.pages, a.statusBar, a.instance.CatalinaBase, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})
	if err := quickTemplatesView.Load(); err != nil {
		a.showMessage("Error", fmt.Sprintf("Failed to load configuration:\n%v", err))
		return
	}
}

type menuItem struct {
	title       string
	description string
	shortcut    rune
}

func (a *App) showSubMenu(title string, items []menuItem) {
	if a.instance == nil {
		a.showMessage("Error", "Please select a Tomcat instance first.\n\nPress 't' from the main menu to detect and select an instance.")
		return
	}

	list := tview.NewList().ShowSecondaryText(true)

	for _, item := range items {
		list.AddItem(item.title, item.description, item.shortcut, nil)
	}

	list.AddItem("[red]Back[-]", "Return to main menu", 0, func() {
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})

	list.SetBorder(true).SetTitle(fmt.Sprintf(" %s ", title))

	pageName := fmt.Sprintf("submenu-%s", title)
	a.pages.AddAndSwitchToPage(pageName, list, true)
	a.app.SetFocus(list)
}

func (a *App) showMessage(title, message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			a.pages.SwitchToPage("main")
			a.app.SetFocus(a.mainMenu)
		})
	modal.SetBorder(true).SetTitle(fmt.Sprintf(" %s ", title))
	a.pages.AddAndSwitchToPage("message", modal, true)
}

// showLanguageSelector shows the language selection dialog
func (a *App) showLanguageSelector() {
	currentLang := i18n.GetLanguage()

	list := tview.NewList().ShowSecondaryText(false)

	for _, lang := range i18n.AvailableLanguages() {
		l := lang // capture for closure
		label := i18n.GetLanguageName(l)
		if l == currentLang {
			label = fmt.Sprintf("[green]► %s (current)[-]", label)
		}
		list.AddItem(label, "", 0, func() {
			i18n.SetLanguage(l)
			// Save language setting to settings.json
			if a.settingsManager != nil {
				a.settingsManager.SetLanguage(string(l))
				a.settingsManager.Save()
			}
			a.setStatus(fmt.Sprintf("Language changed to %s", i18n.GetLanguageName(l)))
			a.pages.RemovePage("lang-selector")
			// Rebuild UI with new language
			a.rebuildMainMenu()
			a.updateInstanceInfo()
			a.pages.SwitchToPage("main")
			a.app.SetFocus(a.mainMenu)
		})
	}

	list.AddItem("", "", 0, nil) // Spacer
	list.AddItem("[red]Cancel[-]", "", 0, func() {
		a.pages.RemovePage("lang-selector")
		a.pages.SwitchToPage("main")
		a.app.SetFocus(a.mainMenu)
	})

	list.SetBorder(true).
		SetTitle(fmt.Sprintf(" %s ", i18n.T("app.lang.select"))).
		SetBorderColor(tcell.ColorYellow)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyF2 {
			a.pages.RemovePage("lang-selector")
			a.pages.SwitchToPage("main")
			a.app.SetFocus(a.mainMenu)
			return nil
		}
		return event
	})

	// Create a centered modal-like layout
	flex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(list, 10, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)

	a.pages.AddAndSwitchToPage("lang-selector", flex, true)
	a.app.SetFocus(list)
}

// Run starts the TUI application
func (a *App) Run() error {
	// Update instance info if already selected
	if a.instance != nil {
		a.updateInstanceInfo()
	}

	// Ensure main menu has focus
	a.app.SetFocus(a.mainMenu)

	return a.app.Run()
}
