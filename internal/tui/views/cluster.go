package views

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// ClusterView handles cluster configuration
type ClusterView struct {
	app           *tview.Application
	pages         *tview.Pages
	mainPages     *tview.Pages
	statusBar     *tview.TextView
	onReturn      func()
	configService *server.ConfigService
}

// NewClusterView creates a new cluster view
func NewClusterView(app *tview.Application, mainPages *tview.Pages, configService *server.ConfigService, statusBar *tview.TextView, onReturn func()) *ClusterView {
	return &ClusterView{
		app:           app,
		pages:         tview.NewPages(),
		mainPages:     mainPages,
		statusBar:     statusBar,
		onReturn:      onReturn,
		configService: configService,
	}
}

// Show displays the cluster view
func (v *ClusterView) Show() {
	v.showMainMenu()
	v.mainPages.AddAndSwitchToPage("cluster", v.pages, true)
}

// getCluster returns the cluster configuration, creating default if needed
func (v *ClusterView) getCluster() *server.Cluster {
	cfg := v.configService.GetConfig()
	if cfg == nil || len(cfg.Services) == 0 {
		return nil
	}
	return cfg.Services[0].Engine.Cluster
}

// showMainMenu shows the main cluster menu
func (v *ClusterView) showMainMenu() {
	menu := tview.NewList().ShowSecondaryText(true)

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	cluster := v.getCluster()
	status := "[red]" + i18n.T("common.disabled") + "[-]"
	if cluster != nil {
		status = "[green]" + i18n.T("common.enabled") + "[-]"
	}

	menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("cluster.status")+"[::-]: %s", status),
		i18n.T("cluster.status.desc"), 's', func() {
			v.showClusterToggle()
		})

	if cluster != nil {
		menu.AddItem("", "", 0, nil)

		menu.AddItem("[::b]"+i18n.T("cluster.settings")+"[::-]",
			i18n.T("cluster.settings.desc"), 'c', func() {
				v.showClusterSettings()
			})

		menu.AddItem("[::b]"+i18n.T("cluster.manager")+"[::-]",
			i18n.T("cluster.manager.desc"), 'm', func() {
				v.showManagerSettings()
			})

		menu.AddItem("[::b]"+i18n.T("cluster.membership")+"[::-]",
			i18n.T("cluster.membership.desc"), 'b', func() {
				v.showMembershipSettings()
			})

		menu.AddItem("[::b]"+i18n.T("cluster.receiver")+"[::-]",
			i18n.T("cluster.receiver.desc"), 'r', func() {
				v.showReceiverSettings()
			})

		menu.AddItem("[::b]"+i18n.T("cluster.sender")+"[::-]",
			i18n.T("cluster.sender.desc"), 'n', func() {
				v.showSenderSettings()
			})

		interceptorCount := 0
		if cluster.Channel != nil {
			interceptorCount = len(cluster.Channel.Interceptors)
		}
		menu.AddItem(fmt.Sprintf("[::b]"+i18n.T("cluster.interceptors")+"[::-] [yellow](%d)[-]", interceptorCount),
			i18n.T("cluster.interceptors.desc"), 'i', func() {
				v.showInterceptorList()
			})

		menu.AddItem("[::b]"+i18n.T("cluster.deployer")+"[::-]",
			i18n.T("cluster.deployer.desc"), 'd', func() {
				v.showDeployerSettings()
			})
	}

	menu.AddItem("", "", 0, nil)
	menu.AddItem("[-:-:-] [white:red] "+i18n.T("common.back")+" [-:-:-]", i18n.T("common.return"), 0, func() {
		v.onReturn()
	})

	// Update help panel when selection changes
	menu.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.cluster.status")
		case 2:
			helpPanel.SetHelpKey("help.cluster.settings")
		case 3:
			helpPanel.SetHelpKey("help.cluster.manager")
		case 4:
			helpPanel.SetHelpKey("help.cluster.membership")
		case 5:
			helpPanel.SetHelpKey("help.cluster.receiver")
		case 6:
			helpPanel.SetHelpKey("help.cluster.sender")
		case 7:
			helpPanel.SetHelpKey("help.cluster.interceptors")
		case 8:
			helpPanel.SetHelpKey("help.cluster.deployer")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.cluster.status")

	menu.SetBorder(true).SetTitle(" " + i18n.T("cluster.title") + " ").SetBorderColor(tcell.ColorDarkCyan)
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

