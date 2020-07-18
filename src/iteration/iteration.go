package iteration

import "time"

func Repeat(c string) string {
	var repeated string
	for i := 0; i < 6; i++ {
		repeated = repeated + c
		time.Sleep(1 * time.Second)
	}
	return repeated
}
