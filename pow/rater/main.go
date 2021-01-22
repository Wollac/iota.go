package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/iotaledger/iota.go/v2/pow"
)

var (
	numWorkers = flag.Int(
		"workers",
		1,
		"number of workers used for mining",
	)
	messageSize = flag.Int(
		"size",
		1024,
		"size of the input message in bytes; should have no impact on the performance")
	duration = flag.Duration(
		"duration",
		5*time.Minute,
		"timeout after which mining is stopped")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), *duration)
	defer cancel()

	msg := make([]byte, *messageSize)
	rand.Read(msg)

	_, err := pow.New(*numWorkers).Mine(ctx, msg, 1e6)
	return err
}
