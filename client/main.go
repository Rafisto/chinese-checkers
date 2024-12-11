package main

import (
	cli "chinese-checkers-client/cli"
	"chinese-checkers-client/web"
)

func main() {
	client := web.NewClient()
	cli := cli.NewCLI(client)

	cli.Start()
}
