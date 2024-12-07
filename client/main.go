package main

import chineseclient "chinese-checkers-client/client"

func main() {
	client := chineseclient.NewClient()
	cli := chineseclient.NewCLI(client)

	cli.CLI()
}
