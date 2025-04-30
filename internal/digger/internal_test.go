package digger

import "testing"

func BenchmarkExtractOrderNumbers(b *testing.B) {
	filename := `test/test.hl7`
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
