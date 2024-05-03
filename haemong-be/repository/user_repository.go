package repository

import (
	"database/sql"
	"errors"
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

func (r *UserRepository) FindUserPassword(userId string) (string, error) {
	var password string
	err := r.db.QueryRow("select password from tbl_user where userId = ?", userId).Scan(&password)
	if err != nil {
		return "", errors.New("유저가 존재하지 않습니다.")
	}
	return password, nil
}

func (r *UserRepository) ChangeUserPassword(userId, newPassword string) error {
	_, err := r.db.Exec("update tbl_user set password = ? where userId = ?", newPassword, userId)
	if err != nil {
		return errors.New("비밀번호 변경에 실패했습니다.")
	}
	return nil
}
