package akamai

import (
	"io"
	"os"
	"testing"
)

func TestParseScriptPath(t *testing.T) {
	script, err := os.Open("scripts/index.html")
	if err != nil {
		t.Fatal(err)
	}

	path, err := ParseScriptPath(script)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(path)
}

func BenchmarkParseScriptPath(b *testing.B) {
	script, err := os.Open("scripts/index.html")
	if err != nil {
		b.Fatal(err)
	}
	defer script.Close()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := ParseScriptPath(script)
		if err != nil {
			b.Fatal(err)
		}

		script.Seek(0, io.SeekStart)
	}
}
