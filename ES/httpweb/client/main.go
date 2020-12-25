package client

func init() {
	var DefaultClient = &Client{}
}

type Client struct {
}

func (c *Client) ServerMux(a, b int) int {
	d := a + b
	return d
}
