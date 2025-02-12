// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _       _   ____                     _          _ 
// | \ | | ___ | |_|  _ \  ___  ___ ___   __| | ___  __| |
// |  \| |/ _ \| __| | | |/ _ \/ __/ _ \ / _` |/ _ \/ _` |
// | |\  | (_) | |_| |_| |  __/ (_| (_) | (_| |  __/ (_| |
// |_| \_|\___/ \__|____/ \___|\___\___/ \__,_|\___|\__,_|

package sis
import "fmt"
import "time"
import "runtime"

// NotDecoded{} is a structure keeping a decoding error at sisimai.Rise()
type NotDecoded struct {
	EmailFile string    // An email file name sisimai tried to decoded
	CalledOff bool      // Unrecoverable error, the decoding process have called off
	BecauseOf string    // An error message of the failure
	WhoCalled string    // Who called the constructor?
	DecodedBy string    // Copy of sis.Fact.DecodedBy
	Timestamp time.Time // When the error occurred
}

// MakeNotDecoded() is a constructor of sis.NotDecoded{}
func MakeNotDecoded(argv0 string, argv1 bool) *NotDecoded {
	// @param    string argv0  Error message
	// @param    bool   argv0  Unrecoverable error or not
	// @return *sis.NotDecoded
	p, _, l, _ := runtime.Caller(1)
	return &NotDecoded{
		BecauseOf: argv0,
		CalledOff: argv1,
		Timestamp: time.Now(),
		WhoCalled: fmt.Sprintf("%s():%d", runtime.FuncForPC(p).Name(), l),
	}
}

// Error() returns the error message as a string
func(this *NotDecoded) Error() string {
	// @param    NONE
	// @return   string  an error message
	if this.BecauseOf == "" { return "" }

	timestring:= this.Timestamp.Format("2006/01/02 15:04:05")
	return fmt.Sprintf("%s %s %s", timestring, this.EmailFile, this.BecauseOf)
}

// Label() returns a label string for printing error message
func(this *NotDecoded) Label() string {
	// @param    NONE
	// @return   string  A label for printing the error message
	if this.CalledOff == true {
		// The error is unrecoverable
		return " *****error: "
	} else {
		// The error is not unrecoverable
		return " ***warning: "
	}
}

// Email() receives a path to email and set it into EmailFile
func(this *NotDecoded) Email(argv1 string) string {
	// @param    string  argv1  A path to an email being set into the EmailFile
	// @return   string         The current value of the EmailFile
	if argv1          == "" { return this.EmailFile  }
	if this.EmailFile == "" { this.EmailFile = argv1 }
	return this.EmailFile
}

