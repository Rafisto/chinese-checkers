package chineseclient

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CLI struct {
	client *Client
}

func NewCLI(client *Client) *CLI {
	cli := &CLI{
		client: client,
	}
	return cli
}

func (c *CLI) CLI() {
	exited := false
	reader := bufio.NewReader(os.Stdin)
	for !exited {
		fmt.Print("[CC] ")
		command, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		var args []string = strings.Fields(command)

		switch args[0] {
		case "exit":
			exited = true
		case "leave":
			err := c.client.LeaveGame()

			if err != nil {
				fmt.Println("There is no escape")
			}

			fmt.Println("Left the game")
		case "create":
			err := c.client.CreateGame()

			if err != nil {
				fmt.Printf("Game creation failed: %s\n", err)
				continue
			}

			fmt.Println("Game successfully created")
		case "join":
			if len(args) == 1 {
				fmt.Println("Usage: join [id]")
				continue
			}

			gameID, err := strconv.Atoi(args[1])

			if err != nil {
				fmt.Println("ID must be a natural number")
				continue
			}

			err = c.client.JoinGame(gameID)

			if err != nil {
				fmt.Printf("Failure joining the game: %s\n", err)
				continue
			}

			fmt.Println("Successfully joined the game")
		case "username":
			if len(args) == 1 {
				fmt.Println("Usage: username [username]")
				continue
			}

			c.client.SetUsername(args[1])
			fmt.Printf("Changed username to %s\n", c.client.GetUsername())
		case "chat":
			if len(args) == 1 {
				fmt.Println("Usage: chat [message]")
				continue
			}

			err := c.client.SendServerMessage(strings.TrimPrefix(command, "chat "))

			if err != nil {
				fmt.Printf("Failure sending the message: %s\n", err)
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}
