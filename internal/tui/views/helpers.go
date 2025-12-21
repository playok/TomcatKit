package views

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// HelpPanel creates a help text view for displaying detailed descriptions
func HelpPanel(helpKey string) *tview.TextView {
	helpText := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true).
		SetScrollable(true)

	helpText.SetText(i18n.T(helpKey))
	helpText.SetBorder(true).
		SetTitle(" " + i18n.T("help.title") + " ").
		SetBorderColor(tcell.ColorDarkGreen)

	return helpText
}

// CreateFormWithHelp creates a flex layout with form on left and help panel on right
func CreateFormWithHelp(form *tview.Form, helpKey string, title string) *tview.Flex {
	helpPanel := HelpPanel(helpKey)

	flex := tview.NewFlex().
		AddItem(form, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	return flex
}

// CreateListWithHelp creates a flex layout with list on left and help panel on right
func CreateListWithHelp(list *tview.List, helpKey string) *tview.Flex {
	helpPanel := HelpPanel(helpKey)

	flex := tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(helpPanel, 0, 1, false)

	return flex
}

// UpdateHelpPanel updates the text of a help panel
func UpdateHelpPanel(helpPanel *tview.TextView, helpKey string) {
	helpPanel.SetText(i18n.T(helpKey))
}

// CreateDynamicHelpPanel creates a help panel that can be updated dynamically
type DynamicHelpPanel struct {
	*tview.TextView
}

// NewDynamicHelpPanel creates a new dynamic help panel
func NewDynamicHelpPanel() *DynamicHelpPanel {
	panel := &DynamicHelpPanel{
		TextView: tview.NewTextView().
			SetDynamicColors(true).
			SetWordWrap(true).
			SetScrollable(true),
	}
	panel.SetBorder(true).
		SetTitle(" " + i18n.T("help.title") + " ").
		SetBorderColor(tcell.ColorDarkGreen)

	return panel
}

// SetHelpKey updates the help panel with the specified help key
func (p *DynamicHelpPanel) SetHelpKey(helpKey string) {
	p.SetText(i18n.T(helpKey))
}

// Form helper functions to reduce cyclomatic complexity

// GetFormText safely gets text from an InputField by label
func GetFormText(form *tview.Form, label string) string {
	if field := form.GetFormItemByLabel(label); field != nil {
		if input, ok := field.(*tview.InputField); ok {
			return input.GetText()
		}
	}
	return ""
}

// GetFormBool safely gets bool from a Checkbox by label
func GetFormBool(form *tview.Form, label string) bool {
	if field := form.GetFormItemByLabel(label); field != nil {
		if checkbox, ok := field.(*tview.Checkbox); ok {
			return checkbox.IsChecked()
		}
	}
	return false
}

// GetFormInt safely gets int from an InputField by label
func GetFormInt(form *tview.Form, label string) int {
	if field := form.GetFormItemByLabel(label); field != nil {
		if input, ok := field.(*tview.InputField); ok {
			val, _ := strconv.Atoi(input.GetText())
			return val
		}
	}
	return 0
}

// GetFormDropDownIndex safely gets the selected index from a DropDown by label
func GetFormDropDownIndex(form *tview.Form, label string) int {
	if field := form.GetFormItemByLabel(label); field != nil {
		if dropdown, ok := field.(*tview.DropDown); ok {
			index, _ := dropdown.GetCurrentOption()
			return index
		}
	}
	return -1
}

// GetFormDropDownText safely gets the selected text from a DropDown by label
func GetFormDropDownText(form *tview.Form, label string) string {
	if field := form.GetFormItemByLabel(label); field != nil {
		if dropdown, ok := field.(*tview.DropDown); ok {
			_, text := dropdown.GetCurrentOption()
			return text
		}
	}
	return ""
}

// BoolToString converts bool to "true" or empty string
func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return ""
}

// StringToBool converts string "true" to bool
func StringToBool(s string) bool {
	return s == "true"
}

// ParseTextAreaLines parses a text area into lines, trimming whitespace and removing empty lines
func ParseTextAreaLines(text string) []string {
	var lines []string
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

// ParseCommaSeparated parses a comma-separated string into a slice
func ParseCommaSeparated(text string) []string {
	var items []string
	for _, item := range strings.Split(text, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			items = append(items, item)
		}
	}
	return items
}

// GetFormTextArea safely gets text from a TextArea by label
func GetFormTextArea(form *tview.Form, label string) string {
	if field := form.GetFormItemByLabel(label); field != nil {
		if textarea, ok := field.(*tview.TextArea); ok {
			return textarea.GetText()
		}
	}
	return ""
}

// ParseKeyValueLines parses lines of "key=value" format into a map
func ParseKeyValueLines(text string) map[string]string {
	result := make(map[string]string)
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			result[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return result
}
