package srcgen

import (
	"strings"
	"unicode"
)

func isUpperCase(str string) bool {
	for _, ch := range str {
		if !unicode.IsUpper(ch) {
			return false
		}
	}
	return true
}

func splitCase(str string) []string {
	if len(str) == 0 {
		return nil
	}

	var split []string

	separators := []string{" ", "_", "-"}

	for _, sep := range separators {
		if strings.Contains(str, sep) {
			split = strings.Split(str, sep)
			break
		}
	}

	if split == nil {
		if isUpperCase(str) {
			split = []string{str}
		} else {
			curr := ""

			for _, ch := range str {
				if unicode.IsUpper(ch) && len(curr) != 0 {
					split = append(split, curr)
					curr = ""
				}

				curr += string(ch)
			}

			if len(curr) != 0 {
				split = append(split, curr)
			}
		}
	}

	var splitNotEmpty []string

	for _, it := range split {
		if len(it) != 0 {
			splitNotEmpty = append(splitNotEmpty, it)
		}
	}

	return splitNotEmpty
}

func joinToSnakeCase(elms []string) string {
	return strings.ToLower(strings.Join(elms, "_"))
}

func joinToKebabCase(elms []string) string {
	return strings.ToLower(strings.Join(elms, "-"))
}

func joinToCamelCase(elms []string) (str string) {
	for _, elm := range elms {
		if len(str) == 0 {
			str = strings.ToLower(elm)
		} else {
			for i, ch := range elm {
				if i == 0 {
					str += string(unicode.ToUpper(ch))
				} else {
					str += string(unicode.ToLower(ch))
				}
			}
		}
	}

	return
}

func joinToGoCamelCase(elms []string) (str string) {
	for _, elm := range elms {
		if len(str) == 0 {
			str = strings.ToLower(elm)
		} else if u := strings.ToUpper(elm); commonInitialisms[u] {
			str += u
		} else {
			for i, ch := range elm {
				if i == 0 {
					str += string(unicode.ToUpper(ch))
				} else {
					str += string(unicode.ToLower(ch))
				}
			}
		}
	}

	return
}

func joinToPascalCase(elms []string) (str string) {
	for _, elm := range elms {
		for i, ch := range elm {
			if i == 0 {
				str += string(unicode.ToUpper(ch))
			} else {
				str += string(unicode.ToLower(ch))
			}
		}
	}

	return
}

var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

func joinToGoPascalCase(elms []string) (str string) {
	for _, elm := range elms {
		if u := strings.ToUpper(elm); commonInitialisms[u] {
			str += u
		} else {
			for i, ch := range elm {
				if i == 0 {
					str += string(unicode.ToUpper(ch))
				} else {
					str += string(unicode.ToLower(ch))
				}
			}
		}
	}

	return
}

func ToSnakeCase(str string) string {
	return joinToSnakeCase(splitCase(str))
}

func ToKebabCase(str string) string {
	return joinToKebabCase(splitCase(str))
}

func ToCamelCase(str string) string {
	return joinToCamelCase(splitCase(str))
}

func ToGoCamelCase(str string) string {
	return joinToGoCamelCase(splitCase(str))
}

func ToPascalCase(str string) string {
	return joinToPascalCase(splitCase(str))
}

func ToGoPascalCase(str string) string {
	return joinToGoPascalCase(splitCase(str))
}
