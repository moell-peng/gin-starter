package helper

import (
	"bytes"
	"strings"
	"unicode"
)

func InSliceUint(needle uint, haystack []uint) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

// 驼峰式写法转为下划线写法
func UnderscoreName(name string) string {
	var buffer bytes.Buffer
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteString("_")
			}
			buffer.WriteString(string(unicode.ToLower(r)))
		} else {
			buffer.WriteString(string(r))
		}
	}

	return buffer.String()
}

// 下划线写法转为驼峰写法
func CamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
