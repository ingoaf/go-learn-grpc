## Generate Go code for test.proto
```bash
protoc -I greet/proto --go_out=. --go_opt=module=github.com/ingoaf/learning-go-grpc/hands-on --go-grpc_out=. --go-grpc_opt=module=github.com/ingoaf/learning-go-grpc/hands-on greet/proto/dummy.proto
```

- `-I` -> import path to proto folder
- `--go_out` -> out dir for client code
- `--go-grpc_out` -> out dir for server code
- `--go_opt=module` -> module path for client
- `--go-grpc_opt=module` -> module path for server
- `greet/proto/dummy.proto` -> path to proto file

## Installation
- https://grpc.io/docs/languages/go/quickstart/

## General notes
- gRPC relies on HTTP2
    - requires SSL by default
    - binary header
- gRPCS server is always async, client can be async or blocking
- we can use interceptors for authentication in gRPC

### Protobuf
- we can import proto file a into protofile b by writing `import a.proto` into protofile b 

## Api types of gRPC
- unary
- server streaming (returns stream)
- client streaming (sends stream)
- bi-directional streaming (sends and returns stream)

They are distinguished by the keyword **stream**

## GRPC client and Server
### Client
- open a gRPC connection to a host and port using SSL
- defer closing connection
- create new client from generated code
- handle response (f.e. print)

## Server
- create tcp listener on address and port
- create Service with custom functionality, add method which implements custom functionality
- create grpc Server
- register custom service 
- serve listener

## Unary API
- client sends one request, server returns one response
- response handling: we call a server method with a function and handle the response immediately

## Server Streaming
- has `stream` in response signature in proto file, f.e. `stream GreetResponse`
- client sends one request, server returns many responses
- handling on client side:
    - create client
    - send request -> function returns stream object
    - make for loop using `stream.Recv()`
    - check for `io.Eof` error -> break loop
    - other error -> handle it
- handling on server side:
    - receive request
    - use for loop and `stream.Send()` to send response

## Client Streaming
- has `stream` in request signature in proto file, f.e. `stream GreetRequest`
- client sends many responses, server returns one response
- handling on client side:
    - create client
    - client sends requests with `.Send(...)`
    - uses `stream.CloseAndRecv()` to close the connection
    - receives server response from the previous function
- handling on server side:
    - receive requests using for loop and `stream.Recv()`
    - check for `io.Eof` error -> break loop
    - other error -> handle it
    - uses `stream.SendAndClose()` to send response