// showClusterToggle shows enable/disable cluster dialog
func (v *ClusterView) showClusterToggle() {
	cluster := v.getCluster()

	if cluster != nil {
		// Cluster is enabled, ask to disable
		modal := tview.NewModal().
			SetText("Disable clustering?\n\nThis will remove all cluster configuration.").
			AddButtons([]string{"Disable", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Disable" {
					cfg := v.configService.GetConfig()
					if cfg != nil && len(cfg.Services) > 0 {
						cfg.Services[0].Engine.Cluster = nil
					}
					if err := v.configService.Save(); err != nil {
						v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					} else {
						v.setStatus("[green]Clustering disabled[-]")
					}
				}
				v.showMainMenu()
			})
		v.pages.AddAndSwitchToPage("toggle", modal, true)
	} else {
		// Cluster is disabled, ask to enable with defaults
		modal := tview.NewModal().
			SetText("Enable clustering?\n\nThis will create a default cluster configuration with:\n- DeltaManager\n- Multicast membership\n- NIO receiver\n- Standard interceptors").
			AddButtons([]string{"Enable", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Enable" {
					cfg := v.configService.GetConfig()
					if cfg != nil && len(cfg.Services) > 0 {
						cfg.Services[0].Engine.Cluster = server.DefaultCluster()
					}
					if err := v.configService.Save(); err != nil {
						v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
					} else {
						v.setStatus("[green]Clustering enabled with default configuration[-]")
					}
				}
				v.showMainMenu()
			})
		v.pages.AddAndSwitchToPage("toggle", modal, true)
	}
}

// showClusterSettings shows cluster basic settings form
func (v *ClusterView) showClusterSettings() {
	cluster := v.getCluster()
	if cluster == nil {
		v.setStatus("[red]Cluster not enabled[-]")
		v.showMainMenu()
		return
	}

	form := tview.NewForm()

	form.AddInputField("Channel Send Options", cluster.ChannelSendOptions, 10, nil, nil)
	form.AddInputField("Channel Start Options", cluster.ChannelStartOptions, 10, nil, nil)
	form.AddCheckbox("Notify Lifecycle Listener On Failure", cluster.NotifyLifecycleListenerOnFailure, nil)

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		cluster.ChannelSendOptions = form.GetFormItemByLabel("Channel Send Options").(*tview.InputField).GetText()
		cluster.ChannelStartOptions = form.GetFormItemByLabel("Channel Start Options").(*tview.InputField).GetText()
		cluster.NotifyLifecycleListenerOnFailure = form.GetFormItemByLabel("Notify Lifecycle Listener On Failure").(*tview.Checkbox).IsChecked()

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Cluster settings saved[-]")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Cluster Settings ").SetBorderColor(tcell.ColorBlue)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("cluster-settings", form, true)
	v.app.SetFocus(form)
}

