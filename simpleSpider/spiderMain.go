package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"practice/simpleSpider/spiderServer"
	"time"
)

var server = spiderServer.NewNovelServer()

func main() {
	var port = flag.String("port", "8001", "Parse service serve port.such as 8001")
	var debug = flag.Bool("debug", false, "Value 'true' is debug model")
	var help = flag.String("help", "", "Parse service helper")
	if (*help) == "" {
		fmt.Fprintf(os.Stdout,
			"Usage: \n %s -port=8001 -debug=false \n", os.Args[0])
		os.Exit(0)
	}
	if (*debug) == true {
		fmt.Println("debug model...")
	}
	flag.Parse()
	//启动一个http服务器，接收请求的数据
	//将数据传递给server
	http.HandleFunc("/", safeHandler(ParseHandler))
	fmt.Println("parse service will run in port:", *port)
	err := http.ListenAndServe(":"+(*port), nil)
	if err != nil {
		log.Println("http server listen error:", err)
	}
}

//解析请求参数
func ParseHandler(w http.ResponseWriter, r *http.Request) {
	//设定等待解析结果的超时时间
	var timeout int64 = 1
	defer r.Body.Close()
	if r.Method != "POST" {
		http.Error(w, "http method not allow.", http.StatusMethodNotAllowed)
		return
	}
	areaRule := r.FormValue("areaRule")
	if len(areaRule) < 1 {
		areaRule = getAreaRule()
	}
	itemRuleStr := r.FormValue("itemRule")
	fmt.Println(itemRuleStr)
	itemRule := map[string]string{}
	if len(itemRuleStr) <= 0 {
		itemRule = getItemRule()
	} else {
		err := json.Unmarshal([]byte(itemRuleStr), &itemRule)
		fmt.Println(itemRule)
		checkError(err)
	}
	dataStr := r.FormValue("dataStr")
	//解析数据，放入spider server处理
	server.HandleReceive(dataStr, areaRule, itemRule)
	select {
	case resData := <-server.DataChan:
		_, err := w.Write([]byte(resData.Text))
		checkError(err)
	case <-time.NewTicker(time.Duration(timeout) * time.Second).C:
		panic(errors.New("wait parsed result had timeout."))
	}
}

//异常发生时的处理
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("WARN:panic in %v-%v", fn, err)
			}
		}()
		fn(w, r)
	}
}

func Handle(conn net.Conn, server spiderServer.NovelServer) {
	var (
		//bNum int
		dataStr string
	)
	defer (conn).Close()
	//var data = make([]byte, 1024)
	var bufReader = bufio.NewReader(conn)
	for {
		tmpStr, err := bufReader.ReadString('\n')
		rAddr := conn.RemoteAddr()
		fmt.Printf("Receive from client %s : %s \n", rAddr, dataStr)
		//bNum,err = (conn).Read(data)
		if err != nil {
			log.Println("tcp server Accept error:", err)
			break
		}
		dataStr = dataStr + tmpStr
		if tmpStr == "quit" {
			break
		}

	}
	fmt.Println(dataStr)
	//解析数据，放入spider server处理
	server.HandleReceive(dataStr, getAreaRule(), getItemRule())
}

func getAreaRule() string {
	areaRule := "#list"
	areaRule = base64.StdEncoding.EncodeToString([]byte(areaRule))
	return areaRule
}

func getItemRule() map[string]string {
	var ruleArr = map[string]string{
		"0": "dd a",
	}
	for index, con := range ruleArr {
		newRule := base64.StdEncoding.EncodeToString([]byte(con))
		ruleArr[index] = string(newRule)
	}
	return ruleArr
}

func checkError(err error) {
	if err != nil {
		panic(err)
		return
	}
}

//var chTask = make(chan beegoOrm.FictionOneOfList)

// 入口函数
//func main() {
//	var wg sync.WaitGroup
//	listUrl := "https://www.biduo.cc/biquge/17_17308"
//	//url := "https://www.biduo.cc/biquge/17_17308/c8698877.html"
//	//status,err := testSpider.GetDetail(url)
//	testSpider.GetBaseUrl(listUrl)
//	go func(wg sync.WaitGroup) {
//		wg.Add(1)
//		status, _, err := testSpider.GetList(chTask, listUrl)
//		if status != nil {
//			log.Fatal(err)
//		}
//		logs.Info(err)
//		wg.Done()
//	}(wg)
//	// 将每一个任务数据放入数据库，
//	// 然后放入任务队列中，获取具体的详情页内容
//	go testSpider.DealTask(chTask,wg)
//
//	wg.Wait()
//}
