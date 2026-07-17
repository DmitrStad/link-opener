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

func TestSumStrings(t *testing.T) {
	numbers := []string{"5", "10", "15"}
	got, err := SumStrings(numbers)
	if err != nil {
		t.Fatalf("SumStrings(%v) returned error: %v", numbers, err)
	}
	want := 30
	if got != want {
		t.Fatalf("SumStrings(%v) = %d, want %d", numbers, got, want)
	}
}

func TestSumStringsEmpty(t *testing.T) {
	got, err := SumStrings([]string{})
	if err != nil {
		t.Fatalf("SumStrings(empty) returned error: %v", err)
	}
	if got != 0 {
		t.Fatalf("SumStrings(empty) = %d, want 0", got)
	}
}

func TestSumStringsInvalidInput(t *testing.T) {
	numbers := []string{"5", "invalid", "15"}
	_, err := SumStrings(numbers)
	if err == nil {
		t.Fatalf("SumStrings(%v) should return error for invalid input", numbers)
	}
}
