package utils

import (
	"encoding/base64"
	"testing"
	"time"
)

func TestUcFirst(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"lowercase", "gin", "Gin"},
		{"empty", "", ""},
		{"multibyte", "åbc", "Åbc"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := UcFirst(tt.input); got != tt.want {
				t.Fatalf("UcFirst(%q)=%q want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestLcFirst(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"uppercase", "Gin", "gin"},
		{"empty", "", ""},
		{"multibyte", "ÅBC", "åBC"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := LcFirst(tt.input); got != tt.want {
				t.Fatalf("LcFirst(%q)=%q want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestBase64Decode(t *testing.T) {
	want := []byte("gin-web-admin")
	encoded := base64.StdEncoding.EncodeToString(want)

	decoded, err := Base64Decode(encoded)
	if err != nil {
		t.Fatalf("Base64Decode returned unexpected error: %v", err)
	}

	if string(decoded) != string(want) {
		t.Fatalf("decoded bytes %q want %q", decoded, want)
	}

	if _, err := Base64Decode("not-base64$$$"); err == nil {
		t.Fatal("expected error for invalid base64 input")
	}
}

func TestTimeToDateTimesString(t *testing.T) {
	ts := time.Date(2024, time.August, 15, 18, 30, 45, 0, time.UTC)
	got := TimeToDateTimesString(ts)
	if got != "2024-08-15 18:30:45" {
		t.Fatalf("TimeToDateTimesString returned %q", got)
	}
}
