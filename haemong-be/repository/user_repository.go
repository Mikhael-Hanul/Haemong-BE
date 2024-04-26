package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserEntity struct {
	userId   string
	password string
	name     string
}

func (r *UserRepository) SaveUser(userId, password, name string) error {
	_, err := r.db.Exec("insert into tbl_user value (?, ?, ?)", userId, password, name)
	if err != nil {
		return fmt.Errorf("유저 등록에 실패함 : " + err.Error())
	}
	return nil
}

func (r *UserRepository) IsUserIdDuplicate(userId string) bool {
	var u UserEntity
	err := r.db.QueryRow("select * from tbl_user where userId = ?", userId).Scan(&u.userId, &u.password, &u.name)
	return err == nil
}
