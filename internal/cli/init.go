package cli

import "flag"

func Init() {
	flag.BoolVar(
		&LSL,
		"lsl",
		false,
		"",
	)
	flag.Parse()
}
