package security

import (
	"html"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	reCtrlChar = regexp.MustCompile(`[\x00-\x08\x0b\x0c\x0e-\x1f\x7f]`)
	reSpaces   = regexp.MustCompile(`[ \t]+`)
)

func normalizeInput(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	// avoid double-escaping on repeated edit/save
	s = html.UnescapeString(s)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	s = reCtrlChar.ReplaceAllString(s, "")
	return strings.TrimSpace(s)
}

func cutRunes(s string, maxLen int) string {
	if maxLen <= 0 {
		return s
	}
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}
	r := []rune(s)
	return string(r[:maxLen])
}

// SanitizePlainText keeps full text semantics and stores HTML-escaped content.
func SanitizePlainText(s string, maxLen int) string {
	s = normalizeInput(s)
	if s == "" {
		return ""
	}
	s = strings.ReplaceAll(s, "\n", " ")
	s = reSpaces.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	s = cutRunes(s, maxLen)
	return html.EscapeString(s)
}

// SanitizeMarkdown preserves markdown/code text but neutralizes all raw HTML by escaping it.
func SanitizeMarkdown(s string, maxLen int) string {
	s = normalizeInput(s)
	if s == "" {
		return ""
	}
	s = strings.TrimSpace(s)
	s = cutRunes(s, maxLen)
	return html.EscapeString(s)
}
