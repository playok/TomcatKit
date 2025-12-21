package views

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/jndi"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// ContextView provides TUI for context.xml configuration
type ContextView struct {
	app           *tview.Application
	pages         *tview.Pages
	mainPages     *tview.Pages
	statusBar     *tview.TextView
	onReturn      func()
	configService *jndi.ContextService
}

// NewContextView creates a new context configuration view
func NewContextView(app *tview.Application, mainPages *tview.Pages, statusBar *tview.TextView, catalinaBase string, onReturn func()) *ContextView {
	return &ContextView{
		app:           app,
		mainPages:     mainPages,
		statusBar:     statusBar,
		onReturn:      onReturn,
		configService: jndi.NewContextService(catalinaBase),
	}
}

// Load initializes the view
func (v *ContextView) Load() error {
	if err := v.configService.Load(); err != nil {
		return fmt.Errorf("failed to load context.xml: %w", err)
	}

	v.pages = tview.NewPages()
	v.showMainMenu()

	v.mainPages.AddAndSwitchToPage("context", v.pages, true)
	return nil
}

// showMainMenu displays the context.xml configuration menu
func (v *ContextView) showMainMenu() {
	ctx := v.configService.GetContext()

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	// Build status info
	resourceCount := len(ctx.Resources)
	envCount := len(ctx.Environments)
	paramCount := len(ctx.Parameters)
	watchedCount := len(ctx.WatchedResources)

	managerStatus := i18n.T("common.notconfigured")
	if ctx.Manager != nil {
		if ctx.Manager.ClassName == jndi.ManagerPersistent {
			managerStatus = "PersistentManager"
		} else {
			managerStatus = "StandardManager"
		}
	}

	list := tview.NewList().
		AddItem(i18n.T("context.settings"), i18n.T("context.settings.desc"), 's', func() {
			v.showContextSettingsForm()
		}).
		AddItem(i18n.T("context.resources"), fmt.Sprintf(i18n.T("context.resources.count"), resourceCount), 'r', func() {
			v.showResourceList()
		}).
		AddItem(i18n.T("context.environment"), fmt.Sprintf(i18n.T("context.environment.count"), envCount), 'e', func() {
			v.showEnvironmentList()
		}).
		AddItem(i18n.T("context.resourcelinks"), fmt.Sprintf(i18n.T("context.resourcelinks.count"), len(ctx.ResourceLinks)), 'l', func() {
			v.showResourceLinkList()
		}).
		AddItem(i18n.T("context.parameters"), fmt.Sprintf(i18n.T("context.parameters.count"), paramCount), 'p', func() {
			v.showParameterList()
		}).
		AddItem(i18n.T("context.watched"), fmt.Sprintf(i18n.T("context.watched.count"), watchedCount), 'w', func() {
			v.showWatchedResourceList()
		}).
		AddItem(i18n.T("context.manager"), managerStatus, 'm', func() {
			v.showManagerForm()
		}).
		AddItem(i18n.T("context.cookie"), i18n.T("context.cookie.desc"), 'c', func() {
			v.showCookieProcessorForm()
		}).
		AddItem(i18n.T("context.jarscanner"), i18n.T("context.jarscanner.desc"), 'j', func() {
			v.showJarScannerForm()
		}).
		AddItem(i18n.T("common.save"), i18n.T("context.save.desc"), 'S', func() {
			v.saveConfiguration()
		}).
		AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
			v.mainPages.RemovePage("context")
			v.onReturn()
		})

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.context.settings")
		case 1:
			helpPanel.SetHelpKey("help.context.resources")
		case 2:
			helpPanel.SetHelpKey("help.context.environment")
		case 3:
			helpPanel.SetHelpKey("help.context.resourcelinks")
		case 4:
			helpPanel.SetHelpKey("help.context.parameters")
		case 5:
			helpPanel.SetHelpKey("help.context.watched")
		case 6:
			helpPanel.SetHelpKey("help.context.manager")
		case 7:
			helpPanel.SetHelpKey("help.context.cookie")
		case 8:
			helpPanel.SetHelpKey("help.context.jarscanner")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.context.settings")

	list.SetBorder(true).SetTitle(" " + i18n.T("context.title") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.mainPages.RemovePage("context")
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
	v.setStatus("Context configuration: " + v.configService.GetFilePath())
}

