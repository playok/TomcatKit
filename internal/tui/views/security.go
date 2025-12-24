package views

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/realm"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// Help key arrays for realm forms
var userDatabaseRealmHelpKeys = []string{
	"help.realm.userdatabase.resource", // 0: Resource Name
}

var dataSourceRealmHelpKeys = []string{
	"help.realm.datasource.name",       // 0: DataSource Name
	"help.realm.datasource.usertable",  // 1: User Table
	"help.realm.datasource.usernameCol", // 2: User Name Column
	"help.realm.datasource.passwordCol", // 3: User Credential Column
	"help.realm.datasource.roletable",  // 4: User Role Table
	"help.realm.datasource.rolenameCol", // 5: Role Name Column
}

var jndiRealmHelpKeys = []string{
	"help.realm.jndi.connectionURL",  // 0: Connection URL
	"help.realm.jndi.connectionName", // 1: Connection Name
	"help.realm.jndi.connectionPwd",  // 2: Connection Password
	"help.realm.jndi.userpattern",    // 3: User Pattern
	"help.realm.jndi.userbase",       // 4: User Base
	"help.realm.jndi.usersearch",     // 5: User Search
	"help.realm.jndi.rolebase",       // 6: Role Base
	"help.realm.jndi.rolename",       // 7: Role Name
	"help.realm.jndi.rolesearch",     // 8: Role Search
}

var genericRealmHelpKeys = []string{
	"help.realm.classname", // 0: Class Name
}

var userFormHelpKeys = []string{
	"help.user.username", // 0: Username
	"help.user.password", // 1: Password
	"help.user.roles",    // 2: Roles
}

var roleFormHelpKeys = []string{
	"help.role.name",        // 0: Role Name
	"help.role.description", // 1: Description
}

// SecurityView handles security and realm configuration UI
type SecurityView struct {
	app           *tview.Application
	pages         *tview.Pages
	configService *server.ConfigService
	usersService  *realm.UsersService
	catalinaBase  string
	onBack        func()
	statusBar     *tview.TextView
}

// NewSecurityView creates a new security configuration view
func NewSecurityView(app *tview.Application, pages *tview.Pages, configService *server.ConfigService, catalinaBase string, statusBar *tview.TextView, onBack func()) *SecurityView {
	return &SecurityView{
		app:           app,
		pages:         pages,
		configService: configService,
		usersService:  realm.NewUsersService(catalinaBase),
		catalinaBase:  catalinaBase,
		onBack:        onBack,
		statusBar:     statusBar,
	}
}

// Show displays the security configuration menu
func (v *SecurityView) Show() {
	list := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	// Get current realm configuration
	srv := v.configService.GetServer()
	realmInfo := "Not configured"
	if srv != nil && len(srv.Services) > 0 {
		if srv.Services[0].Engine.Realm != nil {
			realmInfo = realm.GetShortRealmName(srv.Services[0].Engine.Realm.ClassName)
		}
	}

	list.AddItem(
		fmt.Sprintf("[::b]"+i18n.T("security.realm")+"[::-] [yellow](%s)[-]", realmInfo),
		i18n.T("security.realm.desc"),
		'r',
		func() { v.showRealmConfig() },
	)

	list.AddItem(
		"[::b]"+i18n.T("security.users")+"[::-]",
		i18n.T("security.users.desc"),
		'u',
		func() { v.showUsersConfig() },
	)

	list.AddItem(
		"[::b]"+i18n.T("security.credential")+"[::-]",
		i18n.T("security.credential.desc"),
		'c',
		func() { v.showCredentialHandler() },
	)

	list.AddItem("[-:-:-] [white:red] "+i18n.T("common.back")+" [-:-:-]", i18n.T("common.return"), 'b', v.onBack)

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.security.realm")
		case 1:
			helpPanel.SetHelpKey("help.security.users")
		case 2:
			helpPanel.SetHelpKey("help.security.credential")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.security.realm")

	list.SetBorder(true).SetTitle(" " + i18n.T("security.title") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Create flex layout with list and help panel
	flex := tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("security-config", flex, true)
	v.app.SetFocus(list)
}

// showRealmConfig shows realm configuration
func (v *SecurityView) showRealmConfig() {
	srv := v.configService.GetServer()
	if srv == nil || len(srv.Services) == 0 {
		v.showError("No services configured")
		return
	}

	// Get current realm from Engine level
	currentRealm := srv.Services[0].Engine.Realm

	list := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	menuIndex := 0
	if currentRealm != nil {
		realmName := realm.GetShortRealmName(currentRealm.ClassName)
		realmDesc := realm.GetRealmDescription(currentRealm.ClassName)

		list.AddItem(
			fmt.Sprintf("[white:green]"+i18n.T("security.realm.current")+": %s[-]", realmName),
			realmDesc,
			0,
			func() { v.showRealmDetail(currentRealm, 0) },
		)
		menuIndex++

		// Show nested realms if CombinedRealm or LockOutRealm
		if len(currentRealm.NestedRealms) > 0 {
			list.AddItem("[::b]"+i18n.T("security.realm.nested")+":[::-]", "─────────────────────", 0, nil)
			menuIndex++
			for i, nested := range currentRealm.NestedRealms {
				idx := i
				nestedName := realm.GetShortRealmName(nested.ClassName)
				list.AddItem(
					fmt.Sprintf("  └─ %s", nestedName),
					realm.GetRealmDescription(nested.ClassName),
					0,
					func() { v.showNestedRealmDetail(idx) },
				)
				menuIndex++
			}
		}

		list.AddItem("", "", 0, nil)
		menuIndex++
	}

	setRealmIndex := menuIndex
	list.AddItem("[green]+ "+i18n.T("security.realm.set")+"[-]", i18n.T("security.realm.set.desc"), 's', func() {
		v.showRealmTypeSelector()
	})
	menuIndex++

	removeRealmIndex := -1
	if currentRealm != nil {
		removeRealmIndex = menuIndex
		list.AddItem("[white:red]"+i18n.T("security.realm.remove")+"[-]", i18n.T("security.realm.remove.desc"), 'd', func() {
			v.showConfirm(i18n.T("security.realm.remove"), i18n.T("security.realm.remove.confirm"), func(confirmed bool) {
				if confirmed {
					srv.Services[0].Engine.Realm = nil
					v.configService.UpdateService(0, srv.Services[0])
					if err := v.configService.Save(); err != nil {
						v.showError(fmt.Sprintf("Failed to save: %v", err))
						return
					}
					v.setStatus("[green]" + i18n.T("security.realm.removed") + "[-]")
				}
				v.showRealmConfig()
			})
		})
		menuIndex++
	}

	list.AddItem("[-:-:-] [white:red] "+i18n.T("common.back")+" [-:-:-]", i18n.T("common.return"), 'b', func() {
		v.Show()
	})

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index == 0 && currentRealm != nil {
			helpPanel.SetHelpKey("help.security.realm.current")
		} else if index == setRealmIndex {
			helpPanel.SetHelpKey("help.security.realm.set")
		} else if removeRealmIndex >= 0 && index == removeRealmIndex {
			helpPanel.SetHelpKey("help.security.realm.remove")
		} else {
			helpPanel.SetText("")
		}
	})

	// Initialize help
	if currentRealm != nil {
		helpPanel.SetHelpKey("help.security.realm.current")
	} else {
		helpPanel.SetHelpKey("help.security.realm.set")
	}

	list.SetBorder(true).SetTitle(" " + i18n.T("security.realm.config") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Create flex layout with list and help panel
	flex := tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("realm-config", flex, true)
	v.app.SetFocus(list)
}

