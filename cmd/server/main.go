package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "connectrpc.com/connect"
    "github.com/rs/cors"
    "golang.org/x/net/http2"
    "golang.org/x/net/http2/h2c"

    greetv1 "example/gen/greet/v1"        
    "example/gen/greet/v1/greetv1connect" 
)

type GreetServer struct{}

func (s *GreetServer) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
    log.Printf("Unary Greet called with name: %s\n", req.Msg.Name)
    log.Println("Request headers: ", req.Header())
    res := connect.NewResponse(&greetv1.GreetResponse{
        Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
        GreetingNumber: 1,
    })
    return res, nil
}

func (s *GreetServer) GreetStream(
    ctx context.Context,
    req *connect.Request[greetv1.GreetRequest],
    stream *connect.ServerStream[greetv1.GreetResponse],
) error {
    log.Printf("GreetStream started for client: %s\n", req.Msg.Name)
    log.Println("Stream request headers: ", req.Header())

    name := req.Msg.Name
    greetings := []string{
        "Hello %s! Welcome to the stream",
        "How are you doing today, %s?",
        "Greetings and salutations, %s!",
        "Hope you're having a great day, %s",
        "Thank you for joining us, %s",
    }

    for i, greetFormat := range greetings {
        if ctx.Err() != nil {
            return ctx.Err()
        }

        msg := &greetv1.GreetResponse{
            Greeting:       fmt.Sprintf(greetFormat, name),
            GreetingNumber: int32(i + 1),
        }
        
        err := stream.Send(msg)
        if err != nil {
            return err
        }

        time.Sleep(1 * time.Second)
    }

    return nil
}

func main() {
    greeter := &GreetServer{}
    mux := http.NewServeMux()
    path, handler := greetv1connect.NewGreetServiceHandler(greeter)
    
    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8003"},
        AllowedMethods: []string{
            http.MethodGet,
            http.MethodPost,
            http.MethodPut,
            http.MethodPatch,
            http.MethodDelete,
            http.MethodOptions,
        },
        AllowedHeaders: []string{
            "Accept",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
            "Connect-Protocol-Version",
            "Connect-Timeout-Ms",
            "Connect-Accept-Encoding",
            "Connect-Content-Encoding",
            "Grpc-Timeout",
            "X-Grpc-Web",
            "X-User-Agent",
        },
        ExposedHeaders: []string{
            "Grpc-Status",
            "Grpc-Message",
            "Grpc-Status-Details-Bin",
        },
        AllowCredentials: true,
    })

    mux.Handle(path, corsHandler.Handler(handler))
    
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(
        ":8080",
        h2c.NewHandler(mux, &http2.Server{}),
    ); err != nil {
        log.Fatal(err)
    }
}
