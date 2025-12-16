package main

import "testing"

func TestRenderASCIIConcatenatesLines(t *testing.T) {
	banner := map[rune][]string{
		'A': {
			"A0", "A1", "A2", "A3", "A4", "A5", "A6", "A7",
		},
		'B': {
			"B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7",
		},
	}

	got := RenderASCII("AB", banner)

	var want string
	for i := 0; i < height; i++ {
		want += banner['A'][i] + banner['B'][i] + "\n"
	}

	if got != want {
		t.Fatalf("unexpected render output:\nwant:\n%q\ngot:\n%q", want, got)
	}
}

func TestRenderASCIILiteralNewline(t *testing.T) {
	banner := map[rune][]string{
		'A': {"A0", "A1", "A2", "A3", "A4", "A5", "A6", "A7"},
		'B': {"B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7"},
	}

	got := RenderASCII(`A\nB`, banner)

	var want string
	for i := 0; i < height; i++ {
		want += banner['A'][i] + "\n"
	}
	for i := 0; i < height; i++ {
		want += banner['B'][i] + "\n"
	}

	if got != want {
		t.Fatalf("unexpected render output for literal newline:\nwant:\n%q\ngot:\n%q", want, got)
	}
}

func TestLoadBannerCoversPrintableRange(t *testing.T) {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}
	if len(banner) != runeCount {
		t.Fatalf("expected %d runes, got %d", runeCount, len(banner))
	}
	for r := firstRune; r <= lastRune; r++ {
		lines, ok := banner[rune(r)]
		if !ok {
			t.Fatalf("missing rune %q", rune(r))
		}
		if len(lines) != height {
			t.Fatalf("rune %q has %d lines, want %d", rune(r), len(lines), height)
		}
	}
}
