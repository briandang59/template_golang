package utils

import (
	"net/url"
	"strings"
	"unicode"
)

func kebabToPascal(input string) string {
	parts := strings.Split(input, "-")
	for i, part := range parts {
		if len(part) > 0 {
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			parts[i] = string(runes)
		}
	}
	return strings.Join(parts, "")
}

func capitalizeEachPart(path string) string {
	parts := strings.Split(path, ".")
	for i, part := range parts {
		parts[i] = kebabToPascal(part)
	}
	return strings.Join(parts, ".")
}

func ParsePopulateQuery(values url.Values) []string {
	preloadMap := make(map[string]bool)

	for key := range values {
		if strings.HasPrefix(key, "populate[") {
			str := strings.TrimPrefix(key, "populate[")
			str = strings.TrimSuffix(str, "]")
			str = strings.ReplaceAll(str, "][", ".")

			normalized := capitalizeEachPart(str)
			preloadMap[normalized] = true
		}
	}

	for _, populateValue := range values["populate"] {
		normalized := capitalizeEachPart(populateValue)
		preloadMap[normalized] = true
	}

	var result []string
	for key := range preloadMap {
		result = append(result, key)
	}
	return result
}
