package views

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/jndi"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// JNDIView handles JNDI resource configuration
type JNDIView struct {
	app       *tview.Application
	pages     *tview.Pages
	mainPages *tview.Pages
	statusBar *tview.TextView
	onReturn  func()

	contextService *jndi.ContextService
	catalinaBase   string
}

// NewJNDIView creates a new JNDI view
func NewJNDIView(app *tview.Application, mainPages *tview.Pages, catalinaBase string, statusBar *tview.TextView, onReturn func()) *JNDIView {
	v := &JNDIView{
		app:            app,
		pages:          tview.NewPages(),
		mainPages:      mainPages,
		statusBar:      statusBar,
		onReturn:       onReturn,
		catalinaBase:   catalinaBase,
		contextService: jndi.NewContextService(catalinaBase),
	}
	return v
}

// Show displays the JNDI view
func (v *JNDIView) Show() {
	// Load context.xml
	if err := v.contextService.Load(); err != nil {
		v.showError("Failed to load context.xml", err)
		return
	}

	v.showMainMenu()
	v.mainPages.AddAndSwitchToPage("jndi", v.pages, true)
}

// showMainMenu shows the main JNDI menu
func (v *JNDIView) showMainMenu() {
	menu := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	// Count items for each category
	dsCount := len(v.contextService.GetDataSources())
	mailCount := len(v.contextService.GetMailSessions())
	envCount := len(v.contextService.GetEnvironments())
	linkCount := len(v.contextService.GetResourceLinks())

	menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("jndi.datasource")+"[::-] [yellow](%d)[-]", dsCount),
		i18n.T("jndi.datasource.desc"), 'd', func() {
			v.showDataSourceList()
		})

	menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("jndi.mail")+"[::-] [yellow](%d)[-]", mailCount),
		i18n.T("jndi.mail.desc"), 'm', func() {
			v.showMailSessionList()
		})

	menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("jndi.environment")+"[::-] [yellow](%d)[-]", envCount),
		i18n.T("jndi.environment.desc"), 'e', func() {
			v.showEnvironmentList()
		})

	menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("jndi.resourcelink")+"[::-] [yellow](%d)[-]", linkCount),
		i18n.T("jndi.resourcelink.desc"), 'l', func() {
			v.showResourceLinkList()
		})

	menu.AddItem("", "", 0, nil)
	menu.AddItem("[red]"+i18n.T("common.back")+"[-]", i18n.T("common.return"), 0, func() {
		v.onReturn()
	})

	// Update help panel when selection changes
	menu.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.jndi.datasource")
		case 1:
			helpPanel.SetHelpKey("help.jndi.mail")
		case 2:
			helpPanel.SetHelpKey("help.jndi.environment")
		case 3:
			helpPanel.SetHelpKey("help.jndi.resourcelink")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.jndi.datasource")

	menu.SetBorder(true).SetTitle(" " + i18n.T("jndi.title") + " ").SetBorderColor(tcell.ColorDarkCyan)
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

