# Run Golang Connect GRPC Server

```sh
buf generate
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
go get golang.org/x/net/http2
go get connectrpc.com/connect
go run ./cmd/server/main.go
```

# Run python GRPC Client

```sh
buf generate
python3 -m venv .venv
source .venv/bin/activate
pip install -e .
PYTHONPATH=python_gen python main.py
```

# Run Golang GRPC Client

```sh
go get google.golang.org/grpc
go run main.go
```

# Run Web Client

```sh
python3 -m http.server 8003
```
