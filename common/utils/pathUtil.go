package utils

import (
	"regexp"
	"strings"
)

// url域名正侧
const domainRegex string = `^(https?):\/\/([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,6}(\/.*)?$`

// 检查域名格式是否合法
func IsValidHost(hostPath string) bool {
	reg, err := regexp.Compile(domainRegex)
	if err != nil {
		return false
	}
	return reg.MatchString(hostPath)
}

// 拼接域名和相对路径
func BindUrl(hostname string, path ...string) string {
	if hostname == "" {
		return strings.TrimLeft(strings.Join(path, "/"), "/")
	}
	if len(path) == 0 {
		return strings.TrimRight(hostname, "/")
	}
	pathBuilder := &strings.Builder{}
	for _, v := range path {
		newV := strings.TrimRight(v, "/")
		if !strings.HasPrefix(newV, "/") {
			pathBuilder.WriteString("/")
		}
		pathBuilder.WriteString(newV)
	}
	return strings.TrimRight(hostname, "/") + pathBuilder.String()
}

// 替换资源域名
func ReplaceHost(src string, replaceHost string) string {
	if src == "" || replaceHost == "" {
		return src
	}
	if !IsValidHost(src) || !IsValidHost(replaceHost) {
		return src
	}
	reg, err := regexp.Compile(domainRegex)
	if err != nil {
		return src
	}
	return reg.ReplaceAllString(src, strings.TrimRight(replaceHost, "/"))
}