// showContextSettingsForm displays the context settings form
func (v *ContextView) showContextSettingsForm() {
	ctx := v.configService.GetContext()

	form := tview.NewForm()

	// Basic settings
	form.AddCheckbox("Reloadable", ctx.Reloadable, func(checked bool) {
		ctx.Reloadable = checked
	})

	form.AddCheckbox("CrossContext", ctx.CrossContext, func(checked bool) {
		ctx.CrossContext = checked
	})

	form.AddCheckbox("Privileged", ctx.Privileged, func(checked bool) {
		ctx.Privileged = checked
	})

	// Cookie settings
	cookies := ctx.Cookies
	if cookies == "" {
		cookies = "true"
	}
	cookieOptions := []string{"true", "false"}
	cookieIdx := 0
	if cookies == "false" {
		cookieIdx = 1
	}
	form.AddDropDown("Cookies", cookieOptions, cookieIdx, func(option string, index int) {
		ctx.Cookies = option
	})

	useHttpOnly := ctx.UseHttpOnly
	if useHttpOnly == "" {
		useHttpOnly = "true"
	}
	httpOnlyIdx := 0
	if useHttpOnly == "false" {
		httpOnlyIdx = 1
	}
	form.AddDropDown("UseHttpOnly", []string{"true", "false"}, httpOnlyIdx, func(option string, index int) {
		ctx.UseHttpOnly = option
	})

	form.AddInputField("Session Cookie Name", ctx.SessionCookieName, 30, nil, func(text string) {
		ctx.SessionCookieName = text
	})

	// Caching
	cachingAllowed := ctx.CachingAllowed
	if cachingAllowed == "" {
		cachingAllowed = "true"
	}
	cachingIdx := 0
	if cachingAllowed == "false" {
		cachingIdx = 1
	}
	form.AddDropDown("Caching Allowed", []string{"true", "false"}, cachingIdx, func(option string, index int) {
		ctx.CachingAllowed = option
	})

	form.AddInputField("Cache Max Size (KB)", strconv.Itoa(ctx.CacheMaxSize), 10, acceptNumber, func(text string) {
		if size, err := strconv.Atoi(text); err == nil {
			ctx.CacheMaxSize = size
		}
	})

	// Resource locking
	antiLocking := ctx.AntiResourceLocking
	antiLockingIdx := 1
	if antiLocking == "true" {
		antiLockingIdx = 0
	}
	form.AddDropDown("Anti Resource Locking", []string{"true", "false"}, antiLockingIdx, func(option string, index int) {
		ctx.AntiResourceLocking = option
	})

	// Swallow output
	swallowIdx := 1
	if ctx.SwallowOutput == "true" {
		swallowIdx = 0
	}
	form.AddDropDown("Swallow Output", []string{"true", "false"}, swallowIdx, func(option string, index int) {
		ctx.SwallowOutput = option
	})

	form.AddButton(i18n.T("common.save.short"), func() {
		v.configService.UpdateContextSettings(ctx)
		v.setStatus("Context settings updated")
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMainMenu()
	})

	form.SetBorder(true).SetTitle(" Context Settings ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("settings-form", form, true)
}

// showResourceList displays the JNDI resources list
func (v *ContextView) showResourceList() {
	resources := v.configService.GetResources()

	list := tview.NewList()

	for _, res := range resources {
		r := res // Capture for closure
		resourceType := "Resource"
		if r.Type == string(jndi.ResourceTypeDataSource) {
			resourceType = "DataSource"
		} else if r.Type == string(jndi.ResourceTypeMailSession) {
			resourceType = "MailSession"
		} else if r.Type == string(jndi.ResourceTypeUserDatabase) {
			resourceType = "UserDatabase"
		}
		list.AddItem(
			r.Name,
			fmt.Sprintf("%s | Auth: %s", resourceType, r.Auth),
			0,
			func() {
				v.showResourceForm(&r, false)
			},
		)
	}

	list.AddItem("+ Add DataSource", "Create a new DataSource", 'd', func() {
		ds := jndi.NewDataSourceResource("jdbc/newDB")
		v.showResourceForm(ds, true)
	})

	list.AddItem("+ Add Mail Session", "Create a new Mail Session", 'm', func() {
		ms := jndi.NewMailSessionResource("mail/Session")
		v.showResourceForm(ms, true)
	})

	list.AddItem("Back", "Return to context menu", 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" JNDI Resources ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("resources", list, true)
}

// showResourceForm displays a form for editing a resource
func (v *ContextView) showResourceForm(resource *jndi.Resource, isNew bool) {
	r := *resource

	form := tview.NewForm()

	form.AddInputField("Name", r.Name, 40, nil, func(text string) {
		r.Name = text
	})

	// Auth dropdown
	authOptions := []string{jndi.AuthContainer, jndi.AuthApplication}
	authIdx := 0
	if r.Auth == jndi.AuthApplication {
		authIdx = 1
	}
	form.AddDropDown("Auth", authOptions, authIdx, func(option string, index int) {
		r.Auth = option
	})

	// Type-specific fields
	if r.Type == string(jndi.ResourceTypeDataSource) {
		// Database type selector
		dbTypes := []string{"MySQL", "PostgreSQL", "Oracle", "SQL Server", "MariaDB", "H2", "HSQLDB", "Derby", "SQLite"}
		form.AddDropDown("Database Type", dbTypes, 0, func(option string, index int) {
			if driver, ok := jndi.CommonDrivers[option]; ok {
				r.DriverClassName = driver
			}
			if url, ok := jndi.JDBCURLTemplates[option]; ok {
				r.URL = url
			}
			if query, ok := jndi.ValidationQueries[option]; ok {
				r.ValidationQuery = query
			}
		})

		form.AddInputField("Driver Class", r.DriverClassName, 50, nil, func(text string) {
			r.DriverClassName = text
		})

		form.AddInputField("URL", r.URL, 60, nil, func(text string) {
			r.URL = text
		})

		form.AddInputField("Username", r.Username, 30, nil, func(text string) {
			r.Username = text
		})

		form.AddPasswordField("Password", r.Password, 30, '*', func(text string) {
			r.Password = text
		})

		form.AddInputField("Initial Size", strconv.Itoa(r.InitialSize), 10, acceptNumber, func(text string) {
			r.InitialSize, _ = strconv.Atoi(text)
		})

		form.AddInputField("Max Total", strconv.Itoa(r.MaxTotal), 10, acceptNumber, func(text string) {
			r.MaxTotal, _ = strconv.Atoi(text)
		})

		form.AddInputField("Min Idle", strconv.Itoa(r.MinIdle), 10, acceptNumber, func(text string) {
			r.MinIdle, _ = strconv.Atoi(text)
		})

		form.AddInputField("Max Idle", strconv.Itoa(r.MaxIdle), 10, acceptNumber, func(text string) {
			r.MaxIdle, _ = strconv.Atoi(text)
		})

		form.AddInputField("Validation Query", r.ValidationQuery, 40, nil, func(text string) {
			r.ValidationQuery = text
		})

		form.AddCheckbox("Test On Borrow", r.TestOnBorrow, func(checked bool) {
			r.TestOnBorrow = checked
		})

		form.AddCheckbox("Test While Idle", r.TestWhileIdle, func(checked bool) {
			r.TestWhileIdle = checked
		})

	} else if r.Type == string(jndi.ResourceTypeMailSession) {
		form.AddInputField("SMTP Host", r.MailSmtpHost, 40, nil, func(text string) {
			r.MailSmtpHost = text
		})

		form.AddInputField("SMTP Port", r.MailSmtpPort, 10, nil, func(text string) {
			r.MailSmtpPort = text
		})

		authEnabled := r.MailSmtpAuth == "true"
		form.AddCheckbox("SMTP Auth", authEnabled, func(checked bool) {
			r.MailSmtpAuth = BoolToString(checked)
		})

		starttls := r.MailSmtpStartTLS == "true"
		form.AddCheckbox("STARTTLS", starttls, func(checked bool) {
			r.MailSmtpStartTLS = BoolToString(checked)
		})

		form.AddInputField("SMTP User", r.MailSmtpUser, 40, nil, func(text string) {
			r.MailSmtpUser = text
		})
	}

	form.AddButton(i18n.T("common.save.short"), func() {
		if r.Name == "" {
			v.setStatus("Error: Name is required")
			return
		}

		if isNew {
			if err := v.configService.AddResource(r); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Resource added: " + r.Name)
		} else {
			if err := v.configService.UpdateResource(resource.Name, r); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Resource updated: " + r.Name)
		}
		v.showResourceList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteResource(resource.Name); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Resource deleted: " + resource.Name)
			v.showResourceList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showResourceList()
	})

	title := " Add Resource "
	if !isNew {
		title = " Edit Resource: " + resource.Name + " "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showResourceList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("resource-form", form, true)
}

// showEnvironmentList displays the environment entries list
func (v *ContextView) showEnvironmentList() {
	envs := v.configService.GetEnvironments()

	list := tview.NewList()

	for _, env := range envs {
		e := env // Capture for closure
		list.AddItem(
			e.Name,
			fmt.Sprintf("Type: %s | Value: %s", e.Type, e.Value),
			0,
			func() {
				v.showEnvironmentForm(&e, false)
			},
		)
	}

	list.AddItem("+ Add Environment Entry", "Create a new environment entry", 'a', func() {
		env := jndi.NewEnvironment("newEnv", "", "java.lang.String")
		v.showEnvironmentForm(env, true)
	})

	list.AddItem("Back", "Return to context menu", 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Environment Entries ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("environments", list, true)
}

// showEnvironmentForm displays a form for editing an environment entry
func (v *ContextView) showEnvironmentForm(env *jndi.Environment, isNew bool) {
	e := *env

	form := tview.NewForm()

	form.AddInputField("Name", e.Name, 40, nil, func(text string) {
		e.Name = text
	})

	form.AddInputField("Value", e.Value, 50, nil, func(text string) {
		e.Value = text
	})

	typeIdx := 0
	for i, t := range jndi.EnvironmentTypes {
		if t == e.Type {
			typeIdx = i
			break
		}
	}
	form.AddDropDown("Type", jndi.EnvironmentTypes, typeIdx, func(option string, index int) {
		e.Type = option
	})

	form.AddCheckbox("Override", e.Override, func(checked bool) {
		e.Override = checked
	})

	form.AddInputField("Description", e.Description, 50, nil, func(text string) {
		e.Description = text
	})

	form.AddButton(i18n.T("common.save.short"), func() {
		if e.Name == "" {
			v.setStatus("Error: Name is required")
			return
		}

		if isNew {
			if err := v.configService.AddEnvironment(e); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Environment entry added: " + e.Name)
		} else {
			if err := v.configService.UpdateEnvironment(env.Name, e); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Environment entry updated: " + e.Name)
		}
		v.showEnvironmentList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteEnvironment(env.Name); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Environment entry deleted: " + env.Name)
			v.showEnvironmentList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showEnvironmentList()
	})

	title := " Add Environment Entry "
	if !isNew {
		title = " Edit Environment Entry "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showEnvironmentList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("environment-form", form, true)
}

// showResourceLinkList displays the resource links list
func (v *ContextView) showResourceLinkList() {
	links := v.configService.GetResourceLinks()

	list := tview.NewList()

	for _, link := range links {
		l := link // Capture for closure
		list.AddItem(
			l.Name,
			fmt.Sprintf("Global: %s | Type: %s", l.Global, l.Type),
			0,
			func() {
				v.showResourceLinkForm(&l, false)
			},
		)
	}

	list.AddItem("+ Add Resource Link", "Link to a global resource", 'a', func() {
		link := jndi.NewResourceLink("jdbc/localDB", "jdbc/globalDB", string(jndi.ResourceTypeDataSource))
		v.showResourceLinkForm(link, true)
	})

	list.AddItem("Back", "Return to context menu", 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Resource Links ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("resource-links", list, true)
}

// showResourceLinkForm displays a form for editing a resource link
func (v *ContextView) showResourceLinkForm(link *jndi.ResourceLink, isNew bool) {
	l := *link

	form := tview.NewForm()

	form.AddInputField("Name (local)", l.Name, 40, nil, func(text string) {
		l.Name = text
	})

	form.AddInputField("Global (server.xml)", l.Global, 40, nil, func(text string) {
		l.Global = text
	})

	form.AddInputField("Type", l.Type, 50, nil, func(text string) {
		l.Type = text
	})

	form.AddButton(i18n.T("common.save.short"), func() {
		if l.Name == "" || l.Global == "" {
			v.setStatus("Error: Name and Global are required")
			return
		}

		if isNew {
			if err := v.configService.AddResourceLink(l); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Resource link added: " + l.Name)
		} else {
			if err := v.configService.UpdateResourceLink(link.Name, l); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Resource link updated: " + l.Name)
		}
		v.showResourceLinkList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteResourceLink(link.Name); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Resource link deleted: " + link.Name)
			v.showResourceLinkList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showResourceLinkList()
	})

	title := " Add Resource Link "
	if !isNew {
		title = " Edit Resource Link "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showResourceLinkList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("resource-link-form", form, true)
}

// showParameterList displays the context parameters list
func (v *ContextView) showParameterList() {
	params := v.configService.GetParameters()

	list := tview.NewList()

	for _, param := range params {
		p := param // Capture for closure
		list.AddItem(
			p.Name,
			fmt.Sprintf("Value: %s", p.Value),
			0,
			func() {
				v.showParameterForm(&p, false)
			},
		)
	}

	list.AddItem("+ Add Parameter", "Create a new context parameter", 'a', func() {
		param := jndi.NewContextParameter("paramName", "paramValue")
		v.showParameterForm(param, true)
	})

	list.AddItem("Back", "Return to context menu", 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Context Parameters ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("parameters", list, true)
}

// showParameterForm displays a form for editing a parameter
func (v *ContextView) showParameterForm(param *jndi.ContextParameter, isNew bool) {
	p := *param

	form := tview.NewForm()

	form.AddInputField("Name", p.Name, 40, nil, func(text string) {
		p.Name = text
	})

	form.AddInputField("Value", p.Value, 50, nil, func(text string) {
		p.Value = text
	})

	form.AddCheckbox("Override", p.Override, func(checked bool) {
		p.Override = checked
	})

	form.AddInputField("Description", p.Description, 50, nil, func(text string) {
		p.Description = text
	})

	form.AddButton(i18n.T("common.save.short"), func() {
		if p.Name == "" {
			v.setStatus("Error: Name is required")
			return
		}

		if isNew {
			if err := v.configService.AddParameter(p); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Parameter added: " + p.Name)
		} else {
			if err := v.configService.UpdateParameter(param.Name, p); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Parameter updated: " + p.Name)
		}
		v.showParameterList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			if err := v.configService.DeleteParameter(param.Name); err != nil {
				v.setStatus("Error: " + err.Error())
				return
			}
			v.setStatus("Parameter deleted: " + param.Name)
			v.showParameterList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showParameterList()
	})

	title := " Add Parameter "
	if !isNew {
		title = " Edit Parameter "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showParameterList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("parameter-form", form, true)
}

// showWatchedResourceList displays the watched resources list
func (v *ContextView) showWatchedResourceList() {
	resources := v.configService.GetWatchedResources()

	list := tview.NewList()

	for _, res := range resources {
		r := res // Capture for closure
		list.AddItem(
			r,
			"File or directory to watch for changes",
			0,
			func() {
				v.showWatchedResourceForm(r, false)
			},
		)
	}

	list.AddItem("+ Add Watched Resource", "Add a resource to watch", 'a', func() {
		v.showWatchedResourceForm("", true)
	})

	list.AddItem("Add Default Resources", "WEB-INF/web.xml and conf/web.xml", 'd', func() {
		v.configService.AddWatchedResource("WEB-INF/web.xml")
		v.configService.AddWatchedResource("${catalina.base}/conf/web.xml")
		v.setStatus("Default watched resources added")
		v.showWatchedResourceList()
	})

	list.AddItem("Back", "Return to context menu", 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Watched Resources ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("watched-resources", list, true)
}

// showWatchedResourceForm displays a form for editing a watched resource
func (v *ContextView) showWatchedResourceForm(resource string, isNew bool) {
	newResource := resource

	form := tview.NewForm()

	form.AddInputField("Resource Path", resource, 60, nil, func(text string) {
		newResource = text
	})

	form.AddTextView("Examples", `WEB-INF/web.xml
${catalina.base}/conf/web.xml
WEB-INF/classes
META-INF/context.xml`, 60, 4, true, false)

	form.AddButton(i18n.T("common.save.short"), func() {
		if newResource == "" {
			v.setStatus("Error: Resource path is required")
			return
		}

		if !isNew {
			v.configService.RemoveWatchedResource(resource)
		}
		v.configService.AddWatchedResource(newResource)
		v.setStatus("Watched resource saved: " + newResource)
		v.showWatchedResourceList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			v.configService.RemoveWatchedResource(resource)
			v.setStatus("Watched resource deleted")
			v.showWatchedResourceList()
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showWatchedResourceList()
	})

	title := " Add Watched Resource "
	if !isNew {
		title = " Edit Watched Resource "
	}
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showWatchedResourceList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("watched-resource-form", form, true)
}

// showManagerForm displays the session manager configuration form
func (v *ContextView) showManagerForm() {
	manager := v.configService.GetManager()
	if manager == nil {
		manager = jndi.NewContextManager()
	}

	m := *manager

	form := tview.NewForm()

	// Manager type
	managerTypes := []string{"StandardManager", "PersistentManager", "None (Remove Manager)"}
	managerIdx := 0
	if m.ClassName == jndi.ManagerPersistent {
		managerIdx = 1
	}
	form.AddDropDown("Manager Type", managerTypes, managerIdx, func(option string, index int) {
		if index == 0 {
			m.ClassName = jndi.ManagerStandard
		} else if index == 1 {
			m.ClassName = jndi.ManagerPersistent
		}
	})

	form.AddInputField("Max Active Sessions", strconv.Itoa(m.MaxActiveSessions), 10, nil, func(text string) {
		m.MaxActiveSessions, _ = strconv.Atoi(text)
	})

	form.AddInputField("Session ID Length", strconv.Itoa(m.SessionIdLength), 10, acceptNumber, func(text string) {
		m.SessionIdLength, _ = strconv.Atoi(text)
	})

	form.AddInputField("Max Inactive Interval (sec)", strconv.Itoa(m.MaxInactiveInterval), 10, acceptNumber, func(text string) {
		m.MaxInactiveInterval, _ = strconv.Atoi(text)
	})

	// PersistentManager specific
	saveOnRestart := m.SaveOnRestart == "true"
	form.AddCheckbox("Save On Restart (Persistent)", saveOnRestart, func(checked bool) {
		m.SaveOnRestart = BoolToString(checked)
	})

	// Store configuration
	storeTypes := []string{"FileStore", "JDBCStore"}
	storeIdx := 0
	if m.Store != nil && m.Store.ClassName == jndi.StoreJDBC {
		storeIdx = 1
	}
	form.AddDropDown("Session Store (Persistent)", storeTypes, storeIdx, func(option string, index int) {
		if m.Store == nil {
			m.Store = &jndi.SessionStore{}
		}
		if index == 0 {
			m.Store.ClassName = jndi.StoreFile
		} else {
			m.Store.ClassName = jndi.StoreJDBC
		}
	})

	storeDir := ""
	if m.Store != nil {
		storeDir = m.Store.Directory
	}
	form.AddInputField("Store Directory (FileStore)", storeDir, 40, nil, func(text string) {
		if m.Store == nil {
			m.Store = &jndi.SessionStore{ClassName: jndi.StoreFile}
		}
		m.Store.Directory = text
	})

	form.AddButton(i18n.T("common.save.short"), func() {
		// Check if user selected to remove manager
		dropdown := form.GetFormItem(0).(*tview.DropDown)
		_, option := dropdown.GetCurrentOption()
		if strings.Contains(option, "Remove") {
			v.configService.RemoveManager()
			v.setStatus("Session manager removed")
		} else {
			v.configService.SetManager(&m)
			v.setStatus("Session manager configured: " + m.ClassName)
		}
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMainMenu()
	})

	form.SetBorder(true).SetTitle(" Session Manager Configuration ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("manager-form", form, true)
}

// showCookieProcessorForm displays the cookie processor configuration form
func (v *ContextView) showCookieProcessorForm() {
	processor := v.configService.GetCookieProcessor()
	if processor == nil {
		processor = jndi.NewCookieProcessor()
	}

	p := *processor

	form := tview.NewForm()

	// Processor type
	processorTypes := []string{jndi.CookieProcessorRfc6265 + " (Default)", jndi.CookieProcessorLegacy, "None (Remove)"}
	processorIdx := 0
	if p.ClassName == jndi.CookieProcessorLegacy {
		processorIdx = 1
	}
	form.AddDropDown("Cookie Processor", processorTypes, processorIdx, func(option string, index int) {
		if index == 0 {
			p.ClassName = jndi.CookieProcessorRfc6265
		} else if index == 1 {
			p.ClassName = jndi.CookieProcessorLegacy
		}
	})

	// SameSite
	sameSiteIdx := 0
	for i, val := range jndi.SameSiteCookieValues {
		if val == p.SameSiteCookies {
			sameSiteIdx = i
			break
		}
	}
	form.AddDropDown("SameSite Cookies", jndi.SameSiteCookieValues, sameSiteIdx, func(option string, index int) {
		p.SameSiteCookies = option
	})

	form.AddButton(i18n.T("common.save.short"), func() {
		dropdown := form.GetFormItem(0).(*tview.DropDown)
		_, option := dropdown.GetCurrentOption()
		if strings.Contains(option, "Remove") {
			v.configService.RemoveCookieProcessor()
			v.setStatus("Cookie processor removed")
		} else {
			v.configService.SetCookieProcessor(&p)
			v.setStatus("Cookie processor configured")
		}
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMainMenu()
	})

	form.SetBorder(true).SetTitle(" Cookie Processor Configuration ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("cookie-form", form, true)
}

// showJarScannerForm displays the JAR scanner configuration form
func (v *ContextView) showJarScannerForm() {
	scanner := v.configService.GetJarScanner()
	if scanner == nil {
		scanner = jndi.NewJarScanner()
	}

	s := *scanner

	form := tview.NewForm()

	scanClassPath := s.ScanClassPath == "true"
	form.AddCheckbox("Scan ClassPath", scanClassPath, func(checked bool) {
		if checked {
			s.ScanClassPath = "true"
		} else {
			s.ScanClassPath = "false"
		}
	})

	scanManifest := s.ScanManifest == "true"
	form.AddCheckbox("Scan Manifest", scanManifest, func(checked bool) {
		if checked {
			s.ScanManifest = "true"
		} else {
			s.ScanManifest = "false"
		}
	})

	scanAllFiles := s.ScanAllFiles == "true"
	form.AddCheckbox("Scan All Files", scanAllFiles, func(checked bool) {
		if checked {
			s.ScanAllFiles = "true"
		} else {
			s.ScanAllFiles = "false"
		}
	})

	scanAllDirs := s.ScanAllDirectories == "true"
	form.AddCheckbox("Scan All Directories", scanAllDirs, func(checked bool) {
		if checked {
			s.ScanAllDirectories = "true"
		} else {
			s.ScanAllDirectories = "false"
		}
	})

	scanBootstrap := s.ScanBootstrapClassPath == "true"
	form.AddCheckbox("Scan Bootstrap ClassPath", scanBootstrap, func(checked bool) {
		if checked {
			s.ScanBootstrapClassPath = "true"
		} else {
			s.ScanBootstrapClassPath = "false"
		}
	})

	// JarScanFilter settings
	if s.JarScanFilter == nil {
		s.JarScanFilter = &jndi.JarScanFilter{}
	}

	form.AddInputField("TLD Skip Pattern", s.JarScanFilter.TldSkip, 60, nil, func(text string) {
		s.JarScanFilter.TldSkip = text
	})

	form.AddInputField("Pluggability Skip Pattern", s.JarScanFilter.PluggabilitySkip, 60, nil, func(text string) {
		s.JarScanFilter.PluggabilitySkip = text
	})

	form.AddButton(i18n.T("common.save.short"), func() {
		v.configService.SetJarScanner(&s)
		v.setStatus("JAR scanner configured")
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.remove"), func() {
		v.configService.RemoveJarScanner()
		v.setStatus("JAR scanner removed")
		v.showMainMenu()
	})

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMainMenu()
	})

	form.SetBorder(true).SetTitle(" JAR Scanner Configuration ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("jarscanner-form", form, true)
}

// saveConfiguration saves the context configuration
func (v *ContextView) saveConfiguration() {
	if err := v.configService.Save(); err != nil {
		v.setStatus("Error saving: " + err.Error())
		return
	}
	v.setStatus("Configuration saved to: " + v.configService.GetFilePath())
}

// setStatus updates the status bar
func (v *ContextView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(" " + message)
	}
}
