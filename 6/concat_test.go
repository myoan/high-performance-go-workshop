package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func concatA(id, addr string, ts time.Time) string {
	s := id
	s += " " + addr
	s += " " + ts.String()
	return s
}

func concatB(id, addr string, ts time.Time) string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s %s %v", id, addr, ts)
	return b.String()
}

func concatC(id, addr string, ts time.Time) string {
	return fmt.Sprintf("%s %s %v", id, addr, ts)
}

func concatD(id, addr string, ts time.Time) string {
	b := make([]byte, 0, 40)
	b = append(b, id...)
	b = append(b, ' ')
	b = append(b, addr...)
	b = append(b, ' ')
	b = ts.AppendFormat(b, "2006-01-02 15:04:05.9999999 -0700 MST")
	return string(b)
}

func concatE(id, addr string, ts time.Time) string {
	var b strings.Builder
	b.WriteString(id)
	b.WriteString(" ")
	b.WriteString(addr)
	b.WriteString(" ")
	b.WriteString(ts.String())
	return b.String()
}

func BenchmarkConcatA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatA("id", "addr", time.Now())
	}
}

func BenchmarkConcatB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatB("id", "addr", time.Now())
	}
}

func BenchmarkConcatC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatC("id", "addr", time.Now())
	}
}

func BenchmarkConcatD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatD("id", "addr", time.Now())
	}
}

func BenchmarkConcatE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatE("id", "addr", time.Now())
	}
}
