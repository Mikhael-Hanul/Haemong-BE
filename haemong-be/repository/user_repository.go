package repository

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) SaveUser(userId, userName, password string) error {
	return nil
}

func (r *UserRepository) IsUserIdDuplicate(userId string) bool {
	return false
}
