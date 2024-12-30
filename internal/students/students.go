package students

type Student struct {
	CSID     string `yaml:"csid"`
	Password string `yaml:"pass"`
}

type Students []Student
