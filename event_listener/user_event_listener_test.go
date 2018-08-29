package event_listener

import (
	"testing"
	"github.com/davegarred/woodinville/domain"
	"github.com/stretchr/testify/assert"
	"github.com/davegarred/woodinville/query"
	"github.com/davegarred/woodinville/storage"
)

var (
	userEL   = &UserQueryEventListener{}
	userId   = domain.UserId("a_user_id")
	userName = "Named User"
	aug06    = "2018-08-06T14:56:04-07:00"
)

func TestUserQueryEventListener_OnUserCreated(t *testing.T) {
	storage.ResetTests()

	userEL.OnUserCreated(domain.UserCreated{userId, userName})

	user := storage.FindUser(userId)
	assert.Equal(t, testUser(), user)
}

func TestUserQueryEventListener_OnUserIsSetAsAdmin(t *testing.T) {
	storage.ResetTests()
	userEL.OnUserCreated(domain.UserCreated{userId, userName})

	userEL.OnUserIsSetAsAdmin(domain.UserIsSetAsAdmin{userId})

	expected := testUser()
	expected.Admin = true
	assert.Equal(t, expected, storage.FindUser(userId))

	userEL.OnUserIsSetAsNotAdmin(domain.UserIsSetAsNotAdmin{userId})

	expected.Admin = false
	assert.Equal(t, expected, storage.FindUser(userId))
}

func TestUserQueryEventListener_OnVisitAdded(t *testing.T) {
	storage.ResetTests()
	userEL.OnUserCreated(domain.UserCreated{userId, userName})

	userEL.OnVisitAdded(domain.VisitAdded{userId, aug06, wineryId})

	expected := testUser()
	expected.Visits[wineryId] = []string{aug06}
	assert.Equal(t, expected, storage.FindUser(userId))
}

func testUser() *query.UserQuery {
	return &query.UserQuery{userId, userName, false, make(map[domain.WineryId][]string)}
}