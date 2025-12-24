package views

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/logging"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// LoggingView provides TUI for logging configuration
type LoggingView struct {
	app           *tview.Application
	pages         *tview.Pages
	mainPages     *tview.Pages
	statusBar     *tview.TextView
	onReturn      func()
	configService *logging.ConfigService
}

// NewLoggingView creates a new logging configuration view
func NewLoggingView(app *tview.Application, mainPages *tview.Pages, statusBar *tview.TextView, catalinaBase string, onReturn func()) *LoggingView {
	return &LoggingView{
		app:           app,
		mainPages:     mainPages,
		statusBar:     statusBar,
		onReturn:      onReturn,
		configService: logging.NewConfigService(catalinaBase),
	}
}

// Load initializes the view
func (v *LoggingView) Load() error {
	if err := v.configService.Load(); err != nil {
		return fmt.Errorf("failed to load logging configuration: %w", err)
	}

	v.pages = tview.NewPages()
	v.showMainMenu()

	v.mainPages.AddAndSwitchToPage("logging", v.pages, true)
	return nil
}

// showMainMenu displays the logging configuration main menu
func (v *LoggingView) showMainMenu() {
	config := v.configService.GetConfig()

	// Help panel
	helpPanel := NewDynamicHelpPanel()

	// Count items
	handlerCount := len(config.FileHandlers)
	loggerCount := len(config.Loggers)
	rootHandlerCount := len(config.RootHandlers)

	list := tview.NewList().
		AddItem(i18n.T("logging.filehandlers"), fmt.Sprintf(i18n.T("logging.filehandlers.count"), handlerCount), 'f', func() {
			v.showFileHandlerList()
		}).
		AddItem(i18n.T("logging.console"), fmt.Sprintf("Level: %s", config.ConsoleHandler.Level), 'c', func() {
			v.showConsoleHandlerForm()
		}).
		AddItem(i18n.T("logging.loggers"), fmt.Sprintf(i18n.T("logging.loggers.count"), loggerCount), 'l', func() {
			v.showLoggerList()
		}).
		AddItem(i18n.T("logging.rootlogger"), fmt.Sprintf(i18n.T("logging.rootlogger.count"), rootHandlerCount), 'r', func() {
			v.showRootLoggerForm()
		}).
		AddItem(i18n.T("common.save"), i18n.T("logging.save.desc"), 's', func() {
			v.saveConfiguration()
		}).
		AddItem(i18n.T("common.back"), i18n.T("common.return"), 'b', func() {
			v.mainPages.RemovePage("logging")
			v.onReturn()
		})

	// Update help panel when selection changes
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			helpPanel.SetHelpKey("help.logging.filehandlers")
		case 1:
			helpPanel.SetHelpKey("help.logging.console")
		case 2:
			helpPanel.SetHelpKey("help.logging.loggers")
		case 3:
			helpPanel.SetHelpKey("help.logging.rootlogger")
		default:
			helpPanel.SetText("")
		}
	})

	// Initialize help with first item
	helpPanel.SetHelpKey("help.logging.filehandlers")

	list.SetBorder(true).SetTitle(" " + i18n.T("logging.title") + " ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.mainPages.RemovePage("logging")
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
	v.setStatus("Logging configuration: " + v.configService.GetConfigPath())
}

