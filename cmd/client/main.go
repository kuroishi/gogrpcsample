package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hellopb "kuroishi/grpc/pkg/grpc"
	"fmt"
	"bufio"
	"os"
	"log"
	"context"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")
	scanner = bufio.NewScanner(os.Stdin)
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()
	client = hellopb.NewGreetingServiceClient(conn)
	for {
		fmt.Println("1: sendRequest")
		fmt.Println("2: exit")
		fmt.Print("please enter >")
		scanner.Scan()
		in := scanner.Text()
		switch in {
		case "1":
			Hello()
		case "2":
			fmt.Println("bye.")
			goto M
			
		}
	}
M:
}

func Hello() {
	fmt.Println("Please enter your name.")
	scanner.Scan()
	name := scanner.Text()
	req := &hellopb.HelloRequest{
		Name: name,
	}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.GetMessage())	
}
