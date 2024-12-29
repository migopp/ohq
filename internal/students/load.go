package students

import (
	"fmt"
	"log"
	"os"

	"github.com/migopp/ohq/internal/db"
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
		u := db.User{
			Username: s.CSID,
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
