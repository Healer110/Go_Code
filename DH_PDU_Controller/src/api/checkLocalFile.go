package api

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CheckIPPortFile(path string) (string, int) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// 文件不存在
			return "", -1
		}
	}
	file, _ := os.Open(path)
	defer file.Close()
	content, _ := io.ReadAll(file)
	res := strings.Split(strings.Trim(string(content), " \n"), ":")
	ip := res[0]
	port, _ := strconv.Atoi(res[1])
	return ip, port
}

func SaveIPPoort(ip, port, path string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		PrintErrors(fmt.Sprintf("open file error：%v\n", path))
		return
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		PrintErrors(fmt.Sprintf("write file error：%v\n", path))
	}
}
