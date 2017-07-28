package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", ":8125")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	blue := color.New(color.FgBlue).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for {
		buffer := make([]byte, 1024)
		pc.ReadFrom(buffer)
		str := strings.Split(string(buffer), ":")
		metric := str[0]
		value := strings.Join(str[1:], ":")
		fmt.Printf("[%s] StatsD Metric: %s %s\n", magenta(time.Now().String()), blue(metric), green(value))
	}
}
