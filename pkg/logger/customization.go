package logger

import (
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

// customCallerMarshalFunc tries to marshal a custom function for zerolog.CallerMarshalFunc
func customCallerMarshalFunc(basepath string) func(pc uintptr, file string, line int) string {
	return func(pc uintptr, file string, line int) string {
		parts := strings.Split(file, "/")
		length := len(parts)
		i := slices.Index(parts, basepath)
		if i == notFound {
			return parts[length-1] + ":" + strconv.Itoa(line)
		}

		return strings.Join(parts[i:], "/") + ":" + strconv.Itoa(line)
	}
}
