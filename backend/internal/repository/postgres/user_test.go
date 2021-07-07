package postgres_test

import (
	postgres2 "chatServer/internal/repository/postgres"
	"chatServer/pkg/db/sqlstore"
	"chatServer/pkg/db/sqlstore/postgres"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var store *sqlstore.Store

func TestMain(m *testing.M) {
	c := sqlstore.NewConfig()
	c.User = "Tpostgres"
	c.Password = "Tpostgres"
	c.DBName = "chatServerTest"
	var err error
	store, err = postgres.NewStore(c)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestUserRepo_Create(t *testing.T) {
	type args struct {
		userName string
	}
	r := postgres2.NewUserRepo(store)
	tests := []struct {
		name       string
		args       args
		idExpected int
		wantErr    bool
	}{
		{
			name: "filling 1",
			args: args{
				userName: "1",
			},
			idExpected: 1,
			wantErr:    false,
		},
		{
			name: "filling 2",
			args: args{
				userName: "2",
			},
			idExpected: 2,
			wantErr:    false,
		},
		{
			name: "empty userName",
			args: args{
				userName: "",
			},
			idExpected: 0,
			wantErr:    true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := r.Create(test.args.userName)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.idExpected, got)
			}

		})
	}

}
