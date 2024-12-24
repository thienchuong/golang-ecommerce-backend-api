package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// ur user repo
func (ur *UserRepo) GetInFoUser() string {
	return "GetInFoUser toni the best"
}
