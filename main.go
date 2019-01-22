package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	BuildDate  = ""
	BuildEpoch = ""
	BuildHash  = ""
)

func main() {
	flags := struct {
		helpShort *bool
		helpLong  *bool
		version   *bool
	}{
		helpShort: flag.Bool("h", false, "Print help and exit"),
		helpLong:  flag.Bool("help", false, "Print help and exit"),
		version:   flag.Bool("version", false, "Print version and exit"),
	}
	flag.Parse()

	if *flags.helpShort || *flags.helpLong {
		fmt.Println(msgHelp())
		return
	}

	if *flags.version {
		fmt.Println(msgVersion())
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Println(msgHelp())
		return
	}

	switch flag.Args()[0] {
	case "start":
		cmdStart()

	case "stop":
		cmdStop()

	case "status":
		cmdStatus()

	case "report":
		cmdReport()

	default:
		fmt.Println(msgHelp())
	}
}

func cmdStart()  {}
func cmdStop()   {}
func cmdStatus() {}
func cmdReport() {}

}

func msgHelp() string {
	return strings.TrimSpace(`
owl - A simple work tracker

USE:

	owl [OPTION...] COMMAND

COMMANDS:

	start		Start a session
	stop		Stop a session
	status		Print owl status
	report		Print a report of work done

OPTIONS:

	-h, --help	Print help and exit
	--version	Print version and exit
`)
}

func msgVersion() string {
	if BuildDate == "" && BuildEpoch == "" && BuildHash == "" {
		return "Development version"
	}

	if BuildDate == "" {
		BuildDate = "missing"
	}

	if BuildEpoch == "" {
		BuildEpoch = "missing"
	}

	if BuildHash == "" {
		BuildHash = "missing"
	}

	return strings.TrimSpace(fmt.Sprintf(`
Build date: %s
Build epoch: %s
Build commit hash: %s
`, BuildDate, BuildEpoch, BuildHash))
}
