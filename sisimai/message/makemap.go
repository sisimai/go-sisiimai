// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package message

//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      
import "strings"
import "net/mail"
import sisimoji "sisimai/string"

// makemap() converts a mail.Header struct to a map[string][]string
func makemap(argv0 *mail.Header, argv1 bool) map[string][]string {
	// @param    *mail.Header      argv0 Email header data
	// @param    bool              argv1 Decode "Subject:" header or not
	// @return   map[string]string       Structured email header data
	headermaps := map[string][]string{}
	receivedby := []string{}

	for e, v := range *argv0 {
		// Each key name is the lower-cased string, each value is an array ([]string{})
		// The field name of an email header does not contain " "
		f := strings.ToLower(e)
		if strings.Contains(f, " ")               { continue }
		if strings.Contains(f, "authentication-") { continue } // Authentication-Results:
		if strings.HasPrefix(f, "arc-")           { continue } // ARC-Seal:, ARC-Authentication-Results:, ....
		if strings.HasPrefix(f, "dkim-")          { continue } // DKIM-Sigunature:
		if strings.HasPrefix(f, "-spf")           { continue } // Received-SPF:
		headermaps[f] = v
	}

	if len(headermaps["received"]) > 0 {
		for _, e := range headermaps["received"] {
			// 1. Exclude the Received header including "(qmail ** invoked from network)".
			// 2. Convert all consecutive spaces and line breaks into a single space character.
			if strings.Contains(e, " invoked by uid") || strings.Contains(e, " invoked from network") {
				// Do not include the value of Received: header generated by qmail or qmail-clone
				continue
			}
			for strings.Contains(e, "  ") { e = strings.ReplaceAll(e, "  ", " ") }
			receivedby = append(receivedby, e)
		}
		headermaps["received"] = receivedby
	}

	// The following fields should be exist
	if len(headermaps["from"])       == 0 { headermaps["from"]       = []string{""} }
	if len(headermaps["received"])   == 0 { headermaps["received"]   = []string{""} }
	if len(headermaps["message-id"]) == 0 { headermaps["message-id"] = []string{""} }
	if len(headermaps["subject"])    == 0 { headermaps["subject"]    = []string{""}; return headermaps }
	if argv1 == false { return headermaps }

	if sisimoji.Is8Bit(&(headermaps["subejct"][0])) {
		// The "Subject:" header is including multibyte character, is not a MIME-Encoded text.
		// TODO: Remove all the invalid byte sequence
	} else {
		// The "Subejct:" field is MIME-Encoded text or including only ASCII characters
		// TODO: https://pkg.go.dev/mime
	}

	return headermaps
}

