package testMain

import (
	"encoding/json"
	"fmt"
	"os"
	"practice/simpleSpider/rule"
	"practice/simpleSpider/spiderServer"
	"strconv"
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
	con := make([]byte, 0)
	_, err = fs.Read(con)
	if err != nil {
		b.Error("read file error:", err)
	}
	var text = strings.TrimSpace(string(con))
	var okCount, badCount int
	// 将处理后的数据，从通道中取出
	go func() {
		for {
			select {
			case resData := <-server.DataChan:
				_, err := json.Marshal(resData)
				checkError(err)
				okCount++
			case <-time.NewTicker(time.Duration(1) * time.Second).C:
				b.Error("wait parsed result had timeout.")
				fmt.Println("wait parsed result had timeout.")
				badCount++
			}
		}
		fmt.Println("okCount is:", strconv.Itoa(okCount), "badCount is:", strconv.Itoa(badCount))
	}()
	for i := 0; i < b.N; i++ {
		server.HandleReceive(text, rule.GetAreaRule(), rule.GetItemRule())
		fmt.Println("测试完成》》》》》》》》》》》》》》》》》》》》》》》》")
	}
	//select {}
}

func checkError(err error) {
	if err != nil {
		panic(err)
		return
	}
}
