// Package runes practices the runes concept. Read CONCEPT.md first.
package runes

import "unicode"

// Truncate shortens s to at most max characters (runes). If s is too long,
// it keeps max-1 characters and adds "…". It returns "" if max <= 0.
func Truncate(s string, max int) string {
	if max <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= max {
		return s
	}
	return string(runes[:max-1]) + "…"
}

// Initials returns the upper-cased first character of each whitespace-separated word.
func Initials(name string) string {
	var initials []rune
	atWordStart := true
	for _, r := range name {
		if unicode.IsSpace(r) {
			atWordStart = true
			continue
		}
		if atWordStart {
			initials = append(initials, unicode.ToUpper(r))
			atWordStart = false
		}
	}
	return string(initials)
}
