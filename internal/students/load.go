package students

import (
	"fmt"
	"log"
	"os"

	"github.com/migopp/ohq/internal/db"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

func Load() {
	var students Students
	var err error

	// Open `students.yaml`
	f, err := os.Open("students.yaml")
	if err != nil {
		log.Fatal("Failed to open `students.yaml`")
	}
	defer f.Close()

	// Decode
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&students); err != nil {
		log.Fatal("Failed to decode `students.yaml`")
	}

	// Load student data into the DB
	for _, s := range students {
		// Encrypt `pass`
		epass, err := bcrypt.GenerateFromPassword([]byte(s.Password), 10)
		if err != nil {
			log.Fatalf("Failed to encrypt password: %v", err)
		}
		s.Password = string(epass)

		// Insert data
		u := db.User{
			Username: s.CSID,
			Password: s.Password,
		}
		db.CreateUser(&u)
	}

	// Fetch all to debug
	var u []db.User
	db.FetchAllUsers(&u)
	for _, s := range u {
		fmt.Printf("USER FOUND: %+v\n\n", s)
	}
}
