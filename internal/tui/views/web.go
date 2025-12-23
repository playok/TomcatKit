package views

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/web"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// Servlet form help keys (by form field index)
var servletHelpKeysByIndex = []string{
	"help.servlet.name",         // 0: Servlet Name
	"help.servlet.class",        // 1: Servlet Class
	"help.servlet.jsp",          // 2: JSP File
	"help.servlet.loadonstartup", // 3: Load On Startup
	"help.servlet.async",        // 4: Async Supported
	"help.servlet.initparams",   // 5: Init Params
	"help.servlet.urlpatterns",  // 6: URL Patterns
}

// WebView provides TUI for web.xml configuration
type WebView struct {
	app           *tview.Application
	pages         *tview.Pages
	mainPages     *tview.Pages
	statusBar     *tview.TextView
	onReturn      func()
	configService *web.ConfigService
}

// NewWebView creates a new web.xml configuration view
func NewWebView(app *tview.Application, mainPages *tview.Pages, statusBar *tview.TextView, catalinaBase string, onReturn func()) *WebView {
	return &WebView{
		app:           app,
		mainPages:     mainPages,
		statusBar:     statusBar,
		onReturn:      onReturn,
		configService: web.NewConfigService(catalinaBase),
	}
}

// Load initializes the view
func (v *WebView) Load() error {
	if err := v.configService.Load(); err != nil {
		return fmt.Errorf("failed to load web.xml: %w", err)
	}

	v.pages = tview.NewPages()
	v.showMainMenu()

	v.mainPages.AddAndSwitchToPage("web", v.pages, true)
	return nil
}

// showMainMenu displays the web.xml configuration menu
func (v *WebView) showMainMenu() {
	webapp := v.configService.GetWebApp()

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	servletCount := len(webapp.Servlets)
	filterCount := len(webapp.Filters)
	listenerCount := len(webapp.Listeners)
	errorPageCount := len(webapp.ErrorPages)
	mimeCount := len(webapp.MimeMappings)
	securityCount := len(webapp.SecurityConstraints)

	sessionTimeout := i18n.T("common.notconfigured")
	if webapp.SessionConfig != nil {
		sessionTimeout = fmt.Sprintf("%d %s", webapp.SessionConfig.SessionTimeout, i18n.T("common.minutes"))
	}

	list := tview.NewList().
		AddItem(i18n.T("web.servlets"), fmt.Sprintf(i18n.T("web.servlets.count"), servletCount), 's', func() {
			v.showServletList()
		}).
		AddItem(i18n.T("web.filters"), fmt.Sprintf(i18n.T("web.filters.count"), filterCount), 'f', func() {
			v.showFilterList()
		}).
		AddItem(i18n.T("web.listeners"), fmt.Sprintf(i18n.T("web.listeners.count"), listenerCount), 'l', func() {
			v.showListenerList()
		}).
		AddItem(i18n.T("web.session"), sessionTimeout, 'e', func() {
			v.showSessionConfigForm()
		}).
		AddItem(i18n.T("web.welcomefiles"), i18n.T("web.welcomefiles.desc"), 'w', func() {
			v.showWelcomeFilesForm()
		}).
		AddItem(i18n.T("web.errorpages"), fmt.Sprintf(i18n.T("web.errorpages.count"), errorPageCount), 'r', func() {
			v.showErrorPageList()
		}).
		AddItem(i18n.T("web.mime"), fmt.Sprintf(i18n.T("web.mime.count"), mimeCount), 'm', func() {
			v.showMimeMappingList()
		}).
		AddItem(i18n.T("web.security"), fmt.Sprintf(i18n.T("web.security.count"), securityCount), 'c', func() {
			v.showSecurityConstraintList()
		}).
		AddItem(i18n.T("web.login"), i18n.T("web.login.desc"), 'o', func() {
			v.showLoginConfigForm()
		}).
		AddItem(i18n.T("web.roles"), i18n.T("web.roles.desc"), 'R', func() {
			v.showSecurityRoleList()
		}).
		AddItem(i18n.T("web.contextparams"), i18n.T("web.contextparams.desc"), 'p', func() {
			v.showContextParamList()
		}).
		AddItem(i18n.T("common.save"), i18n.T("web.save.desc"), 'S', func() {
			v.saveConfiguration()
		}).
		AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
			v.mainPages.RemovePage("web")
			v.onReturn()
		})

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.web.servlets")
		case 1:
			helpPanel.SetHelpKey("help.web.filters")
		case 2:
			helpPanel.SetHelpKey("help.web.listeners")
		case 3:
			helpPanel.SetHelpKey("help.web.session")
		case 4:
			helpPanel.SetHelpKey("help.web.welcomefiles")
		case 5:
			helpPanel.SetHelpKey("help.web.errorpages")
		case 6:
			helpPanel.SetHelpKey("help.web.mime")
		case 7:
			helpPanel.SetHelpKey("help.web.security")
		case 8:
			helpPanel.SetHelpKey("help.web.login")
		case 9:
			helpPanel.SetHelpKey("help.web.roles")
		case 10:
			helpPanel.SetHelpKey("help.web.contextparams")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.web.servlets")

	list.SetBorder(true).SetTitle(" " + i18n.T("web.title") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.mainPages.RemovePage("web")
			v.onReturn()
			return nil
		}
		return event
	})

	// Create flex layout with list and help panel
	flex := tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("menu", flex, true)
	v.setStatus("Web configuration: " + v.configService.GetFilePath())
}

