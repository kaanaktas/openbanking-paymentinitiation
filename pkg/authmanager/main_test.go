package authmanager

import (
	"github.com/kaanaktas/openbanking-paymentinitiation/api"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/store"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("DRIVER_NAME") == "" {
		_ = os.Setenv("DRIVER_NAME", "sqlite3")

	}
	if os.Getenv("DATASOURCE_URL") == "" {
		_ = os.Setenv("DATASOURCE_URL", "../../testdata/paymentinitiation.sqlite")
	}

	log.Print("AUTHMANAGER START")
	dbx := store.LoadDBConnection()

	api.RunSql(dbx, "../../testdata/insert_data.down.sql")
	api.RunSql(dbx, "../../testdata/insert_data.up.sql")

	exitCode := m.Run()
	log.Print("AUTHMANAGER END")
	os.Exit(exitCode)
}
