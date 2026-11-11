package slicespractice

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		amounts []int
		want    int
	}{
		{name: "several amounts", amounts: []int{100, 250, 75}, want: 425},
		{name: "single amount", amounts: []int{42}, want: 42},
		{name: "includes refunds (negatives)", amounts: []int{500, -200}, want: 300},
		{name: "empty slice", amounts: []int{}, want: 0},
		{name: "nil slice", amounts: nil, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.amounts); got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.amounts, got, tt.want)
			}
		})
	}
}

func TestAbove(t *testing.T) {
	tests := []struct {
		name    string
		amounts []int
		min     int
		want    []int
	}{
		{name: "keeps larger amounts in order", amounts: []int{50, 500, 20, 900}, min: 100, want: []int{500, 900}},
		{name: "equal to min is excluded", amounts: []int{100, 101}, min: 100, want: []int{101}},
		{name: "negative min", amounts: []int{-5, 0, -1}, min: -2, want: []int{0, -1}},
		{name: "nothing matches", amounts: []int{1, 2, 3}, min: 10, want: []int{}},
		{name: "empty input", amounts: nil, min: 0, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Above(tt.amounts, tt.min)
			if got == nil || !slices.Equal(got, tt.want) {
				t.Errorf("Above(%v, %d) = %#v, want %v", tt.amounts, tt.min, got, tt.want)
			}
		})
	}
}

func TestAboveDoesNotModifyInput(t *testing.T) {
	input := []int{5, 50, 500}
	original := slices.Clone(input)
	got := Above(input, 10)
	if len(got) > 0 {
		got[0] = -1
	}
	if !slices.Equal(input, original) {
		t.Errorf("Above modified its input: got %v, want %v", input, original)
	}
}

func TestPages(t *testing.T) {
	tests := []struct {
		name     string
		items    []string
		pageSize int
		want     [][]string
	}{
		{name: "last page shorter", items: []string{"a", "b", "c", "d", "e"}, pageSize: 2, want: [][]string{{"a", "b"}, {"c", "d"}, {"e"}}},
		{name: "fits exactly", items: []string{"a", "b", "c", "d"}, pageSize: 2, want: [][]string{{"a", "b"}, {"c", "d"}}},
		{name: "one page", items: []string{"a", "b"}, pageSize: 5, want: [][]string{{"a", "b"}}},
		{name: "page size one", items: []string{"a", "b", "c"}, pageSize: 1, want: [][]string{{"a"}, {"b"}, {"c"}}},
		{name: "no items", items: nil, pageSize: 3, want: [][]string{}},
		{name: "zero page size", items: []string{"a"}, pageSize: 0, want: [][]string{}},
		{name: "negative page size", items: []string{"a"}, pageSize: -2, want: [][]string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pages(tt.items, tt.pageSize)
			if got == nil || !slices.EqualFunc(got, tt.want, slices.Equal) {
				t.Errorf("Pages(%q, %d) = %q, want %q", tt.items, tt.pageSize, got, tt.want)
			}
		})
	}
}

func TestPagesAreIndependentCopies(t *testing.T) {
	items := []string{"tx1", "tx2", "tx3"}
	pages := Pages(items, 2)
	if len(pages) == 0 || len(pages[0]) == 0 {
		t.Fatalf("Pages(%q, 2) = %q, want two pages", items, pages)
	}
	pages[0][0] = "CHANGED"
	if items[0] != "tx1" {
		t.Errorf("changing pages[0][0] also changed items[0] to %q; each page should be a copy", items[0])
	}
}