// showRealmTypeSelector shows realm type selection
func (v *SecurityView) showRealmTypeSelector() {
	list := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	realmTypes := realm.AvailableRealmTypes()
	for i, className := range realmTypes {
		cn := className
		name := realm.GetShortRealmName(className)
		desc := realm.GetRealmDescription(className)
		idx := i
		list.AddItem(name, desc, 0, func() {
			v.createRealm(cn)
		})
		_ = idx // used for help panel logic
	}

	list.AddItem("[white:red]"+i18n.T("common.cancel")+"[-]", i18n.T("common.return"), 0, func() {
		v.showRealmConfig()
	})

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index < len(realmTypes) {
			switch realmTypes[index] {
			case realm.ClassUserDatabaseRealm:
				helpPanel.SetHelpKey("help.security.realm.userdatabase")
			case realm.ClassDataSourceRealm:
				helpPanel.SetHelpKey("help.security.realm.datasource")
			case realm.ClassJNDIRealm:
				helpPanel.SetHelpKey("help.security.realm.jndi")
			case realm.ClassLockOutRealm:
				helpPanel.SetHelpKey("help.security.realm.lockout")
			case realm.ClassCombinedRealm:
				helpPanel.SetHelpKey("help.security.realm.combined")
			default:
				helpPanel.SetText("")
			}
		} else {
			helpPanel.SetText("")
		}
	})

	// Initialize help
	if len(realmTypes) > 0 {
		helpPanel.SetHelpKey("help.security.realm.userdatabase")
	}

	list.SetBorder(true).SetTitle(" " + i18n.T("security.realm.selecttype") + " ").SetBorderColor(tcell.ColorGreen)

	// Create flex layout with list and help panel
	flex := tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("realm-type-selector", flex, true)
	v.app.SetFocus(list)
}

// createRealm creates a new realm of the specified type
func (v *SecurityView) createRealm(className string) {
	srv := v.configService.GetServer()

	var newRealm server.Realm
	switch className {
	case realm.ClassUserDatabaseRealm:
		newRealm = realm.DefaultUserDatabaseRealm()
		v.showUserDatabaseRealmForm(&newRealm, true)
		return
	case realm.ClassDataSourceRealm:
		newRealm = realm.DefaultDataSourceRealm()
		v.showDataSourceRealmForm(&newRealm, true)
		return
	case realm.ClassJNDIRealm:
		newRealm = realm.DefaultJNDIRealm()
		v.showJNDIRealmForm(&newRealm, true)
		return
	case realm.ClassLockOutRealm:
		newRealm = realm.DefaultLockOutRealm()
	case realm.ClassCombinedRealm:
		newRealm = realm.DefaultCombinedRealm()
	default:
		newRealm = server.Realm{ClassName: className}
	}

	srv.Services[0].Engine.Realm = &newRealm
	v.configService.UpdateService(0, srv.Services[0])
	if err := v.configService.Save(); err != nil {
		v.showError(fmt.Sprintf("Failed to save: %v", err))
		return
	}
	v.setStatus("[green]Realm configured[-]")
	v.showRealmConfig()
}

