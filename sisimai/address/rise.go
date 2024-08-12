// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address
import "fmt"
import "strings"

type EmailAddress struct {
	Address string // Email address
	User    string // Local part of the email addres
	Host    string // Domain part of the email address
	Verp    string // VERP
	Alias   string // Alias of the email address
	Name    string // Display name
	Comment string // (Comment)
}

// Rise() is a constructor of Sisimai::Address
func Rise(argvs [3]string) EmailAddress {
	// @param    [3]string argvs  ["Email address", "name", "comment"]
	// @return   EmailAddress     EmailAddress struct when the email address was not valid
	if len(argvs[0]) == 0 { return EmailAddress{} }

	thing := new(EmailAddress)
	heads := "<"
	tails := ">,.;"
	point := strings.LastIndex(argvs[0], "@")

	if point > 0 {
		// Get the local part and the domain part from the email address
		lpart := argvs[0][:point]
		dpart := argvs[0][point+1:]
		email := ExpandVERP(argvs[0]);
		alias := false

		if len(email) == 0 {
			// Is not a VERP address, try to expand the address as an alias
			email = ExpandAlias(argvs[0])
			if len(email) > 0 { alias = true }
		}

		if strings.Contains(email, "@") {
			// The address is a VERP or an alias
			if alias {
				// The address is an alias like "neko+nyaan@example.jp"
				thing.Alias = argvs[0]

			} else {
				// The address is a VERP: b+neko=example.jp@example.org
				thing.Verp  = argvs[0]
			}
		}

		// Remove the folowing characters: "<", ">", ",", ".", and ";" from the email address
		lpart = strings.TrimLeft(lpart, heads)
		dpart = strings.TrimRight(dpart, tails)

		thing.User    = lpart
		thing.Host    = dpart
		thing.Address = fmt.Sprintf("%s@%s", lpart, dpart)

	} else {
		// The argument does not include "@"
		if IsMailerDaemon(argvs[0]) == false { return EmailAddress{} }
		if strings.Contains(argvs[0], " ")   { return EmailAddress{} }

		// The argument does not include " "
		thing.User    = argvs[0]
		thing.Address = argvs[0]
	}

	thing.Name    = argvs[1]
	thing.Comment = argvs[2]
	fmt.Printf("THING = [%q]\n", thing)
	return *thing
}

// *EmailAddress.Void() returns true if it does not include a valid email address
func(this *EmailAddress) Void() bool {
	if len(this.Address) == 0 { return false }
	return true
}

