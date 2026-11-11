// Package stringspractice practices the strings concept. Read CONCEPT.md first.
package stringspractice

// MaskAccount replaces every character except the last 4 with "*".
// Numbers with 4 or fewer characters are returned unchanged.
func MaskAccount(number string) string {
	if len(number) <= 4 {
		return number
	}
	visible := number[len(number)-4:]
	masked := ""
	for i := 0; i < len(number)-4; i++ {
		masked += "*"
	}
	return masked + visible
}

// ShortID returns the first 4 characters, "...", and the last 4 characters of id.
// IDs with 11 or fewer characters are returned unchanged.
func ShortID(id string) string {
	if len(id) <= 11 {
		return id
	}
	return id[:4] + "..." + id[len(id)-4:]
}
