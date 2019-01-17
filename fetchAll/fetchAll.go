package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	// 监控时间跨度
	start := time.Now()

	// 创建一个管道用于协程之间通信
	ch := make(chan string)

	for num := 0; num < 2; num++ {

		// 开启多个协程并发访问web资源
		for _, v := range os.Args[1:] {

			go fetch(v, ch)

		}

		// 监听通道输出，并阻塞主协程
		for range os.Args[1:] {

			fmt.Println(<-ch)

		}

	}

	// 程序结束时输出程序运行时间
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

// 获取网络数据
func fetch(url string, ch chan<- string) {

	start := time.Now()

	res, err := http.Get(url)

	if err != nil {

		ch <- fmt.Sprint(err)

		return

	}

	defer res.Body.Close()

	webLen, err := io.Copy(ioutil.Discard, res.Body)

	if err != nil {

		ch <- fmt.Sprintf("while reading %s: %v", url, err)

		return

	}

	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), webLen, url)

}
