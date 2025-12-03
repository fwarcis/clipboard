package cli

import "os"

var (
	ExecPath = os.Args[0]
	Args     = os.Args[1:]
	LastArg  = os.Args[len(os.Args)-1]
)

