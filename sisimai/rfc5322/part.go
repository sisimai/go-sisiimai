// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  ____  _____ ____ ____ _________  ____  
// |  _ \|  ___/ ___| ___|___ /___ \|___ \ 
// | |_) | |_ | |   |___ \ |_ \ __) | __) |
// |  _ <|  _|| |___ ___) |__) / __/ / __/ 
// |_| \_\_|   \____|____/____/_____|_____|
import "strings"

// Part() split the entire message body given as the 1st argument into error message lines and the
// original message part only include email headers.
func Part(email *string, cutby []string, keeps bool) [2]string {
	// @param    *string  email    Entire message body
	// @param    []string cutby    String list of the message/rfc822 or the beginning of the original message part
	// @param    bool     keeps    Flag for keeping strings after "\n\n"
	// @return   []string          { "Error message lines", "The original message" }
	// @since    v5.0.0
	if len(*email) == 0 { return [2]string{"", ""} }
	if len(cutby)  == 0 { return [2]string{"", ""} }

	positionor := -1 // A position of the boudary string
	formerpart := "" // The error message part
	latterpart := "" // The original message part

	for _, e := range cutby {
		// Find a boundary string(2nd argument)] from the 1st argument
		positionor = strings.Index(*email, e); if positionor == -1 { continue }
		break
	}

	if positionor > 0 {
		// There is the boundary string in the message body
		formerpart  = (*email)[:positionor]
		latterpart  = (*email)[positionor:]

	} else {
		// Substitute the entire message to the former part when the boundary string is not included the *email
		formerpart = *email
		latterpart = ""
	}

	if len(latterpart) > 0 {
		// Remove blank lines, the message body of the original message, and append "\n" at the end
		// of the original message headers
		// 1. Remove leading blank lines
		// 2. Remove text after the first blank line: \n\n
		// 3. Append "\n" at the end of test block when the last character is not "\n"
		for _, e := range strings.Split(latterpart, "") {
			// Remove leading blank lines
			if e == " " || e == "\n" || e == "\r" { continue }

			p := strings.Index(latterpart, e)
			if p > 0 {
				// There is leading space characters at the head of parts[1]
				latterpart = latterpart[p:len(latterpart)]
			}
			break
		}

		if keeps == true && strings.Contains(latterpart, "\n\n") {
			// Remove text after the first blank line when "keeps" is true
			latterpart = latterpart[0:strings.Index(latterpart, "\n\n") + 1]
		}

		if strings.HasSuffix(latterpart, "\n") == false {
			// Append "\n" at the end of the original message
			latterpart += "\n"
		}
	}
	return [2]string{formerpart, latterpart}
}

