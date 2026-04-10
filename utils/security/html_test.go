package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeRichHTML_RemovesDangerousContent(t *testing.T) {
	raw := `<script>alert('x')</script><p onclick="evil()">Hello</p><a href="javascript:bad()">link</a>`
	clean := SanitizeRichHTML(raw, 0)

	assert.NotContains(t, clean, "script")
	assert.NotContains(t, clean, "onclick")
	assert.NotContains(t, clean, "javascript:")
	assert.Contains(t, clean, "Hello")
}

func TestSanitizeRichHTML_TruncatesAndNeutralizesSQL(t *testing.T) {
	raw := `<p>UNION select * from users -- </p>`
	clean := SanitizeRichHTML(raw, 0)
	assert.NotContains(t, clean, "UNION")
	assert.NotContains(t, clean, "--")
}
