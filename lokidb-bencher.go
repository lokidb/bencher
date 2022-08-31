package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/lokidb/server/client"
)

var (
	addr       = flag.String("addr", "localhost:50051", "Server address")
	iterations = flag.Int("it", 10000, "Number of iterations")
	command    = flag.String("c", "get", "Command to benchmark options: [get, set, del]")
)

var validCommands = map[string]struct{}{
	"get": {},
	"set": {},
	"del": {},
}

func main() {
	flag.Parse()
	c := client.New(*addr, time.Second*10)

	if _, validCommand := validCommands[*command]; !validCommand {
		fmt.Println("Invalid command. options [get, set, del]")
		os.Exit(1)
	}

	startTime := time.Now()

	exec(c, *command)

	duration := time.Since(startTime)
	nanoPerOp := duration.Nanoseconds() / int64(*iterations)
	secPerOp := 1000000000 / nanoPerOp

	fmt.Printf("Ran %d iterations in %s\n", *iterations, duration)
	fmt.Printf("op/s %d\n", secPerOp)
}

func exec(c *client.Client, command string) {
	lCh := make(chan string)
	wg := new(sync.WaitGroup)

	// Adding routines to workgroup and running then
	for i := 0; i < 250; i++ {
		wg.Add(1)
		go worker(c, lCh, wg, command)
	}

	// Processing all links by spreading them to `free` goroutines
	for i := 0; i < *iterations; i++ {
		key := strconv.Itoa(i)
		lCh <- key
	}

	// Closing channel (waiting in goroutines won't continue any more)
	close(lCh)

	// Waiting for all goroutines to finish (otherwise they die as main routine dies)
	wg.Wait()
}

func worker(c *client.Client, linkChan chan string, wg *sync.WaitGroup, command string) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	for key := range linkChan {
		var err error

		switch command {
		case "get":
			_, err = c.Get(key)
		case "set":
			err = c.Set(key, "asdasdajlsdjaksdjalkjsdajsdkj")
		case "del":
			_, err = c.Del(key)
		}

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
