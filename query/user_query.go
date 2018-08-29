package query

import (
	"github.com/davegarred/woodinville/domain"
)

type UserQuery struct {
	domain.UserId                       `json:"id"`
	Name   string                       `json:"name"`
	Admin  bool                         `json:"admin"`
	Visits map[domain.WineryId][]string `json:"visits"`
}
