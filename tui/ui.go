package tui

import (
	"fmt"
	"time"
)

func LoadingText(ch chan bool) {
	for {
		select {
		case <-ch:
			break
		default:
			for _, v := range `-\|/` {
				time.Sleep(time.Millisecond * 500)
				fmt.Printf("\r压测中...%v", string(v))
			}
		}
	}
}
