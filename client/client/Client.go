package chineseclient

type Client struct {
	// socket   string // TODO: change later to socket
	gameID   int
	username string
}

func NewClient() *Client {
	client := &Client{}
	return client
}

func (c *Client) SetGameID(gameID int) {
	c.gameID = gameID
}

func (c *Client) SetUsername(username string) {
	c.username = username
}

func (c *Client) GetGameID() int {
	return c.gameID
}

func (c *Client) GetUsername() string {
	return c.username
}

func (c *Client) JoinGame(gameID int) error {
	// TODO
	return nil
}

func (c *Client) CreateGame() error {
	// TODO
	return nil
}

func (c *Client) LeaveGame() error {
	// TODO
	return nil
}

func (c *Client) ChangeUsername(newUsername string) error {
	// TODO: change username on server
	c.username = newUsername
	return nil
}

func (c *Client) SendServerMessage(message string) error {
	// TODO
	return nil
}
