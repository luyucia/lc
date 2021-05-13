package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"golang.org/x/net/websocket"
	"io"
	"os"
)

type LogInfo struct {
	Level    string `json:"level"`
	Time     string `json:"time"`
	FileLine string `json:"file_line"`
	Tag      string `json:"tag"`
	TraceId  string `json:"trace_id"`
	SpanId   string `json:"span_id"`
	Method   string `json:"method"`
	Host     string `json:"host"`
	Uri      string `json:"uri"`
	Params   string `json:"params"`
	ClientIp string `json:"client_ip"`
	Msg      string `json:"msg"`
}

type App struct {
	filepath   string
	configPath string
	server     *echo.Echo
	watcher    *fsnotify.Watcher
	logbuff    []LogInfo
	logbuffLen int
}

var A App
var newLineReady = make(chan int)

func parseParam() {
	flag.StringVar(&A.configPath, "c", ".", "config path")
	flag.StringVar(&A.filepath, "f", "test.log", "log path")
	flag.Parse()
}

func loadConfig() {
	viper.SetConfigName("config.ini")
	viper.SetConfigType("ini")
	viper.AddConfigPath(A.configPath)
	viper.ReadInConfig()

	//println(viper.GetString("color.red"))
}

func startServer() {
	A.server = echo.New()
	A.server.Use(middleware.Logger())
	A.server.Use(middleware.Recover())
	A.server.Static("/", "./web")
	A.server.GET("/ws", hello)
	A.server.Logger.Fatal(A.server.Start(":8022"))
}

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		//for {
		// Write
		//err := websocket.Message.Send(ws, "Hello, Client!")
		//if err != nil {
		//	log.Fatal(err)
		//}

		//// Read
		//msg := ""
		//err = websocket.Message.Receive(ws, &msg)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("%s\n", msg)

		//fileInfo, _ := os.Stdin.Stat()
		//if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		//	log.Fatal("The command is intended to work with pipes.")
		//}
		//s := bufio.NewScanner(os.Stdin)
		//
		//for s.Scan() {
		//	str := s.Text()
		//	err := websocket.Message.Send(ws, str)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//}

		//}

		//initLogReader(func(line []byte) {
		//	err := websocket.Message.Send(ws, string(line))
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//})
		offset := 0
		for {
			if offset < A.logbuffLen {
				sendData, _ := json.Marshal(A.logbuff[offset])
				err := websocket.Message.Send(ws, string(sendData))
				offset++
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

//func initLogReader(hookfn func([]byte)) {
//
//	f, err := os.Open(A.filepath)
//	if err != nil {
//		println(err)
//	}
//
//	bfRd := bufio.NewReader(f)
//	for {
//		line, err := bfRd.ReadBytes('\n')
//		hookfn(line)
//		if err == io.EOF {
//			<-newLineReady
//		}
//		if err != nil {
//			println(err)
//		}
//	}
//}

func logReaderStart() {
	f, err := os.Open(A.filepath)
	if err != nil {
		println(err)
	}
	bfRd := bufio.NewReader(f)
	for {
		//从头读到尾部
		line, err := bfRd.ReadBytes('\n')
		A.logbuff = append(A.logbuff, Dparser(string(line)))
		A.logbuffLen++
		//读到末尾，阻塞，收到文件变更消息，则继续读取
		if err == io.EOF {
			<-newLineReady
		}
		if err != nil {
			println(err)
		}
	}

}

func watcherStart() {

	var err error
	A.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		println(err)
	}

	err = A.watcher.Add(A.filepath)
	if err != nil {
		println(err)
	}

	for {
		select {
		case event := <-A.watcher.Events:
			{
				if event.Op&fsnotify.Write == fsnotify.Write {
					newLineReady <- 1
				}
			}
		}
	}

}

func main() {

	parseParam()
	loadConfig()
	go watcherStart()
	go logReaderStart()
	startServer()

}
