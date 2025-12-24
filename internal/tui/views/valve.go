package views

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// Help key arrays for each valve type form (indexed by form item order)
var accessLogValveHelpKeys = []string{
	"help.valve.classname",                // 0: Type (TextView)
	"help.valve.accesslog.directory",      // 1: Directory
	"help.valve.accesslog.prefix",         // 2: Prefix
	"help.valve.accesslog.suffix",         // 3: Suffix
	"help.valve.accesslog.pattern",        // 4: Pattern
	"help.valve.accesslog.filedateformat", // 5: File Date Format
	"help.valve.accesslog.rotate",         // 6: Rotatable
	"help.valve.accesslog.renameonrotate", // 7: Rename On Rotate
	"help.valve.accesslog.encoding",       // 8: Encoding
	"help.valve.accesslog.buffered",       // 9: Buffered
	"help.valve.accesslog.conditionif",    // 10: Condition If
	"help.valve.accesslog.conditionunless", // 11: Condition Unless
}

var remoteAddrValveHelpKeys = []string{
	"help.valve.classname",              // 0: Type (TextView)
	"help.valve.remoteaddr.allow",       // 1: Allow (regex)
	"help.valve.remoteaddr.deny",        // 2: Deny (regex)
	"help.valve.remoteaddr.denystatus",  // 3: Deny Status
	"help.valve.remoteaddr.addport",     // 4: Add Connector Port
	"help.valve.remoteaddr.invalidauth", // 5: Invalid Auth When Deny
}

var remoteIpValveHelpKeys = []string{
	"help.valve.classname",                // 0: Type (TextView)
	"help.valve.remoteip.header",          // 1: Remote IP Header
	"help.valve.remoteip.protocolheader",  // 2: Protocol Header
	"help.valve.remoteip.protocolhttps",   // 3: Protocol Header HTTPS Value
	"help.valve.remoteip.portheader",      // 4: Port Header
	"help.valve.remoteip.internalproxies", // 5: Internal Proxies
	"help.valve.remoteip.trustedproxies",  // 6: Trusted Proxies
	"help.valve.remoteip.changelocalport", // 7: Change Local Port
	"help.valve.remoteip.changelocalname", // 8: Change Local Name
}

var errorReportValveHelpKeys = []string{
	"help.valve.classname",              // 0: Type (TextView)
	"help.valve.errorreport.serverinfo", // 1: Show Server Info
	"help.valve.errorreport.report",     // 2: Show Report
}

var singleSignOnValveHelpKeys = []string{
	"help.valve.classname",            // 0: Type (TextView)
	"help.valve.sso.cookiedomain",     // 1: Cookie Domain
	"help.valve.sso.cookiename",       // 2: Cookie Name
	"help.valve.sso.requirereauth",    // 3: Require Reauthentication
}

var stuckThreadValveHelpKeys = []string{
	"help.valve.classname",                 // 0: Type (TextView)
	"help.valve.stuckthread.threshold",     // 1: Threshold (seconds)
	"help.valve.stuckthread.interruptthreshold", // 2: Interrupt Thread Threshold
}

var crawlerValveHelpKeys = []string{
	"help.valve.classname",                    // 0: Type (TextView)
	"help.valve.crawler.useragents",           // 1: Crawler User Agents
	"help.valve.crawler.sessioninactiveinterval", // 2: Session Inactive Interval
}

var semaphoreValveHelpKeys = []string{
	"help.valve.classname",         // 0: Type (TextView)
	"help.valve.semaphore.concurrency", // 1: Concurrency
	"help.valve.semaphore.fairness",    // 2: Fairness
	"help.valve.semaphore.block",       // 3: Block
}

var replicationValveHelpKeys = []string{
	"help.valve.classname",                 // 0: Type (TextView)
	"help.valve.replication.filter",        // 1: Filter
	"help.valve.replication.primaryindicator", // 2: Primary Indicator
	"help.valve.replication.primaryindicatorname", // 3: Primary Indicator Name
}

