package p2p

import (
	"fmt"

	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/rlp"
)

type (
	GetBlockHeaders = eth.GetBlockHeadersPacket
	BlockHeaders    = eth.BlockHeadersPacket
)

const (
	getBlockHeadersMsg = 0x13
	blockHeadersMsg    = 0x14
)

func (c *Client) handleBlockHeaders(payload []byte) error {
	var bh *BlockHeaders
	if err := rlp.DecodeBytes(payload, &bh); err != nil {
		return fmt.Errorf("error decoding block headers: %w", err)
	}
	c.logg.Debug("received block headers", "bh", bh)
	for _, header := range bh.BlockHeadersRequest {
		c.logg.Info("received block header", "number", header.Number, "hash", header.Hash().Hex())
	}

	if err := c.sendMessage(blockHeadersMsg, &bh); err != nil {
		return err
	}

	return nil
}

func (c *Client) SendGetBlockHeaders(bh GetBlockHeaders) error {
	c.logg.Debug("sending get block headers request", "request", bh)

	if err := c.sendMessage(getBlockHeadersMsg, bh); err != nil {
		return fmt.Errorf("failed to send get block headers: %w", err)
	}

	c.logg.Debug("successfully sent get block headers request")
	return nil
}
