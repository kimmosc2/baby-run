package tui

import (
	"fmt"
	"time"
)


func LoadingText(t time.Time) {
	for {
		for _, v := range `-\|/` {
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("\r压测中...%v", string(v))
		}
	}
}
