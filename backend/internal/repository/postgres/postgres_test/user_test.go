package postgres_test

import (
	"chatServer/internal/models"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createUser(t *testing.T, user *models.User) {
	if user.Id != 0 {
		return
	}
	var err error
	user.Id, err = repo.User.Create(user.UserName)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Id)
}

func TestUserRepo_GetById(t *testing.T) {
	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		{
			name: "get not existed user",
			user: &models.User{
				Id: 100,
			},
			wantErr: true,
		},
		{
			name: "wrong incoming data",
			user: &models.User{
				Id: -10,
			},
			wantErr: true,
		},
		{
			name: "get full user info",
			user: &models.User{
				UserName: fake.UserName(),
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			createUser(t, test.user)
			got, err := repo.User.GetById(test.user.Id)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.user.UserName, got.UserName)
				assert.Equal(t, test.user.Id, got.Id)
			}
		})
	}

}

func TestUserRepo_Create(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "filling 1",
			args: args{
				userName: fake.UserName(),
			},
			wantErr: false,
		},
		{
			name: "filling 2",
			args: args{
				userName: fake.CharactersN(100),
			},
			wantErr: false,
		},
		{
			name: "empty userName",
			args: args{
				userName: "",
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := repo.User.Create(test.args.userName)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}

		})
	}
}
