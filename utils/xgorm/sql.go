package xgorm

import "fmt"

func ArgSliceString(strs []string) string {
	str := ""
	for i, v := range strs {
		if i == 0 {
			str = fmt.Sprintf("'%v'", v)
		} else {
			str += "," + fmt.Sprintf("'%v'", v)
		}
	}
	return str
}