// showManagerSettings shows session manager settings
func (v *ClusterView) showManagerSettings() {
	cluster := v.getCluster()
	if cluster == nil {
		return
	}

	if cluster.Manager == nil {
		cluster.Manager = &server.ClusterManager{
			ClassName: server.ClusterManagerDelta,
		}
	}

	form := tview.NewForm()

	managerTypes := []string{server.ClusterManagerDelta, server.ClusterManagerBackup}
	managerIdx := 0
	for i, m := range managerTypes {
		if m == cluster.Manager.ClassName {
			managerIdx = i
			break
		}
	}
	form.AddDropDown("Manager Type", []string{"DeltaManager", "BackupManager"}, managerIdx, nil)

	form.AddCheckbox("Expire Sessions On Shutdown", cluster.Manager.ExpireSessionsOnShutdown, nil)
	form.AddCheckbox("Notify Listeners On Replication", cluster.Manager.NotifyListenersOnReplication, nil)

	// DeltaManager specific
	form.AddInputField("State Transfer Timeout (ms)", strconv.Itoa(cluster.Manager.StateTransferTimeout), 10, nil, nil)
	form.AddCheckbox("Send All Sessions", cluster.Manager.SendAllSessions, nil)
	form.AddInputField("Send All Sessions Size", strconv.Itoa(cluster.Manager.SendAllSessionsSize), 10, nil, nil)

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		idx, _ := form.GetFormItemByLabel("Manager Type").(*tview.DropDown).GetCurrentOption()
		cluster.Manager.ClassName = managerTypes[idx]
		cluster.Manager.ExpireSessionsOnShutdown = form.GetFormItemByLabel("Expire Sessions On Shutdown").(*tview.Checkbox).IsChecked()
		cluster.Manager.NotifyListenersOnReplication = form.GetFormItemByLabel("Notify Listeners On Replication").(*tview.Checkbox).IsChecked()
		cluster.Manager.StateTransferTimeout, _ = strconv.Atoi(form.GetFormItemByLabel("State Transfer Timeout (ms)").(*tview.InputField).GetText())
		cluster.Manager.SendAllSessions = form.GetFormItemByLabel("Send All Sessions").(*tview.Checkbox).IsChecked()
		cluster.Manager.SendAllSessionsSize, _ = strconv.Atoi(form.GetFormItemByLabel("Send All Sessions Size").(*tview.InputField).GetText())

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Session manager saved[-]")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Session Manager ").SetBorderColor(tcell.ColorGreen)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("manager-settings", form, true)
	v.app.SetFocus(form)
}

// showMembershipSettings shows membership settings
func (v *ClusterView) showMembershipSettings() {
	cluster := v.getCluster()
	if cluster == nil || cluster.Channel == nil {
		return
	}

	if cluster.Channel.Membership == nil {
		cluster.Channel.Membership = &server.Membership{
			ClassName: server.MembershipMcastService,
			Address:   "228.0.0.4",
			Port:      45564,
			Frequency: 500,
			DropTime:  3000,
		}
	}

	m := cluster.Channel.Membership
	form := tview.NewForm()

	form.AddInputField("Multicast Address", m.Address, 20, nil, nil)
	form.AddInputField("Multicast Port", strconv.Itoa(m.Port), 10, nil, nil)
	form.AddInputField("Frequency (ms)", strconv.Itoa(m.Frequency), 10, nil, nil)
	form.AddInputField("Drop Time (ms)", strconv.Itoa(m.DropTime), 10, nil, nil)
	form.AddInputField("Bind Address", m.Bind, 20, nil, nil)
	form.AddCheckbox("Recovery Enabled", m.RecoveryEnabled, nil)
	form.AddInputField("Recovery Counter", strconv.Itoa(m.RecoveryCounter), 10, nil, nil)
	form.AddInputField("Recovery Sleep Time (ms)", strconv.Itoa(m.RecoverySleepTime), 10, nil, nil)
	form.AddCheckbox("Local Loopback Disabled", m.LocalLoopbackDisabled, nil)

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		m.Address = form.GetFormItemByLabel("Multicast Address").(*tview.InputField).GetText()
		m.Port, _ = strconv.Atoi(form.GetFormItemByLabel("Multicast Port").(*tview.InputField).GetText())
		m.Frequency, _ = strconv.Atoi(form.GetFormItemByLabel("Frequency (ms)").(*tview.InputField).GetText())
		m.DropTime, _ = strconv.Atoi(form.GetFormItemByLabel("Drop Time (ms)").(*tview.InputField).GetText())
		m.Bind = form.GetFormItemByLabel("Bind Address").(*tview.InputField).GetText()
		m.RecoveryEnabled = form.GetFormItemByLabel("Recovery Enabled").(*tview.Checkbox).IsChecked()
		m.RecoveryCounter, _ = strconv.Atoi(form.GetFormItemByLabel("Recovery Counter").(*tview.InputField).GetText())
		m.RecoverySleepTime, _ = strconv.Atoi(form.GetFormItemByLabel("Recovery Sleep Time (ms)").(*tview.InputField).GetText())
		m.LocalLoopbackDisabled = form.GetFormItemByLabel("Local Loopback Disabled").(*tview.Checkbox).IsChecked()

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Membership settings saved[-]")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Membership (Multicast) ").SetBorderColor(tcell.ColorYellow)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("membership-settings", form, true)
	v.app.SetFocus(form)
}