// showDataSourceList shows the list of DataSource resources
func (v *JNDIView) showDataSourceList() {
	list := tview.NewList().ShowSecondaryText(true)

	dataSources := v.contextService.GetDataSources()
	for _, ds := range dataSources {
		resource := ds // capture for closure
		info := fmt.Sprintf("%s - %s", resource.DriverClassName, resource.URL)
		if len(info) > 60 {
			info = info[:57] + "..."
		}
		list.AddItem(resource.Name, info, 0, func() {
			v.showDataSourceForm(&resource, false)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add DataSource[-]", "Create new JDBC DataSource", 'a', func() {
		newDS := jndi.NewDataSourceResource("jdbc/NewDataSource")
		v.showDataSourceForm(newDS, true)
	})
	list.AddItem("[yellow]► Quick Add (Template)[-]", "Create from database template", 't', func() {
		v.showDataSourceTemplates()
	})
	list.AddItem("[red]Back[-]", "Return to JNDI menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" DataSource Resources ").SetBorderColor(tcell.ColorGreen)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("datasource-list", list, true)
	v.app.SetFocus(list)
}

// showDataSourceTemplates shows database templates for quick DataSource creation
func (v *JNDIView) showDataSourceTemplates() {
	list := tview.NewList().ShowSecondaryText(true)

	// Sort database names for consistent display
	var dbNames []string
	for name := range jndi.CommonDrivers {
		dbNames = append(dbNames, name)
	}
	sort.Strings(dbNames)

	for _, dbName := range dbNames {
		name := dbName
		driver := jndi.CommonDrivers[name]
		list.AddItem(name, driver, 0, func() {
			ds := jndi.NewDataSourceResource("jdbc/MyDB")
			ds.DriverClassName = jndi.CommonDrivers[name]
			ds.URL = jndi.JDBCURLTemplates[name]
			ds.ValidationQuery = jndi.ValidationQueries[name]
			v.showDataSourceForm(ds, true)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[red]Back[-]", "Return to DataSource list", 0, func() {
		v.showDataSourceList()
	})

	list.SetBorder(true).SetTitle(" Select Database Type ").SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showDataSourceList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("datasource-templates", list, true)
	v.app.SetFocus(list)
}

// Known JDBC drivers for common databases
var jdbcDrivers = []string{
	"com.mysql.cj.jdbc.Driver",                     // MySQL 8.x
	"com.mysql.jdbc.Driver",                        // MySQL 5.x (legacy)
	"org.mariadb.jdbc.Driver",                      // MariaDB
	"org.postgresql.Driver",                        // PostgreSQL
	"oracle.jdbc.OracleDriver",                     // Oracle
	"com.microsoft.sqlserver.jdbc.SQLServerDriver", // SQL Server
	"org.h2.Driver",                                // H2
	"org.hsqldb.jdbc.JDBCDriver",                   // HSQLDB
	"org.sqlite.JDBC",                              // SQLite
	"org.apache.derby.jdbc.ClientDriver",           // Apache Derby (Client)
	"org.apache.derby.jdbc.EmbeddedDriver",         // Apache Derby (Embedded)
	"com.ibm.db2.jcc.DB2Driver",                    // IBM DB2
	"org.firebirdsql.jdbc.FBDriver",                // Firebird
}

// DataSource property help i18n keys
var dataSourceHelpKeys = map[string]string{
	"Name (JNDI)":      "help.ds.name",
	"Auth":             "help.ds.auth",
	"Factory":          "help.ds.factory",
	"Driver Class":     "help.ds.driver",
	"URL":              "help.ds.url",
	"Username":         "help.ds.username",
	"Password":         "help.ds.password",
	"Initial Size":     "help.ds.initialsize",
	"Max Total":        "help.ds.maxtotal",
	"Max Idle":         "help.ds.maxidle",
	"Min Idle":         "help.ds.minidle",
	"Max Wait (ms)":    "help.ds.maxwait",
	"Validation Query": "help.ds.validationquery",
	"Test On Borrow":   "help.ds.testonborrow",
	"Test While Idle":  "help.ds.testwhileidle",
}

// Mail Session property help i18n keys
var mailSessionHelpKeys = map[string]string{
	"Name (JNDI)": "help.mail.name",
	"Auth":        "help.mail.auth",
	"SMTP Host":   "help.mail.host",
	"SMTP Port":   "help.mail.port",
	"SMTP User":   "help.mail.user",
	"Protocol":    "help.mail.protocol",
	"SMTP Auth":   "help.mail.smtpauth",
	"StartTLS":    "help.mail.starttls",
	"Debug Mode":  "help.mail.debug",
}

// Environment property help i18n keys
var environmentHelpKeys = map[string]string{
	"Name (JNDI)": "help.env.name",
	"Type":        "help.env.type",
	"Value":       "help.env.value",
	"Override":    "help.env.override",
	"Description": "help.env.description",
}

// ResourceLink property help i18n keys
var resourceLinkHelpKeys = map[string]string{
	"Local Name":  "help.reslink.name",
	"Global Name": "help.reslink.global",
	"Type":        "help.reslink.type",
}

// showDataSourceForm shows the DataSource edit form
func (v *JNDIView) showDataSourceForm(resource *jndi.Resource, isNew bool) {
	form := tview.NewForm()
	preview := NewPreviewPanel()
	formReady := false

	// Help panel on the right
	helpPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	helpPanel.SetBorder(true).SetTitle(" " + i18n.T("help.title") + " ").SetBorderColor(tcell.ColorBlue)

	// Function to update help text
	updateHelp := func(label string) {
		if key, ok := dataSourceHelpKeys[label]; ok {
			helpPanel.SetText(i18n.T(key))
		} else {
			helpPanel.SetText(i18n.T("help.default"))
		}
	}

	// updatePreview reads form fields and generates XML preview
	updatePreview := func() {
		if !formReady {
			return
		}
		tempRes := jndi.Resource{
			Name:            GetFormText(form, "Name (JNDI)"),
			Type:            resource.Type,
			Auth:            GetFormDropDownText(form, "Auth"),
			Factory:         GetFormText(form, "Factory"),
			DriverClassName: GetFormDropDownText(form, "Driver Class"),
			URL:             GetFormText(form, "URL"),
			Username:        GetFormText(form, "Username"),
			Password:        GetFormText(form, "Password"),
			InitialSize:     GetFormInt(form, "Initial Size"),
			MaxTotal:        GetFormInt(form, "Max Total"),
			MaxIdle:         GetFormInt(form, "Max Idle"),
			MinIdle:         GetFormInt(form, "Min Idle"),
			MaxWaitMillis:   GetFormInt(form, "Max Wait (ms)"),
			ValidationQuery: GetFormText(form, "Validation Query"),
			TestOnBorrow:    GetFormBool(form, "Test On Borrow"),
			TestWhileIdle:   GetFormBool(form, "Test While Idle"),
		}
		preview.SetXMLPreview(GenerateJNDIResourceXML(&tempRes))
	}

	// Track focused field for help panel
	lastFocusedLabel := ""
	form.SetFocusFunc(func() {
		idx, _ := form.GetFocusedItemIndex()
		if idx >= 0 && idx < form.GetFormItemCount() {
			label := form.GetFormItem(idx).GetLabel()
			if label != lastFocusedLabel {
				lastFocusedLabel = label
				updateHelp(label)
			}
		}
	})

	// Basic settings
	form.AddInputField("Name (JNDI)", resource.Name, 40, nil, func(text string) { updatePreview() })
	form.AddDropDown("Auth", []string{"Container", "Application"}, getDropDownIndex([]string{"Container", "Application"}, resource.Auth), func(option string, index int) { updatePreview() })
	form.AddInputField("Factory", resource.Factory, 50, nil, func(text string) { updatePreview() })

	// Connection settings - Driver Class as DropDown
	driverIndex := getDropDownIndex(jdbcDrivers, resource.DriverClassName)
	if driverIndex < 0 && resource.DriverClassName != "" {
		// If current driver is not in the list, add it
		jdbcDrivers = append([]string{resource.DriverClassName}, jdbcDrivers...)
		driverIndex = 0
	}
	form.AddDropDown("Driver Class", jdbcDrivers, driverIndex, func(option string, index int) { updatePreview() })
	form.AddInputField("URL", resource.URL, 60, nil, func(text string) { updatePreview() })
	form.AddInputField("Username", resource.Username, 30, nil, func(text string) { updatePreview() })
	form.AddPasswordField("Password", resource.Password, 30, '*', func(text string) { updatePreview() })

	// Pool settings
	form.AddInputField("Initial Size", strconv.Itoa(resource.InitialSize), 10, acceptNumber, func(text string) { updatePreview() })
	form.AddInputField("Max Total", strconv.Itoa(resource.MaxTotal), 10, acceptNumber, func(text string) { updatePreview() })
	form.AddInputField("Max Idle", strconv.Itoa(resource.MaxIdle), 10, acceptNumber, func(text string) { updatePreview() })
	form.AddInputField("Min Idle", strconv.Itoa(resource.MinIdle), 10, acceptNumber, func(text string) { updatePreview() })
	form.AddInputField("Max Wait (ms)", strconv.Itoa(resource.MaxWaitMillis), 10, acceptNumber, func(text string) { updatePreview() })

	// Validation settings
	form.AddInputField("Validation Query", resource.ValidationQuery, 50, nil, func(text string) { updatePreview() })
	form.AddCheckbox("Test On Borrow", resource.TestOnBorrow, func(checked bool) { updatePreview() })
	form.AddCheckbox("Test While Idle", resource.TestWhileIdle, func(checked bool) { updatePreview() })

	form.AddButton(i18n.T("common.save.short"), func() {
		resource.Name = GetFormText(form, "Name (JNDI)")
		resource.Auth = GetFormDropDownText(form, "Auth")
		resource.Factory = GetFormText(form, "Factory")
		resource.DriverClassName = GetFormDropDownText(form, "Driver Class")
		resource.URL = GetFormText(form, "URL")
		resource.Username = GetFormText(form, "Username")
		resource.Password = GetFormText(form, "Password")
		resource.InitialSize = GetFormInt(form, "Initial Size")
		resource.MaxTotal = GetFormInt(form, "Max Total")
		resource.MaxIdle = GetFormInt(form, "Max Idle")
		resource.MinIdle = GetFormInt(form, "Min Idle")
		resource.MaxWaitMillis = GetFormInt(form, "Max Wait (ms)")
		resource.ValidationQuery = GetFormText(form, "Validation Query")
		resource.TestOnBorrow = GetFormBool(form, "Test On Borrow")
		resource.TestWhileIdle = GetFormBool(form, "Test While Idle")

		var err error
		if isNew {
			err = v.contextService.AddResource(*resource)
		} else {
			err = v.contextService.UpdateResource(resource.Name, *resource)
		}

		if err != nil {
			v.setStatus(fmt.Sprintf("[red]Error: %v[-]", err))
			return
		}

		if err := v.contextService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}

		v.setStatus("[green]DataSource saved successfully[-]")
		v.showDataSourceList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			v.confirmDelete("DataSource", resource.Name, func() {
				if err := v.contextService.DeleteResource(resource.Name); err != nil {
					v.setStatus(fmt.Sprintf("[red]Error: %v[-]", err))
					return
				}
				if err := v.contextService.Save(); err != nil {
					v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					return
				}
				v.setStatus("[green]DataSource deleted[-]")
				v.showDataSourceList()
			})
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showDataSourceList()
	})

	// Initial preview
	formReady = true
	updatePreview()
	updateHelp("Name (JNDI)")

	title := " Edit DataSource "
	if isNew {
		title = " Add DataSource "
	}
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorGreen)

	// Handle key events and update help on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showDataSourceList()
			return nil
		}
		// Update help after navigation
		go func() {
			v.app.QueueUpdateDraw(func() {
				idx, _ := form.GetFocusedItemIndex()
				if idx >= 0 && idx < form.GetFormItemCount() {
					label := form.GetFormItem(idx).GetLabel()
					if label != lastFocusedLabel {
						lastFocusedLabel = label
						updateHelp(label)
					}
				}
			})
		}()
		return event
	})

	// Create layout: left side (form + preview), right side (help)
	leftPane := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 2, true).
		AddItem(helpPanel, 40, 0, false)

	v.pages.AddAndSwitchToPage("datasource-form", layout, true)
	v.app.SetFocus(form)
}