// ValveView handles valve configuration
type ValveView struct {
	app           *tview.Application
	pages         *tview.Pages
	mainPages     *tview.Pages
	statusBar     *tview.TextView
	onReturn      func()
	configService *server.ConfigService
}

// NewValveView creates a new valve view
func NewValveView(app *tview.Application, mainPages *tview.Pages, configService *server.ConfigService, statusBar *tview.TextView, onReturn func()) *ValveView {
	return &ValveView{
		app:           app,
		pages:         tview.NewPages(),
		mainPages:     mainPages,
		statusBar:     statusBar,
		onReturn:      onReturn,
		configService: configService,
	}
}

// Show displays the valve view
func (v *ValveView) Show() {
	v.showMainMenu()
	v.mainPages.AddAndSwitchToPage("valve", v.pages, true)
}

// showMainMenu shows the main valve menu (select scope)
func (v *ValveView) showMainMenu() {
	menu := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	cfg := v.configService.GetConfig()
	engineValveCount := 0
	if cfg != nil && len(cfg.Services) > 0 {
		engineValveCount = len(cfg.Services[0].Engine.Valves)
	}

	menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("valve.engine")+"[::-] [yellow](%d)[-]", engineValveCount),
		i18n.T("valve.engine.desc"), 'e', func() {
			v.showEngineValves()
		})

	menu.AddItem("[::b]"+i18n.T("valve.host")+"[::-]",
		i18n.T("valve.host.desc"), 'h', func() {
			v.showHostSelector()
		})

	menu.AddItem("[::b]"+i18n.T("valve.context")+"[::-]",
		i18n.T("valve.context.desc"), 'c', func() {
			v.showContextHostSelector()
		})

	menu.AddItem("", "", 0, nil)
	menu.AddItem("[yellow]► "+i18n.T("valve.quickadd")+"[-]",
		i18n.T("valve.quickadd.desc"), 'q', func() {
			v.showQuickAddMenu()
		})

	menu.AddItem("", "", 0, nil)
	menu.AddItem("[-:-:-] [white:red] "+i18n.T("common.back")+" [-:-:-]", i18n.T("common.return"), 0, func() {
		v.onReturn()
	})

	// Update help panel when selection changes
	menu.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.valve.engine")
		case 1:
			helpPanel.SetHelpKey("help.valve.host")
		case 2:
			helpPanel.SetHelpKey("help.valve.context")
		case 4:
			helpPanel.SetHelpKey("help.valve.quickadd")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.valve.engine")

	menu.SetBorder(true).SetTitle(" " + i18n.T("valve.title") + " ").SetBorderColor(tcell.ColorDarkCyan)
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

// showQuickAddMenu shows quick add options for common valves
func (v *ValveView) showQuickAddMenu() {
	menu := tview.NewList().ShowSecondaryText(true)

	menu.AddItem("Access Log Valve",
		"Log HTTP requests in customizable format", 'a', func() {
			valve := server.DefaultAccessLogValve()
			v.showValveForm(&valve, true, "engine", nil, nil)
		})

	menu.AddItem("Remote Address Valve",
		"IP-based access control (allow/deny)", 'r', func() {
			valve := server.DefaultRemoteAddrValve()
			v.showValveForm(&valve, true, "engine", nil, nil)
		})

	menu.AddItem("Remote IP Valve",
		"Handle proxy headers (X-Forwarded-For)", 'i', func() {
			valve := server.DefaultRemoteIpValve()
			v.showValveForm(&valve, true, "engine", nil, nil)
		})

	menu.AddItem("Error Report Valve",
		"Customize error pages", 'e', func() {
			valve := server.DefaultErrorReportValve()
			v.showValveForm(&valve, true, "engine", nil, nil)
		})

	menu.AddItem("Single Sign-On Valve",
		"Enable SSO across applications", 's', func() {
			valve := server.DefaultSingleSignOnValve()
			v.showValveForm(&valve, true, "engine", nil, nil)
		})

	menu.AddItem("Stuck Thread Detection Valve",
		"Monitor for stuck request threads", 't', func() {
			valve := server.DefaultStuckThreadDetectionValve()
			v.showValveForm(&valve, true, "engine", nil, nil)
		})

	menu.AddItem("", "", 0, nil)
	menu.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to valve menu", 0, func() {
		v.showMainMenu()
	})

	menu.SetBorder(true).SetTitle(" Quick Add Valve ").SetBorderColor(tcell.ColorGreen)
	menu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("quick-add", menu, true)
	v.app.SetFocus(menu)
}