// showReceiverSettings shows receiver settings
func (v *ClusterView) showReceiverSettings() {
	cluster := v.getCluster()
	if cluster == nil || cluster.Channel == nil {
		return
	}

	if cluster.Channel.Receiver == nil {
		cluster.Channel.Receiver = &server.Receiver{
			ClassName:       server.ReceiverNioReceiver,
			Address:         "auto",
			Port:            4000,
			AutoBind:        100,
			SelectorTimeout: 5000,
			MaxThreads:      6,
		}
	}

	r := cluster.Channel.Receiver
	form := tview.NewForm()

	receiverTypes := []string{"NioReceiver", "BioReceiver"}
	receiverIdx := 0
	if r.ClassName == server.ReceiverBioReceiver {
		receiverIdx = 1
	}
	form.AddDropDown("Receiver Type", receiverTypes, receiverIdx, nil)

	form.AddInputField("Address", r.Address, 20, nil, nil)
	form.AddInputField("Port", strconv.Itoa(r.Port), 10, nil, nil)
	form.AddInputField("Auto Bind Range", strconv.Itoa(r.AutoBind), 10, nil, nil)
	form.AddInputField("Selector Timeout (ms)", strconv.Itoa(r.SelectorTimeout), 10, nil, nil)
	form.AddInputField("Max Threads", strconv.Itoa(r.MaxThreads), 10, nil, nil)
	form.AddInputField("Min Threads", strconv.Itoa(r.MinThreads), 10, nil, nil)
	form.AddInputField("RX Buffer Size", strconv.Itoa(r.RxBufSize), 10, nil, nil)
	form.AddInputField("TX Buffer Size", strconv.Itoa(r.TxBufSize), 10, nil, nil)
	form.AddInputField("Timeout (ms)", strconv.Itoa(r.Timeout), 10, nil, nil)

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		idx, _ := form.GetFormItemByLabel("Receiver Type").(*tview.DropDown).GetCurrentOption()
		if idx == 0 {
			r.ClassName = server.ReceiverNioReceiver
		} else {
			r.ClassName = server.ReceiverBioReceiver
		}
		r.Address = form.GetFormItemByLabel("Address").(*tview.InputField).GetText()
		r.Port, _ = strconv.Atoi(form.GetFormItemByLabel("Port").(*tview.InputField).GetText())
		r.AutoBind, _ = strconv.Atoi(form.GetFormItemByLabel("Auto Bind Range").(*tview.InputField).GetText())
		r.SelectorTimeout, _ = strconv.Atoi(form.GetFormItemByLabel("Selector Timeout (ms)").(*tview.InputField).GetText())
		r.MaxThreads, _ = strconv.Atoi(form.GetFormItemByLabel("Max Threads").(*tview.InputField).GetText())
		r.MinThreads, _ = strconv.Atoi(form.GetFormItemByLabel("Min Threads").(*tview.InputField).GetText())
		r.RxBufSize, _ = strconv.Atoi(form.GetFormItemByLabel("RX Buffer Size").(*tview.InputField).GetText())
		r.TxBufSize, _ = strconv.Atoi(form.GetFormItemByLabel("TX Buffer Size").(*tview.InputField).GetText())
		r.Timeout, _ = strconv.Atoi(form.GetFormItemByLabel("Timeout (ms)").(*tview.InputField).GetText())

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Receiver settings saved[-]")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Receiver ").SetBorderColor(tcell.ColorPurple)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("receiver-settings", form, true)
	v.app.SetFocus(form)
}

