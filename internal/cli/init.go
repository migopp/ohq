package cli

import "flag"

func Init() {
	flag.BoolVar(
		&LSL,
		"lsl",
		false,
		"Load student data from `students.yaml`",
	)
	flag.Parse()
}
