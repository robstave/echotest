package models

var users []User

func init() {

	var u1 = NewUser("rstave", "rob", "stave", 1)
	var u2 = NewUser("kali", "kali", "stave", 2)
	var u3 = NewUser("ducky", "duck", "the cat", 3)
	var u4 = NewUser("fish", "asdf", "ssga", 4)
	users = append(users, *u1)
	users = append(users, *u2)
	users = append(users, *u3)
	users = append(users, *u4)
}

func GetUser(index int) User {
	return users[index]
}

func GetUsers() []User {
	return users
}
