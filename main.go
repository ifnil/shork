package main

import (
	"context"

	"github.com/ifnil/shork/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd.Execute(ctx)
}
