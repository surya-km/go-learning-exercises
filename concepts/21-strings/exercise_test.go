package stringspractice

import "testing"

func TestMaskAccount(t *testing.T) {
	tests := []struct {
		number string
		want   string
	}{
		{number: "1234567890", want: "******7890"},
		{number: "4111111111111111", want: "************1111"},
		{number: "12345", want: "*2345"},
		{number: "GB29NWBK60161331926819", want: "******************6819"},
		{number: "1234", want: "1234"},
		{number: "12", want: "12"},
		{number: "", want: ""},
	}

	for _, tt := range tests {
		t.Run("number_"+tt.number, func(t *testing.T) {
			got := MaskAccount(tt.number)
			if got != tt.want {
				t.Errorf("MaskAccount(%q) = %q, want %q", tt.number, got, tt.want)
			}
		})
	}
}

func TestShortID(t *testing.T) {
	tests := []struct {
		id   string
		want string
	}{
		{id: "txn_9f8e7d6c5b4a", want: "txn_...5b4a"},
		{id: "abcdefghijkl", want: "abcd...ijkl"},
		{id: "ch_3MmlLrLkdIwHu7ix0snN0B15", want: "ch_3...0B15"},
		{id: "abcdefghijk", want: "abcdefghijk"},
		{id: "ab", want: "ab"},
		{id: "", want: ""},
	}

	for _, tt := range tests {
		t.Run("id_"+tt.id, func(t *testing.T) {
			got := ShortID(tt.id)
			if got != tt.want {
				t.Errorf("ShortID(%q) = %q, want %q", tt.id, got, tt.want)
			}
		})
	}
}
