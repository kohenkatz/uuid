package uuid

import (
	"testing"
)

func BenchmarkGetTimeFromTimestamp(b *testing.B) {
	u, err := NewV7()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t, _ := TimestampFromV7(u)
		t.Time()
	}
}

func BenchmarkGetNativeTime(b *testing.B) {
	u, err := NewV7()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TimeFromV7(u)
	}
}
