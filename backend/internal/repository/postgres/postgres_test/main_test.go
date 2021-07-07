package postgres_test

import (
	"chatServer/internal/repository"
	"chatServer/pkg/db/sqlstore"
	"chatServer/pkg/db/sqlstore/postgres"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

var store *sqlstore.Store
var repo *repository.Repository

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	c := sqlstore.NewConfig()
	c.User = "Tpostgres"
	c.Password = "Tpostgres"
	c.DBName = "chatServerTest"
	var err error
	store, err = postgres.NewStore(c)
	if err != nil {
		log.Fatal(err)
	}
	repo = repository.NewRepository(store)
	os.Exit(m.Run())
}
