package com

import "testing"

func TestStrToConversions(t *testing.T) {
	s := StrTo("42")

	if got, err := s.Int(); err != nil || got != 42 {
		t.Fatalf("Int()=%d, err=%v", got, err)
	}

	if got, err := s.Int64(); err != nil || got != 42 {
		t.Fatalf("Int64()=%d, err=%v", got, err)
	}

	if got, err := s.Uint8(); err != nil || got != 42 {
		t.Fatalf("Uint8()=%d, err=%v", got, err)
	}

	if got, err := s.Float64(); err != nil || got != 42 {
		t.Fatalf("Float64()=%f, err=%v", got, err)
	}
}

func TestStrToHelpers(t *testing.T) {
	zero := StrTo(string(rune(0x1E)))
	if zero.Exist() {
		t.Fatal("sentinel value should not exist")
	}
	if zero.String() != "" {
		t.Fatalf("String() should be empty, got %q", zero.String())
	}

	s := StrTo("7")
	if !s.Exist() || s.String() != "7" {
		t.Fatalf("expected valid StrTo")
	}
	if s.MustInt() != 7 {
		t.Fatalf("MustInt() expected 7")
	}
	if s.MustInt64() != 7 {
		t.Fatalf("MustInt64() expected 7")
	}
	if s.MustUint8() != 7 {
		t.Fatalf("MustUint8() expected 7")
	}
	if s.MustFloat64() != 7 {
		t.Fatalf("MustFloat64() expected 7")
	}
}

func TestToStr(t *testing.T) {
	if got := ToStr(true); got != "true" {
		t.Fatalf("ToStr(true)=%q", got)
	}
	if got := ToStr(float32(3.14159), 2, 32); got != "3.14" {
		t.Fatalf("ToStr(float32)=%q", got)
	}
	if got := ToStr(int64(255), 16); got != "ff" {
		t.Fatalf("ToStr(int64, base 16)=%q", got)
	}
	if got := ToStr([]byte("gin")); got != "gin" {
		t.Fatalf("ToStr([]byte)=%q", got)
	}
}

func TestHexConversions(t *testing.T) {
	cases := map[string]int{
		"0":    0,
		"1a3f": 0x1a3f,
		"ff":   255,
	}
	for input, want := range cases {
		got, err := HexStr2int(input)
		if err != nil || got != want {
			t.Fatalf("HexStr2int(%q)=%d, err=%v", input, got, err)
		}
		if back := Int2HexStr(want); back != input {
			t.Fatalf("Int2HexStr(%d)=%q want %q", want, back, input)
		}
	}

	if _, err := HexStr2int("1g"); err == nil {
		t.Fatal("expected error for invalid hex")
	}
}

func TestPowInt(t *testing.T) {
	if got := PowInt(2, 5); got != 32 {
		t.Fatalf("PowInt(2,5)=%d", got)
	}
	if got := PowInt(2, 0); got != 1 {
		t.Fatalf("PowInt(2,0)=%d", got)
	}
	if got := PowInt(2, -3); got != 1 {
		t.Fatalf("PowInt negative exponent should return 1, got %d", got)
	}
	if got := PowInt(0, 3); got != 0 {
		t.Fatalf("PowInt(0,3)=%d", got)
	}
}
