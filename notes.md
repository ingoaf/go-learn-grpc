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

### Server
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

## Bi-directional Streaming
- has `stream` in request and response signature in proto file
- client sends many responses, server returns many responses
- handling on server side:
    - reseive requests using for loop and `stream.Recv()`
    - check for `io.Eof` error -> break loop
    - other error -> handle it
    - use `stream.Send()` to send response
- handling on client side:
    - create client
    - client sends and receive response each in a goroutine
    - for this a channel is created
    - in sending goroutine: 
        - use `stream.Send()` to send
        - after finish close with `stream.CloseSend()`
    - in receiving goroutine:
        - use `stream.Recv()` to receive responses from server
        - check for `io.Eof` error -> break loop
        - other error -> handle it
        - close previously created channel
    - channel blocks at the end until closed

## Grpc Errors and Error Codes
- we can send error codes along with errors to specify certain errors (similar idea to http status codes)
- on server side:
    - use `status.Errorf()` from `google.golang.org/grpc/status` to provide an error along with an error code
    - use `codes.InvalidArgument` from `"google.golang.org/grpc/codes` to specify what went wrong
- on client side:
    - make a call, handle the error
    - check error with `e, ok := status.FromError()`
        - if ok -> grpc Error
        - check the code with `e.Codes()` 

## Deadlines/Timeouts
- we can introduce timeouts on gRPC calls using `context.Context`
- on client side:
    - introduce a context with timeout and pass it to the client
    - call the gRPC server, pass context
    - analyze the error via `e.Codes()` and `codes.DeadlineExceeded` to check if the server timed out
- on server side:
    - analyze the error via `ctx.Error()` and `codes.DeadlineExceeded` to check if the client canceled

## TLS
Idea: secure connection by f.e. loading certs from files
- on server side:
    - create `[]grpc.ServerOption{}`
    - load cert and private key of server with `credentials.NewServerTLSFromFile`
    - add the options while calling `grpc.NewServer`
- on client side:
    - create `[]grpc.DialOption{}`
    - load CA cert with `NewClientTLSFromFile`
    - add the options while callind `grpc.Dial`

## Reflection
- gRPC reflection allows to list all Methods/Requests/Responses which are available on the server (similar to automatic swagger generation in REST)
- to enable reflection:
    - import `google.golang.org/grpc/reflection`
    - use `reflection.Register(server)` after registering your service in the grpc Server

### Evans
- Evans is a grpc client CLI tool which makes use of reflection
- example command: `evans --host localhost --port 50051 --reflection repl`