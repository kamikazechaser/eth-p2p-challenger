package p2p

import (
	"fmt"
)

const (
	pingMsg = 0x02
	pongMsg = 0x03
)

func (c *Client) handlePing() error {
	c.logg.Debug("received ping message")

	if err := c.sendPong(); err != nil {
		return fmt.Errorf("failed to send pong: %w", err)
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
