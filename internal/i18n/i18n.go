package i18n

import "sync"

// Language represents a supported language
type Language string

const (
	English  Language = "en"
	Korean   Language = "ko"
	Japanese Language = "ja"
)

// Current language (default: English)
var (
	currentLang Language = English
	langMutex   sync.RWMutex
)

// SetLanguage sets the current language
func SetLanguage(lang Language) {
	langMutex.Lock()
	defer langMutex.Unlock()
	currentLang = lang
}

// GetLanguage returns the current language
func GetLanguage() Language {
	langMutex.RLock()
	defer langMutex.RUnlock()
	return currentLang
}

// GetLanguageName returns the display name for a language
func GetLanguageName(lang Language) string {
	switch lang {
	case English:
		return "English"
	case Korean:
		return "한국어"
	case Japanese:
		return "日本語"
	default:
		return string(lang)
	}
}

// AvailableLanguages returns all supported languages
func AvailableLanguages() []Language {
	return []Language{English, Korean, Japanese}
}

// T translates a message key to the current language
func T(key string) string {
	return Translate(key, GetLanguage())
}

// Translate translates a message key to the specified language
func Translate(key string, lang Language) string {
	if messages, ok := translations[lang]; ok {
		if msg, ok := messages[key]; ok {
			return msg
		}
	}
	// Fallback to English
	if messages, ok := translations[English]; ok {
		if msg, ok := messages[key]; ok {
			return msg
		}
	}
	// Return key if not found
	return key
}
