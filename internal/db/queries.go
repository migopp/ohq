package db

// `CreateUser` adds `u` to the SQLite DB, updating
// `u` if successful, and returning an error if not.
func CreateUser(u *User) error {
	r := db.Create(u)
	return r.Error
}

// `FetchUserWithID` fetches `u` with name `un` from the
// SQLite DB, updating `u` if successful, and returning
// an error if not.
func FetchUserWithName(un string) (User, error) {
	var u User
	r := db.Where("username = ?", un).First(&u)
	return u, r.Error
}

// `FetchUserWithID` fetches `u` with id `id` from the
// SQLite DB, updating `u` if successful, and returning
// an error if not.
func FetchUserWithID(id uint) (User, error) {
	var u User
	r := db.First(&u, id)
	return u, r.Error
}

// `FetchAllUsers` fetches all users from the SQLite DB
// and stores them in `u`.
func FetchAllUsers() ([]User, error) {
	var u []User
	r := db.Find(&u)
	return u, r.Error
}
