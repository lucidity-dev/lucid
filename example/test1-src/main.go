package main

import (
	"fmt"
	"os"
	"time"

	"nanomsg.org/go-mangos"
	"nanomsg.org/go-mangos/protocol/pub"
	"nanomsg.org/go-mangos/protocol/req"
	"nanomsg.org/go-mangos/protocol/sub"
	"nanomsg.org/go-mangos/transport/ipc"
	"nanomsg.org/go-mangos/transport/tcp"

	"github.com/golang/protobuf/proto"
	pb "github.com/lucidity-dev/bulletin/protobuf"
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

func main() {
	fmt.Println("test1 service")
	url := "tcp://127.0.0.1:40899"
	var err error
	var server mangos.Socket
	var data []byte
	var msg []byte

	server, err = req.NewSocket()
	if err != nil {
		die("error creating socket: %s", err.Error())
	}

	server.AddTransport(tcp.NewTransport())

	if err = server.Dial(url); err != nil {
		die("cant dial on socket: %s", err)
	}

	request := &pb.Message{
		Cmd:  pb.Message_GET,
		Args: "test1-pub1",
	}

	data, err = proto.Marshal(request)

	if err != nil {
		die("error creating protobuf: %s", err)
	}

	if err = server.Send(data); err != nil {
		die("error sending message: %s", err)
	}

	if msg, err = server.Recv(); err != nil {
		die("didn't receive any data: %s", err)
	}

	var body pb.Topic
	proto.Unmarshal(msg, &body)

	var pub_sock, sub_sock mangos.Socket

	if pub_sock, err = pub.NewSocket(); err != nil {
		die("can't get new pub socket: %s", err)
	}
	pub_sock.AddTransport(ipc.NewTransport())
	pub_sock.AddTransport(tcp.NewTransport())
	fmt.Println("URL: " + body.Url)
	if err = pub_sock.Listen(body.Url); err != nil {
		die("can't listen on pub socket: %s", err.Error())
	}

	request = &pb.Message{
		Cmd:  pb.Message_GET,
		Args: "test-pub1",
	}

	data, err = proto.Marshal(request)

	if err != nil {
		die("error creating protobuf: %s", err)
	}

	if err = server.Send(data); err != nil {
		die("error sending message: %s", err)
	}

	if msg, err = server.Recv(); err != nil {
		die("didn't receive any data: %s", err)
	}

	proto.Unmarshal(msg, &body)

	if sub_sock, err = sub.NewSocket(); err != nil {
		die("can't get new sub socket: %s", err.Error())
	}

	sub_sock.AddTransport(ipc.NewTransport())
	sub_sock.AddTransport(tcp.NewTransport())
	if err = sub_sock.Dial(body.Url); err != nil {
		die("can't dial on sub socket: %s", err.Error())
	}
	// Empty byte array effectively subscribes to everything
	err = sub_sock.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil {
		die("cannot subscribe: %s", err.Error())
	}

	go func() {
		for {
			if msg, err = sub_sock.Recv(); err != nil {
				die("Cannot recv: %s", err.Error())
			}
			fmt.Printf("CLIENT test1: RECEIVED %s\n", string(msg))
		}
	}()

	for {
		d := date()
		fmt.Printf("SERVER test1: PUBLISHING DATE %s\n", d)
		if err = pub_sock.Send([]byte(d)); err != nil {
			die("Failed publishing: %s", err.Error())
		}

		time.Sleep(time.Second)
	}
}