// showFileHandlerList displays the list of file handlers
func (v *LoggingView) showFileHandlerList() {
	config := v.configService.GetConfig()

	list := tview.NewList()

	for _, handler := range config.FileHandlers {
		h := handler // Capture for closure
		handlerType := "Sync"
		if h.IsAsync() {
			handlerType = "Async"
		}
		list.AddItem(
			h.Prefix,
			fmt.Sprintf("%s | Level: %s | File: %s%s", handlerType, h.Level, h.FilePrefix, h.Suffix),
			0,
			func() {
				v.showFileHandlerForm(&h, false)
			},
		)
	}

	list.AddItem("+ Add File Handler", "Create a new file handler", 'a', func() {
		v.showFileHandlerForm(nil, true)
	})

	list.AddItem("Back", "Return to logging menu", 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" File Handlers ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("handlers", list, true)
	v.setStatus("Select a file handler to configure")
}

// showFileHandlerForm displays a form for editing a file handler
func (v *LoggingView) showFileHandlerForm(handler *logging.FileHandler, isNew bool) {
	var h logging.FileHandler
	if handler != nil {
		h = *handler
	} else {
		// Generate a new prefix
		prefix := fmt.Sprintf("%dcustom", len(v.configService.GetConfig().FileHandlers)+1)
		h = *logging.NewAsyncFileHandler(prefix, "custom.")
	}

	form := tview.NewForm()

	// Prefix
	form.AddInputField("Prefix", h.Prefix, 20, nil, func(text string) {
		h.Prefix = text
	})

	// Handler type
	handlerTypes := []string{"AsyncFileHandler (Recommended)", "FileHandler"}
	initialType := 0
	if h.ClassName == logging.HandlerFileHandler {
		initialType = 1
	}
	form.AddDropDown("Handler Type", handlerTypes, initialType, func(option string, index int) {
		if index == 0 {
			h.ClassName = logging.HandlerAsyncFileHandler
		} else {
			h.ClassName = logging.HandlerFileHandler
		}
	})

	// Level
	levels := make([]string, 0)
	levelDescriptions := make([]string, 0)
	initialLevel := 0
	for i, level := range logging.AvailableLogLevels() {
		levels = append(levels, string(level))
		levelDescriptions = append(levelDescriptions, fmt.Sprintf("%s - %s", level, logging.GetLogLevelDescription(level)))
		if level == h.Level {
			initialLevel = i
		}
	}
	form.AddDropDown("Level", levelDescriptions, initialLevel, func(option string, index int) {
		h.Level = logging.AvailableLogLevels()[index]
	})

	// Directory
	form.AddInputField("Directory", h.Directory, 40, nil, func(text string) {
		h.Directory = text
	})

	// File prefix
	form.AddInputField("File Prefix", h.FilePrefix, 30, nil, func(text string) {
		h.FilePrefix = text
	})

	// Suffix
	form.AddInputField("Suffix", h.Suffix, 10, nil, func(text string) {
		h.Suffix = text
	})

	// Max days
	form.AddInputField("Max Days", strconv.Itoa(h.MaxDays), 10, acceptNumber, func(text string) {
		h.MaxDays, _ = strconv.Atoi(text)
	})

	// Encoding
	form.AddInputField("Encoding", h.Encoding, 15, nil, func(text string) {
		h.Encoding = text
	})

	// Buffer size (only for async)
	form.AddInputField("Buffer Size", strconv.Itoa(h.BufferSize), 10, acceptNumber, func(text string) {
		h.BufferSize, _ = strconv.Atoi(text)
	})

	// Formatter
	formatters := []string{"(Default)", logging.FormatterOneLineFormatter, logging.FormatterVerboseFormatter, logging.FormatterSimpleFormatter}
	initialFormatter := 0
	for i, f := range formatters {
		if f == h.Formatter {
			initialFormatter = i
			break
		}
	}
	form.AddDropDown("Formatter", formatters, initialFormatter, func(option string, index int) {
		if index == 0 {
			h.Formatter = ""
		} else {
			h.Formatter = formatters[index]
		}
	})

	// Save button
	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		if h.Prefix == "" {
			v.setStatus("Error: Prefix is required")
			return
		}

		if isNew {
			v.configService.AddFileHandler(&h)
			v.setStatus("File handler added: " + h.Prefix)
		} else {
			// Update existing
			v.configService.RemoveFileHandler(handler.Prefix)
			v.configService.AddFileHandler(&h)
			v.setStatus("File handler updated: " + h.Prefix)
		}
		v.showFileHandlerList()
	})

	// Delete button (only for existing)
	if !isNew {
		form.AddButton("[white:red]"+i18n.T("common.delete")+"[-:-]", func() {
			v.configService.RemoveFileHandler(handler.Prefix)
			v.setStatus("File handler deleted: " + handler.Prefix)
			v.showFileHandlerList()
		})
	}

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showFileHandlerList()
	})

	title := " Add File Handler "
	if !isNew {
		title = " Edit File Handler: " + handler.Prefix + " "
	}
	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showFileHandlerList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("handler-form", form, true)
}