// showMailSessionList shows the list of Mail Session resources
func (v *JNDIView) showMailSessionList() {
	list := tview.NewList().ShowSecondaryText(true)

	mailSessions := v.contextService.GetMailSessions()
	for _, ms := range mailSessions {
		resource := ms
		info := fmt.Sprintf("%s:%s", resource.MailSmtpHost, resource.MailSmtpPort)
		list.AddItem(resource.Name, info, 0, func() {
			v.showMailSessionForm(&resource, false)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add Mail Session[-]", "Create new Mail Session", 'a', func() {
		newMS := jndi.NewMailSessionResource("mail/Session")
		v.showMailSessionForm(newMS, true)
	})
	list.AddItem("[red]Back[-]", "Return to JNDI menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Mail Session Resources ").SetBorderColor(tcell.ColorBlue)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("mailsession-list", list, true)
	v.app.SetFocus(list)
}

// showMailSessionForm shows the Mail Session edit form
func (v *JNDIView) showMailSessionForm(resource *jndi.Resource, isNew bool) {
	form := tview.NewForm()
	preview := NewPreviewPanel()
	formReady := false

	// Help panel on the right
	helpPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	helpPanel.SetBorder(true).SetTitle(" " + i18n.T("help.title") + " ").SetBorderColor(tcell.ColorBlue)

	// Function to update help text
	updateHelp := func(label string) {
		if key, ok := mailSessionHelpKeys[label]; ok {
			helpPanel.SetText(i18n.T(key))
		} else {
			helpPanel.SetText(i18n.T("help.default"))
		}
	}

	// updatePreview reads form fields and generates XML preview
	updatePreview := func() {
		if !formReady {
			return
		}
		tempRes := jndi.Resource{
			Name:                  GetFormText(form, "Name (JNDI)"),
			Type:                  resource.Type,
			Auth:                  GetFormDropDownText(form, "Auth"),
			MailSmtpHost:          GetFormText(form, "SMTP Host"),
			MailSmtpPort:          GetFormText(form, "SMTP Port"),
			MailSmtpUser:          GetFormText(form, "SMTP User"),
			MailTransportProtocol: GetFormDropDownText(form, "Protocol"),
			MailSmtpAuth:          BoolToString(GetFormBool(form, "SMTP Auth")),
			MailSmtpStartTLS:      BoolToString(GetFormBool(form, "StartTLS")),
			MailDebug:             BoolToString(GetFormBool(form, "Debug Mode")),
		}
		preview.SetXMLPreview(GenerateJNDIResourceXML(&tempRes))
	}

	// Track focused field for help panel
	lastFocusedLabel := ""
	form.SetFocusFunc(func() {
		idx, _ := form.GetFocusedItemIndex()
		if idx >= 0 && idx < form.GetFormItemCount() {
			label := form.GetFormItem(idx).GetLabel()
			if label != lastFocusedLabel {
				lastFocusedLabel = label
				updateHelp(label)
			}
		}
	})

	// Basic settings
	form.AddInputField("Name (JNDI)", resource.Name, 40, nil, func(text string) { updatePreview() })
	form.AddDropDown("Auth", []string{"Container", "Application"}, getDropDownIndex([]string{"Container", "Application"}, resource.Auth), func(option string, index int) { updatePreview() })

	// SMTP settings
	form.AddInputField("SMTP Host", resource.MailSmtpHost, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("SMTP Port", resource.MailSmtpPort, 10, nil, func(text string) { updatePreview() })
	form.AddInputField("SMTP User", resource.MailSmtpUser, 30, nil, func(text string) { updatePreview() })
	form.AddDropDown("Protocol", []string{"smtp", "smtps"}, getDropDownIndex([]string{"smtp", "smtps"}, resource.MailTransportProtocol), func(option string, index int) { updatePreview() })

	// Authentication and TLS
	authEnabled := resource.MailSmtpAuth == "true"
	form.AddCheckbox("SMTP Auth", authEnabled, func(checked bool) { updatePreview() })
	tlsEnabled := resource.MailSmtpStartTLS == "true"
	form.AddCheckbox("StartTLS", tlsEnabled, func(checked bool) { updatePreview() })
	debugEnabled := resource.MailDebug == "true"
	form.AddCheckbox("Debug Mode", debugEnabled, func(checked bool) { updatePreview() })

	form.AddButton(i18n.T("common.save.short"), func() {
		resource.Name = GetFormText(form, "Name (JNDI)")
		resource.Auth = GetFormDropDownText(form, "Auth")
		resource.MailSmtpHost = GetFormText(form, "SMTP Host")
		resource.MailSmtpPort = GetFormText(form, "SMTP Port")
		resource.MailSmtpUser = GetFormText(form, "SMTP User")
		resource.MailTransportProtocol = GetFormDropDownText(form, "Protocol")
		resource.MailSmtpAuth = BoolToString(GetFormBool(form, "SMTP Auth"))
		resource.MailSmtpStartTLS = BoolToString(GetFormBool(form, "StartTLS"))
		resource.MailDebug = BoolToString(GetFormBool(form, "Debug Mode"))

		var err error
		if isNew {
			err = v.contextService.AddResource(*resource)
		} else {
			err = v.contextService.UpdateResource(resource.Name, *resource)
		}

		if err != nil {
			v.setStatus(fmt.Sprintf("[red]Error: %v[-]", err))
			return
		}

		if err := v.contextService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}

		v.setStatus("[green]Mail Session saved successfully[-]")
		v.showMailSessionList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			v.confirmDelete("Mail Session", resource.Name, func() {
				if err := v.contextService.DeleteResource(resource.Name); err != nil {
					v.setStatus(fmt.Sprintf("[red]Error: %v[-]", err))
					return
				}
				if err := v.contextService.Save(); err != nil {
					v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					return
				}
				v.setStatus("[green]Mail Session deleted[-]")
				v.showMailSessionList()
			})
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showMailSessionList()
	})

	// Initial preview and help
	formReady = true
	updatePreview()
	updateHelp("Name (JNDI)")

	title := " Edit Mail Session "
	if isNew {
		title = " Add Mail Session "
	}
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorBlue)

	// Handle key events and update help on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMailSessionList()
			return nil
		}
		// Update help after navigation
		go func() {
			v.app.QueueUpdateDraw(func() {
				idx, _ := form.GetFocusedItemIndex()
				if idx >= 0 && idx < form.GetFormItemCount() {
					label := form.GetFormItem(idx).GetLabel()
					if label != lastFocusedLabel {
						lastFocusedLabel = label
						updateHelp(label)
					}
				}
			})
		}()
		return event
	})

	// Create layout: left side (form + preview), right side (help)
	leftPane := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 2, true).
		AddItem(helpPanel, 40, 0, false)

	v.pages.AddAndSwitchToPage("mailsession-form", layout, true)
	v.app.SetFocus(form)
}

