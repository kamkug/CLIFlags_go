package flags

import (
	"fmt"
	"log"
	"os"
)

type Flag struct {
	//	Flags []string //list of supported flags
	Names    []string    // name or names for this flag
	Defaults interface{} // default value that this flag takes
	Help     string      // help sentence

}

type Flags struct {
	SetName string
	Flags   []Flag
	Count   int
}

//CreateFlagStr creates a single flag
func CreateFlag(n []string, d interface{}, h string) Flag {
	flag := Flag{
		Names:    n,
		Defaults: d,
		Help:     h,
	}
	return flag
}

//PopulateSet adds a single flag to a set
func PopulateSet(fSet Flags, flag Flag) Flags {

	length := len(fSet.Flags) + 1
	set := append(fSet.Flags, flag)

	flagSet := Flags{
		SetName: flag.Names[0],
		Flags:   set,
		Count:   length,
	}
	return flagSet
}

//LengthError creates a custom error struct
type LengthError struct {
	Length int
	Err    error
}

//Error allows LengthError to implement an Error method becoming an Error Interface type as well
func (e *LengthError) Error() string {
	return "[-] Need at least " + string(e.Length+48) + " arguments: " + ": too short"
}

//Verify the arg length
func (f Flags) VerifyLength(n int) (bool, error) {
	if len(os.Args) < (n * 2) {
		return false, &LengthError{Length: n}
	}

	return true, nil

}

//GetArgs returns the slice of arguments
func (f Flags) GetArgs() []string {
	return os.Args
}

//CheckFlag returns the value assigned to the flag
func (f Flags) CheckFlag(s string) (string, bool) {
	argsList := f.GetArgs()
	for i, arg := range argsList {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[-]Please provide a value for the %v\n", s)
				log.Fatalf("[Error] Please provide a value for the %v\n", s)
			}
		}()
		if arg == s {
			return argsList[i+1], true
		}
	}
	return "[-] No value assigned to  " + s, false
}
