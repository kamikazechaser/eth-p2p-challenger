package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kamikazechaser/eth-p2p-challenger/internal/challenger"
	"github.com/kamikazechaser/eth-p2p-challenger/internal/p2p"
	"github.com/kamikazechaser/eth-p2p-challenger/internal/util"
	"github.com/knadh/koanf/v2"
)

const defaultGracefulShutdownPeriod = time.Second * 15

var (
	build = "dev"

	confFlag string

	lo *slog.Logger
	ko *koanf.Koanf
)

func init() {
	flag.StringVar(&confFlag, "config", "config.toml", "Config file location")
	flag.Parse()

	lo = util.InitLogger()
	ko = util.InitConfig(lo, confFlag)

	lo.Info("starting eth-p2p-challenger", "build", build)
}

func main() {
	var (
		wg        sync.WaitGroup
		closeOnce sync.Once
	)

	ctx, stop := notifyShutdown()

	client, err := p2p.NewClient(p2p.ClientOpts{
		EnodeURL:      ko.MustString("client.enode"),
		PrivateKeyHex: ko.MustString("client.private_key"),
		Logg:          lo,
	})
	if err != nil {
		lo.Error("failed to create client", "err", err)
		os.Exit(1)
	}

	if err := client.Connect(ctx); err != nil {
		lo.Error("failed to connect to node", "err", err)
		os.Exit(1)
	}

	challnger := challenger.NewChallenger(challenger.ChallengerOpts{
		Logg:      lo,
		P2PClient: client,
	})

	wg.Add(1)
	go func() {
		defer wg.Done()
		client.ReadProcess()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		challnger.Start()
	}()

	<-ctx.Done()
	lo.Info("shutdown signal received")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), defaultGracefulShutdownPeriod)

	wg.Add(1)
	go func() {
		defer wg.Done()
		closeOnce.Do(func() {
			challnger.Stop()
			lo.Info("closing client connection")
			client.Close()
		})
	}()

	go func() {
		wg.Wait()
		stop()
		cancel()
		os.Exit(0)
	}()

	<-shutdownCtx.Done()
	if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
		stop()
		cancel()
		lo.Error("graceful shutdown period exceeded, forcefully shutting down")
	}
	os.Exit(1)

}

func notifyShutdown() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
}
