package main

import (
	cli "chinese-checkers-client/client"
	"chinese-checkers-client/web"
)

func main() {
	client := web.NewClient()
	cli := cli.NewCLI(client)

	cli.CLI()
}
