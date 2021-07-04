package service

import "yousync/database"

const (
	PublicUid      = "-1"
	PublicUsername = "public"
)

func GetUserById(uid string) (*database.User, error) {
	var user database.User
	err := database.Instance.Where(map[string]string{"uid": uid}).FirstOrCreate(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
