package p2p

const (
	pingMsg = 0x02
	pongMsg = 0x03
)

func (c *Client) handlePing(payload []byte) error {
	c.logg.Debug("received ping")
	if err := c.sendPong(); err != nil {
		return err
	}

	return nil
}

func (c *Client) sendPong() error {
	_, err := c.rlpxConn.Write(pongMsg, []byte{})
	if err != nil {
		c.logg.Error("failed to write pong")
		return err
	}
	c.logg.Debug("successfully sent pong")

	return nil
}
