package db

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

func LoadFromFile(lf string) {
	var users []User
	var err error

	// Open `students.yaml`
	f, err := os.Open(lf)
	if err != nil {
		log.Fatalf("Failed to open `%s`", lf)
	}
	defer f.Close()

	// Decode
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&users); err != nil {
		log.Fatalf("Failed to decode `%s`", lf)
	}

	// Load student data into the DB
	for _, u := range users {
		// Encrypt `pass`
		epass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			log.Fatalf("Failed to encrypt password: %v", err)
		}
		u.Password = string(epass)

		// Insert data
		CreateUser(&u)
	}
}
