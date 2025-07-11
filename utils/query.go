package utils

import (
	"net/url"
	"strings"
	"unicode"
)

// Chuyển từ "equipment-type" => "EquipmentType"
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

// Chuyển từ "equipment-type.department-head" => "EquipmentType.DepartmentHead"
func capitalizeEachPart(path string) string {
	parts := strings.Split(path, ".")
	for i, part := range parts {
		parts[i] = kebabToPascal(part)
	}
	return strings.Join(parts, ".")
}

// Phân tích query populate từ URL và trả về slice field preload
// Ví dụ: populate=equipment-type.department-head => []string{"EquipmentType.DepartmentHead"}
func ParsePopulateQuery(values url.Values) []string {
	preloadMap := make(map[string]bool)

	// Hỗ trợ cả 2 dạng: populate[...][...] và populate=...
	for key := range values {
		// 1. Dạng populate[xxx][yyy]
		if strings.HasPrefix(key, "populate[") {
			// parse: populate[departments][manager] -> departments.manager
			str := strings.TrimPrefix(key, "populate[")
			str = strings.TrimSuffix(str, "]")
			str = strings.ReplaceAll(str, "][", ".")

			normalized := capitalizeEachPart(str)
			preloadMap[normalized] = true
		}
	}

	// 2. Dạng populate=xxx.yyy
	for _, populateValue := range values["populate"] {
		normalized := capitalizeEachPart(populateValue)
		preloadMap[normalized] = true
	}

	// Trả về unique preload fields
	var result []string
	for key := range preloadMap {
		result = append(result, key)
	}
	return result
}
