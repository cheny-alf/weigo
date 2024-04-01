package dao

import (
	"wego/apps/service/user/model/ent"
)

type DBDAO struct {
	User *UserDao
}

func NewDBDAO(db *ent.Client) *DBDAO {
	return &DBDAO{
		User: NewUserDao(db),
	}
}