// showServletList displays the servlet list
func (v *WebView) showServletList() {
	servlets := v.configService.GetServlets()

	list := tview.NewList()

	for _, servlet := range servlets {
		s := servlet // Capture for closure
		info := s.ServletClass
		if s.JspFile != "" {
			info = "JSP: " + s.JspFile
		}
		list.AddItem(
			s.ServletName,
			info,
			0,
			func() {
				v.showServletForm(&s, false)
			},
		)
	}

	list.AddItem("+ "+i18n.T("web.servlet.add"), i18n.T("web.servlet.add.desc"), 'a', func() {
		s := web.NewServlet("newServlet", "")
		v.showServletForm(s, true)
	})

	list.AddItem("+ "+i18n.T("web.servlet.quickdefault"), i18n.T("web.servlet.quickdefault.desc"), 'd', func() {
		s := &web.Servlet{
			ServletName:   "default",
			ServletClass:  "org.apache.catalina.servlets.DefaultServlet",
			LoadOnStartup: "1",
			InitParams: []web.InitParam{
				{ParamName: "debug", ParamValue: "0"},
				{ParamName: "listings", ParamValue: "false"},
			},
		}
		v.showServletForm(s, true)
	})

	list.AddItem("+ "+i18n.T("web.servlet.quickjsp"), i18n.T("web.servlet.quickjsp.desc"), 'j', func() {
		s := &web.Servlet{
			ServletName:   "jsp",
			ServletClass:  "org.apache.jasper.servlet.JspServlet",
			LoadOnStartup: "3",
			InitParams: []web.InitParam{
				{ParamName: "fork", ParamValue: "false"},
				{ParamName: "xpoweredBy", ParamValue: "false"},
			},
		}
		v.showServletForm(s, true)
	})

	list.AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("web.servlets") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("servlets", list, true)
}

// showServletForm displays a form for editing a servlet
func (v *WebView) showServletForm(servlet *web.Servlet, isNew bool) {
	s := *servlet

	form := tview.NewForm()

	// Help panel on the right
	helpPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	helpPanel.SetBorder(true).SetTitle(" " + i18n.T("help.title") + " ").SetBorderColor(tcell.ColorBlue)

	// Function to update help text based on focused field index
	lastFocusedIndex := -1
	updateHelp := func(index int) {
		if index >= 0 && index < len(servletHelpKeysByIndex) {
			helpPanel.SetText(i18n.T(servletHelpKeysByIndex[index]))
		} else {
			helpPanel.SetText(i18n.T("help.web.servlets"))
		}
	}

	form.AddInputField(i18n.T("web.servlet.name"), s.ServletName, 30, nil, func(text string) {
		s.ServletName = text
	})

	form.AddInputField(i18n.T("web.servlet.class"), s.ServletClass, 60, nil, func(text string) {
		s.ServletClass = text
	})

	form.AddInputField(i18n.T("web.servlet.jsp"), s.JspFile, 40, nil, func(text string) {
		s.JspFile = text
	})

	form.AddInputField(i18n.T("web.servlet.loadonstartup"), s.LoadOnStartup, 5, nil, func(text string) {
		s.LoadOnStartup = text
	})

	asyncSupported := s.AsyncSupported == "true"
	form.AddCheckbox(i18n.T("web.servlet.async"), asyncSupported, func(checked bool) {
		s.AsyncSupported = BoolToString(checked)
	})

	// Init params as text
	initParamsStr := ""
	for _, p := range s.InitParams {
		initParamsStr += fmt.Sprintf("%s=%s\n", p.ParamName, p.ParamValue)
	}
	form.AddTextArea(i18n.T("web.servlet.initparams"), initParamsStr, 60, 4, 0, func(text string) {
		s.InitParams = nil
		for _, line := range strings.Split(text, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				s.InitParams = append(s.InitParams, web.InitParam{
					ParamName:  strings.TrimSpace(parts[0]),
					ParamValue: strings.TrimSpace(parts[1]),
				})
			}
		}
	})

	// URL mappings
	mappingsStr := ""
	for _, m := range v.configService.GetServletMappings() {
		if m.ServletName == servlet.ServletName {
			mappingsStr = strings.Join(m.URLPatterns, "\n")
			break
		}
	}
	form.AddTextArea(i18n.T("web.servlet.urlpatterns"), mappingsStr, 40, 3, 0, nil)

	form.AddButton(i18n.T("common.save"), func() {
		if s.ServletName == "" {
			v.setStatus("Error: " + i18n.T("web.servlet.error.name"))
			return
		}
		if s.ServletClass == "" && s.JspFile == "" {
			v.setStatus("Error: " + i18n.T("web.servlet.error.class"))
			return
		}

		urlPatterns := ParseTextAreaLines(GetFormTextArea(form, i18n.T("web.servlet.urlpatterns")))

		var err error
		if isNew {
			err = v.configService.AddServlet(s)
			if err == nil && len(urlPatterns) > 0 {
				v.configService.AddServletMapping(web.ServletMapping{
					ServletName: s.ServletName,
					URLPatterns: urlPatterns,
				})
			}
		} else {
			err = v.configService.UpdateServlet(servlet.ServletName, s)
		}

		if err != nil {
			v.setStatus("Error: " + err.Error())
			return
		}

		if isNew {
			v.setStatus(i18n.T("web.servlet.added") + ": " + s.ServletName)
		} else {
			v.setStatus(i18n.T("web.servlet.updated") + ": " + s.ServletName)
		}
		v.showServletList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteServlet(servlet.ServletName); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.servlet.deleted") + ": " + servlet.ServletName)
			v.showServletList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showServletList()
	})

	title := " " + i18n.T("web.servlet.add") + " "
	if !isNew {
		title = " " + i18n.T("web.servlet.edit") + ": " + servlet.ServletName + " "
	}
	form.SetBorder(true).SetTitle(title)

	// Initial help
	updateHelp(0)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showServletList()
			return nil
		}
		// Update help after navigation
		go func() {
			v.app.QueueUpdateDraw(func() {
				idx, _ := form.GetFocusedItemIndex()
				if idx != lastFocusedIndex {
					lastFocusedIndex = idx
					updateHelp(idx)
				}
			})
		}()
		return event
	})

	// Create layout with help panel
	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(form, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("servlet-form", layout, true)
}

