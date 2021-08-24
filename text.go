package eyes

import "fmt"

func Format(v ...interface{}) (s string) {
	for key, val := range v {
		if key%2 == 0 {
			s += fmt.Sprintf("%+v", val) + `=`
		} else {
			s += fmt.Sprintf("%+v", val) + " "
		}
	}
	return
}
