package token

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("DRIVER_NAME") == "" {
		_ = os.Setenv("DRIVER_NAME", "sqlite3")

	}
	//if os.Getenv("DATASOURCE_URL") == "" {
	//	_ = os.Setenv("DATASOURCE_URL", "../../testdata/paymentinitiation.sqlite")
	//}

	_ = os.Setenv("CLIENT_CA_CERT_PEM", "../../certs/hsbc_sandbox_intermediate.cer")
	_ = os.Setenv("CLIENT_CERT_PEM", "../../certs/server.pem")
	_ = os.Setenv("CLIENT_KEY_PEM", "../../certs/server.key")

	log.Print("TOKEN START")
	//dbx := store.LoadDBConnection()
	//
	//api.RunSql(dbx, "../../testdata/insert_data.down.sql")
	//api.RunSql(dbx, "../../testdata/insert_data.up.sql")

	exitCode := m.Run()
	log.Print("TOKEN END")
	os.Exit(exitCode)
}
