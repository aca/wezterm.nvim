package main

import (
	"os"

	"github.com/neovim/go-client/nvim"
)

func main() {
	direction := os.Args[1]
	addr := os.Getenv("NVIM_LISTEN_ADDRESS")
	c, err := nvim.Dial(addr)
	if err != nil {
		os.Exit(1)
	}

	var output int
	err = c.Eval("winnr() == "+"winnr('"+direction+"')", &output)
	if err != nil {
		os.Exit(1)
	}

	if output == 1 {
		os.Exit(1)
	}
}
