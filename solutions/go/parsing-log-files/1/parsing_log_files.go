package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	if len(text) < 5 {
		return false
	}
	re := regexp.MustCompile(`(\[TRC\])|(\[DBG\])|(\[INF\])|(\[WRN\])|(\[ERR\])|(\[FTL\])`)
	return re.MatchString(text[:5])
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(~|\*|=|-)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	var count int
	for _, line := range lines {
		count += len(re.FindAllString(line, -1))
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+[^ ]+`)

	for i, line := range lines {
		res := re.FindString(line)
		if len(res) == 0 {
			continue
		}
		prefix := regexp.MustCompile(`\s+`).Split(res, -1)[1]
		lines[i] = "[USR]" + " " + prefix + " " + line
	}
	return lines
}
