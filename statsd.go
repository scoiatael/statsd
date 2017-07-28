package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	blue    = color.New(color.FgBlue).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
)

func Format(metric, value string) string {
	if strings.HasPrefix(metric, "_e") {
		parts := strings.Split(value, "|")
		eventName := parts[0]
		eventPayload := strings.Join(parts[1:], "|")
		return fmt.Sprintf("Event:  %s %s", cyan(eventName), green(eventPayload))
	}
	return fmt.Sprintf("Metric: %s %s", blue(metric), green(value))
}

func main() {
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", ":8125")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buffer := make([]byte, 1024)
		pc.ReadFrom(buffer)
		str := strings.Split(string(buffer), ":")
		metric := str[0]
		value := strings.Join(str[1:], ":")
		fmt.Printf("[%s] StatsD %s\n", magenta(time.Now().String()), Format(metric, value))
	}
}
