package main

import (
	"context"

	"github.com/ifnil/shork/cmd"
)

// TODO: cmd flag
// TODO: passphrase support (+ tui)
// TODO: logging
// TODO: testing
// TODO: bitwarden support

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd.Execute(ctx)
}
