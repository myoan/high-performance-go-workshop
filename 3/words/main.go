package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

/*
// ❯ time ./words moby.txt
// "moby.txt": 181320 words
// ./words moby.txt  0.03s user 0.00s system 114% cpu 0.023 total

type ReadBuf struct {
	data []byte
	cur  int
}

func NewReadBuf(r io.Reader) *ReadBuf {
	buf := make([]byte, 5_000_000)
	r.Read(buf)
	return &ReadBuf{
		data: buf,
		cur:  0,
	}
}

func (rb *ReadBuf) Next() (rune, error) {
	if rb.cur >= len(rb.data) {
		return 0, io.EOF
	}
	ret := rune(rb.data[rb.cur])
	rb.cur++
	return ret, nil
}
*/

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	// defer profile.Start().Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}

	b := bufio.NewReader(f)

	words := 0
	inword := false

	// rb := NewReadBuf(b)

	for {
		r, err := readbyte(b)
		// r, err := rb.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}

		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words\n", os.Args[1], words)
}
