package testMain

import (
	"fmt"
	"os"
	"practice/simpleSpider/rule"
	"practice/simpleSpider/spiderServer"
	"strings"
	"testing"
	"time"
)

func Benchmark_server(b *testing.B) {
	b.ResetTimer()
	var server = spiderServer.NewNovelServer()
	fs, err := os.Open("/home/www/go/src/practice/simpleSpider/spiderServer/test.txt")
	if err != nil {
		b.Error("open file error:", err)
	}
	con := make([]byte, 1024000)
	_, err = fs.Read(con)
	if err != nil {
		b.Error("read file error:", err)
	}
	var text = strings.TrimSpace(string(con))
	// 将处理后的数据，从通道中取出
	go func() {
		for {
			select {
			case resData := <-server.DataChan:
				_ = resData
				//fmt.Println("after:", resData)
			case <-time.NewTicker(time.Duration(5) * time.Second).C:
				fmt.Println("wait parsed result had timeout.")
			}
		}
	}()
	for i := 0; i < b.N; i++ {
		server.HandleReceive(text, rule.GetAreaRule(), rule.GetItemRule())
		//fmt.Println("测试完成》》》》》》》》》》》》》》》》》》》》》》》》")
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
		return
	}
}
