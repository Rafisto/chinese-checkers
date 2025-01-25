package game

type Game interface {
	AddPlayer(playerID int) error
	SetPlayerNum(playerNum int) error
	GetID() int
	GetBoard() Board
	GetPlayers() []int
	GetPlayerNum() int
	GetCurrentPlayerNum() int
	GetTurn() int
	GetPlayerTurn() int
	GetProgress() []int
	GetVariant() string
	GetEnded() bool
	SetTurn(int)
	SetProgress([]int)
	SetEnded(bool)
	Move(playerID, oldX, oldY, x, y int) error
	SkipTurn(playerID int) error
	AddBot(botID int) error
}
