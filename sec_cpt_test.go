package akamai

import (
	"strings"
	"testing"
)

func TestParseSecCpt(t *testing.T) {
	input := strings.NewReader(`<iframe id="sec-cpt-if" provider="crypto" class="crypto" data-key="" data-duration=5 src="/_sec/cp_challenge/ak-challenge-4-0.htm"></iframe>`)

	duration, path, err := ParseSecCpt(input)
	if err != nil {
		t.Fatal(err)
	}

	if duration != 5 {
		t.FailNow()
	}

	if path != "/_sec/cp_challenge/ak-challenge-4-0.htm" {
		t.FailNow()
	}

	t.Log(duration, path)
}

func BenchmarkParseSecCpt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, _, err := ParseSecCpt(strings.NewReader(`<iframe id="sec-cpt-if" provider="crypto" class="crypto" data-key="" data-duration=5 src="/_sec/cp_challenge/ak-challenge-4-0.htm"></iframe>`)); err != nil {
			b.Error(err)
		}
	}
}
