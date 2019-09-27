package run

import (
	"baby-run/conf"
	"baby-run/tui"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Result struct {
	sync.Mutex               // 锁
	Counter    uint64        // 总请求数
	Duration   time.Duration // 总消耗时间
	ErrCounter uint64        // 总错误数
}

var BabyRes Result
var counterChan = make(chan uint64)
var durationChan = make(chan time.Duration)
var errCountChan = make(chan uint64)

func Start(c conf.BabyConfig) {
	if me := strings.ToUpper(c.Method); me == http.MethodGet {
		for cnum := c.Client; cnum > 0; cnum-- {
			go func() {
				for {
					if t, ok := Get(c.Url); ok {
						BabyRes.Lock()
						BabyRes.Counter++
						BabyRes.Duration += t
						BabyRes.Unlock()
					} else {
						BabyRes.Lock()
						BabyRes.Counter++
						BabyRes.Duration += t
						BabyRes.ErrCounter++
						BabyRes.Unlock()
					}
				}
			}()
		}
	} else if me == http.MethodPost {
		for cnum := c.Client; cnum > 0; cnum-- {
			go func() {
				for {
					if t, ok := Post(c.Url, c.ContentType, c.Body); ok {
						BabyRes.Lock()
						BabyRes.Counter++
						BabyRes.Duration += t
						BabyRes.Unlock()
					} else {
						BabyRes.Lock()
						BabyRes.Counter++
						BabyRes.Duration += t
						BabyRes.ErrCounter++
						BabyRes.Unlock()
					}
				}
			}()
		}
	} else {
		log.Fatal("no such method:", c.Method)
	}

	var ch = make(chan bool)
	go tui.LoadingText(ch)
	select {
	case <-time.After(time.Duration(int64(c.Times)) * time.Second):
		ch <- true
		break
	}

	fmt.Printf("\r== Result ============================================\n")
	fmt.Printf("总协程数:%d\n", c.Client)
	fmt.Printf("单协程持续时间:%vs\n", c.Times)
	fmt.Printf("总请求次数:%d\n", BabyRes.Counter)
	fmt.Printf("成功数:%d\n", BabyRes.Counter-BabyRes.ErrCounter)
	fmt.Printf("失败数:%d\n", BabyRes.ErrCounter)
	fmt.Printf("成功占比:%.2f%%\n", float64((BabyRes.Counter-BabyRes.ErrCounter)/BabyRes.Counter)*100)
	fmt.Printf("总时间:%v\n", BabyRes.Duration)
	fmt.Printf("平均时间:%vms\n", uint64(BabyRes.Duration.Nanoseconds()/1e6)/BabyRes.Counter)
	fmt.Printf("\r=======================================================\n")
}

// 如果返回true代表请求成功,否则false
func Get(url string) (time.Duration, bool) {
	start := time.Now()
	resp, err := http.Get(url)
	return time.Now().Sub(start), err == nil && resp.StatusCode == http.StatusOK
}

func Post(url, contentType, body string) (time.Duration, bool) {
	start := time.Now()
	resp, err := http.Post(url, contentType, strings.NewReader(body))
	return time.Now().Sub(start), err == nil && resp.StatusCode == http.StatusOK
}
