package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"
	"golang.org/x/net/http2"
)

// loggingHandler 是一个http.Handler，它会打印请求的信息，然后调用下一个handler。
type loggingHandler struct {
	next http.Handler
}

func (h *loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s for %s", r.RemoteAddr, r.URL.Path)

	// 创建一个responseWriter来捕获错误
	lrw := &loggingResponseWriter{ResponseWriter: w}

	// 调用下一个handler
	h.next.ServeHTTP(lrw, r)

	// 如果有错误，打印错误信息
	if lrw.statusCode >= 400 {
		log.Printf("Error serving %s: %d", r.URL.Path, lrw.statusCode)
	}
}

// loggingResponseWriter 是一个http.ResponseWriter，它会记录状态码。
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// go build -ldflags="-s -w" -o http2.exe
// http2.exe -d d:\dist -p 8000 -v 2 -c cert.pem -k key.pem
func main() {
	// 设置默认的文件夹路径和端口
	defaultDir, _ := os.Getwd()
	dir := flag.String("d", defaultDir, "启动服务所在的文件目录（默认是当前目录）")
	port := flag.String("p", "8000", "监听的端口（默认是8000）")
	httpVersion := flag.String("v", "1.1 或 2", "http版本（默认是1.1）")
	certPath := flag.String("c", "", "证书文件路径（默认是空）")
	keyPath := flag.String("k", "", "私钥文件路径（默认是空）")

	// 解析命令行参数
	flag.Parse()

	// 创建文件服务器
	fs := http.FileServer(http.Dir(*dir))

	// 创建loggingHandler
	lh := &loggingHandler{next: fs}

	// 设置路由
	http.Handle("/", lh)

	log.Printf("启动HTTP服务端口为 %s ，服务所在的文件目录是 %s \n", *port, *dir)

	if *httpVersion == "2" && *certPath != "" && *keyPath != "" {
		// 启动HTTP/2的服务器
		server := &http.Server{
			Addr:    ":" + *port,
			Handler: lh,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		}
		http2.ConfigureServer(server, &http2.Server{})
		log.Fatal(server.ListenAndServeTLS(*certPath, *keyPath))
	} else {
		// 启动HTTP/1.1的服务器
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}
}
