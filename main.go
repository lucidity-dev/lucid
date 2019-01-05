// Copyright © 2018 Rishi Desai
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lucidity-dev/lucid/cmd"

	"nanomsg.org/go-mangos/protocol/req"
	"nanomsg.org/go-mangos/transport/tcp"

	"github.com/golang/protobuf/proto"
	pb "github.com/lucidity-dev/bulletin/protobuf"
)

func cleanup() {
	url := "tcp://127.0.0.1:40899"
	server, err := req.NewSocket()
	if err != nil {
		log.Fatalf("Error: %v \n", err)
		os.Exit(1)
	}

	server.AddTransport(tcp.NewTransport())
	if err = server.Dial(url); err != nil {
		log.Fatalf("Error: %v \n", err)
	}
	request := &pb.Message{
		Cmd:  pb.Message_FLUSH_ALL,
		Args: "",
	}

	data, err := proto.Marshal(request)

	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}

	if err = server.Send(data); err != nil {
		log.Fatalf("Error: %v \n", err)
		os.Exit(1)
	}

	if _, err = server.Recv(); err != nil {
		log.Fatalf("Error: %v \n", err)
		os.Exit(1)
	}

}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
	cmd.Execute()
}
