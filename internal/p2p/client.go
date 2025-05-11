package p2p

import (
	"context"
	"crypto/ecdsa"
	"log/slog"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/rlpx"
	"github.com/kamikazechaser/eth-p2p-challenger/internal/util"
)

type (
	ClientOpts struct {
		EnodeURL      string
		PrivateKeyHex string
		UserAgent     string
		Logg          *slog.Logger
	}

	Client struct {
		userAgent  string
		conn       net.Conn
		stopCh     chan struct{}
		privateKey *ecdsa.PrivateKey
		rlpxConn   *rlpx.Conn
		logg       *slog.Logger
		enode      *enode.Node
	}
)

const timeout = 10 * time.Second

func NewClient(o ClientOpts) (*Client, error) {
	privateKey, err := util.DumpKeyFromHex(o.PrivateKeyHex)
	if err != nil {
		return nil, err
	}

	return &Client{
		userAgent:  o.UserAgent,
		logg:       o.Logg,
		privateKey: privateKey,
		enode:      enode.MustParseV4(o.EnodeURL),
	}, nil
}

func (c *Client) Connect(ctx context.Context) error {
	netDialer := net.Dialer{
		Timeout: timeout,
	}
	tcpEndpoint, _ := c.enode.TCPEndpoint()

	conn, err := netDialer.DialContext(ctx, "tcp", tcpEndpoint.String())
	if err != nil {
		c.logg.Error("failed to connect to node tcp endpoint", "tcpEndpoint", tcpEndpoint.String(), "err", err)
		return err
	}
	c.logg.Debug("successfully established connection with node tcp endpoint", "tcpEndpoint", tcpEndpoint.String())
	c.conn = conn

	c.rlpxConn = rlpx.NewConn(c.conn, c.enode.Pubkey())
	c.logg.Debug("starting handhsake with node", "nodeID", c.enode.ID().String())

	_, err = c.rlpxConn.Handshake(c.privateKey)
	if err != nil {
		c.logg.Error("failed to handshake with node", "nodeID", c.enode.ID().String(), "err", err)
		return err
	}

	return nil
}

func (c *Client) Close() error {
	c.logg.Debug("closing connection")
	if c.conn != nil {
		c.conn.Close()
	}
	return nil
}

// TODO: Refactor to break out of loop and support individual connections held in a pool
func (c *Client) ReadProcess() {
	if err := c.sendHello(); err != nil {
		c.logg.Error("failed to send hello message", "err", err)
		return
	}

	for {
		code, payload, _, err := c.rlpxConn.Read()
		if err != nil {
			c.logg.Error("failed to read message", "err", err)
			return
		}

		switch int(code) {
		case helloMsg:
			if err := c.handleHello(payload); err != nil {
				c.logg.Error("failed to handle hello message", "err", err)
				return
			}
		case statusMsg:
			if err := c.handleStatus(payload); err != nil {
				c.logg.Error("failed to handle status message", "err", err)
				return
			}
		case pingMsg:
			if err := c.handlePing(payload); err != nil {
				c.logg.Error("failed to handle ping message", "err", err)
				return
			}
		default:
			c.logg.Debug("received message", "code", code, "payload", payload)
		}
	}

}
