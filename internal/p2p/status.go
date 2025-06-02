package p2p

import (
	"fmt"

	ethproto "github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/rlp"
)

type Status = ethproto.StatusPacket

const statusMsg = 0x10

func (c *Client) sendStatus(decodedStatusPacket *Status) error {
	c.logg.Debug("prepared status message with same payload")

	if err := c.sendMessage(statusMsg, decodedStatusPacket); err != nil {
		return err
	}
	c.logg.Debug("successfully sent status message")

	return nil
}

func (c *Client) handleStatus(payload []byte) error {
	var s Status
	if err := rlp.DecodeBytes(payload, &s); err != nil {
		return fmt.Errorf("failed to decode status message: %w", err)
	}

	c.logg.Debug("received status message", "status", s)

	if err := c.sendStatus(&s); err != nil {
		return fmt.Errorf("failed to send status message: %w", err)
	}

	c.ready = true
	return nil
}

func (c *Client) Ready() bool {
	return c.ready
}
