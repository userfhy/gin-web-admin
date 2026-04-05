package security

import (
	"html"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	reScriptTag = regexp.MustCompile(`(?is)<\s*(script|style|iframe|object|embed|link|meta)[^>]*>.*?<\s*/\s*(script|style|iframe|object|embed|link|meta)\s*>`)
	reTag       = regexp.MustCompile(`(?is)<[^>]+>`)
	reEvtAttr   = regexp.MustCompile(`(?is)\son[a-z]+\s*=\s*(".*?"|'.*?'|[^\s>]+)`)
	reJSProto   = regexp.MustCompile(`(?is)(javascript|vbscript|data)\s*:`)
	reCtrlChar  = regexp.MustCompile(`[\x00-\x08\x0b\x0c\x0e-\x1f\x7f]`)
	reSpaces    = regexp.MustCompile(`[ \t]+`)
)

// SanitizePlainText cleans a plain text field and escapes html entities.
func SanitizePlainText(s string, maxLen int) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	s = reCtrlChar.ReplaceAllString(s, "")
	s = reScriptTag.ReplaceAllString(s, "")
	s = reTag.ReplaceAllString(s, "")
	s = reJSProto.ReplaceAllString(s, "")
	s = reSpaces.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	s = html.EscapeString(s)
	if maxLen > 0 && utf8.RuneCountInString(s) > maxLen {
		r := []rune(s)
		s = string(r[:maxLen])
	}
	return s
}

// SanitizeMarkdown removes dangerous html/script payload while keeping markdown text.
func SanitizeMarkdown(s string, maxLen int) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	s = reCtrlChar.ReplaceAllString(s, "")
	s = reScriptTag.ReplaceAllString(s, "")
	s = reEvtAttr.ReplaceAllString(s, "")
	s = reJSProto.ReplaceAllString(s, "")
	// block inline raw html in markdown
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.TrimSpace(s)
	if maxLen > 0 && utf8.RuneCountInString(s) > maxLen {
		r := []rune(s)
		s = string(r[:maxLen])
	}
	return s
}
