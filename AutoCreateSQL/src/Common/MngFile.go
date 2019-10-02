package Common

import "strings"

func TrimRightForCuset(s, cuset string) string {
	var str string
	if len(s) == 0 {
		return str
	}
	sConcat := strings.Split(s,cuset)
	for i:= 0 ; i < len(sConcat) - 1 ; i++ {
		str += sConcat[i]
	}
	return str
}
