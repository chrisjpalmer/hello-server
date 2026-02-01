package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"example.com/htmx-test/web"
)

func main() {
	port := flag.Int("port", 8080, "--port=8080")

	flag.Parse()

	srv := web.NewServer(*port)

	go func() {
		if err := srv.Serve(); err != nil {
			fmt.Println("error while serving", err)
		}
	}()

	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	fmt.Printf("server running... on port %d", *port)

	<-ctx.Done()

	fmt.Println("closing server down")
	srv.Close()
}
