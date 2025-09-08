package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	mapping := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001F50D': "search",
		'\u2600':     "weather",
	}
	for _, c := range log {
		if _, exists := mapping[c]; exists {
			return mapping[c]
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""
	for _, c := range log {
		if c == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(c)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
