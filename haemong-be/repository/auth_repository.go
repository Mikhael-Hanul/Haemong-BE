package repository

import (
	"database/sql"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
func (r *AuthRepository) IsUserPasswordMatching(userId, password string) bool {
	var u UserEntity
	err := r.db.QueryRow("select * from tbl_user where userId = ? and password = ?", userId, password).Scan(&u.UserId, &u.Password, &u.Name)
	return err == nil
}
