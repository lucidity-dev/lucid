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
package builder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/lucidity-dev/lucid/parser"

	"nanomsg.org/go-mangos"
	"nanomsg.org/go-mangos/protocol/req"
	"nanomsg.org/go-mangos/transport/tcp"

	"github.com/golang/protobuf/proto"
	pb "github.com/lucidity-dev/bulletin/protobuf"
)

var url = "tcp://127.0.0.1:40899"
var server mangos.Socket
var project_topics = make(map[string]string)

func setupServer() {
	var err error
	server, err = req.NewSocket()
	if err != nil {
		log.Fatalf("Error: %v \n", err)
		os.Exit(1)
	}

	server.AddTransport(tcp.NewTransport())
	if err = server.Dial(url); err != nil {
		log.Fatalf("Error: %v \n", err)
	}
}

func BuildProject(project parser.Project) {
	// TODO: make parallel
	setupServer()
	fmt.Println("Registering Topics with bulletin")
	for _, service := range project.Services {
		registerPublishers(service.Publishers)
	}

	//fmt.Println("Starting Services: ")
	//for _, service := range project.Services {
	//	startService(service.Src, service.Run)
	//}
}

func registerPublishers(topics []string) {
	for _, t := range topics {
		registerPublisher(t)
	}
}

func registerPublisher(topic string) {
	request := &pb.Message{
		Cmd:  pb.Message_REGISTER,
		Args: topic,
	}

	data, err := proto.Marshal(request)

	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}

	if err = server.Send(data); err != nil {
		log.Fatalf("Error: %v \n", err)
		os.Exit(1)
	}

	var msg []byte

	if msg, err = server.Recv(); err != nil {
		log.Fatalf("Error: %v \n", err)
		os.Exit(1)
	}

	var body pb.Topic
	proto.Unmarshal(msg, &body)

	fmt.Printf("topic: %s \n", topic)
	fmt.Printf("url: %s \n", body.Url)

	project_topics[topic] = body.Url
}

func startService(dir string, run string) {
	run_cmd := strings.Fields(run)
	run = run_cmd[0]
	cmd := exec.Command(run, run_cmd[1:]...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	//err := cmd.Run()
	if err != nil {
		log.Fatalf("Error: %v \n", err)
	}
}