// showEngineValves shows valves at the engine level
func (v *ValveView) showEngineValves() {
	cfg := v.configService.GetConfig()
	if cfg == nil || len(cfg.Services) == 0 {
		v.setStatus("[red]No service configuration found[-]")
		return
	}

	engine := &cfg.Services[0].Engine
	v.showValveList(engine.Valves, "Engine Valves", "engine", nil, nil)
}

// showHostSelector shows host selection for host-level valves
func (v *ValveView) showHostSelector() {
	cfg := v.configService.GetConfig()
	if cfg == nil || len(cfg.Services) == 0 {
		v.setStatus("[red]No service configuration found[-]")
		return
	}

	list := tview.NewList().ShowSecondaryText(true)

	for i := range cfg.Services[0].Engine.Hosts {
		host := &cfg.Services[0].Engine.Hosts[i]
		info := fmt.Sprintf("%d valves", len(host.Valves))
		list.AddItem(host.Name, info, 0, func() {
			v.showValveList(host.Valves, fmt.Sprintf("Valves for %s", host.Name), "host", host, nil)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to valve menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Select Host ").SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("host-selector", list, true)
	v.app.SetFocus(list)
}

// showContextHostSelector shows host selection for context-level valves
func (v *ValveView) showContextHostSelector() {
	cfg := v.configService.GetConfig()
	if cfg == nil || len(cfg.Services) == 0 {
		v.setStatus("[red]No service configuration found[-]")
		return
	}

	list := tview.NewList().ShowSecondaryText(true)

	for i := range cfg.Services[0].Engine.Hosts {
		host := &cfg.Services[0].Engine.Hosts[i]
		info := fmt.Sprintf("%d contexts", len(host.Contexts))
		list.AddItem(host.Name, info, 0, func() {
			v.showContextSelector(host)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to valve menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Select Host for Context Valves ").SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("context-host-selector", list, true)
	v.app.SetFocus(list)
}

// showContextSelector shows context selection for context-level valves
func (v *ValveView) showContextSelector(host *server.Host) {
	list := tview.NewList().ShowSecondaryText(true)

	for i := range host.Contexts {
		ctx := &host.Contexts[i]
		path := ctx.Path
		if path == "" {
			path = "/"
		}
		info := fmt.Sprintf("%d valves", len(ctx.Valves))
		list.AddItem(path, info, 0, func() {
			v.showValveList(ctx.Valves, fmt.Sprintf("Valves for %s%s", host.Name, path), "context", host, ctx)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to host selector", 0, func() {
		v.showContextHostSelector()
	})

	list.SetBorder(true).SetTitle(fmt.Sprintf(" Contexts for %s ", host.Name)).SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showContextHostSelector()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("context-selector", list, true)
	v.app.SetFocus(list)
}

// showValveList shows the list of valves
func (v *ValveView) showValveList(valves []server.Valve, title, scope string, host *server.Host, ctx *server.Context) {
	list := tview.NewList().ShowSecondaryText(true)

	for i := range valves {
		valve := &valves[i]
		shortName := server.GetValveShortName(valve.ClassName)
		desc := server.GetValveDescription(valve.ClassName)
		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}
		list.AddItem(shortName, desc, 0, func() {
			v.showValveForm(valve, false, scope, host, ctx)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add Valve[-]", "Add new valve", 'a', func() {
		v.showValveTypeSelector(scope, host, ctx)
	})
	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return", 0, func() {
		v.returnFromValveList(scope, host)
	})

	list.SetBorder(true).SetTitle(fmt.Sprintf(" %s ", title)).SetBorderColor(tcell.ColorGreen)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.returnFromValveList(scope, host)
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("valve-list", list, true)
	v.app.SetFocus(list)
}

// returnFromValveList returns to the appropriate menu based on scope
func (v *ValveView) returnFromValveList(scope string, host *server.Host) {
	switch scope {
	case "engine":
		v.showMainMenu()
	case "host":
		v.showHostSelector()
	case "context":
		if host != nil {
			v.showContextSelector(host)
		} else {
			v.showContextHostSelector()
		}
	default:
		v.showMainMenu()
	}
}

// showValveTypeSelector shows valve type selection
func (v *ValveView) showValveTypeSelector(scope string, host *server.Host, ctx *server.Context) {
	list := tview.NewList().ShowSecondaryText(true)

	valveTypes := server.AvailableValveTypes()
	for _, valveClass := range valveTypes {
		vc := valveClass
		name := server.GetValveShortName(vc)
		desc := server.GetValveDescription(vc)
		list.AddItem(name, desc, 0, func() {
			newValve := &server.Valve{ClassName: vc}
			// Set defaults based on type
			switch vc {
			case server.ValveAccessLog:
				*newValve = server.DefaultAccessLogValve()
			case server.ValveRemoteAddr:
				*newValve = server.DefaultRemoteAddrValve()
			case server.ValveRemoteIp:
				*newValve = server.DefaultRemoteIpValve()
			case server.ValveErrorReport:
				*newValve = server.DefaultErrorReportValve()
			case server.ValveSingleSignOn:
				*newValve = server.DefaultSingleSignOnValve()
			case server.ValveStuckThreadDetection:
				*newValve = server.DefaultStuckThreadDetectionValve()
			}
			v.showValveForm(newValve, true, scope, host, ctx)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to valve list", 0, func() {
		v.showValveListForScope(scope, host, ctx)
	})

	list.SetBorder(true).SetTitle(" Select Valve Type ").SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showValveListForScope(scope, host, ctx)
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("valve-type-selector", list, true)
	v.app.SetFocus(list)
}

// showValveListForScope shows the valve list for the given scope
func (v *ValveView) showValveListForScope(scope string, host *server.Host, ctx *server.Context) {
	cfg := v.configService.GetConfig()
	if cfg == nil || len(cfg.Services) == 0 {
		return
	}

	switch scope {
	case "engine":
		v.showValveList(cfg.Services[0].Engine.Valves, "Engine Valves", scope, nil, nil)
	case "host":
		if host != nil {
			v.showValveList(host.Valves, fmt.Sprintf("Valves for %s", host.Name), scope, host, nil)
		}
	case "context":
		if host != nil && ctx != nil {
			path := ctx.Path
			if path == "" {
				path = "/"
			}
			v.showValveList(ctx.Valves, fmt.Sprintf("Valves for %s%s", host.Name, path), scope, host, ctx)
		}
	}
}

// showValveForm shows the valve edit form
func (v *ValveView) showValveForm(valve *server.Valve, isNew bool, scope string, host *server.Host, ctx *server.Context) {
	form := tview.NewForm()
	helpPanel := NewDynamicHelpPanel()
	previewPanel := NewPreviewPanel()

	// Show valve type
	valveName := server.GetValveShortName(valve.ClassName)
	form.AddTextView("Type", valveName, 40, 1, true, false)

	// Create update preview function
	updatePreview := func() {
		tempValve := *valve
		v.readValveFromForm(form, &tempValve)
		previewPanel.SetXMLPreview(GenerateValveXML(&tempValve))
	}

	// Determine which help keys to use based on valve type
	var helpKeys []string
	switch valve.ClassName {
	case server.ValveAccessLog, server.ValveExtendedAccessLog:
		helpKeys = accessLogValveHelpKeys
		v.addAccessLogFields(form, valve, updatePreview)
	case server.ValveRemoteAddr, server.ValveRemoteCIDR, server.ValveRemoteHost:
		helpKeys = remoteAddrValveHelpKeys
		v.addRemoteAddrFields(form, valve, updatePreview)
	case server.ValveRemoteIp:
		helpKeys = remoteIpValveHelpKeys
		v.addRemoteIpFields(form, valve, updatePreview)
	case server.ValveErrorReport:
		helpKeys = errorReportValveHelpKeys
		v.addErrorReportFields(form, valve, updatePreview)
	case server.ValveSingleSignOn:
		helpKeys = singleSignOnValveHelpKeys
		v.addSingleSignOnFields(form, valve, updatePreview)
	case server.ValveStuckThreadDetection:
		helpKeys = stuckThreadValveHelpKeys
		v.addStuckThreadFields(form, valve, updatePreview)
	case server.ValveCrawlerSessionManager:
		helpKeys = crawlerValveHelpKeys
		v.addCrawlerFields(form, valve, updatePreview)
	case server.ValveSemaphore:
		helpKeys = semaphoreValveHelpKeys
		v.addSemaphoreFields(form, valve, updatePreview)
	case server.ValveReplication:
		helpKeys = replicationValveHelpKeys
		v.addReplicationFields(form, valve, updatePreview)
	default:
		// Generic valve - just show class name
		helpKeys = []string{"help.valve.classname"}
		form.AddInputField("Class Name", valve.ClassName, 60, nil, func(text string) {
			updatePreview()
		})
	}

	// Initialize help with first item
	if len(helpKeys) > 0 {
		helpPanel.SetHelpKey(helpKeys[0])
	}

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		v.saveValveFromForm(form, valve, isNew, scope, host, ctx)
	})

	if !isNew {
		form.AddButton("[white:red]"+i18n.T("common.delete")+"[-:-]", func() {
			v.confirmDeleteValve(valve, scope, host, ctx)
		})
	}

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showValveListForScope(scope, host, ctx)
	})

	title := fmt.Sprintf(" Edit %s ", valveName)
	if isNew {
		title = fmt.Sprintf(" Add %s ", valveName)
	}
	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorGreen)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showValveListForScope(scope, host, ctx)
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(helpKeys) {
						helpPanel.SetHelpKey(helpKeys[idx])
					}
				})
			}()
		}
		return event
	})

	// Initial preview
	updatePreview()

	// Layout: left side (form top + preview bottom), right side (help)
	leftPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(previewPanel, 0, 1, false)

	layout := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("valve-form", layout, true)
	v.app.SetFocus(form)
}

