package akamai

import (
	"io"
	"os"
	"testing"
)

func TestParsePixelHtmlVar(t *testing.T) {
	script, err := os.Open("scripts/pixel.html")
	if err != nil {
		t.Fatal(err)
	}

	if v, err := ParsePixelHtmlVar(script); err != nil {
		t.Fatal(err)
	} else {
		t.Log(v)
	}
}

func TestParsePixelScriptURL(t *testing.T) {
	script, err := os.Open("scripts/pixel.html")
	if err != nil {
		t.Fatal(err)
	}

	if scriptUrl, postUrl, err := ParsePixelScriptURL(script); err != nil {
		t.Fatal(err)
	} else {
		t.Log(scriptUrl, postUrl)
	}
}

func BenchmarkParsePixelScriptURL(b *testing.B) {
	script, err := os.Open("scripts/pixel.html")
	if err != nil {
		b.Fatal(err)
	}
	defer script.Close()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := ParsePixelScriptURL(script)
		if err != nil {
			b.Fatal(err)
		}

		script.Seek(0, io.SeekStart)
	}
}

func TestParsePixelScriptVar(t *testing.T) {
	script, err := os.Open("scripts/pixel.js")
	if err != nil {
		t.Fatal(err)
	}

	if v, err := ParsePixelScriptVar(script); err != nil {
		t.Fatal(err)
	} else {
		t.Log(v)
	}
}
