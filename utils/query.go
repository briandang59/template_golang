package utils

import (
	"net/url"
	"strings"
	"unicode"
)

// viết hoa chữ cái đầu của mỗi phần
func capitalizeEachPart(path string) string {
	parts := strings.Split(path, ".")
	for i, part := range parts {
		if len(part) > 0 {
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			parts[i] = string(runes)
		}
	}
	return strings.Join(parts, ".")
}

func ParsePopulateQuery(values url.Values) []string {
	preloadMap := make(map[string]bool)

	for key := range values {
		if strings.HasPrefix(key, "populate[") {
			// parse: populate[departments][manager] -> departments.manager
			str := strings.TrimPrefix(key, "populate[")
			str = strings.TrimSuffix(str, "]")
			str = strings.ReplaceAll(str, "][", ".")

			// viết hoa field đúng với struct Go
			normalized := capitalizeEachPart(str)
			preloadMap[normalized] = true
		}
	}

	var result []string
	for key := range preloadMap {
		result = append(result, key)
	}
	return result
}
