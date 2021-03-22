package main

import (
	"github.com/joho/godotenv"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/cache"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/config"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/store"
	"github.com/kaanaktas/openbanking-paymentinitiation/pkg/authmanager"
	cfg "github.com/kaanaktas/openbanking-paymentinitiation/pkg/configmanager"
	"github.com/kaanaktas/openbanking-paymentinitiation/pkg/consent"
	"github.com/kaanaktas/openbanking-paymentinitiation/pkg/domesticpayments"
	"github.com/kaanaktas/openbanking-paymentinitiation/pkg/session"
	"github.com/kaanaktas/openbanking-paymentinitiation/pkg/token"
	"log"
	"os"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := config.NewEchoEngine()
	dbx := store.LoadDBConnection()
	chInMemory := cache.LoadInMemory()

	configRepository := cfg.NewRepository(dbx)
	configService := cfg.NewService(configRepository, chInMemory)
	sessionRepository := session.NewRepository(dbx)
	sessionService := session.NewService(sessionRepository)
	tokenService := token.NewService(configService)
	consentRepositoryRead := consent.NewRepositoryRead(dbx)
	consentServiceRead := consent.NewServiceRead(consentRepositoryRead)
	consentRepositoryWrite := consent.NewRepositoryWrite(dbx)
	consentServiceWrite := consent.NewServiceWrite(consentRepositoryWrite)
	consentProxyService := consent.NewFacade(consentServiceRead, consentServiceWrite, tokenService, configService)
	consentManagerService := authmanager.NewAuthManager(consentServiceRead, consentServiceWrite, tokenService)
	domesticPaymentsService := domesticpayments.NewService(consentManagerService, configService)

	// Routes
	session.RegisterHandler(e, sessionService)
	consent.RegisterHandler(e, sessionService, consentServiceRead, consentProxyService)
	domesticpayments.RegisterHandler(e, domesticPaymentsService)

	log.Printf("starting server at :%s", port)

	if err := e.Start(":" + port); err != nil {
		log.Fatalf("error while starting server at :%s, %v", port, err)
	}
}
