package p2p

import (
	"fmt"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rlp"
)

const disconnectMsg = 0x01

func (c *Client) handleDisconnect(payload []byte) error {
	var disc *p2p.DiscReason

	if len(payload) > 0 {
		reason := payload[0:1]
		if len(payload) > 1 {
			reason = payload[1:2]
		}
		if err := rlp.DecodeBytes(reason, &disc); err != nil {
			return err
		}
	}
	c.logg.Debug("received disconnect message", "reason", disc)
	return fmt.Errorf("disconnect message received: %v", disc)
}
