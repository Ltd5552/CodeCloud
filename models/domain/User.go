package domain

type User struct {
	userID   string
	username string
	password string
}

func (u *User) UserID() string {
	return u.userID
}

func (u *User) SetUserID(userID string) {
	u.userID = userID
}

func (u *User) Username() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}