// showConsoleHandlerForm displays the console handler configuration
func (v *LoggingView) showConsoleHandlerForm() {
	config := v.configService.GetConfig()
	h := *config.ConsoleHandler

	form := tview.NewForm()

	// Level
	levels := make([]string, 0)
	initialLevel := 0
	for i, level := range logging.AvailableLogLevels() {
		levels = append(levels, fmt.Sprintf("%s - %s", level, logging.GetLogLevelDescription(level)))
		if level == h.Level {
			initialLevel = i
		}
	}
	form.AddDropDown("Level", levels, initialLevel, func(option string, index int) {
		h.Level = logging.AvailableLogLevels()[index]
	})

	// Formatter
	formatters := []string{logging.FormatterOneLineFormatter, logging.FormatterVerboseFormatter, logging.FormatterSimpleFormatter}
	initialFormatter := 0
	for i, f := range formatters {
		if f == h.Formatter {
			initialFormatter = i
			break
		}
	}
	form.AddDropDown("Formatter", formatters, initialFormatter, func(option string, index int) {
		h.Formatter = formatters[index]
	})

	// Encoding
	form.AddInputField("Encoding", h.Encoding, 15, nil, func(text string) {
		h.Encoding = text
	})

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		v.configService.SetConsoleHandler(&h)
		v.setStatus("Console handler configuration saved")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Console Handler Configuration ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("console-form", form, true)
}

// showLoggerList displays the list of configured loggers
func (v *LoggingView) showLoggerList() {
	config := v.configService.GetConfig()

	list := tview.NewList()

	for _, logger := range config.Loggers {
		l := logger // Capture for closure
		handlerInfo := "default"
		if len(l.Handlers) > 0 {
			handlerInfo = strings.Join(l.Handlers, ", ")
		}
		list.AddItem(
			l.Name,
			fmt.Sprintf("Level: %s | Handlers: %s", l.Level, handlerInfo),
			0,
			func() {
				v.showLoggerForm(&l, false)
			},
		)
	}

	list.AddItem("+ Add Logger", "Create a new logger configuration", 'a', func() {
		v.showLoggerForm(nil, true)
	})

	list.AddItem("Quick Add Common Logger", "Add from common Tomcat loggers", 'q', func() {
		v.showQuickAddLoggerMenu()
	})

	list.AddItem("Back", "Return to logging menu", 'b', func() {
		v.showMainMenu()
	})

	list.SetBorder(true).SetTitle(" Loggers ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("loggers", list, true)
	v.setStatus("Select a logger to configure")
}

// showQuickAddLoggerMenu shows common loggers to add
func (v *LoggingView) showQuickAddLoggerMenu() {
	config := v.configService.GetConfig()
	existingLoggers := make(map[string]bool)
	for _, l := range config.Loggers {
		existingLoggers[l.Name] = true
	}

	list := tview.NewList()

	for _, loggerName := range logging.CommonLoggers {
		name := loggerName // Capture for closure
		if existingLoggers[name] {
			continue // Skip already configured
		}

		// Short name for display
		shortName := name
		if len(shortName) > 50 {
			shortName = "..." + shortName[len(shortName)-47:]
		}

		list.AddItem(shortName, name, 0, func() {
			logger := &logging.Logger{
				Name:              name,
				Level:             logging.LogLevelInfo,
				UseParentHandlers: true,
			}
			v.showLoggerForm(logger, true)
		})
	}

	list.AddItem("Back", "Return to logger list", 'b', func() {
		v.showLoggerList()
	})

	list.SetBorder(true).SetTitle(" Quick Add Common Logger ")

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showLoggerList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("quick-add", list, true)
}

// showLoggerForm displays a form for editing a logger
func (v *LoggingView) showLoggerForm(logger *logging.Logger, isNew bool) {
	var l logging.Logger
	if logger != nil {
		l = *logger
	} else {
		l = logging.Logger{
			Name:              "",
			Level:             logging.LogLevelInfo,
			UseParentHandlers: true,
		}
	}

	form := tview.NewForm()

	// Logger name
	form.AddInputField("Logger Name", l.Name, 60, nil, func(text string) {
		l.Name = text
	})

	// Level
	levels := make([]string, 0)
	initialLevel := 0
	for i, level := range logging.AvailableLogLevels() {
		levels = append(levels, fmt.Sprintf("%s - %s", level, logging.GetLogLevelDescription(level)))
		if level == l.Level {
			initialLevel = i
		}
	}
	form.AddDropDown("Level", levels, initialLevel, func(option string, index int) {
		l.Level = logging.AvailableLogLevels()[index]
	})

	// Handlers
	handlersStr := strings.Join(l.Handlers, ", ")
	form.AddInputField("Handlers", handlersStr, 60, nil, func(text string) {
		handlers := []string{}
		for _, h := range strings.Split(text, ",") {
			h = strings.TrimSpace(h)
			if h != "" {
				handlers = append(handlers, h)
			}
		}
		l.Handlers = handlers
	})

	// Use parent handlers
	form.AddCheckbox("Use Parent Handlers", l.UseParentHandlers, func(checked bool) {
		l.UseParentHandlers = checked
	})

	// Available handlers hint
	config := v.configService.GetConfig()
	var availableHandlers []string
	for _, h := range config.FileHandlers {
		availableHandlers = append(availableHandlers, h.GetHandlerName())
	}
	availableHandlers = append(availableHandlers, logging.HandlerConsoleHandler)
	form.AddTextView("Available Handlers", strings.Join(availableHandlers, "\n"), 60, 4, true, false)

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		if l.Name == "" {
			v.setStatus("Error: Logger name is required")
			return
		}

		if isNew {
			v.configService.AddLogger(&l)
			v.setStatus("Logger added: " + l.Name)
		} else {
			v.configService.UpdateLogger(logger.Name, &l)
			v.setStatus("Logger updated: " + l.Name)
		}
		v.showLoggerList()
	})

	if !isNew {
		form.AddButton("[white:red]"+i18n.T("common.delete")+"[-:-]", func() {
			v.configService.RemoveLogger(logger.Name)
			v.setStatus("Logger deleted: " + logger.Name)
			v.showLoggerList()
		})
	}

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showLoggerList()
	})

	title := " Add Logger "
	if !isNew {
		title = " Edit Logger "
	}
	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(title)

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showLoggerList()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("logger-form", form, true)
}

