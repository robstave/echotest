package models

// User
type User struct {
	UserName  string
	FirstName string
	LastName  string
	ID        int32
}

// NewUser
func NewUser(userName string, firstName string, lastName string, id int32) *User {
	u := &User{
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
		ID:        id,
	}

	return u

}

func (u *User) GetId() int32 {
	return u.ID
}
