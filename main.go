package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	for {
		nc, _ := nats.Connect(nats.DefaultURL)
		sub, err := nc.SubscribeSync("update")
		if err != nil {
			log.Println(err)
		}

		timeout := 525960 * time.Minute // 1 year
		m, err := sub.NextMsg(timeout)
		if err != nil {
			log.Println(err)
		}
		if m == nil {
			continue
		}
		if string(m.Data) == "hehehe" {
			cmd := exec.Command("echo", "hehehe")
			output, _ := cmd.CombinedOutput()
			fmt.Println(string(output))
			nc.Publish(m.Reply, output)
		}
	}
}
