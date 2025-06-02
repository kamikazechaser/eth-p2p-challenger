package p2p

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
)

type Hello = protoHandshake

const helloMsg = 0x00

func (c *Client) sendHello() error {
	hello := &Hello{
		Version: p2pProtocolVersion,
		Caps:    c.supportedCaps(),
		ID:      c.ID(),
	}
	c.logg.Debug("prepared hello message", "hello", hello)

	if err := c.sendMessage(helloMsg, hello); err != nil {
		return err
	}
	c.logg.Debug("successfully sent hello message")

	return nil
}

func (c *Client) handleHello(payload []byte) error {
	var h Hello
	if err := rlp.DecodeBytes(payload, &h); err != nil {
		return fmt.Errorf("failed to decode hello message: %w", err)
	}
	c.logg.Debug("received hello message", "hello", h)

	if h.Version < p2pProtocolVersion {
		return fmt.Errorf("unsupported protocol version: %d", h.Version)
	}

	for _, cap := range h.Caps {
		c.logg.Debug("supported eth caps", "cap_name", cap.Name, "cap_version", cap.Version)
		if cap.Name == "eth" && cap.Version < ethProtocolVersion {
			return fmt.Errorf("unsupported eth protocol version: %d", cap.Version)
		}
		if cap.Name == "snap" && cap.Version == 1 {
			c.rlpxConn.SetSnappy(true)
			c.logg.Debug("setting snap to true")
		}
	}

	return nil

}
