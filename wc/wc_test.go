package wc_test

import (
	"testing"

	"github.com/carantes/wc-go/wc"
)

func TestCountBytes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		content []byte
		want    int
	}{
		{
			name:    "empty",
			content: []byte(""),
			want:    0,
		},
		{
			name:    "one line with eight bytes",
			content: []byte("one line"),
			want:    8,
		},
		{
			name:    "two lines with eighteen bytes",
			content: []byte("one line\ntwo lines"),
			want:    18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wc.CountBytes(tt.content); got != tt.want {
				t.Errorf("CountBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		content []byte
		want    int
	}{
		{
			name:    "empty",
			content: []byte(""),
			want:    0,
		},
		{
			name:    "one line",
			content: []byte("one line"),
			want:    1,
		},
		{
			name:    "two lines",
			content: []byte("one line\ntwo lines"),
			want:    2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wc.CountLines(tt.content); got != tt.want {
				t.Errorf("CountLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		content []byte
		want    int
	}{
		{
			name:    "empty",
			content: []byte(""),
			want:    0,
		},
		{
			name:    "one line with two words",
			content: []byte("two words"),
			want:    2,
		},
		{
			name:    "two lines with four words",
			content: []byte("one two\nthree four"),
			want:    4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wc.CountWords(tt.content); got != tt.want {
				t.Errorf("CountWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountCharacters(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		content []byte
		want    int
	}{
		{
			name:    "empty",
			content: []byte(""),
			want:    0,
		},
		{
			name:    "one line with eight bytes",
			content: []byte("one line"),
			want:    8,
		},
		{
			name:    "two lines with eighteen bytes",
			content: []byte("one line\ntwo lines"),
			want:    18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wc.CountCharacters(tt.content); got != tt.want {
				t.Errorf("CountCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