// showRealmDetail shows realm detail based on type
func (v *SecurityView) showRealmDetail(r *server.Realm, serviceIndex int) {
	switch r.ClassName {
	case realm.ClassUserDatabaseRealm:
		v.showUserDatabaseRealmForm(r, false)
	case realm.ClassDataSourceRealm:
		v.showDataSourceRealmForm(r, false)
	case realm.ClassJNDIRealm:
		v.showJNDIRealmForm(r, false)
	case realm.ClassLockOutRealm, realm.ClassCombinedRealm:
		v.showWrapperRealmForm(r)
	default:
		v.showGenericRealmForm(r)
	}
}

// showUserDatabaseRealmForm shows UserDatabaseRealm configuration form
func (v *SecurityView) showUserDatabaseRealmForm(r *server.Realm, isNew bool) {
	// Create help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey(userDatabaseRealmHelpKeys[0])

	// Create preview panel
	previewPanel := NewPreviewPanel()

	form := tview.NewForm()

	updatePreview := func() {
		tempRealm := server.Realm{
			ClassName:    r.ClassName,
			ResourceName: form.GetFormItem(0).(*tview.InputField).GetText(),
		}
		previewPanel.SetXMLPreview(GenerateRealmXML(&tempRealm))
	}

	form.AddInputField("Resource Name", r.ResourceName, 30, nil, func(text string) {
		updatePreview()
	})

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		r.ResourceName = form.GetFormItem(0).(*tview.InputField).GetText()

		if isNew {
			srv := v.configService.GetServer()
			srv.Services[0].Engine.Realm = r
			v.configService.UpdateService(0, srv.Services[0])
		}

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]UserDatabaseRealm configured[-]")
		v.showRealmConfig()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showRealmConfig()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" UserDatabaseRealm ").SetBorderColor(tcell.ColorDarkCyan)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showRealmConfig()
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(userDatabaseRealmHelpKeys) {
						helpPanel.SetHelpKey(userDatabaseRealmHelpKeys[idx])
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

	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("userdatabase-realm-form", flex, true)
	v.app.SetFocus(form)
}

