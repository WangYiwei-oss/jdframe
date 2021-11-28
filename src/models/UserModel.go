package models

type UserModel struct {
	UserName string
	Password string
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (u *UserModel)JModelName()string{
	return "UserModel"
}
