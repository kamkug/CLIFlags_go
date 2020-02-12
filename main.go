package flags

import (
	"fmt"
	"log"
	"os"
)

type Flag struct {
	//	Flags []string //list of supported flags
	Id       string
	Names    []string    // name or names for this flag
	Defaults interface{} // default value that this flag takes
	Help     string      // help sentence
	Verify   func(string) bool
}

type FlagSet struct {
	SetName string
	Flags   []Flag
	Count   int
}

//FlagsValues collects all user defined values for flags and uses default
//values for not defined|mentioned
func (fs FlagSet) FlagsValues(args []string) (bool, map[string]interface{}) {
	fVals := make(map[string]interface{})
	flagi := fs.Flags
	for i, _ := range flagi {
		for _, v1 := range flagi[i].Names {
			var value string
			for j, _ := range args {
				if nextCLIArg := args[j]; nextCLIArg == v1 {
					if j+1 <= len(args)-1 && args[j+1][0] != 45 {
						value = args[j+1]
						//fVals[flagi[i].Id] = args[j+1]          //verify that the particular flag ID has corect input

						//} else {
						//	fVals[flagi[i].Id] = flagi[i].Defaults
						//}
					}

				}
				fmt.Println(value)
				if flagi[i].Verify(value) {
					fVals[flagi[i].Id] = value
					return true, fVals
				} else {
					fVals[flagi[i].Id] = flagi[i].Defaults
					return true, fVals
				}
			}
		}
	}
	//if len(fVals) < len(fs.Flags) {
	//	return false, fVals
	//}
	return true, fVals
}

//CreateFlag creates a single flag
func CreateFlag(id string, n []string, d interface{}, h string, v func(string) bool) Flag {
	flag := Flag{
		Id:       id,
		Names:    n,
		Defaults: d,
		Help:     h,
		Verify:   v,
	}
	return flag
}

//PopulateSet adds a single flag to a set
func PopulateSet(fSet FlagSet, flag Flag) FlagSet {

	length := len(fSet.Flags) + 1
	set := append(fSet.Flags, flag)

	flagSet := FlagSet{
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
func (f FlagSet) VerifyLength(n int) (bool, error) {
	if len(os.Args) < (n * 2) {
		return false, &LengthError{Length: n}
	}

	return true, nil

}

//GetArgs returns the slice of arguments
func (f FlagSet) GetArgs() []string {
	return os.Args
}

//CheckFlag returns the value assigned to the flag
func (f FlagSet) CheckFlag(s string) (string, bool) {
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
