package network

import (
	"fmt"
	"github.com/crytjy/jkd-leaf/conf"
	"github.com/crytjy/jkd-leaf/log"
	"net/http"
	"server/common/nlog"
	"server/common/response"
)

var mux *http.ServeMux

func StartHttp(routeHandle func()) {
	// 指定监听地址和端口
	addr := conf.Server.HttpAddr
	if addr != "" {
		// 创建一个新的ServeMux
		mux = http.NewServeMux()

		//注册路由
		routeHandle()

		// 启动HTTP服务器
		fmt.Printf("Server is running on http://%s\n", addr)
		err := http.ListenAndServe(addr, mux)
		if err != nil {
			panic(err)
		}

		log.Debug("Start Http!")
	}
}

func HandleFunc(pattern string, callback func(params any) interface{}) {
	mux.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			writer.WriteHeader(405)
		} else {
			req := request.URL.Query()
			// 调用方法，传入空参数
			res := callback(req)

			rep := response.Success(writer, res)
			saveRes := map[string]interface{}{
				"Req": req,
				"Rep": rep,
			}

			fmt.Println("saveRes", saveRes)

			nlog.Log("Http响应:", saveRes)
		}
	})
}