// showSenderSettings shows sender/transport settings
func (v *ClusterView) showSenderSettings() {
	cluster := v.getCluster()
	if cluster == nil || cluster.Channel == nil {
		return
	}

	if cluster.Channel.Sender == nil {
		cluster.Channel.Sender = &server.Sender{
			ClassName: server.SenderReplicationTransmitter,
			Transport: &server.Transport{
				ClassName: server.TransportPooledParallelSender,
			},
		}
	}

	if cluster.Channel.Sender.Transport == nil {
		cluster.Channel.Sender.Transport = &server.Transport{
			ClassName: server.TransportPooledParallelSender,
		}
	}

	t := cluster.Channel.Sender.Transport
	form := tview.NewForm()

	form.AddInputField("RX Buffer Size", strconv.Itoa(t.RxBufSize), 10, nil, nil)
	form.AddInputField("TX Buffer Size", strconv.Itoa(t.TxBufSize), 10, nil, nil)
	form.AddCheckbox("Direct Buffer", t.DirectBuffer, nil)
	form.AddInputField("Keep Alive Count", strconv.Itoa(t.KeepAliveCount), 10, nil, nil)
	form.AddInputField("Keep Alive Time (ms)", strconv.Itoa(t.KeepAliveTime), 10, nil, nil)
	form.AddInputField("Timeout (ms)", strconv.Itoa(t.Timeout), 10, nil, nil)
	form.AddInputField("Max Retry Attempts", strconv.Itoa(t.MaxRetryAttempts), 10, nil, nil)
	form.AddCheckbox("TCP No Delay", t.TcpNoDelay, nil)
	form.AddCheckbox("SO Keep Alive", t.SoKeepAlive, nil)
	form.AddCheckbox("Throw On Failed Ack", t.ThrowOnFailedAck, nil)

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		t.RxBufSize, _ = strconv.Atoi(form.GetFormItemByLabel("RX Buffer Size").(*tview.InputField).GetText())
		t.TxBufSize, _ = strconv.Atoi(form.GetFormItemByLabel("TX Buffer Size").(*tview.InputField).GetText())
		t.DirectBuffer = form.GetFormItemByLabel("Direct Buffer").(*tview.Checkbox).IsChecked()
		t.KeepAliveCount, _ = strconv.Atoi(form.GetFormItemByLabel("Keep Alive Count").(*tview.InputField).GetText())
		t.KeepAliveTime, _ = strconv.Atoi(form.GetFormItemByLabel("Keep Alive Time (ms)").(*tview.InputField).GetText())
		t.Timeout, _ = strconv.Atoi(form.GetFormItemByLabel("Timeout (ms)").(*tview.InputField).GetText())
		t.MaxRetryAttempts, _ = strconv.Atoi(form.GetFormItemByLabel("Max Retry Attempts").(*tview.InputField).GetText())
		t.TcpNoDelay = form.GetFormItemByLabel("TCP No Delay").(*tview.Checkbox).IsChecked()
		t.SoKeepAlive = form.GetFormItemByLabel("SO Keep Alive").(*tview.Checkbox).IsChecked()
		t.ThrowOnFailedAck = form.GetFormItemByLabel("Throw On Failed Ack").(*tview.Checkbox).IsChecked()

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Sender settings saved[-]")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Sender / Transport ").SetBorderColor(tcell.ColorBlue)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("sender-settings", form, true)
	v.app.SetFocus(form)
}

// showInterceptorList shows the list of interceptors
func (v *ClusterView) showInterceptorList() {
	cluster := v.getCluster()
	if cluster == nil || cluster.Channel == nil {
		return
	}

	list := tview.NewList().ShowSecondaryText(true)

	for i := range cluster.Channel.Interceptors {
		interceptor := &cluster.Channel.Interceptors[i]
		name := server.GetInterceptorShortName(interceptor.ClassName)
		desc := server.GetInterceptorDescription(interceptor.ClassName)
		list.AddItem(name, desc, 0, func() {
			v.showInterceptorForm(interceptor, false)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[green]► Add Interceptor[-]", "Add new interceptor", 'a', func() {
		v.showInterceptorTypeSelector()
	})
	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to cluster menu", 0, func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Channel Interceptors ").SetBorderColor(tcell.ColorGreen)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("interceptor-list", list, true)
	v.app.SetFocus(list)
}