// showEnvironmentList shows the list of Environment entries
func (v *JNDIView) showEnvironmentList() {
	list := tview.NewList().ShowSecondaryText(true)

	environments := v.contextService.GetEnvironments()
	for _, env := range environments {
		entry := env
		info := fmt.Sprintf("%s = %s", entry.Type, entry.Value)
		if len(info) > 50 {
			info = info[:47] + "..."
		}
		list.AddItem(entry.Name, info, 0, func() {
			v.showEnvironmentForm(&entry, false)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add Environment Entry[-]", "Create new environment variable", 'a', func() {
		newEnv := jndi.NewEnvironment("myapp/config", "", "java.lang.String")
		v.showEnvironmentForm(newEnv, true)
	})
	list.AddItem("[red]Back[-]", "Return to JNDI menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Environment Entries ").SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("environment-list", list, true)
	v.app.SetFocus(list)
}

// showEnvironmentForm shows the Environment edit form
func (v *JNDIView) showEnvironmentForm(env *jndi.Environment, isNew bool) {
	form := tview.NewForm()
	preview := NewPreviewPanel()
	formReady := false

	// Help panel on the right
	helpPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	helpPanel.SetBorder(true).SetTitle(" " + i18n.T("help.title") + " ").SetBorderColor(tcell.ColorBlue)

	// Function to update help text
	updateHelp := func(label string) {
		if key, ok := environmentHelpKeys[label]; ok {
			helpPanel.SetText(i18n.T(key))
		} else {
			helpPanel.SetText(i18n.T("help.default"))
		}
	}

	// updatePreview reads form fields and generates XML preview
	updatePreview := func() {
		if !formReady {
			return
		}
		tempEnv := jndi.Environment{
			Name:        GetFormText(form, "Name (JNDI)"),
			Value:       GetFormText(form, "Value"),
			Type:        GetFormDropDownText(form, "Type"),
			Override:    GetFormBool(form, "Override"),
			Description: GetFormText(form, "Description"),
		}
		preview.SetXMLPreview(GenerateEnvironmentXML(&tempEnv))
	}

	// Track focused field for help panel
	lastFocusedLabel := ""
	form.SetFocusFunc(func() {
		idx, _ := form.GetFocusedItemIndex()
		if idx >= 0 && idx < form.GetFormItemCount() {
			label := form.GetFormItem(idx).GetLabel()
			if label != lastFocusedLabel {
				lastFocusedLabel = label
				updateHelp(label)
			}
		}
	})

	form.AddInputField("Name (JNDI)", env.Name, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("Value", env.Value, 50, nil, func(text string) { updatePreview() })
	form.AddDropDown("Type", jndi.EnvironmentTypes, getDropDownIndex(jndi.EnvironmentTypes, env.Type), func(option string, index int) { updatePreview() })
	form.AddCheckbox("Override", env.Override, func(checked bool) { updatePreview() })
	form.AddInputField("Description", env.Description, 50, nil, func(text string) { updatePreview() })

	form.AddButton(i18n.T("common.save.short"), func() {
		env.Name = GetFormText(form, "Name (JNDI)")
		env.Value = GetFormText(form, "Value")
		env.Type = GetFormDropDownText(form, "Type")
		env.Override = GetFormBool(form, "Override")
		env.Description = GetFormText(form, "Description")

		var err error
		if isNew {
			err = v.contextService.AddEnvironment(*env)
		} else {
			err = v.contextService.UpdateEnvironment(env.Name, *env)
		}

		if err != nil {
			v.setStatus(fmt.Sprintf("[red]Error: %v[-]", err))
			return
		}

		if err := v.contextService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}

		v.setStatus("[green]Environment entry saved successfully[-]")
		v.showEnvironmentList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			v.confirmDelete("Environment", env.Name, func() {
				if err := v.contextService.DeleteEnvironment(env.Name); err != nil {
					v.setStatus(fmt.Sprintf("[red]Error: %v[-]", err))
					return
				}
				if err := v.contextService.Save(); err != nil {
					v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					return
				}
				v.setStatus("[green]Environment entry deleted[-]")
				v.showEnvironmentList()
			})
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showEnvironmentList()
	})

	// Initial preview and help
	formReady = true
	updatePreview()
	updateHelp("Name (JNDI)")

	title := " Edit Environment Entry "
	if isNew {
		title = " Add Environment Entry "
	}
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorYellow)

	// Handle key events and update help on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showEnvironmentList()
			return nil
		}
		// Update help after navigation
		go func() {
			v.app.QueueUpdateDraw(func() {
				idx, _ := form.GetFocusedItemIndex()
				if idx >= 0 && idx < form.GetFormItemCount() {
					label := form.GetFormItem(idx).GetLabel()
					if label != lastFocusedLabel {
						lastFocusedLabel = label
						updateHelp(label)
					}
				}
			})
		}()
		return event
	})

	// Create layout: left side (form + preview), right side (help)
	leftPane := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 2, true).
		AddItem(helpPanel, 40, 0, false)

	v.pages.AddAndSwitchToPage("environment-form", layout, true)
	v.app.SetFocus(form)
}

