package sum

import "testing"

func TestSum(t *testing.T) {
	counts := map[string]int{
		"https://example.com": 2,
		"https://example.org": 3,
		"https://example.net": 5,
	}

	got := Sum(counts)
	want := 10
	if got != want {
		t.Fatalf("Sum(%v) = %d, want %d", counts, got, want)
	}
}

func TestSumEmptyMap(t *testing.T) {
	got := Sum(map[string]int{})
	if got != 0 {
		t.Fatalf("Sum(empty map) = %d, want 0", got)
	}
}
