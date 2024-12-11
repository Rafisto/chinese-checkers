package CLI

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"chinese-checkers-client/web"
)

type CLI struct {
	client              *web.Client
	webSocketConnection *web.WebSocketConnection
}

func NewCLI(client *web.Client) *CLI {
	cli := &CLI{
		client:              client,
		webSocketConnection: web.NewWebSocketConnection(),
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

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "exit":
			exited = true
		case "leave":
			err := c.client.LeaveGame()

			if err != nil {
				fmt.Println("There is no escape")
			}

			fmt.Println("Left the game")
		case "games":
			body, err := c.client.ListGamesHandler()

			if err != nil {
				fmt.Printf("Error getting available games: %s\n", err)
			}

			fmt.Println(string(body))
		case "create":
			if len(args) == 1 {
				fmt.Println("Usage create [playerNumber]")
				continue
			}

			if c.client.GetUsername() == "" {
				fmt.Println("Please set a username first with: username [username]")
				continue
			}

			playerNum, err := strconv.Atoi(args[1])

			if err != nil {
				fmt.Println("Player must be a natural number")
				continue
			}

			if !slices.Contains([]int{2, 3, 4, 6}, playerNum) {
				fmt.Println("Player number has to be one of: 2, 3, 4, 6")
			}

			gameID, err := c.client.CreateGame(playerNum)

			if err != nil {
				fmt.Printf("Game creation failed: %s\n", err)
				continue
			}
			fmt.Printf("Game successfully created with ID: %d\n", gameID)
		case "join":
			if len(args) == 1 {
				fmt.Println("Usage: join [GameID]")
				continue
			}

			if c.client.GetUsername() == "" {
				fmt.Println("Please set a username first with: username [username]")
				continue
			}

			gameID, err := strconv.Atoi(args[1])

			if err != nil {
				fmt.Println("ID must be a natural number")
				continue
			}

			playerID, err := c.client.JoinGame(gameID)

			if err != nil {
				fmt.Printf("Failure joining the game: %s\n", err)
				continue
			}

			fmt.Println("Successfully joined the game")
			fmt.Println("Connecting to the socket")

			err = c.webSocketConnection.EstablishConnection(gameID, playerID)

			if err != nil {
				fmt.Printf("Failure connecting to the socket: %s\n", err)
				continue
			}

			fmt.Println("Connected to the socket")

			go func() {
				for {
					message, err := c.webSocketConnection.ReceiveMessage()
					if err != nil {
						fmt.Printf("Failure receiving the message: %s\n", err)
						break
					}
					fmt.Printf("Received message: %s\n", message)
				}
			}()
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

			if args[1] == "" {
				continue
			}

			if c.webSocketConnection == nil {
				fmt.Println("You need to join a game first")
				continue
			}

			message := strings.Join(args[1:], " ")
			err := c.webSocketConnection.EmitMessage(message)
			if err != nil {
				fmt.Printf("Failure sending the message: %s\n", err)
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}