// showInterceptorTypeSelector shows interceptor type selection
func (v *ClusterView) showInterceptorTypeSelector() {
	list := tview.NewList().ShowSecondaryText(true)

	for _, className := range server.AvailableInterceptorTypes() {
		cn := className
		name := server.GetInterceptorShortName(cn)
		desc := server.GetInterceptorDescription(cn)
		list.AddItem(name, desc, 0, func() {
			newInt := &server.Interceptor{ClassName: cn}
			v.showInterceptorForm(newInt, true)
		})
	}

	list.AddItem("", "", 0, nil)
	list.AddItem("[-:-:-] [white:red] Back [-:-:-]", "Return to interceptor list", 0, func() {
		v.showInterceptorList()
	})

	list.SetBorder(true).SetTitle(" Select Interceptor Type ").SetBorderColor(tcell.ColorYellow)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showInterceptorList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("interceptor-selector", list, true)
	v.app.SetFocus(list)
}

// showInterceptorForm shows interceptor form
func (v *ClusterView) showInterceptorForm(interceptor *server.Interceptor, isNew bool) {
	form := tview.NewForm()

	name := server.GetInterceptorShortName(interceptor.ClassName)
	form.AddTextView("Type", name, 40, 1, true, false)

	// Add fields based on interceptor type
	switch interceptor.ClassName {
	case server.InterceptorTcpFailureDetector:
		form.AddInputField("Connect Timeout (ms)", strconv.Itoa(interceptor.ConnectTimeout), 10, nil, nil)
		form.AddCheckbox("Perform Send Test", interceptor.PerformSendTest, nil)
		form.AddCheckbox("Perform Read Test", interceptor.PerformReadTest, nil)
		form.AddInputField("Read Test Timeout (ms)", strconv.Itoa(interceptor.ReadTestTimeout), 10, nil, nil)
		form.AddInputField("Remove Suspects Timeout (ms)", strconv.Itoa(interceptor.RemoveSuspectsTimeout), 10, nil, nil)

	case server.InterceptorMessageDispatch:
		form.AddInputField("Max Queue Size", strconv.Itoa(interceptor.MaxQueueSize), 10, nil, nil)
		form.AddCheckbox("Optional Queue", interceptor.OptionalQueue, nil)
		form.AddCheckbox("Always Send", interceptor.AlwaysSend, nil)

	case server.InterceptorThroughput:
		form.AddInputField("Interval (seconds)", strconv.Itoa(interceptor.Interval), 10, nil, nil)

	case server.InterceptorEncrypt:
		form.AddInputField("Encryption Algorithm", interceptor.EncryptionAlgorithm, 30, nil, nil)
		form.AddInputField("Encryption Key", interceptor.EncryptionKey, 40, nil, nil)
		form.AddInputField("Encryption Key File", interceptor.EncryptionKeyFile, 50, nil, nil)
	}

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		// Extract values based on type
		switch interceptor.ClassName {
		case server.InterceptorTcpFailureDetector:
			interceptor.ConnectTimeout = GetFormInt(form, "Connect Timeout (ms)")
			interceptor.PerformSendTest = GetFormBool(form, "Perform Send Test")
			interceptor.PerformReadTest = GetFormBool(form, "Perform Read Test")
			interceptor.ReadTestTimeout = GetFormInt(form, "Read Test Timeout (ms)")
			interceptor.RemoveSuspectsTimeout = GetFormInt(form, "Remove Suspects Timeout (ms)")

		case server.InterceptorMessageDispatch:
			interceptor.MaxQueueSize = GetFormInt(form, "Max Queue Size")
			interceptor.OptionalQueue = GetFormBool(form, "Optional Queue")
			interceptor.AlwaysSend = GetFormBool(form, "Always Send")

		case server.InterceptorThroughput:
			interceptor.Interval = GetFormInt(form, "Interval (seconds)")

		case server.InterceptorEncrypt:
			interceptor.EncryptionAlgorithm = GetFormText(form, "Encryption Algorithm")
			interceptor.EncryptionKey = GetFormText(form, "Encryption Key")
			interceptor.EncryptionKeyFile = GetFormText(form, "Encryption Key File")
		}

		if isNew {
			cluster := v.getCluster()
			if cluster != nil && cluster.Channel != nil {
				cluster.Channel.Interceptors = append(cluster.Channel.Interceptors, *interceptor)
			}
		}

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Interceptor saved[-]")
		v.showInterceptorList()
	})

	if !isNew {
		form.AddButton("[white:red]"+i18n.T("common.delete")+"[-:-]", func() {
			v.confirmDeleteInterceptor(interceptor)
		})
	}

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showInterceptorList()
	})

	title := fmt.Sprintf(" Edit %s ", name)
	if isNew {
		title = fmt.Sprintf(" Add %s ", name)
	}
	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(title).SetBorderColor(tcell.ColorGreen)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showInterceptorList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("interceptor-form", form, true)
	v.app.SetFocus(form)
}

