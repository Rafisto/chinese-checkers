package game

var GameTypes = [...]func(int, int) (Game, error){NewClassicGame}
