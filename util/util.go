package util

import (
	"encoding/json"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

// GenerateUUID generates a new UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// IsValidUUID validates UUID format
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// ToSnakeCase converts string to snake_case
func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// ToCamelCase converts string to camelCase
func ToCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

// SanitizeFilename sanitizes a filename
func SanitizeFilename(filename string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9._-]`)
	return reg.ReplaceAllString(filename, "_")
}

// GetFileExtension returns file extension
func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

// IsAllowedFileType checks if file type is allowed
func IsAllowedFileType(filename string, allowedTypes []string) bool {
	ext := GetFileExtension(filename)
	for _, t := range allowedTypes {
		if "."+t == ext || t == ext {
			return true
		}
	}
	return false
}

// FormatFileSize formats file size to human readable
func FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return string(rune(size)) + " B"
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return string(rune(size/div)) + " " + string("KMGTPE"[exp]) + "B"
}

// ParseDate parses date string to time
func ParseDate(dateStr string) (time.Time, error) {
	layouts := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		"02/01/2006",
		"02-01-2006",
		time.RFC3339,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t, nil
		}
	}
	return time.Time{}, nil
}

// FormatDate formats time to string
func FormatDate(t time.Time, format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return t.Format(format)
}

// Contains checks if slice contains element
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Unique returns unique elements from slice
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// ToJSON converts struct to JSON string
func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// FromJSON converts JSON string to struct
func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

// Ptr returns pointer to value
func Ptr[T any](v T) *T {
	return &v
}

// Deref returns value from pointer with default
func Deref[T any](ptr *T, defaultVal T) T {
	if ptr == nil {
		return defaultVal
	}
	return *ptr
}

// Min returns minimum of two values
func Min[T ~int | ~int64 | ~float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns maximum of two values
func Max[T ~int | ~int64 | ~float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