// addAccessLogFields adds form fields for AccessLogValve
func (v *ValveView) addAccessLogFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddInputField("Directory", valve.Directory, 40, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Prefix", valve.Prefix, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Suffix", valve.Suffix, 20, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Pattern", valve.Pattern, 60, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("File Date Format", valve.FileDateFormat, 20, nil, func(text string) {
		updatePreview()
	})
	form.AddCheckbox("Rotatable", valve.Rotatable, func(checked bool) {
		updatePreview()
	})
	form.AddCheckbox("Rename On Rotate", valve.RenameOnRotate, func(checked bool) {
		updatePreview()
	})
	form.AddInputField("Encoding", valve.Encoding, 20, nil, func(text string) {
		updatePreview()
	})
	form.AddCheckbox("Buffered", valve.Buffered, func(checked bool) {
		updatePreview()
	})
	form.AddInputField("Condition If", valve.ConditionIf, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Condition Unless", valve.ConditionUnless, 30, nil, func(text string) {
		updatePreview()
	})
}

// addRemoteAddrFields adds form fields for RemoteAddrValve/RemoteCIDRValve
func (v *ValveView) addRemoteAddrFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddInputField("Allow (regex)", valve.Allow, 60, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Deny (regex)", valve.Deny, 60, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Deny Status", strconv.Itoa(valve.DenyStatus), 10, nil, func(text string) {
		updatePreview()
	})
	form.AddCheckbox("Add Connector Port", valve.AddConnectorPort, func(checked bool) {
		updatePreview()
	})
	form.AddCheckbox("Invalid Auth When Deny", valve.InvalidAuthenticationWhenDeny, func(checked bool) {
		updatePreview()
	})
}