// showFilterList displays the filter list
func (v *WebView) showFilterList() {
	filters := v.configService.GetFilters()

	list := tview.NewList()

	for _, filter := range filters {
		f := filter // Capture for closure
		list.AddItem(
			f.FilterName,
			f.FilterClass,
			0,
			func() {
				v.showFilterForm(&f, false)
			},
		)
	}

	list.AddItem("+ "+i18n.T("web.filter.add"), i18n.T("web.filter.add.desc"), 'a', func() {
		f := web.NewFilter("newFilter", "")
		v.showFilterForm(f, true)
	})

	list.AddItem("+ "+i18n.T("web.filter.quickcors"), i18n.T("web.filter.quickcors.desc"), 'c', func() {
		f := &web.Filter{
			FilterName:  "CorsFilter",
			FilterClass: "org.apache.catalina.filters.CorsFilter",
			InitParams: []web.InitParam{
				{ParamName: "cors.allowed.origins", ParamValue: "*"},
				{ParamName: "cors.allowed.methods", ParamValue: "GET,POST,PUT,DELETE,OPTIONS"},
				{ParamName: "cors.allowed.headers", ParamValue: "Content-Type,Authorization"},
			},
		}
		v.showFilterForm(f, true)
	})

	list.AddItem("+ "+i18n.T("web.filter.quickencoding"), i18n.T("web.filter.quickencoding.desc"), 'e', func() {
		f := &web.Filter{
			FilterName:  "setCharacterEncodingFilter",
			FilterClass: "org.apache.catalina.filters.SetCharacterEncodingFilter",
			InitParams: []web.InitParam{
				{ParamName: "encoding", ParamValue: "UTF-8"},
				{ParamName: "ignore", ParamValue: "false"},
			},
		}
		v.showFilterForm(f, true)
	})

	list.AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("web.filters") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("filters", list, true)
}

