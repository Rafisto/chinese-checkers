package main

import (
	"chinese-checkers/game"
	"chinese-checkers/server"
)

//	@title			Chinese Checkers API
//	@version		1.0-sprint1
//	@description	This is the API for the Chinese Checkers game. It allows you to create and join games, and play the game with other players.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Project Documentation
//	@contact.url	https://github.com/rafisto/chinese-checkers
//	@contact.email	rvrelay@gmail.com

//	@host		localhost:8080
//	@BasePath	/

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	board, err := game.NewClassicBoard(6)
	if err != nil {
		print("oops")
	}
	board.PrintBoard()

	board.GetPawns().PrintPawns()

	s := server.NewServer()

	s.RunServer(8080)

}
