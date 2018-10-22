## 安装代码生成工具
pip install grpcio
pip install grpcio_tools

## 生成go文件
python -m grpc_tools.protoc -I.  --go_out=plugins=grpc:. helloworld.proto -I /Users/congzaifeng/go/src/googleapis-master --grpc-gateway_out=logtostderr=true:.

## grpc gateway 测试
http://127.0.0.1:8080/v1/hello?name=gogogo

