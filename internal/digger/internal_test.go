package digger

import "testing"

func BenchmarkExtractOrderNumbers(b *testing.B) {
	filename := `C:\tmp\007\MedgenInput\{00d10210-af6f-4ca9-804d-0ee859e7a8ba}.hl7`
	expected := "203334805"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, err := extractOrderNumber(filename)
		if err != nil {
			b.Fatalf("unexpected error: %s", err.Error())
		}
		if result != expected {
			b.Fatalf("expected: %s, but got: [%s]", expected, result)
		}
	}
}
