package cmdline

import (
	"flag"
)

// SubCommand zzz
type SubCommand interface {
	Init(flagSet *flag.FlagSet)
}

// ServerCmd zzz
type ServerCmd struct {
	Port int
}

// Init zzz
func (serverCmd *ServerCmd) Init(flagSet *flag.FlagSet) {
	flagSet.IntVar(&serverCmd.Port, "port", 8080, "listening port")
}

// CmdLine zzz
type CmdLine struct {
	Server *ServerCmd
}

// Parse zzz
func (cmdLine *CmdLine) Parse(args []string) {
	switch args[1] {
	case "server":
		cmdLine.Server = &ServerCmd{}
		cmdLine.Server.Init(flag.NewFlagSet("server", flag.ExitOnError))
	}
}

// NewCmdLine zzz
func NewCmdLine(args []string) *CmdLine {
	cmdLine := &CmdLine{}
	cmdLine.Parse(args)
	return cmdLine
}