// showResourceLinkList shows the list of Resource Links
func (v *JNDIView) showResourceLinkList() {
	list := tview.NewList().ShowSecondaryText(true)

	links := v.contextService.GetResourceLinks()
	for _, link := range links {
		l := link
		info := fmt.Sprintf("→ %s (%s)", l.Global, l.Type)
		list.AddItem(l.Name, info, 0, func() {
			v.showResourceLinkForm(&l, false)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add Resource Link[-]", "Create new link to global resource", 'a', func() {
		newLink := jndi.NewResourceLink("jdbc/LocalDB", "jdbc/GlobalDB", string(jndi.ResourceTypeDataSource))
		v.showResourceLinkForm(newLink, true)
	})
	list.AddItem("[red]Back[-]", "Return to JNDI menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Resource Links ").SetBorderColor(tcell.ColorPurple)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("resourcelink-list", list, true)
	v.app.SetFocus(list)
}

// showResourceLinkForm shows the Resource Link edit form
func (v *JNDIView) showResourceLinkForm(link *jndi.ResourceLink, isNew bool) {
	form := tview.NewForm()
	preview := NewPreviewPanel()
	formReady := false

	// Help panel on the right
	helpPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	helpPanel.SetBorder(true).SetTitle(" " + i18n.T("help.title") + " ").SetBorderColor(tcell.ColorBlue)

	// Function to update help text
	updateHelp := func(label string) {
		if key, ok := resourceLinkHelpKeys[label]; ok {
			helpPanel.SetText(i18n.T(key))
		} else {
			helpPanel.SetText(i18n.T("help.default"))
		}
	}

	// updatePreview reads form fields and generates XML preview
	updatePreview := func() {
		if !formReady {
			return
		}
		tempLink := jndi.ResourceLink{
			Name:   GetFormText(form, "Local Name"),
			Global: GetFormText(form, "Global Name"),
			Type:   GetFormDropDownText(form, "Type"),
		}
		preview.SetXMLPreview(GenerateResourceLinkXML(&tempLink))
	}

	// Track focused field for help panel
	lastFocusedLabel := ""
	form.SetFocusFunc(func() {
		idx, _ := form.GetFocusedItemIndex()
		if idx >= 0 && idx < form.GetFormItemCount() {
			label := form.GetFormItem(idx).GetLabel()
			if label != lastFocusedLabel {
				lastFocusedLabel = label
				updateHelp(label)
			}
		}
	})

	form.AddInputField("Local Name", link.Name, 40, nil, func(text string) { updatePreview() })
	form.AddInputField("Global Name", link.Global, 40, nil, func(text string) { updatePreview() })

	types := []string{
		string(jndi.ResourceTypeDataSource),
		string(jndi.ResourceTypeMailSession),
		string(jndi.ResourceTypeUserDatabase),
		string(jndi.ResourceTypeBean),
	}
	form.AddDropDown("Type", types, getDropDownIndex(types, link.Type), func(option string, index int) { updatePreview() })

	form.AddButton(i18n.T("common.save.short"), func() {
		link.Name = GetFormText(form, "Local Name")
		link.Global = GetFormText(form, "Global Name")
		link.Type = GetFormDropDownText(form, "Type")

		var err error
		if isNew {
			err = v.contextService.AddResourceLink(*link)
		} else {
			err = v.contextService.UpdateResourceLink(link.Name, *link)
		}

		if err != nil {
			v.setStatus(fmt.Sprintf("[red]Error: %v[-]", err))
			return
		}

		if err := v.contextService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}

		v.setStatus("[green]Resource Link saved successfully[-]")
		v.showResourceLinkList()
	})

	if !isNew {
		form.AddButton(i18n.T("common.delete"), func() {
			v.confirmDelete("Resource Link", link.Name, func() {
				if err := v.contextService.DeleteResourceLink(link.Name); err != nil {
					v.setStatus(fmt.Sprintf("[red]Error: %v[-]", err))
					return
				}
				if err := v.contextService.Save(); err != nil {
					v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					return
				}
				v.setStatus("[green]Resource Link deleted[-]")
				v.showResourceLinkList()
			})
		})
	}

	form.AddButton(i18n.T("common.cancel"), func() {
		v.showResourceLinkList()
	})

	// Initial preview and help
	formReady = true
	updatePreview()
	updateHelp("Local Name")

	title := " Edit Resource Link "
	if isNew {
		title = " Add Resource Link "
	}
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorPurple)

	// Handle key events and update help on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showResourceLinkList()
			return nil
		}
		// Update help after navigation
		go func() {
			v.app.QueueUpdateDraw(func() {
				idx, _ := form.GetFocusedItemIndex()
				if idx >= 0 && idx < form.GetFormItemCount() {
					label := form.GetFormItem(idx).GetLabel()
					if label != lastFocusedLabel {
						lastFocusedLabel = label
						updateHelp(label)
					}
				}
			})
		}()
		return event
	})

	// Create layout: left side (form + preview), right side (help)
	leftPane := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 2, true).
		AddItem(helpPanel, 40, 0, false)

	v.pages.AddAndSwitchToPage("resourcelink-form", layout, true)
	v.app.SetFocus(form)
}

