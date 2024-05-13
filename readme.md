# http2

http2是一个强大而易用的HTTP服务器，它支持HTTP/1.1和HTTP/2。这个服务器不仅可以作为一个静态文件服务器，还可以用于调试和测试。它的设计目标是提供一个简单、快速、可配置的HTTP服务，让开发者可以快速地启动一个HTTP服务，而不需要编写任何代码。它是基于Go语言的net/http包开发的，所以它的性能和稳定性是非常可靠的。

## 特性

- 支持HTTP/1.1和HTTP/2
- 可以通过命令行参数来配置HTTP服务
- 支持TLS，可以通过命令行参数来指定证书和私钥的路径
- 内置日志记录，可以记录每个请求的信息和错误
- 作为一个静态文件服务器，可以通过命令行参数来指定服务所在的文件目录

## 使用

下载http2.exe文件，然后将它放到你的PATH环境变量中。

然后，你可以使用以下命令来启动HTTP服务：

```
http2.exe -h

http2.exe -d d:\dist

http2.exe -d d:\dist -p 8001 -v 2 -c cert.pem -k key.pem
```

在这个命令中，`-d`参数用来指定服务所在的文件目录，`-p`参数用来指定监听的端口，`-v`参数用来指定http版本，`-c`参数用来指定证书文件路径，`-k`参数用来指定私钥文件路径。

## 构建

go build -ldflags="-s -w" -o http2.exe

## 生成证书

生成证书的步骤，打开命令提示符或 PowerShell 窗口，然后执行以下步骤来生成证书和私钥：

1. 生成私钥（key.pem）：
   ```
   openssl genrsa -out key.pem 2048
   ```
2. 生成证书请求（CSR）：
   ```
   openssl req -new -key key.pem -out csr.pem
   ```
   当执行这个命令时，OpenSSL 会要求你输入一些信息。在 "Common Name" 提示时，你应该输入 "localhost"。

3. 使用私钥和 CSR 生成证书（cert.pem）：
   ```
   openssl x509 -req -days 3650 -in csr.pem -signkey key.pem -out cert.pem
   ```

## 贡献

欢迎任何形式的贡献，包括报告问题、提供反馈、改进代码等。如果你想对这个项目做出贡献，你可以通过GitHub上的issue和pull request来进行交流。

## 许可

这个项目使用MIT许可证，详细信息请见[LICENSE](LICENSE)文件。