// showFilterForm displays a form for editing a filter
func (v *WebView) showFilterForm(filter *web.Filter, isNew bool) {
	f := *filter

	form := tview.NewForm()

	form.AddInputField(i18n.T("web.filter.name"), f.FilterName, 30, nil, func(text string) {
		f.FilterName = text
	})

	form.AddInputField(i18n.T("web.filter.class"), f.FilterClass, 60, nil, func(text string) {
		f.FilterClass = text
	})

	asyncSupported := f.AsyncSupported == "true"
	form.AddCheckbox(i18n.T("web.filter.async"), asyncSupported, func(checked bool) {
		f.AsyncSupported = BoolToString(checked)
	})

	// Init params
	initParamsStr := ""
	for _, p := range f.InitParams {
		initParamsStr += fmt.Sprintf("%s=%s\n", p.ParamName, p.ParamValue)
	}
	form.AddTextArea(i18n.T("web.filter.initparams"), initParamsStr, 60, 4, 0, func(text string) {
		f.InitParams = nil
		for _, line := range strings.Split(text, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				f.InitParams = append(f.InitParams, web.InitParam{
					ParamName:  strings.TrimSpace(parts[0]),
					ParamValue: strings.TrimSpace(parts[1]),
				})
			}
		}
	})

	// URL patterns for mapping
	mappingsStr := ""
	for _, m := range v.configService.GetFilterMappings() {
		if m.FilterName == filter.FilterName {
			mappingsStr = strings.Join(m.URLPatterns, "\n")
			break
		}
	}
	form.AddTextArea(i18n.T("web.filter.urlpatterns"), mappingsStr, 40, 3, 0, nil)

	form.AddButton(i18n.T("common.save"), func() {
		if f.FilterName == "" || f.FilterClass == "" {
			v.setStatus("Error: " + i18n.T("web.filter.error.required"))
			return
		}

		urlPatterns := ParseTextAreaLines(GetFormTextArea(form, i18n.T("web.filter.urlpatterns")))

		var err error
		if isNew {
			err = v.configService.AddFilter(f)
			if err == nil && len(urlPatterns) > 0 {
				v.configService.AddFilterMapping(web.FilterMapping{
					FilterName:  f.FilterName,
					URLPatterns: urlPatterns,
					Dispatchers: []string{web.DispatcherRequest},
				})
			}
		} else {
			err = v.configService.UpdateFilter(filter.FilterName, f)
		}

		if err != nil {
			v.setStatus("Error: " + err.Error())
			return
		}

		if isNew {
			v.setStatus(i18n.T("web.filter.added") + ": " + f.FilterName)
		} else {
			v.setStatus(i18n.T("web.filter.updated") + ": " + f.FilterName)
		}
		v.showFilterList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteFilter(filter.FilterName); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.filter.deleted") + ": " + filter.FilterName)
			v.showFilterList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showFilterList()
	})

	title := " " + i18n.T("web.filter.add") + " "
	if !isNew {
		title = " " + i18n.T("web.filter.edit") + ": " + filter.FilterName + " "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showFilterList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("filter-form", form, true)
}

// showListenerList displays the listener list
func (v *WebView) showListenerList() {
	listeners := v.configService.GetListeners()

	list := tview.NewList()

	for _, listener := range listeners {
		l := listener // Capture for closure
		// Show short class name
		shortName := l.ListenerClass
		if lastDot := strings.LastIndex(shortName, "."); lastDot > 0 {
			shortName = shortName[lastDot+1:]
		}
		list.AddItem(
			shortName,
			l.ListenerClass,
			0,
			func() {
				v.showListenerForm(&l, false)
			},
		)
	}

	list.AddItem("+ "+i18n.T("web.listener.add"), i18n.T("web.listener.add.desc"), 'a', func() {
		l := &web.Listener{}
		v.showListenerForm(l, true)
	})

	list.AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("web.listeners") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("listeners", list, true)
}