// confirmDelete shows a confirmation dialog
func (v *JNDIView) confirmDelete(itemType, name string, onConfirm func()) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Delete %s '%s'?", itemType, name)).
		AddButtons([]string{"Delete", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Delete" {
				onConfirm()
			} else {
				v.pages.SwitchToPage(v.getCurrentListPage(itemType))
			}
		})
	v.pages.AddAndSwitchToPage("confirm-delete", modal, true)
}

// getCurrentListPage returns the list page name for the item type
func (v *JNDIView) getCurrentListPage(itemType string) string {
	switch itemType {
	case "DataSource":
		return "datasource-list"
	case "Mail Session":
		return "mailsession-list"
	case "Environment":
		return "environment-list"
	case "Resource Link":
		return "resourcelink-list"
	default:
		return "menu"
	}
}

// showError shows an error message
func (v *JNDIView) showError(title string, err error) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("%s:\n%v", title, err)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			v.onReturn()
		})
	modal.SetBorder(true).SetBorderColor(tcell.ColorRed)
	v.mainPages.AddAndSwitchToPage("jndi-error", modal, true)
}

// setStatus updates the status bar
func (v *JNDIView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(fmt.Sprintf(" %s", message))
	}
}

// getDropDownIndex returns the index of a value in a slice
func getDropDownIndex(options []string, value string) int {
	for i, opt := range options {
		if opt == value {
			return i
		}
	}
	return 0
}

// getFormFieldValue is a helper to get text from an InputField
func getFormFieldValue(form *tview.Form, label string) string {
	item := form.GetFormItemByLabel(label)
	if field, ok := item.(*tview.InputField); ok {
		return field.GetText()
	}
	return ""
}