// showRootLoggerForm displays the root logger configuration
func (v *LoggingView) showRootLoggerForm() {
	config := v.configService.GetConfig()

	form := tview.NewForm()

	// Show all available handlers with checkboxes
	selectedHandlers := make(map[string]bool)
	for _, h := range config.RootHandlers {
		selectedHandlers[h] = true
	}

	// Get all available handlers
	var allHandlers []string
	for _, h := range config.FileHandlers {
		allHandlers = append(allHandlers, h.GetHandlerName())
	}
	allHandlers = append(allHandlers, logging.HandlerConsoleHandler)

	// Add checkboxes for each handler
	for _, handlerName := range allHandlers {
		name := handlerName // Capture for closure
		form.AddCheckbox(name, selectedHandlers[name], func(checked bool) {
			selectedHandlers[name] = checked
		})
	}

	form.AddButton("[white:green]"+i18n.T("common.save.short")+"[-:-]", func() {
		var newHandlers []string
		for _, h := range allHandlers {
			if selectedHandlers[h] {
				newHandlers = append(newHandlers, h)
			}
		}
		v.configService.SetRootHandlers(newHandlers)
		v.setStatus("Root logger handlers updated")
		v.showMainMenu()
	})

	form.AddButton("[black:yellow]"+i18n.T("common.cancel")+"[-:-]", func() {
		v.showMainMenu()
	})

	form.SetButtonBackgroundColor(tcell.ColorDefault)
	form.SetBorder(true).SetTitle(" Root Logger Handlers ")

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.showMainMenu()
			return nil
		}
		return event
	})

	v.pages.AddAndSwitchToPage("root-form", form, true)
	v.setStatus("Select handlers for root logger (.handlers)")
}

// saveConfiguration saves the logging configuration
func (v *LoggingView) saveConfiguration() {
	if err := v.configService.Save(); err != nil {
		v.setStatus("Error saving: " + err.Error())
		return
	}
	v.setStatus("Configuration saved to: " + v.configService.GetConfigPath())
}

// setStatus updates the status bar
func (v *LoggingView) setStatus(message string) {
	if v.statusBar != nil {
		v.statusBar.SetText(" " + message)
	}
}