// showDataSourceRealmForm shows DataSourceRealm configuration form
func (v *SecurityView) showDataSourceRealmForm(r *server.Realm, isNew bool) {
	// Create help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey(dataSourceRealmHelpKeys[0])

	// Create preview panel
	previewPanel := NewPreviewPanel()

	form := tview.NewForm()

	updatePreview := func() {
		tempRealm := server.Realm{
			ClassName:      r.ClassName,
			DataSourceName: form.GetFormItem(0).(*tview.InputField).GetText(),
			UserTable:      form.GetFormItem(1).(*tview.InputField).GetText(),
			UserNameCol:    form.GetFormItem(2).(*tview.InputField).GetText(),
			UserCredCol:    form.GetFormItem(3).(*tview.InputField).GetText(),
			UserRoleTable:  form.GetFormItem(4).(*tview.InputField).GetText(),
			RoleNameCol:    form.GetFormItem(5).(*tview.InputField).GetText(),
		}
		previewPanel.SetXMLPreview(GenerateRealmXML(&tempRealm))
	}

	form.AddInputField("DataSource Name", r.DataSourceName, 40, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("User Table", r.UserTable, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("User Name Column", r.UserNameCol, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("User Credential Column", r.UserCredCol, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("User Role Table", r.UserRoleTable, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Role Name Column", r.RoleNameCol, 30, nil, func(text string) {
		updatePreview()
	})

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		r.DataSourceName = form.GetFormItem(0).(*tview.InputField).GetText()
		r.UserTable = form.GetFormItem(1).(*tview.InputField).GetText()
		r.UserNameCol = form.GetFormItem(2).(*tview.InputField).GetText()
		r.UserCredCol = form.GetFormItem(3).(*tview.InputField).GetText()
		r.UserRoleTable = form.GetFormItem(4).(*tview.InputField).GetText()
		r.RoleNameCol = form.GetFormItem(5).(*tview.InputField).GetText()

		if isNew {
			srv := v.configService.GetServer()
			srv.Services[0].Engine.Realm = r
			v.configService.UpdateService(0, srv.Services[0])
		}

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]DataSourceRealm configured[-]")
		v.showRealmConfig()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showRealmConfig()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" DataSourceRealm ").SetBorderColor(tcell.ColorDarkCyan)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showRealmConfig()
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(dataSourceRealmHelpKeys) {
						helpPanel.SetHelpKey(dataSourceRealmHelpKeys[idx])
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

	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("datasource-realm-form", flex, true)
	v.app.SetFocus(form)
}

// showJNDIRealmForm shows JNDIRealm (LDAP) configuration form
func (v *SecurityView) showJNDIRealmForm(r *server.Realm, isNew bool) {
	// Create help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey(jndiRealmHelpKeys[0])

	// Create preview panel
	previewPanel := NewPreviewPanel()

	form := tview.NewForm()

	updatePreview := func() {
		tempRealm := server.Realm{
			ClassName:          r.ClassName,
			ConnectionURL:      form.GetFormItem(0).(*tview.InputField).GetText(),
			ConnectionName:     form.GetFormItem(1).(*tview.InputField).GetText(),
			ConnectionPassword: form.GetFormItem(2).(*tview.InputField).GetText(),
			UserPattern:        form.GetFormItem(3).(*tview.InputField).GetText(),
			UserBase:           form.GetFormItem(4).(*tview.InputField).GetText(),
			UserSearch:         form.GetFormItem(5).(*tview.InputField).GetText(),
			RoleBase:           form.GetFormItem(6).(*tview.InputField).GetText(),
			RoleName:           form.GetFormItem(7).(*tview.InputField).GetText(),
			RoleSearch:         form.GetFormItem(8).(*tview.InputField).GetText(),
		}
		previewPanel.SetXMLPreview(GenerateRealmXML(&tempRealm))
	}

	form.AddInputField("Connection URL", r.ConnectionURL, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Connection Name (DN)", r.ConnectionName, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddPasswordField("Connection Password", r.ConnectionPassword, 30, '*', func(text string) {
		updatePreview()
	})
	form.AddInputField("User Pattern", r.UserPattern, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("User Base", r.UserBase, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("User Search", r.UserSearch, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Role Base", r.RoleBase, 50, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Role Name", r.RoleName, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Role Search", r.RoleSearch, 50, nil, func(text string) {
		updatePreview()
	})

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		r.ConnectionURL = form.GetFormItem(0).(*tview.InputField).GetText()
		r.ConnectionName = form.GetFormItem(1).(*tview.InputField).GetText()
		r.ConnectionPassword = form.GetFormItem(2).(*tview.InputField).GetText()
		r.UserPattern = form.GetFormItem(3).(*tview.InputField).GetText()
		r.UserBase = form.GetFormItem(4).(*tview.InputField).GetText()
		r.UserSearch = form.GetFormItem(5).(*tview.InputField).GetText()
		r.RoleBase = form.GetFormItem(6).(*tview.InputField).GetText()
		r.RoleName = form.GetFormItem(7).(*tview.InputField).GetText()
		r.RoleSearch = form.GetFormItem(8).(*tview.InputField).GetText()

		if isNew {
			srv := v.configService.GetServer()
			srv.Services[0].Engine.Realm = r
			v.configService.UpdateService(0, srv.Services[0])
		}

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]JNDIRealm configured[-]")
		v.showRealmConfig()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showRealmConfig()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" JNDIRealm (LDAP) ").SetBorderColor(tcell.ColorDarkCyan)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showRealmConfig()
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(jndiRealmHelpKeys) {
						helpPanel.SetHelpKey(jndiRealmHelpKeys[idx])
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

	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("jndi-realm-form", flex, true)
	v.app.SetFocus(form)
}

// showWrapperRealmForm shows LockOutRealm/CombinedRealm configuration
func (v *SecurityView) showWrapperRealmForm(r *server.Realm) {
	list := tview.NewList().ShowSecondaryText(true)

	realmName := realm.GetShortRealmName(r.ClassName)

	list.AddItem(fmt.Sprintf("[::b]%s[::-]", realmName), realm.GetRealmDescription(r.ClassName), 0, nil)
	list.AddItem("", "", 0, nil)

	// Show nested realms
	if len(r.NestedRealms) > 0 {
		list.AddItem("[::b]Nested Realms:[::-]", "", 0, nil)
		for i, nested := range r.NestedRealms {
			idx := i
			nestedName := realm.GetShortRealmName(nested.ClassName)
			list.AddItem(
				fmt.Sprintf("  [yellow]%s[-]", nestedName),
				realm.GetRealmDescription(nested.ClassName),
				0,
				func() { v.editNestedRealm(r, idx) },
			)
		}
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]+ Add Nested Realm[-]", "Add a realm to this wrapper", 'a', func() {
		v.showAddNestedRealmSelector(r)
	})

	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to realm config", 'b', func() {
		v.showRealmConfig()
	})

	list.SetBorder(true).SetTitle(fmt.Sprintf(" %s ", realmName)).SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("wrapper-realm-form", list, true)
	v.app.SetFocus(list)
}

