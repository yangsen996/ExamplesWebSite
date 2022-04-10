package main

import (
	"context"
	"github.com/yangsen996/ExamplesWebSite/global"
	"github.com/yangsen996/ExamplesWebSite/initialize"
	"github.com/yangsen996/ExamplesWebSite/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	global.G_VP = initialize.Viper()
	global.G_DB = initialize.GormDB()
	if global.G_DB != nil {
		initialize.RegisterTables(global.G_DB)
		db, _ := global.G_DB.DB()
		defer db.Close()
	}
	r := router.Router()

	srv := &http.Server{
		Addr:              ":9999",
		Handler:           r,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen:", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit)
	<-quit
	log.Println("server shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	select {
	case <-ctx.Done():
	}
	log.Println("server exiting")
}
