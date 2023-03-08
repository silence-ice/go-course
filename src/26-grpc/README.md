# grpc使用

## 1、grpc基本概念
### 1）什么是grpc？
1. gRPC 的基本思想是定义一个服务，指定其可以被远程调用的方法及其参数和返回类型。
2. gRPC 默认使用 protocol buffers 作为接口定义语言，来描述服务接口和有效载荷消息结构。

### 2）什么是Protocol Buffers？
1. Protocol Buffers(protobuf)：与编程语言无关，与程序运行平台无关的数据序列化协议以及接口定义语言(IDL: interface definition language)。
2. Protocol Buffers 使用 编译器protoc 来编译.proto文件。

### 3）安装protoc编译器？
1. protoc-gen-go插件：用于生成xx.pb.go文件

   ```go install google.golang.org/protobuf/cmd/protoc-gen-go@latest```
2. protoc-gen-go-grpc插件：用于生成xx_grpc.pb.go文件

    ```go install google.golang.org/protobuf/cmd/protoc-gen-go@latest```
3. `go install` 默认是安装在 `go path`路径下（可使用`go env`）查看，然后讲`go path`放在环境变量即可

### 4）为什么要使用grpc？
gRPC 帮你解决了不同语言及环境间通信的复杂性，我们可以一次性的在一个.proto 文件中定义服务并使用任何支持它的语言去实现客户端和服务器

### 5）grpc编写流程
1. 在一个 .proto 文件内定义服务。
2. 用 protocol buffer 编译器生成服务器和客户端代码。
3. 使用 gRPC 的 Go API 为你的服务创建一个简单的客户端和服务器。

### 6）帮助文档
1. grpc概念： https://doc.oschina.net/grpc?t=58009
2. grpc概念： https://juejin.cn/post/7048486841728630798
3. protocol buffer语法规则：https://colobu.com/2015/01/07/Protobuf-language-guide/
4. grpc官方文档：https://grpc.io/docs/languages/go/quickstart/

## 2、protoc 命令使用

### 1） 项目结构如下
```goland
├── 26-grpc
│   ├── helloworld
│   │   ├── greeter.proto
│   │   └── greeter.pb.go
│   └── greeter_client
│   │   ├──main.go
│   └── greeter_server
│   │   ├──main.go
```

### 2） 定义流程
1. 在`greeter.proto`中定义`option go_package = "/helloworld";`，表示会创建 `helloworld`目录，再写入文件
2. `cd`当到路径`26-grpc`路径下
3. 执行`protoc --proto_path=. --go_out=. --go-grpc_out=. helloworld/helloworld.proto`
4. `--proto_path` 表示的是我们要在哪个路径下搜索.proto文件，这个参数既可以用-I指定，也可以使用--proto_path=指定
5. `--go_out=.` 表示生成 目标文件的所在路径，.表示当前路径
6. `helloworld/helloworld.proto` 表示当前路径（`pwd`）下需要编译的文件

### 3） 帮助文档
1. https://juejin.cn/post/6949927882126966820