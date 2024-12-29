package db

// `CreateUser` adds `u` to the SQLite DB, updating
// `u` if successful, and returning an error if not.
func CreateUser(u *User) error {
	r := db.Create(u)
	return r.Error
}

// `FetchUserWithID` fetches `u` from the SQLite DB,
// updating `u` if successful, and returning an error if not.
func FetchUserWithID(u *User, id uint) error {
	r := db.First(u, id)
	return r.Error
}

// `FetchAllUsers` fetches all users from the SQLite DB
// and stores them in `u`.
func FetchAllUsers(u *[]User) error {
	r := db.Find(u)
	return r.Error
}