// addRemoteIpFields adds form fields for RemoteIpValve
func (v *ValveView) addRemoteIpFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddInputField("Remote IP Header", valve.RemoteIpHeader, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Protocol Header", valve.ProtocolHeader, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Protocol Header HTTPS Value", valve.ProtocolHeaderHttpsValue, 20, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Port Header", valve.PortHeader, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Internal Proxies (regex)", valve.InternalProxies, 60, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Trusted Proxies (regex)", valve.TrustedProxies, 60, nil, func(text string) {
		updatePreview()
	})
	form.AddCheckbox("Change Local Port", valve.ChangeLocalPort, func(checked bool) {
		updatePreview()
	})
	form.AddCheckbox("Change Local Name", valve.ChangeLocalName, func(checked bool) {
		updatePreview()
	})
}

// addErrorReportFields adds form fields for ErrorReportValve
func (v *ValveView) addErrorReportFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddCheckbox("Show Server Info", valve.ShowServerInfo, func(checked bool) {
		updatePreview()
	})
	form.AddCheckbox("Show Report", valve.ShowReport, func(checked bool) {
		updatePreview()
	})
}

// addSingleSignOnFields adds form fields for SingleSignOn
func (v *ValveView) addSingleSignOnFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddInputField("Cookie Domain", valve.CookieDomain, 40, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Cookie Name", valve.CookieName, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddCheckbox("Require Reauthentication", valve.RequireReauthentication, func(checked bool) {
		updatePreview()
	})
}

