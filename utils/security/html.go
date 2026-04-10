package security

import (
	"html"
	"regexp"
	"strings"
)

var (
	reScriptTag   = regexp.MustCompile(`(?is)<script.*?>.*?</script>`)
	reStyleTag    = regexp.MustCompile(`(?is)<style.*?>.*?</style>`)
	reEventAttr   = regexp.MustCompile(`(?i)on[a-z]+\s*=`)
	reJsProtocol  = regexp.MustCompile(`(?i)javascript:`)
	reIframeTag   = regexp.MustCompile(`(?is)<iframe.*?>.*?</iframe>`)
	reSQLKeywords = regexp.MustCompile(`(?i)(union\s+select|drop\s+table|delete\s+from|insert\s+into|update\s+\w+)`)
)

// SanitizeRichHTML removes risky tags/attributes and escapes the remaining text.
func SanitizeRichHTML(input string, maxLen int) string {
	input = normalizeInput(input)
	if input == "" {
		return ""
	}
	if maxLen > 0 {
		input = cutRunes(input, maxLen)
	}
	clean := reScriptTag.ReplaceAllString(input, "")
	clean = reStyleTag.ReplaceAllString(clean, "")
	clean = reIframeTag.ReplaceAllString(clean, "")
	clean = reEventAttr.ReplaceAllString(clean, "")
	clean = reJsProtocol.ReplaceAllString(clean, "")
	clean = reSQLKeywords.ReplaceAllString(clean, "")
	clean = strings.ReplaceAll(clean, "`", "")
	return neutralizeSQLMeta(html.EscapeString(clean))
}

func neutralizeSQLMeta(s string) string {
	if s == "" {
		return ""
	}
	replacements := []struct {
		old string
		new string
	}{
		{"--", "—"},
		{"/*", "/&#42;"},
		{"*/", "&#42;/"},
		{"'", "&#39;"},
		{"\"", "&quot;"},
	}
	for _, repl := range replacements {
		s = strings.ReplaceAll(s, repl.old, repl.new)
	}
	return s
}