// showAddNestedRealmSelector shows realm type selector for nested realms
func (v *SecurityView) showAddNestedRealmSelector(parentRealm *server.Realm) {
	list := tview.NewList().ShowSecondaryText(true)

	// Only show non-wrapper realms for nesting
	nestedTypes := []string{
		realm.ClassUserDatabaseRealm,
		realm.ClassDataSourceRealm,
		realm.ClassJNDIRealm,
		realm.ClassJAASRealm,
	}

	for _, className := range nestedTypes {
		cn := className
		name := realm.GetShortRealmName(className)
		desc := realm.GetRealmDescription(className)
		list.AddItem(name, desc, 0, func() {
			var newRealm server.Realm
			switch cn {
			case realm.ClassUserDatabaseRealm:
				newRealm = realm.DefaultUserDatabaseRealm()
			case realm.ClassDataSourceRealm:
				newRealm = realm.DefaultDataSourceRealm()
			case realm.ClassJNDIRealm:
				newRealm = realm.DefaultJNDIRealm()
			default:
				newRealm = server.Realm{ClassName: cn}
			}
			parentRealm.NestedRealms = append(parentRealm.NestedRealms, newRealm)

			if err := v.configService.Save(); err != nil {
				v.showError(fmt.Sprintf("Failed to save: %v", err))
				return
			}
			v.setStatus("[green]Nested realm added[-]")
			v.showWrapperRealmForm(parentRealm)
		})
	}

	list.AddItem("[red]Cancel[-]", "Return", 0, func() {
		v.showWrapperRealmForm(parentRealm)
	})

	list.SetBorder(true).SetTitle(" Add Nested Realm ").SetBorderColor(tcell.ColorGreen)
	v.pages.AddAndSwitchToPage("add-nested-realm", list, true)
	v.app.SetFocus(list)
}

// editNestedRealm allows editing a nested realm
func (v *SecurityView) editNestedRealm(parentRealm *server.Realm, index int) {
	if index >= len(parentRealm.NestedRealms) {
		return
	}

	nested := &parentRealm.NestedRealms[index]

	list := tview.NewList().ShowSecondaryText(true)

	list.AddItem("Edit Settings", "Configure this realm", 'e', func() {
		v.showRealmDetail(nested, 0)
	})

	list.AddItem("[red]Remove[-]", "Remove this nested realm", 'd', func() {
		v.showConfirm("Remove Nested Realm", "Remove this nested realm?", func(confirmed bool) {
			if confirmed {
				parentRealm.NestedRealms = append(parentRealm.NestedRealms[:index], parentRealm.NestedRealms[index+1:]...)
				if err := v.configService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]Nested realm removed[-]")
			}
			v.showWrapperRealmForm(parentRealm)
		})
	})

	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return", 'b', func() {
		v.showWrapperRealmForm(parentRealm)
	})

	nestedName := realm.GetShortRealmName(nested.ClassName)
	list.SetBorder(true).SetTitle(fmt.Sprintf(" Nested Realm: %s ", nestedName))
	v.pages.AddAndSwitchToPage("edit-nested-realm", list, true)
	v.app.SetFocus(list)
}

// showNestedRealmDetail shows detail of nested realm
func (v *SecurityView) showNestedRealmDetail(index int) {
	srv := v.configService.GetServer()
	if srv == nil || len(srv.Services) == 0 || srv.Services[0].Engine.Realm == nil {
		return
	}

	parentRealm := srv.Services[0].Engine.Realm
	if index >= len(parentRealm.NestedRealms) {
		return
	}

	v.editNestedRealm(parentRealm, index)
}

// showGenericRealmForm shows a generic realm form
func (v *SecurityView) showGenericRealmForm(r *server.Realm) {
	// Create help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey(genericRealmHelpKeys[0])

	// Create preview panel
	previewPanel := NewPreviewPanel()

	form := tview.NewForm()

	updatePreview := func() {
		tempRealm := server.Realm{
			ClassName: form.GetFormItem(0).(*tview.InputField).GetText(),
		}
		previewPanel.SetXMLPreview(GenerateRealmXML(&tempRealm))
	}

	form.AddInputField("Class Name", r.ClassName, 60, nil, func(text string) {
		updatePreview()
	})

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		r.ClassName = form.GetFormItem(0).(*tview.InputField).GetText()
		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]Realm updated[-]")
		v.showRealmConfig()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showRealmConfig()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Realm Configuration ").SetBorderColor(tcell.ColorDarkCyan)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showRealmConfig()
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(genericRealmHelpKeys) {
						helpPanel.SetHelpKey(genericRealmHelpKeys[idx])
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

	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("generic-realm-form", flex, true)
	v.app.SetFocus(form)
}

