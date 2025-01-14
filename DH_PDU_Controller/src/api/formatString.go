package api

import "strings"

// 格式化返回值
func FormatResponse(res string) string {
	res = strings.Trim(res, " VA\n")
	return res
}
