package main

import (
	"fmt"
	"net"
	"time"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/api"
	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
)

const ConnectTimeout = time.Second * 3

func main() {
	go func() {
		listen, err := net.Listen("tcp", "127.0.0.1:8020")
		if err != nil {
			panic(err)
		}
		defer listen.Close()
		defer fmt.Println("I closed")
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		body := make([]byte, 0)
		conn.Read(body)
		conn.Write([]byte("jsldifjasoidfjiol"))
		fmt.Println("body: ", string(body))
		time.Sleep(time.Second)
	}()
	time.Sleep(time.Second)
	sdk_on := api.NewClient("127.0.0.1:8020", "f067c8c3dce2f165d595f1c82e3d3b61dd502130157f5b503b49fcf4784afc7f", false)

	err := sdk_on.Connect(ConnectTimeout)
	if err != nil {
		fmt.Println(err)
	}

	sdk_on.CreateAccount(&types.CreateAccountRequest{Name: "ssss"})
	time.Sleep(time.Second)

}