// addStuckThreadFields adds form fields for StuckThreadDetectionValve
func (v *ValveView) addStuckThreadFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddInputField("Threshold (seconds)", strconv.Itoa(valve.Threshold), 10, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Interrupt Thread Threshold", strconv.Itoa(valve.InterruptThreadThreshold), 10, nil, func(text string) {
		updatePreview()
	})
}

// addCrawlerFields adds form fields for CrawlerSessionManagerValve
func (v *ValveView) addCrawlerFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddInputField("Crawler User Agents", valve.CrawlerUserAgents, 60, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Session Inactive Interval", strconv.Itoa(valve.SessionInactiveInterval), 10, nil, func(text string) {
		updatePreview()
	})
}

// addSemaphoreFields adds form fields for SemaphoreValve
func (v *ValveView) addSemaphoreFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddInputField("Concurrency", strconv.Itoa(valve.Concurrency), 10, nil, func(text string) {
		updatePreview()
	})
	form.AddCheckbox("Fairness", valve.Fairness, func(checked bool) {
		updatePreview()
	})
	form.AddCheckbox("Block", valve.Block, func(checked bool) {
		updatePreview()
	})
}

// addReplicationFields adds form fields for ReplicationValve
func (v *ValveView) addReplicationFields(form *tview.Form, valve *server.Valve, updatePreview func()) {
	form.AddInputField("Filter", valve.Filter, 60, nil, func(text string) {
		updatePreview()
	})
	form.AddCheckbox("Primary Indicator", valve.PrimaryIndicator, func(checked bool) {
		updatePreview()
	})
	form.AddInputField("Primary Indicator Name", valve.PrimaryIndicatorName, 30, nil, func(text string) {
		updatePreview()
	})
}

