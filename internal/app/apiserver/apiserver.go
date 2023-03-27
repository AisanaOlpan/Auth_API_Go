package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/AisanaOlpan/GoProject/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

//Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	sessionStrore := sessions.NewCookieStore([]byte(config.SessionKey))
	store := sqlstore.New(db)
	srv := newServer(store, sessionStrore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
