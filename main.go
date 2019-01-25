package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
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

	if err := initFolders(); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't initialize directories: %s\n", err)
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "start":
		cmdStart()

	case "stop":
		msg, err := cmdStop()
		switch err.(type) {
		case nil:
			fmt.Printf("%s\n", msg)

		case ErrCorruptedState:
			if e := os.Remove(fileState()); e != nil {
				fmt.Fprintf(os.Stderr, "Error: Couldn't remove corrupted file: %s\n", e)
			}
			fmt.Fprintf(os.Stderr, "%s\n", err)

		default:
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	case "status":
		msg, err := cmdStatus()
		switch err.(type) {
		case nil:
			fmt.Printf("%s\n", msg)

		case ErrCorruptedState:
			if e := os.Remove(fileState()); e != nil {
				fmt.Fprintf(os.Stderr, "Error: Couldn't remove corrupted file: %s\n", e)
			}
			fmt.Fprintf(os.Stderr, "%s\n", err)

		default:
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}

	case "report":
		cmdReport()

	default:
		fmt.Println(msgHelp())
	}
}

type ErrCorruptedState string

func (e ErrCorruptedState) Error() string {
	return string(e)
}

func errCorruptedState(format string, args ...interface{}) ErrCorruptedState {
	return ErrCorruptedState(fmt.Sprintf(format, args...))
}

func cmdStart() {}

func cmdStop() (string, error) {
	if !fileExists(fileState()) {
		return fmt.Sprintf("Not working"), nil
	}

	fh, err := os.Open(fileState())
	if err != nil {
		return "", fmt.Errorf("Error: %s", err)
	}

	defer fh.Close()
	s, err := LoadState(fh)
	if err != nil {
		return "", errCorruptedState("Error: State corrupted: %s", err)
	}

	if s.EpochStart == 0 {
		return "Not working", nil
	} else if s.EpochStart < 0 {
		return "", errCorruptedState("Error: State corrupted: Contains negative epoch")
	}

	start := time.Unix(s.EpochStart, 0)

	stop := time.Now()
	if err := os.Remove(fileState()); err != nil {
		// ?
	}

	return "OK", nil
}

func cmdStatus() (string, error) {
	if !fileExists(fileState()) {
		return fmt.Sprintf("Not working"), nil
	}

	fh, err := os.Open(fileState())
	if err != nil {
		return "", fmt.Errorf("Error: %s", err)
	}

	defer fh.Close()
	s, err := LoadState(fh)
	if err != nil {
		return "", errCorruptedState("Error: State corrupted: %s", err)
	}

	if s.EpochStart == 0 {
		return "Not working", nil
	} else if s.EpochStart < 0 {
		return "", errCorruptedState("Error: State corrupted: Contains negative epoch")
	}

	start := time.Unix(s.EpochStart, 0)

	return fmt.Sprintf("Working: Started at %s", start), nil
}

func cmdReport() {}

type State struct {
	EpochStart int64 `json:"epoch_start"`
}

func LoadState(r io.Reader) (State, error) {
	s := State{}
	err := json.NewDecoder(r).Decode(&s)

	return s, err
}

func DumpState(w io.Writer, s State) error {
	return json.NewEncoder(w).Encode(s)
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