// showCredentialHandler shows credential handler configuration
func (v *SecurityView) showCredentialHandler() {
	srv := v.configService.GetServer()
	if srv == nil || len(srv.Services) == 0 || srv.Services[0].Engine.Realm == nil {
		v.showError("No realm configured. Configure a realm first.")
		return
	}

	r := srv.Services[0].Engine.Realm
	if r.CredentialHandler == nil {
		r.CredentialHandler = &server.CredentialHandler{
			ClassName: "org.apache.catalina.realm.MessageDigestCredentialHandler",
			Algorithm: "SHA-256",
		}
	}

	form := tview.NewForm()
	form.AddDropDown("Handler Class", realm.CredentialHandlerClasses(),
		indexOf(r.CredentialHandler.ClassName, realm.CredentialHandlerClasses()), nil)
	form.AddDropDown("Algorithm", realm.CredentialHandlerAlgorithms(),
		indexOf(r.CredentialHandler.Algorithm, realm.CredentialHandlerAlgorithms()), nil)
	form.AddInputField("Iterations", strconv.Itoa(r.CredentialHandler.Iterations), 10, acceptDigits, nil)
	form.AddInputField("Salt Length", strconv.Itoa(r.CredentialHandler.SaltLength), 10, acceptDigits, nil)

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		_, r.CredentialHandler.ClassName = form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()
		_, r.CredentialHandler.Algorithm = form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
		r.CredentialHandler.Iterations, _ = strconv.Atoi(form.GetFormItem(2).(*tview.InputField).GetText())
		r.CredentialHandler.SaltLength, _ = strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())

		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]Credential handler updated[-]")
		v.Show()
	})

	form.AddButton("[white:red]"+i18n.T("common.remove")+"[-:-]", func() {
		r.CredentialHandler = nil
		if err := v.configService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]Credential handler removed[-]")
		v.Show()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.Show()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Credential Handler (Password Hashing) ").SetBorderColor(tcell.ColorDarkCyan)

	helpText := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[yellow]Credential Handler[-] defines how passwords are hashed.

[white]Recommended:[-] SHA-256 or SHA-512 with iterations >= 1000
[white]Generate hash:[-] $CATALINA_HOME/bin/digest.sh -a <algorithm> <password>`)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(helpText, 4, 0, false)

	layout.SetBorder(true).SetTitle(" Credential Handler ")
	v.pages.AddAndSwitchToPage("credential-handler", layout, true)
	v.app.SetFocus(form)
}

// showUsersConfig shows tomcat-users.xml configuration
func (v *SecurityView) showUsersConfig() {
	if err := v.usersService.Load(); err != nil {
		v.showError(fmt.Sprintf("Failed to load tomcat-users.xml:\n%v", err))
		return
	}

	list := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	users := v.usersService.GetUsers()
	roles := v.usersService.GetRoles()

	list.AddItem(
		fmt.Sprintf("[::b]"+i18n.T("security.users.list")+"[::-] [yellow](%d)[-]", len(users)),
		i18n.T("security.users.list.desc"),
		'u',
		func() { v.showUsersList() },
	)

	list.AddItem(
		fmt.Sprintf("[::b]"+i18n.T("security.roles.list")+"[::-] [yellow](%d)[-]", len(roles)),
		i18n.T("security.roles.list.desc"),
		'r',
		func() { v.showRolesList() },
	)

	list.AddItem("[-:-:-] [white:red] "+i18n.T("common.back")+" [-:-:-]", i18n.T("common.return"), 'b', func() {
		v.Show()
	})

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.security.users.list")
		case 1:
			helpPanel.SetHelpKey("help.security.roles.list")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help
	helpPanel.SetHelpKey("help.security.users.list")

	list.SetBorder(true).SetTitle(" " + i18n.T("security.users.title") + " ").SetBorderColor(tcell.ColorDarkCyan)

	// Create flex layout with list and help panel
	flex := tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("users-config", flex, true)
	v.app.SetFocus(list)
}

// showUsersList shows list of users
func (v *SecurityView) showUsersList() {
	list := tview.NewList().ShowSecondaryText(true)

	users := v.usersService.GetUsers()
	for _, user := range users {
		u := user
		list.AddItem(
			fmt.Sprintf("[yellow]%s[-]", u.Username),
			fmt.Sprintf("Roles: %s", u.Roles),
			0,
			func() { v.showUserDetail(u.Username) },
		)
	}

	list.AddItem("[green]+ Add User[-]", "Create a new user", 'a', func() {
		v.showAddUser()
	})

	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to users config", 'b', func() {
		v.showUsersConfig()
	})

	list.SetBorder(true).SetTitle(" Users ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("users-list", list, true)
	v.app.SetFocus(list)
}

// showUserDetail shows user detail form
func (v *SecurityView) showUserDetail(username string) {
	user := v.usersService.GetUser(username)
	if user == nil {
		return
	}

	// Create help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey(userFormHelpKeys[0])

	// Create preview panel
	previewPanel := NewPreviewPanel()

	form := tview.NewForm()

	updatePreview := func() {
		previewPanel.SetXMLPreview(GenerateUserXML(
			form.GetFormItem(0).(*tview.InputField).GetText(),
			form.GetFormItem(1).(*tview.InputField).GetText(),
			form.GetFormItem(2).(*tview.InputField).GetText(),
		))
	}

	form.AddInputField("Username", user.Username, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddPasswordField("Password", user.Password, 30, '*', func(text string) {
		updatePreview()
	})
	form.AddInputField("Roles", user.Roles, 50, nil, func(text string) {
		updatePreview()
	})

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		newUsername := form.GetFormItem(0).(*tview.InputField).GetText()
		user.Username = newUsername
		user.Password = form.GetFormItem(1).(*tview.InputField).GetText()
		user.Roles = form.GetFormItem(2).(*tview.InputField).GetText()

		if err := v.usersService.UpdateUser(username, *user); err != nil {
			v.showError(fmt.Sprintf("Failed to update: %v", err))
			return
		}

		if err := v.usersService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]User updated[-]")
		v.showUsersList()
	})

	form.AddButton("[white:red]"+i18n.T("common.delete")+"[-:-]", func() {
		v.showConfirm("Delete User", fmt.Sprintf("Delete user '%s'?", username), func(confirmed bool) {
			if confirmed {
				if err := v.usersService.DeleteUser(username); err != nil {
					v.showError(fmt.Sprintf("Failed to delete: %v", err))
					return
				}
				if err := v.usersService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]User deleted[-]")
			}
			v.showUsersList()
		})
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showUsersList()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(fmt.Sprintf(" User: %s ", username)).SetBorderColor(tcell.ColorDarkCyan)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showUsersList()
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(userFormHelpKeys) {
						helpPanel.SetHelpKey(userFormHelpKeys[idx])
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

	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("user-detail", flex, true)
	v.app.SetFocus(form)
}

// showAddUser shows add user form
func (v *SecurityView) showAddUser() {
	// Create help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey(userFormHelpKeys[0])

	// Create preview panel
	previewPanel := NewPreviewPanel()

	form := tview.NewForm()

	updatePreview := func() {
		previewPanel.SetXMLPreview(GenerateUserXML(
			form.GetFormItem(0).(*tview.InputField).GetText(),
			form.GetFormItem(1).(*tview.InputField).GetText(),
			form.GetFormItem(2).(*tview.InputField).GetText(),
		))
	}

	form.AddInputField("Username", "", 30, nil, func(text string) {
		updatePreview()
	})
	form.AddPasswordField("Password", "", 30, '*', func(text string) {
		updatePreview()
	})
	form.AddInputField("Roles", "", 50, nil, func(text string) {
		updatePreview()
	})

	form.AddButton("[white:green]"+i18n.T("common.add")+"[-:-]", func() {
		user := realm.User{
			Username: form.GetFormItem(0).(*tview.InputField).GetText(),
			Password: form.GetFormItem(1).(*tview.InputField).GetText(),
			Roles:    form.GetFormItem(2).(*tview.InputField).GetText(),
		}

		if user.Username == "" {
			v.showError("Username is required")
			return
		}

		if err := v.usersService.AddUser(user); err != nil {
			v.showError(fmt.Sprintf("Failed to add: %v", err))
			return
		}

		if err := v.usersService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]User added[-]")
		v.showUsersList()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showUsersList()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Add User ").SetBorderColor(tcell.ColorGreen)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showUsersList()
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(userFormHelpKeys) {
						helpPanel.SetHelpKey(userFormHelpKeys[idx])
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

	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("add-user", flex, true)
	v.app.SetFocus(form)
}

// showRolesList shows list of roles
func (v *SecurityView) showRolesList() {
	list := tview.NewList().ShowSecondaryText(true)

	roles := v.usersService.GetRoles()
	for _, role := range roles {
		r := role
		desc := r.Description
		if desc == "" {
			desc = "(no description)"
		}
		list.AddItem(
			fmt.Sprintf("[yellow]%s[-]", r.RoleName),
			desc,
			0,
			func() { v.showRoleDetail(r.RoleName) },
		)
	}

	list.AddItem("[green]+ Add Role[-]", "Create a new role", 'a', func() {
		v.showAddRole()
	})

	list.AddItem("[green]+ Add Common Roles[-]", "Add standard Tomcat roles", 'c', func() {
		v.showAddCommonRoles()
	})

	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to users config", 'b', func() {
		v.showUsersConfig()
	})

	list.SetBorder(true).SetTitle(" Roles ").SetBorderColor(tcell.ColorDarkCyan)
	v.pages.AddAndSwitchToPage("roles-list", list, true)
	v.app.SetFocus(list)
}

// showRoleDetail shows role detail
func (v *SecurityView) showRoleDetail(roleName string) {
	role := v.usersService.GetRole(roleName)
	if role == nil {
		return
	}

	// Create help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey(roleFormHelpKeys[0])

	// Create preview panel
	previewPanel := NewPreviewPanel()

	form := tview.NewForm()

	updatePreview := func() {
		previewPanel.SetXMLPreview(GenerateRoleXML(
			form.GetFormItem(0).(*tview.InputField).GetText(),
			form.GetFormItem(1).(*tview.InputField).GetText(),
		))
	}

	form.AddInputField("Role Name", role.RoleName, 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Description", role.Description, 50, nil, func(text string) {
		updatePreview()
	})

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		// Delete old role and add new one if name changed
		newRoleName := form.GetFormItem(0).(*tview.InputField).GetText()
		if newRoleName != roleName {
			v.usersService.DeleteRole(roleName)
		}

		newRole := realm.Role{
			RoleName:    newRoleName,
			Description: form.GetFormItem(1).(*tview.InputField).GetText(),
		}
		v.usersService.AddRole(newRole)

		if err := v.usersService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]Role updated[-]")
		v.showRolesList()
	})

	form.AddButton("[white:red]"+i18n.T("common.delete")+"[-:-]", func() {
		v.showConfirm("Delete Role", fmt.Sprintf("Delete role '%s'?", roleName), func(confirmed bool) {
			if confirmed {
				if err := v.usersService.DeleteRole(roleName); err != nil {
					v.showError(fmt.Sprintf("Failed to delete: %v", err))
					return
				}
				if err := v.usersService.Save(); err != nil {
					v.showError(fmt.Sprintf("Failed to save: %v", err))
					return
				}
				v.setStatus("[green]Role deleted[-]")
			}
			v.showRolesList()
		})
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showRolesList()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(fmt.Sprintf(" Role: %s ", roleName)).SetBorderColor(tcell.ColorDarkCyan)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showRolesList()
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(roleFormHelpKeys) {
						helpPanel.SetHelpKey(roleFormHelpKeys[idx])
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

	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("role-detail", flex, true)
	v.app.SetFocus(form)
}

// showAddRole shows add role form
func (v *SecurityView) showAddRole() {
	// Create help panel
	helpPanel := NewDynamicHelpPanel()
	helpPanel.SetHelpKey(roleFormHelpKeys[0])

	// Create preview panel
	previewPanel := NewPreviewPanel()

	form := tview.NewForm()

	updatePreview := func() {
		previewPanel.SetXMLPreview(GenerateRoleXML(
			form.GetFormItem(0).(*tview.InputField).GetText(),
			form.GetFormItem(1).(*tview.InputField).GetText(),
		))
	}

	form.AddInputField("Role Name", "", 30, nil, func(text string) {
		updatePreview()
	})
	form.AddInputField("Description", "", 50, nil, func(text string) {
		updatePreview()
	})

	form.AddButton("[white:green]"+i18n.T("common.add")+"[-:-]", func() {
		role := realm.Role{
			RoleName:    form.GetFormItem(0).(*tview.InputField).GetText(),
			Description: form.GetFormItem(1).(*tview.InputField).GetText(),
		}

		if role.RoleName == "" {
			v.showError("Role name is required")
			return
		}

		if err := v.usersService.AddRole(role); err != nil {
			v.showError(fmt.Sprintf("Failed to add: %v", err))
			return
		}

		if err := v.usersService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus("[green]Role added[-]")
		v.showRolesList()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showRolesList()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Add Role ").SetBorderColor(tcell.ColorGreen)

	// Update help panel on navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showRolesList()
			return nil
		}

		// Update help on Tab/Enter/Up/Down navigation
		if event.Key() == tcell.KeyTab || event.Key() == tcell.KeyEnter ||
			event.Key() == tcell.KeyUp || event.Key() == tcell.KeyDown {
			go func() {
				v.app.QueueUpdateDraw(func() {
					idx, _ := form.GetFocusedItemIndex()
					if idx >= 0 && idx < len(roleFormHelpKeys) {
						helpPanel.SetHelpKey(roleFormHelpKeys[idx])
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

	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	v.pages.AddAndSwitchToPage("add-role", flex, true)
	v.app.SetFocus(form)
}

// showAddCommonRoles shows common roles selector
func (v *SecurityView) showAddCommonRoles() {
	list := tview.NewList().ShowSecondaryText(true)

	for _, role := range realm.CommonRoles {
		r := role
		existing := v.usersService.GetRole(r.RoleName)
		status := ""
		if existing != nil {
			status = " [gray](exists)[-]"
		}
		list.AddItem(
			fmt.Sprintf("%s%s", r.RoleName, status),
			r.Description,
			0,
			func() {
				if err := v.usersService.AddRole(r); err == nil {
					if err := v.usersService.Save(); err != nil {
						v.showError(fmt.Sprintf("Failed to save: %v", err))
						return
					}
					v.setStatus(fmt.Sprintf("[green]Role '%s' added[-]", r.RoleName))
					v.showAddCommonRoles() // Refresh
				}
			},
		)
	}

	list.AddItem("[green]Add All[-]", "Add all common roles", 'a', func() {
		added := 0
		for _, role := range realm.CommonRoles {
			if err := v.usersService.AddRole(role); err == nil {
				added++
			}
		}
		if err := v.usersService.Save(); err != nil {
			v.showError(fmt.Sprintf("Failed to save: %v", err))
			return
		}
		v.setStatus(fmt.Sprintf("[green]Added %d roles[-]", added))
		v.showRolesList()
	})

	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to roles list", 'b', func() {
		v.showRolesList()
	})

	list.SetBorder(true).SetTitle(" Add Common Roles ").SetBorderColor(tcell.ColorGreen)
	v.pages.AddAndSwitchToPage("add-common-roles", list, true)
	v.app.SetFocus(list)
}

// Helper functions
func (v *SecurityView) showError(message string) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			v.Show()
		})
	modal.SetBackgroundColor(tcell.ColorDarkRed)
	v.pages.AddAndSwitchToPage("error", modal, true)
}

func (v *SecurityView) showConfirm(title, message string, callback func(bool)) {
	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			callback(buttonIndex == 0)
		})
	modal.SetBorder(true).SetTitle(fmt.Sprintf(" %s ", title))
	v.pages.AddAndSwitchToPage("confirm", modal, true)
}

func (v *SecurityView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(fmt.Sprintf(" %s", message))
	}
}
