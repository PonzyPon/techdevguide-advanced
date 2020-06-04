package decompression

import (
	"strings"
	"testing"
)

func TestDecompress1(t *testing.T) {
	// arrange
	want := "abcabcabcababababc"
	argument := "3[abc]4[ab]c"
	// act
	got := Decompress(argument)
	// assert
	if got != want {
		t.Errorf("decompress(%q) = %q; want %q", argument, got, want)
	}
}

func TestDecompress2(t *testing.T) {
	// arrange
	want := "aaaaaaaaaa"
	argument := "10[a]"
	// act
	got := Decompress(argument)
	// assert
	if got != want {
		t.Errorf("decompress(%q) = %q; want %q", argument, got, want)
	}
}

func TestDecompress3(t *testing.T) {
	// arrange
	want := "aaabaaab"
	argument := "2[3[a]b]"
	// act
	got := Decompress(argument)
	// assert
	if got != want {
		t.Errorf("decompress(%q) = %q; want %q", argument, got, want)
	}
}

func TestDecompress4(t *testing.T) {
	// arrange
	want := "aaa"
	argument := "aaa"
	// act
	got := Decompress(argument)
	// assert
	if got != want {
		t.Errorf("decompress(%q) = %q; want %q", argument, got, want)
	}
}

func TestDecompress5(t *testing.T) {
	// arrange
	want := "aabcaabcd"
	argument := "2[aabc]d"
	// act
	got := Decompress(argument)
	// assert
	if got != want {
		t.Errorf("decompress(%q) = %q; want %q", argument, got, want)
	}
}

func TestDecompress6(t *testing.T) {
	// arrange
	want := "abbbabbbcabbbabbbc"
	argument := "2[2[abbb]c]"
	// act
	got := Decompress(argument)
	// assert
	if got != want {
		t.Errorf("decompress(%q) = %q; want %q", argument, got, want)
	}
}

func TestDecompress7(t *testing.T) {
	// arrange
	want := ""
	argument := "0[abc]"
	// act
	got := Decompress(argument)
	// assert
	if got != want {
		t.Errorf("decompress(%q) = %q; want %q", argument, got, want)
	}
}

func TestDecompress8(t *testing.T) {
	// arrange
	want := "ab"
	argument := "a[]b"
	// act
	got := Decompress(argument)
	// assert
	if got != want {
		t.Errorf("decompress(%q) = %q; want %q", argument, got, want)
	}
}

func TestTokenize1(t *testing.T) {
	// arrange
	want := []string{"3", "[", "abc", "]", "4", "[", "ab", "]", "c"}
	argument := "3[abc]4[ab]c"
	// act
	got := tokenize(argument)
	// assert
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("tokenize(%q) = %q; want %q", argument, got, want)
	}
}
