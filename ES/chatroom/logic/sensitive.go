package logic

import "strings"

func FilterSensitive(content string) string {
	for _, word := range global.SensitiveWord {
		content = strings.ReplaceAll(content, word, "**")
	}
	return content

}