// readAccessLogValve reads AccessLogValve fields from form
func readAccessLogValve(form *tview.Form, valve *server.Valve) {
	valve.Directory = GetFormText(form, "Directory")
	valve.Prefix = GetFormText(form, "Prefix")
	valve.Suffix = GetFormText(form, "Suffix")
	valve.Pattern = GetFormText(form, "Pattern")
	valve.FileDateFormat = GetFormText(form, "File Date Format")
	valve.Rotatable = GetFormBool(form, "Rotatable")
	valve.RenameOnRotate = GetFormBool(form, "Rename On Rotate")
	valve.Encoding = GetFormText(form, "Encoding")
	valve.Buffered = GetFormBool(form, "Buffered")
	valve.ConditionIf = GetFormText(form, "Condition If")
	valve.ConditionUnless = GetFormText(form, "Condition Unless")
}

// readRemoteAddrValve reads RemoteAddrValve fields from form
func readRemoteAddrValve(form *tview.Form, valve *server.Valve) {
	valve.Allow = GetFormText(form, "Allow (regex)")
	valve.Deny = GetFormText(form, "Deny (regex)")
	valve.DenyStatus = GetFormInt(form, "Deny Status")
	valve.AddConnectorPort = GetFormBool(form, "Add Connector Port")
	valve.InvalidAuthenticationWhenDeny = GetFormBool(form, "Invalid Auth When Deny")
}

// readRemoteIpValve reads RemoteIpValve fields from form
func readRemoteIpValve(form *tview.Form, valve *server.Valve) {
	valve.RemoteIpHeader = GetFormText(form, "Remote IP Header")
	valve.ProtocolHeader = GetFormText(form, "Protocol Header")
	valve.ProtocolHeaderHttpsValue = GetFormText(form, "Protocol Header HTTPS Value")
	valve.PortHeader = GetFormText(form, "Port Header")
	valve.InternalProxies = GetFormText(form, "Internal Proxies (regex)")
	valve.TrustedProxies = GetFormText(form, "Trusted Proxies (regex)")
	valve.ChangeLocalPort = GetFormBool(form, "Change Local Port")
	valve.ChangeLocalName = GetFormBool(form, "Change Local Name")
}

// readErrorReportValve reads ErrorReportValve fields from form
func readErrorReportValve(form *tview.Form, valve *server.Valve) {
	valve.ShowServerInfo = GetFormBool(form, "Show Server Info")
	valve.ShowReport = GetFormBool(form, "Show Report")
}

// readSingleSignOnValve reads SingleSignOnValve fields from form
func readSingleSignOnValve(form *tview.Form, valve *server.Valve) {
	valve.CookieDomain = GetFormText(form, "Cookie Domain")
	valve.CookieName = GetFormText(form, "Cookie Name")
	valve.RequireReauthentication = GetFormBool(form, "Require Reauthentication")
}

// readStuckThreadValve reads StuckThreadDetectionValve fields from form
func readStuckThreadValve(form *tview.Form, valve *server.Valve) {
	valve.Threshold = GetFormInt(form, "Threshold (seconds)")
	valve.InterruptThreadThreshold = GetFormInt(form, "Interrupt Thread Threshold")
}

// readCrawlerSessionValve reads CrawlerSessionManagerValve fields from form
func readCrawlerSessionValve(form *tview.Form, valve *server.Valve) {
	valve.CrawlerUserAgents = GetFormText(form, "Crawler User Agents")
	valve.SessionInactiveInterval = GetFormInt(form, "Session Inactive Interval")
}

// readSemaphoreValve reads SemaphoreValve fields from form
func readSemaphoreValve(form *tview.Form, valve *server.Valve) {
	valve.Concurrency = GetFormInt(form, "Concurrency")
	valve.Fairness = GetFormBool(form, "Fairness")
	valve.Block = GetFormBool(form, "Block")
}

