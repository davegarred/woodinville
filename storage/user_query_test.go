package storage

import (
	"testing"
	"github.com/davegarred/woodinville/domain"
	"github.com/stretchr/testify/assert"
)

var (
	userEL   = &UserQueryEventListener{}
	userId   = domain.UserId("a_user_id")
	userName = "Named User"
	aug06    = "2018-08-06T14:56:04-07:00"
)

func TestUserQueryEventListener_OnUserCreated(t *testing.T) {
	users = make(map[domain.UserId]*UserQuery)

	userEL.OnUserCreated(domain.UserCreated{userId, userName})

	user := FindUser(userId)
	assert.Equal(t, testUser(), user)
}

func TestUserQueryEventListener_OnUserIsSetAsAdmin(t *testing.T) {
	users = make(map[domain.UserId]*UserQuery)
	userEL.OnUserCreated(domain.UserCreated{userId, userName})

	userEL.OnUserIsSetAsAdmin(domain.UserIsSetAsAdmin{userId})

	expected := testUser()
	expected.Admin = true
	assert.Equal(t, expected, FindUser(userId))

	userEL.OnUserIsSetAsNotAdmin(domain.UserIsSetAsNotAdmin{userId})

	expected.Admin = false
	assert.Equal(t, expected, FindUser(userId))
}

func TestUserQueryEventListener_OnVisitAdded(t *testing.T) {
	users = make(map[domain.UserId]*UserQuery)
	userEL.OnUserCreated(domain.UserCreated{userId, userName})

	userEL.OnVisitAdded(domain.VisitAdded{userId, aug06, wineryId})

	expected := testUser()
	expected.Visits[wineryId] = []string{aug06}
	assert.Equal(t, expected, FindUser(userId))
}

func testUser() *UserQuery {
	return &UserQuery{userId, userName, false, make(map[domain.WineryId][]string)}
}