package postgres_test

import (
	"chatServer/internal/models"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createChat(t *testing.T, chat *models.Chat, users []models.User, isCreateUsers bool) (id int, err error) {
	if isCreateUsers {
		for i := range users {
			createUser(t, &users[i])
			assert.NoError(t, err)
			assert.NotEmpty(t, users[i].Id)
		}
	}
	chat.UsersId = make([]int64, 0)
	for _, user := range users {
		chat.UsersId = append(chat.UsersId, int64(user.Id))
	}
	if chat.Id != 0 {
		return
	}
	chat.Id, err = repo.Chat.Create(chat.Name, users)
	return chat.Id, err
}

func TestChatRepo_Create(t *testing.T) {
	tests := []struct {
		name        string
		chat        *models.Chat
		users       []models.User
		createUsers bool
		wantErr     bool
	}{
		{
			name: "create empty chat",
			chat: &models.Chat{
				Name: fake.CharactersN(15),
			},
			users:       nil,
			createUsers: false,
			wantErr:     false,
		},
		{
			name: "create chat with two users",
			chat: &models.Chat{
				Name: fake.CharactersN(15),
			},
			users:       []models.User{{UserName: fake.UserName()}, {UserName: fake.UserName()}, {UserName: fake.UserName()}},
			createUsers: true,
			wantErr:     false,
		},
		{
			name:        "not existed user",
			chat:        &models.Chat{Name: fake.CharactersN(15)},
			users:       []models.User{{Id: 100, UserName: fake.UserName()}},
			createUsers: false,
			wantErr:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := createChat(t, test.chat, test.users, test.createUsers)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}
		})
	}
}

func TestChatRepo_GetById(t *testing.T) {
	tests := []struct {
		name       string
		chat       *models.Chat
		users      []models.User
		createChat bool
		wantErr    bool
	}{
		{
			name:       "get normal chat",
			chat:       &models.Chat{Name: fake.CharactersN(15)},
			users:      []models.User{{UserName: fake.UserName()}, {UserName: fake.UserName()}},
			wantErr:    false,
			createChat: true,
		},
		{
			name: "get not existed chat",
			chat: &models.Chat{Name: fake.CharactersN(15),
				Id: 1000,
			},
			users:      nil,
			wantErr:    true,
			createChat: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			test.chat.Id, err = createChat(t, test.chat, test.users, true)
			assert.NoError(t, err)
			got, err := repo.Chat.GetById(test.chat.Id)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.chat.Name, got.Name)
				assert.Equal(t, test.chat.Id, got.Id)
			}

		})
	}
}
