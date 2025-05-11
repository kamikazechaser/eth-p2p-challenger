package p2p

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rlp"
)

type protoHandshake struct {
	Version    uint64
	Name       string
	Caps       []p2p.Cap
	ListenPort uint64
	ID         []byte
	Rest       []rlp.RawValue `rlp:"tail"`
}

const (
	p2pProtocolVersion = 5
	ethProtoclVersion  = 68
)

func (c *Client) supportedCaps() []p2p.Cap {
	return []p2p.Cap{
		{
			Name:    "eth",
			Version: ethProtoclVersion,
		},
		{
			Name:    "snap",
			Version: 1,
		},
	}
}

func (c *Client) ID() []byte {
	return crypto.FromECDSAPub(&c.privateKey.PublicKey)[1:]
}

func (c *Client) sendMessage(code uint64, payload any) error {
	if c.rlpxConn == nil {
		c.logg.Error("rlpxConn is nil, cannot send hello message")
		return nil
	}

	encodedPayload, err := rlp.EncodeToBytes(payload)
	if err != nil {
		c.logg.Error("failed to encode payload", "payload", payload, "err", err)
		return err
	}

	_, err = c.rlpxConn.Write(code, encodedPayload)
	if err != nil {
		c.logg.Error("failed to write message", "code", code, "payload", payload, "err", err)
		return err
	}
	return nil
}
