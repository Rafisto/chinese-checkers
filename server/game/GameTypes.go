package game

var GameTypes = map[string]func(int, int) (Game, error){
	"classic": NewClassicGame,
	"chaos":   NewChaosGame,
}
