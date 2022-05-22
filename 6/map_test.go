package main

import "testing"

func AccessByte(m map[string]string, k []byte) string {
	return m[string(k)]
}

func AccessString(m map[string]string, k []byte) string {
	kstr := string(k)
	return m[kstr]
}

func BenchmarkAccessMapUsingBytes(b *testing.B) {
	m := map[string]string{
		"foo": "bar",
		"maa": "bow",
	}
	bytes := []byte("foo")
	for i := 0; i < b.N; i++ {
		AccessByte(m, bytes)
	}
}

func BenchmarkAccessMapUsingString(b *testing.B) {
	m := map[string]string{
		"foo": "bar",
		"maa": "bow",
	}
	bytes := []byte("foo")
	for i := 0; i < b.N; i++ {
		AccessString(m, bytes)
	}
}
