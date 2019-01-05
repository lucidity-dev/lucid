package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lucidity-dev/luciditygo/lucidity"

	"nanomsg.org/go-mangos"
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

func main() {
	lucidity.Connect()
	pub_sock, err := lucidity.GetTopicSocket("test2-pub", lucidity.PUB_SOCKET)

	if err != nil {
		die("%s", err)
	}

	var sub_sock mangos.Socket
	sub_sock, err = lucidity.GetTopicSocket("test1-pub1", lucidity.SUB_SOCKET)

	go func() {
		for {
			var msg []byte
			if msg, err = sub_sock.Recv(); err != nil {
				die("Cannot recv: %s", err.Error())
			}
			fmt.Printf("CLIENT test2: RECEIVED %s\n", string(msg))
		}
	}()

	for {
		d := date()
		fmt.Printf("SERVER test2: PUBLISHING DATE %s\n", d)
		if err = pub_sock.Send([]byte(d)); err != nil {
			die("Failed publishing: %s", err.Error())
		}

		time.Sleep(time.Second)
	}

}
