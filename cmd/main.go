package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kamikazechaser/eth-p2p-challenger/internal/p2p"
	"github.com/kamikazechaser/eth-p2p-challenger/internal/util"
	"github.com/knadh/koanf/v2"
)

const defaultGracefulShutdownPeriod = time.Second * 5

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
	var wg sync.WaitGroup
	ctx, stop := notifyShutdown()

	client, err := p2p.NewClient(p2p.ClientOpts{
		EnodeURL:      ko.MustString("test.enode"),
		PrivateKeyHex: ko.MustString("client.private_key"),
		UserAgent:     fmt.Sprintf("%s/%s", ko.MustString("client.useragent"), build),
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
	wg.Add(1)
	go func() {
		client.ReadProcess()
		defer wg.Done()
	}()

	<-ctx.Done()
	lo.Info("shutdown signal received")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), defaultGracefulShutdownPeriod)

	wg.Add(1)
	go func() {
		client.Close()
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
