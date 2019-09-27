package conf

import (
	"errors"
	"flag"
	"os"
)

type BabyConfig struct {
	Client      int
	Url         string
	Times       int
	Body        string
	ContentType string
	Method      string
}

var Config BabyConfig

var (
	client      int    // goroutine数量   默认1
	url         string // 地址            默认空
	times       int    // 测试时间        单位秒
	help        bool
	method      string // 方法            默认GET
	contentType string // multipart/form-data
	body        string // body
)

func init() {
	flag.IntVar(&client, "c", 1, "How many client quantities do you want to open")
	flag.IntVar(&times, "t", 10, "how long do you want to test")
	flag.StringVar(&url, "u", "", "target url")
	flag.StringVar(&method, "m", "GET", "request method")
	flag.StringVar(&contentType, "cType", "multipart/form-data", "content type")
	flag.StringVar(&body, "b", "", "request body")
	flag.BoolVar(&help, "h", false, "baby-run help info")
}

func Init() error {
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if url == "" {
		return errors.New("blank url")
	}

	Config.Client = client
	Config.Times = times
	Config.Url = url
	Config.Method = method
	Config.Body = body
	Config.ContentType = contentType
	return nil
}
