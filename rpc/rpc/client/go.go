package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//client, err := proto.DialHelloService("tcp", "localhost:1234")
	//if err != nil {
	//	log.Fatal("dialing:", err)
	//}

	// codec
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	//err = client.Hello("cxy", &reply)
	err = client.Call("path/to/pkg.HelloService.Hello", "cxy", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
