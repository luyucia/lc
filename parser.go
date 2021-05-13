package main

import (
	"regexp"
	"strings"
)

var didireg = regexp.MustCompile(`\\[{.*?}\\]`)

func parseBasic(s string) []string {
	i := 0
	var result []string
	tmp := ""

	for {
		if i >= len(s) {
			break
		}
		if s[i] == '[' {

		} else if s[i] == ']' {
			result = append(result, tmp)
			tmp = ""
		} else {
			tmp += string(s[i])
		}
		i++
	}
	result = append(result, tmp)
	return result

}

func parserInfo(kv []string) map[string]string {
	l := len(kv)
	i := 0
	infos := map[string]string{}
	for {
		if i >= l-1 {
			break
		}
		kvarr := strings.Split(kv[i], "=")
		infos[kvarr[0]] = kvarr[1]
		i++
	}
	return infos
}

func Dparser(log string) LogInfo {

	l := LogInfo{}
	arr := strings.Split(log, "||")
	if len(arr) < 3 {
		println("解析日志错误:", log)
		return l
	}

	basic := parseBasic(arr[0])
	info := parserInfo(arr[1:])
	msg := arr[len(arr)-1]
	l.Level = basic[0]
	l.Time = basic[1]
	l.FileLine = basic[2]
	l.Tag = basic[3]
	l.TraceId = info["traceid"]
	l.Host = info["host"]
	l.ClientIp = info["from"]
	l.SpanId = info["spanid"]
	l.Uri = info["uri"]
	l.Method = info["method"]
	l.Params = info["params"]
	l.Msg = msg

	return l

}
