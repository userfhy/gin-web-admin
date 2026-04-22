package setting

import "testing"

func TestDurationUnmarshalText(t *testing.T) {
	var d Duration
	if err := d.UnmarshalText([]byte("7d")); err != nil {
		t.Fatalf("expected 7d to parse: %v", err)
	}
	if d.Std().Hours() != 168 {
		t.Fatalf("expected 168 hours, got %v", d.Std())
	}

	if err := d.UnmarshalText([]byte("2h30m")); err != nil {
		t.Fatalf("expected go duration to parse: %v", err)
	}
	if d.Std().Minutes() != 150 {
		t.Fatalf("expected 150 minutes, got %v", d.Std())
	}
}

func TestDurationUnmarshalTextInvalid(t *testing.T) {
	var d Duration
	if err := d.UnmarshalText([]byte("7x")); err == nil {
		t.Fatal("expected invalid duration error")
	}
}