// showListenerForm displays a form for editing a listener
func (v *WebView) showListenerForm(listener *web.Listener, isNew bool) {
	l := *listener

	form := tview.NewForm()

	form.AddInputField(i18n.T("web.listener.class"), l.ListenerClass, 60, nil, func(text string) {
		l.ListenerClass = text
	})

	form.AddInputField(i18n.T("web.listener.description"), l.Description, 50, nil, func(text string) {
		l.Description = text
	})

	form.AddButton(i18n.T("common.save"), func() {
		if l.ListenerClass == "" {
			v.setStatus("Error: " + i18n.T("web.listener.error.class"))
			return
		}

		if isNew {
			if err := v.configService.AddListener(l); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.listener.added"))
		}
		v.showListenerList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteListener(listener.ListenerClass); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.listener.deleted"))
			v.showListenerList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showListenerList()
	})

	title := " " + i18n.T("web.listener.add") + " "
	if !isNew {
		title = " " + i18n.T("web.listener.edit") + " "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showListenerList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("listener-form", form, true)
}

// showSessionConfigForm displays the session configuration form
func (v *WebView) showSessionConfigForm() {
	config := v.configService.GetSessionConfig()
	if config == nil {
		config = &web.SessionConfig{SessionTimeout: 30}
	}

	c := *config
	if c.CookieConfig == nil {
		c.CookieConfig = &web.CookieConfig{}
	}

	form := tview.NewForm()

	form.AddInputField(i18n.T("web.session.timeout"), strconv.Itoa(c.SessionTimeout), 10, acceptNumber, func(text string) {
		if val, err := strconv.Atoi(text); err == nil {
			c.SessionTimeout = val
		}
	})

	// Tracking modes
	trackingModes := []string{web.TrackingCookie, web.TrackingURL, web.TrackingSSL}
	currentModes := make(map[string]bool)
	for _, m := range c.TrackingModes {
		currentModes[m] = true
	}

	form.AddCheckbox(i18n.T("web.session.tracking.cookie"), currentModes[web.TrackingCookie], func(checked bool) {
		currentModes[web.TrackingCookie] = checked
	})

	form.AddCheckbox(i18n.T("web.session.tracking.url"), currentModes[web.TrackingURL], func(checked bool) {
		currentModes[web.TrackingURL] = checked
	})

	form.AddCheckbox(i18n.T("web.session.tracking.ssl"), currentModes[web.TrackingSSL], func(checked bool) {
		currentModes[web.TrackingSSL] = checked
	})

	// Cookie config
	form.AddInputField(i18n.T("web.session.cookie.name"), c.CookieConfig.Name, 20, nil, func(text string) {
		c.CookieConfig.Name = text
	})

	form.AddInputField(i18n.T("web.session.cookie.domain"), c.CookieConfig.Domain, 30, nil, func(text string) {
		c.CookieConfig.Domain = text
	})

	form.AddInputField(i18n.T("web.session.cookie.path"), c.CookieConfig.Path, 20, nil, func(text string) {
		c.CookieConfig.Path = text
	})

	httpOnly := c.CookieConfig.HttpOnly == "true"
	form.AddCheckbox(i18n.T("web.session.cookie.httponly"), httpOnly, func(checked bool) {
		if checked {
			c.CookieConfig.HttpOnly = "true"
		} else {
			c.CookieConfig.HttpOnly = ""
		}
	})

	secure := c.CookieConfig.Secure == "true"
	form.AddCheckbox(i18n.T("web.session.cookie.secure"), secure, func(checked bool) {
		if checked {
			c.CookieConfig.Secure = "true"
		} else {
			c.CookieConfig.Secure = ""
		}
	})

	form.AddButton(i18n.T("common.save"), func() {
		// Update tracking modes
		c.TrackingModes = nil
		for _, mode := range trackingModes {
			if currentModes[mode] {
				c.TrackingModes = append(c.TrackingModes, mode)
			}
		}

		v.configService.SetSessionConfig(&c)
		v.setStatus(i18n.T("web.session.saved"))
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMainMenu()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("web.session.title") + " ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("session-form", form, true)
}

// showWelcomeFilesForm displays the welcome files form
func (v *WebView) showWelcomeFilesForm() {
	files := v.configService.GetWelcomeFiles()

	form := tview.NewForm()

	filesStr := strings.Join(files, "\n")
	form.AddTextArea(i18n.T("web.welcomefiles.perline"), filesStr, 40, 6, 0, nil)

	form.AddButton(i18n.T("common.save"), func() {
		field := form.GetFormItemByLabel(i18n.T("web.welcomefiles.perline")).(*tview.TextArea)
		text := field.GetText()
		var newFiles []string
		for _, line := range strings.Split(text, "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				newFiles = append(newFiles, line)
			}
		}
		v.configService.SetWelcomeFiles(newFiles)
		v.setStatus(i18n.T("web.welcomefiles.saved"))
		v.showMainMenu()
	})

	form.AddButton(i18n.T("web.welcomefiles.adddefaults"), func() {
		v.configService.SetWelcomeFiles([]string{"index.html", "index.htm", "index.jsp"})
		v.setStatus(i18n.T("web.welcomefiles.defaultadded"))
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMainMenu()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("web.welcomefiles.title") + " ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("welcome-form", form, true)
}

// showErrorPageList displays the error page list
func (v *WebView) showErrorPageList() {
	errorPages := v.configService.GetErrorPages()

	list := tview.NewList()

	for i, ep := range errorPages {
		idx := i // Capture for closure
		e := ep
		code := e.ErrorCode
		if code == "" {
			code = "Exception: " + e.ExceptionType
		}
		list.AddItem(
			code,
			e.Location,
			0,
			func() {
				v.showErrorPageForm(&e, idx, false)
			},
		)
	}

	list.AddItem("+ "+i18n.T("web.errorpage.add"), i18n.T("web.errorpage.add.desc"), 'a', func() {
		e := &web.ErrorPage{}
		v.showErrorPageForm(e, -1, true)
	})

	list.AddItem("+ "+i18n.T("web.errorpage.quickcommon"), i18n.T("web.errorpage.quickcommon.desc"), 'c', func() {
		v.configService.AddErrorPage(*web.NewErrorPage("404", "/error/404.html"))
		v.configService.AddErrorPage(*web.NewErrorPage("500", "/error/500.html"))
		v.setStatus(i18n.T("web.errorpage.commonadded"))
		v.showErrorPageList()
	})

	list.AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("web.errorpages") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("error-pages", list, true)
}

// showErrorPageForm displays a form for editing an error page
func (v *WebView) showErrorPageForm(errorPage *web.ErrorPage, index int, isNew bool) {
	e := *errorPage

	form := tview.NewForm()

	// Error type selection
	errorTypes := []string{"Error Code", "Exception Type"}
	errorTypeIdx := 0
	if e.ExceptionType != "" {
		errorTypeIdx = 1
	}
	form.AddDropDown(i18n.T("web.errorpage.type"), errorTypes, errorTypeIdx, nil)

	form.AddInputField(i18n.T("web.errorpage.code"), e.ErrorCode, 10, nil, func(text string) {
		e.ErrorCode = text
	})

	form.AddInputField(i18n.T("web.errorpage.exception"), e.ExceptionType, 50, nil, func(text string) {
		e.ExceptionType = text
	})

	form.AddInputField(i18n.T("web.errorpage.location"), e.Location, 40, nil, func(text string) {
		e.Location = text
	})

	form.AddButton(i18n.T("common.save"), func() {
		if e.Location == "" {
			v.setStatus("Error: " + i18n.T("web.errorpage.error.location"))
			return
		}

		// Clear the non-selected type
		dropdown := form.GetFormItem(0).(*tview.DropDown)
		_, option := dropdown.GetCurrentOption()
		if option == "Error Code" {
			e.ExceptionType = ""
		} else {
			e.ErrorCode = ""
		}

		if isNew {
			v.configService.AddErrorPage(e)
			v.setStatus(i18n.T("web.errorpage.added"))
		}
		v.showErrorPageList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteErrorPage(index); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.errorpage.deleted"))
			v.showErrorPageList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showErrorPageList()
	})

	title := " " + i18n.T("web.errorpage.add") + " "
	if !isNew {
		title = " " + i18n.T("web.errorpage.edit") + " "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showErrorPageList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("error-page-form", form, true)
}

// showMimeMappingList displays the MIME mapping list
func (v *WebView) showMimeMappingList() {
	mappings := v.configService.GetMimeMappings()

	list := tview.NewList()

	for _, m := range mappings {
		mapping := m // Capture for closure
		list.AddItem(
			"."+mapping.Extension,
			mapping.MimeType,
			0,
			func() {
				v.showMimeMappingForm(&mapping, false)
			},
		)
	}

	list.AddItem("+ "+i18n.T("web.mime.add"), i18n.T("web.mime.add.desc"), 'a', func() {
		m := &web.MimeMapping{}
		v.showMimeMappingForm(m, true)
	})

	list.AddItem("+ "+i18n.T("web.mime.quickcommon"), i18n.T("web.mime.quickcommon.desc"), 'c', func() {
		for ext, mime := range web.CommonMimeMappings {
			v.configService.AddMimeMapping(web.MimeMapping{Extension: ext, MimeType: mime})
		}
		v.setStatus(i18n.T("web.mime.commonadded"))
		v.showMimeMappingList()
	})

	list.AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("web.mime") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("mime-mappings", list, true)
}

// showMimeMappingForm displays a form for editing a MIME mapping
func (v *WebView) showMimeMappingForm(mapping *web.MimeMapping, isNew bool) {
	m := *mapping

	form := tview.NewForm()

	form.AddInputField(i18n.T("web.mime.extension"), m.Extension, 15, nil, func(text string) {
		m.Extension = strings.TrimPrefix(text, ".")
	})

	form.AddInputField(i18n.T("web.mime.type"), m.MimeType, 40, nil, func(text string) {
		m.MimeType = text
	})

	form.AddButton(i18n.T("common.save"), func() {
		if m.Extension == "" || m.MimeType == "" {
			v.setStatus("Error: " + i18n.T("web.mime.error.required"))
			return
		}

		if isNew {
			if err := v.configService.AddMimeMapping(m); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.mime.added"))
		}
		v.showMimeMappingList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteMimeMapping(mapping.Extension); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.mime.deleted"))
			v.showMimeMappingList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMimeMappingList()
	})

	title := " " + i18n.T("web.mime.add") + " "
	if !isNew {
		title = " " + i18n.T("web.mime.edit") + " "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMimeMappingList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("mime-form", form, true)
}

// showSecurityConstraintList displays the security constraint list
func (v *WebView) showSecurityConstraintList() {
	constraints := v.configService.GetSecurityConstraints()

	list := tview.NewList()

	for i, c := range constraints {
		idx := i // Capture for closure
		constraint := c
		name := "Constraint"
		if len(constraint.WebResourceCollections) > 0 {
			name = constraint.WebResourceCollections[0].WebResourceName
		}
		roles := "No roles"
		if constraint.AuthConstraint != nil && len(constraint.AuthConstraint.RoleNames) > 0 {
			roles = strings.Join(constraint.AuthConstraint.RoleNames, ", ")
		}
		list.AddItem(
			name,
			fmt.Sprintf("Roles: %s", roles),
			0,
			func() {
				v.showSecurityConstraintForm(&constraint, idx, false)
			},
		)
	}

	list.AddItem("+ "+i18n.T("web.securityconstraint.add"), i18n.T("web.securityconstraint.add.desc"), 'a', func() {
		c := web.NewSecurityConstraint("Protected Resource", []string{"/*"})
		v.showSecurityConstraintForm(c, -1, true)
	})

	list.AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("web.security") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("security-constraints", list, true)
}

// showSecurityConstraintForm displays a form for editing a security constraint
func (v *WebView) showSecurityConstraintForm(constraint *web.SecurityConstraint, index int, isNew bool) {
	c := *constraint

	form := tview.NewForm()

	// Resource name
	resourceName := ""
	if len(c.WebResourceCollections) > 0 {
		resourceName = c.WebResourceCollections[0].WebResourceName
	}
	form.AddInputField(i18n.T("web.securityconstraint.resourcename"), resourceName, 30, nil, func(text string) {
		if len(c.WebResourceCollections) == 0 {
			c.WebResourceCollections = []web.WebResourceCollection{{}}
		}
		c.WebResourceCollections[0].WebResourceName = text
	})

	// Helper to ensure WebResourceCollection exists
	ensureCollection := func() {
		if len(c.WebResourceCollections) == 0 {
			c.WebResourceCollections = []web.WebResourceCollection{{}}
		}
	}

	// URL patterns
	urlPatterns := ""
	if len(c.WebResourceCollections) > 0 {
		urlPatterns = strings.Join(c.WebResourceCollections[0].URLPatterns, "\n")
	}
	form.AddTextArea(i18n.T("web.securityconstraint.urlpatterns"), urlPatterns, 40, 3, 0, func(text string) {
		ensureCollection()
		c.WebResourceCollections[0].URLPatterns = ParseTextAreaLines(text)
	})

	// HTTP methods
	httpMethods := ""
	if len(c.WebResourceCollections) > 0 {
		httpMethods = strings.Join(c.WebResourceCollections[0].HTTPMethods, ", ")
	}
	form.AddInputField(i18n.T("web.securityconstraint.httpmethods"), httpMethods, 40, nil, func(text string) {
		ensureCollection()
		c.WebResourceCollections[0].HTTPMethods = ParseCommaSeparated(text)
	})

	// Role names
	roles := ""
	if c.AuthConstraint != nil {
		roles = strings.Join(c.AuthConstraint.RoleNames, "\n")
	}
	form.AddTextArea(i18n.T("web.securityconstraint.roles"), roles, 30, 3, 0, func(text string) {
		roleNames := ParseTextAreaLines(text)
		if len(roleNames) > 0 {
			if c.AuthConstraint == nil {
				c.AuthConstraint = &web.AuthConstraint{}
			}
			c.AuthConstraint.RoleNames = roleNames
		} else {
			c.AuthConstraint = nil
		}
	})

	// Transport guarantee
	transportOptions := []string{web.TransportNone, web.TransportIntegral, web.TransportConfidential}
	transportIdx := 0
	if c.UserDataConstraint != nil {
		for i, t := range transportOptions {
			if t == c.UserDataConstraint.TransportGuarantee {
				transportIdx = i
				break
			}
		}
	}
	form.AddDropDown(i18n.T("web.securityconstraint.transport"), transportOptions, transportIdx, func(option string, index int) {
		if option == web.TransportNone {
			c.UserDataConstraint = nil
		} else {
			c.UserDataConstraint = &web.UserDataConstraint{TransportGuarantee: option}
		}
	})

	form.AddButton(i18n.T("common.save"), func() {
		if isNew {
			v.configService.AddSecurityConstraint(c)
			v.setStatus(i18n.T("web.securityconstraint.added"))
		} else {
			if err := v.configService.UpdateSecurityConstraint(index, c); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.securityconstraint.updated"))
		}
		v.showSecurityConstraintList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteSecurityConstraint(index); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.securityconstraint.deleted"))
			v.showSecurityConstraintList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showSecurityConstraintList()
	})

	title := " " + i18n.T("web.securityconstraint.add") + " "
	if !isNew {
		title = " " + i18n.T("web.securityconstraint.edit") + " "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showSecurityConstraintList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("security-constraint-form", form, true)
}

// showLoginConfigForm displays the login configuration form
func (v *WebView) showLoginConfigForm() {
	config := v.configService.GetLoginConfig()
	if config == nil {
		config = &web.LoginConfig{}
	}

	c := *config
	if c.FormLoginConfig == nil {
		c.FormLoginConfig = &web.FormLoginConfig{}
	}

	form := tview.NewForm()

	// Auth method
	authMethods := []string{"(None)", web.AuthMethodBasic, web.AuthMethodDigest, web.AuthMethodForm, web.AuthMethodClientCert}
	authIdx := 0
	for i, m := range authMethods {
		if m == c.AuthMethod {
			authIdx = i
			break
		}
	}
	form.AddDropDown(i18n.T("web.login.authmethod"), authMethods, authIdx, func(option string, index int) {
		if option == "(None)" {
			c.AuthMethod = ""
		} else {
			c.AuthMethod = option
		}
	})

	form.AddInputField(i18n.T("web.login.realmname"), c.RealmName, 30, nil, func(text string) {
		c.RealmName = text
	})

	// Form login config (only for FORM auth)
	form.AddInputField(i18n.T("web.login.formloginpage"), c.FormLoginConfig.FormLoginPage, 40, nil, func(text string) {
		c.FormLoginConfig.FormLoginPage = text
	})

	form.AddInputField(i18n.T("web.login.formerrorpage"), c.FormLoginConfig.FormErrorPage, 40, nil, func(text string) {
		c.FormLoginConfig.FormErrorPage = text
	})

	form.AddButton(i18n.T("common.save"), func() {
		if c.AuthMethod == "" {
			v.configService.RemoveLoginConfig()
			v.setStatus(i18n.T("web.login.removed"))
		} else {
			// Only include form config for FORM auth
			if c.AuthMethod != web.AuthMethodForm {
				c.FormLoginConfig = nil
			}
			v.configService.SetLoginConfig(&c)
			v.setStatus(i18n.T("web.login.saved"))
		}
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMainMenu()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("web.login.title") + " ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("login-form", form, true)
}

// showSecurityRoleList displays the security role list
func (v *WebView) showSecurityRoleList() {
	roles := v.configService.GetSecurityRoles()

	list := tview.NewList()

	for _, role := range roles {
		r := role // Capture for closure
		list.AddItem(
			r.RoleName,
			r.Description,
			0,
			func() {
				v.showSecurityRoleForm(&r, false)
			},
		)
	}

	list.AddItem("+ "+i18n.T("web.role.add"), i18n.T("web.role.add.desc"), 'a', func() {
		r := &web.SecurityRole{}
		v.showSecurityRoleForm(r, true)
	})

	list.AddItem("+ "+i18n.T("web.role.quickcommon"), i18n.T("web.role.quickcommon.desc"), 'c', func() {
		v.configService.AddSecurityRole(web.SecurityRole{RoleName: "admin", Description: "Administrator"})
		v.configService.AddSecurityRole(web.SecurityRole{RoleName: "user", Description: "Standard user"})
		v.configService.AddSecurityRole(web.SecurityRole{RoleName: "manager", Description: "Manager role"})
		v.setStatus(i18n.T("web.role.commonadded"))
		v.showSecurityRoleList()
	})

	list.AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("web.roles") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("security-roles", list, true)
}

// showSecurityRoleForm displays a form for editing a security role
func (v *WebView) showSecurityRoleForm(role *web.SecurityRole, isNew bool) {
	r := *role

	form := tview.NewForm()

	form.AddInputField(i18n.T("web.role.name"), r.RoleName, 30, nil, func(text string) {
		r.RoleName = text
	})

	form.AddInputField(i18n.T("web.role.description"), r.Description, 50, nil, func(text string) {
		r.Description = text
	})

	form.AddButton(i18n.T("common.save"), func() {
		if r.RoleName == "" {
			v.setStatus("Error: " + i18n.T("web.role.error.name"))
			return
		}

		if isNew {
			if err := v.configService.AddSecurityRole(r); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.role.added"))
		}
		v.showSecurityRoleList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteSecurityRole(role.RoleName); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.role.deleted"))
			v.showSecurityRoleList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showSecurityRoleList()
	})

	title := " " + i18n.T("web.role.add") + " "
	if !isNew {
		title = " " + i18n.T("web.role.edit") + " "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showSecurityRoleList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("security-role-form", form, true)
}

// showContextParamList displays the context parameter list
func (v *WebView) showContextParamList() {
	params := v.configService.GetContextParams()

	list := tview.NewList()

	for _, param := range params {
		p := param // Capture for closure
		list.AddItem(
			p.ParamName,
			p.ParamValue,
			0,
			func() {
				v.showContextParamForm(&p, false)
			},
		)
	}

	list.AddItem("+ "+i18n.T("web.contextparam.add"), i18n.T("web.contextparam.add.desc"), 'a', func() {
		p := &web.ContextParam{}
		v.showContextParamForm(p, true)
	})

	list.AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" " + i18n.T("web.contextparams") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("context-params", list, true)
}

// showContextParamForm displays a form for editing a context parameter
func (v *WebView) showContextParamForm(param *web.ContextParam, isNew bool) {
	p := *param

	form := tview.NewForm()

	form.AddInputField(i18n.T("web.contextparam.name"), p.ParamName, 30, nil, func(text string) {
		p.ParamName = text
	})

	form.AddInputField(i18n.T("web.contextparam.value"), p.ParamValue, 50, nil, func(text string) {
		p.ParamValue = text
	})

	form.AddInputField(i18n.T("web.contextparam.description"), p.Description, 50, nil, func(text string) {
		p.Description = text
	})

	form.AddButton(i18n.T("common.save"), func() {
		if p.ParamName == "" {
			v.setStatus("Error: " + i18n.T("web.contextparam.error.name"))
			return
		}

		if isNew {
			if err := v.configService.AddContextParam(p); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.contextparam.added"))
		} else {
			if err := v.configService.UpdateContextParam(param.ParamName, p); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.contextparam.updated"))
		}
		v.showContextParamList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteContextParam(param.ParamName); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus(i18n.T("web.contextparam.deleted"))
			v.showContextParamList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showContextParamList()
	})

	title := " " + i18n.T("web.contextparam.add") + " "
	if !isNew {
		title = " " + i18n.T("web.contextparam.edit") + " "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showContextParamList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("context-param-form", form, true)
}

// saveConfiguration saves the web configuration
func (v *WebView) saveConfiguration() {
	if err := v.configService.Save(); err != nil {
		v.setStatus("Error saving: " + err.Error())
		return
	}
	v.setStatus("Configuration saved to: " + v.configService.GetFilePath())
}

// setStatus updates the status bar
func (v *WebView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(" " + message)
	}
}