// confirmDeleteInterceptor shows delete confirmation
func (v *ClusterView) confirmDeleteInterceptor(interceptor *server.Interceptor) {
	name := server.GetInterceptorShortName(interceptor.ClassName)
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Delete %s?", name)).
		AddButtons([]string{"Delete", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Delete" {
				cluster := v.getCluster()
				if cluster != nil && cluster.Channel != nil {
					for i := range cluster.Channel.Interceptors {
						if &cluster.Channel.Interceptors[i] == interceptor {
							cluster.Channel.Interceptors = append(cluster.Channel.Interceptors[:i], cluster.Channel.Interceptors[i+1:]...)
							break
						}
					}
				}
				if err := v.configService.Save(); err != nil {
					v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
				} else {
					v.setStatus("[green]Interceptor deleted[-]")
				}
			}
			v.showInterceptorList()
		})
	v.pages.AddAndSwitchToPage("confirm-delete", modal, true)
}

// showDeployerSettings shows farm deployer settings
func (v *ClusterView) showDeployerSettings() {
	cluster := v.getCluster()
	if cluster == nil {
		return
	}

	if cluster.Deployer == nil {
		cluster.Deployer = &server.FarmWarDeployer{
			ClassName:    server.DeployerFarmWarDeployer,
			TempDir:      "/tmp/war-temp/",
			DeployDir:    "/tmp/war-deploy/",
			WatchDir:     "/tmp/war-listen/",
			WatchEnabled: false,
		}
	}

	d := cluster.Deployer
	form := tview.NewForm()

	form.AddCheckbox("Watch Enabled", d.WatchEnabled, nil)
	form.AddInputField("Temp Dir", d.TempDir, 50, nil, nil)
	form.AddInputField("Deploy Dir", d.DeployDir, 50, nil, nil)
	form.AddInputField("Watch Dir", d.WatchDir, 50, nil, nil)
	form.AddInputField("Process Deploy Frequency", strconv.Itoa(d.ProcessDeployFrequency), 10, nil, nil)

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		d.WatchEnabled = form.GetFormItemByLabel("Watch Enabled").(*tview.Checkbox).IsChecked()
		d.TempDir = form.GetFormItemByLabel("Temp Dir").(*tview.InputField).GetText()
		d.DeployDir = form.GetFormItemByLabel("Deploy Dir").(*tview.InputField).GetText()
		d.WatchDir = form.GetFormItemByLabel("Watch Dir").(*tview.InputField).GetText()
		d.ProcessDeployFrequency, _ = strconv.Atoi(form.GetFormItemByLabel("Process Deploy Frequency").(*tview.InputField).GetText())

		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Farm deployer settings saved[-]")
		v.showMainMenu()
	})

	form.AddButton("[white:red]"+i18n.T("cluster.deployer.remove")+"[-:-]", func() {
		cluster.Deployer = nil
		if err := v.configService.Save(); err != nil {
			v.setStatus(fmt.Sprintf("[red]Failed to save: %v[-]", err))
			return
		}
		v.setStatus("[green]Farm deployer removed[-]")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Farm War Deployer ").SetBorderColor(tcell.ColorYellow)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("deployer-settings", form, true)
	v.app.SetFocus(form)
}

// setStatus updates the status bar
func (v *ClusterView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(fmt.Sprintf(" %s", message))
	}
}
