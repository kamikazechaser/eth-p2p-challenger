package challenger

import (
	"log/slog"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	ethproto "github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/kamikazechaser/eth-p2p-challenger/internal/p2p"
)

type (
	ChallengerOpts struct {
		Logg      *slog.Logger
		P2PClient *p2p.Client
	}

	Challenger struct {
		stopCh    chan struct{}
		interval  time.Duration
		p2pClient *p2p.Client
		ticker    *time.Ticker
		logg      *slog.Logger
	}
)

const interval = 10 * time.Second

func NewChallenger(o ChallengerOpts) *Challenger {
	return &Challenger{
		stopCh:    make(chan struct{}),
		logg:      o.Logg,
		interval:  interval,
		p2pClient: o.P2PClient,
		ticker:    time.NewTicker(interval),
	}
}

func (b *Challenger) Stop() {
	b.ticker.Stop()
	b.stopCh <- struct{}{}
}

func (c *Challenger) Start() {
	for {
		select {
		case <-c.stopCh:
			c.logg.Debug("challnger shutting down")
			return
		case <-c.ticker.C:
			c.logg.Debug("challenger tick")
			c.challenge()
		}
	}
}

func (c *Challenger) challenge() error {
	if !c.p2pClient.Ready() {
		c.logg.Warn("p2p client not ready, skipping challenge")
		return nil
	}

	getHeadersRequest := p2p.GetBlockHeaders{
		GetBlockHeadersRequest: &ethproto.GetBlockHeadersRequest{
			Origin: eth.HashOrNumber{
				Number: uint64(35000000 + rand.Intn(1000000)),
			},
			Amount:  5,
			Skip:    0,
			Reverse: rand.Intn(2) == 1,
		},
	}
	c.logg.Info("sending random block headers request", "start_block", getHeadersRequest.GetBlockHeadersRequest.Origin.Number, "reverse", getHeadersRequest.GetBlockHeadersRequest.Reverse)
	if err := c.p2pClient.SendGetBlockHeaders(getHeadersRequest); err != nil {
		c.logg.Error("failed to request block headers after ping", "err", err)
		return err
	}

	return nil
}