// readReplicationValve reads ReplicationValve fields from form
func readReplicationValve(form *tview.Form, valve *server.Valve) {
	valve.Filter = GetFormText(form, "Filter")
	valve.PrimaryIndicator = GetFormBool(form, "Primary Indicator")
	valve.PrimaryIndicatorName = GetFormText(form, "Primary Indicator Name")
}

// readValveFromForm reads form values into a valve struct
func (v *ValveView) readValveFromForm(form *tview.Form, valve *server.Valve) {
	switch valve.ClassName {
	case server.ValveAccessLog, server.ValveExtendedAccessLog:
		readAccessLogValve(form, valve)
	case server.ValveRemoteAddr, server.ValveRemoteCIDR, server.ValveRemoteHost:
		readRemoteAddrValve(form, valve)
	case server.ValveRemoteIp:
		readRemoteIpValve(form, valve)
	case server.ValveErrorReport:
		readErrorReportValve(form, valve)
	case server.ValveSingleSignOn:
		readSingleSignOnValve(form, valve)
	case server.ValveStuckThreadDetection:
		readStuckThreadValve(form, valve)
	case server.ValveCrawlerSessionManager:
		readCrawlerSessionValve(form, valve)
	case server.ValveSemaphore:
		readSemaphoreValve(form, valve)
	case server.ValveReplication:
		readReplicationValve(form, valve)
	default:
		valve.ClassName = GetFormText(form, "Class Name")
	}
}

// saveValveFromForm saves valve data from form
func (v *ValveView) saveValveFromForm(form *tview.Form, valve *server.Valve, isNew bool, scope string, host *server.Host, ctx *server.Context) {
	// Extract values using the readValveFromForm helper
	v.readValveFromForm(form, valve)

	// Add valve if new
	if isNew {
		cfg := v.configService.GetConfig()
		if cfg != nil && len(cfg.Services) > 0 {
			switch scope {
			case "engine":
				cfg.Services[0].Engine.Valves = append(cfg.Services[0].Engine.Valves, *valve)
			case "host":
				if host != nil {
					host.Valves = append(host.Valves, *valve)
				}
			case "context":
				if ctx != nil {
					ctx.Valves = append(ctx.Valves, *valve)
				}
			}
		}
	}

	if err := v.configService.Save(); err != nil {
		v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
		return
	}

	v.setStatus("[green]Valve saved successfully[-]")
	v.showValveListForScope(scope, host, ctx)
}

// confirmDeleteValve shows delete confirmation
func (v *ValveView) confirmDeleteValve(valve *server.Valve, scope string, host *server.Host, ctx *server.Context) {
	valveName := server.GetValveShortName(valve.ClassName)
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Delete %s?", valveName)).
		AddButtons([]string{"Delete", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Delete" {
				v.deleteValve(valve, scope, host, ctx)
			} else {
				v.showValveListForScope(scope, host, ctx)
			}
		})
	v.pages.AddAndSwitchToPage("confirm-delete", modal, true)
}

// deleteValve deletes a valve
func (v *ValveView) deleteValve(valve *server.Valve, scope string, host *server.Host, ctx *server.Context) {
	cfg := v.configService.GetConfig()
	if cfg == nil || len(cfg.Services) == 0 {
		return
	}

	var valves *[]server.Valve
	switch scope {
	case "engine":
		valves = &cfg.Services[0].Engine.Valves
	case "host":
		if host != nil {
			valves = &host.Valves
		}
	case "context":
		if ctx != nil {
			valves = &ctx.Valves
		}
	}

	if valves != nil {
		for i := range *valves {
			if &(*valves)[i] == valve {
				*valves = append((*valves)[:i], (*valves)[i+1:]...)
				break
			}
		}
	}

	if err := v.configService.Save(); err != nil {
		v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
		return
	}

	v.setStatus("[green]Valve deleted[-]")
	v.showValveListForScope(scope, host, ctx)
}

// setStatus updates the status bar
func (v *ValveView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(fmt.Sprintf(" %s", message))
	}
}
