package app

import (
	"context"
	"fmt"
	"frontdev333/wishlist/configs"
	"frontdev333/wishlist/database"
	"frontdev333/wishlist/internal"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg = &sync.WaitGroup{}

func Start() {
	cfg := configs.NewConfig()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)
	defer cancel()
	db, err := database.New(cfg.DataBaseCredentials)
	if err != nil {
		panic(err)
	}
	_ = db

	router := internal.NewRouter()

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	go func() {
		slog.Info("starting server on", slog.String("port", cfg.Addr))
		if err = server.ListenAndServe(); err != http.ErrServerClosed {
			slog.Error("server error", err.Error())
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down server")
	if err = server.Shutdown(ctx); err != nil {
		fmt.Printf("shutdown with error: %v", err)
	}
	slog.Info("shutdown complete")
	slog.Info("waiting for goroutines")
	wg.Wait()
	slog.Info("goroutines are finished")
}
