package main

import (
	"context"

	"github.com/d1gt/musigo/internal/tui"
)

func main() {
	t := tui.NewTui()
	t.Run(context.Background())
}
