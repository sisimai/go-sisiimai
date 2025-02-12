// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 

package rfc2045
import "io"
import "fmt"
import "mime"
import "strings"
import "mime/quotedprintable"
import "golang.org/x/net/html/charset"

// IsEncoded() checks that the argument is MIME encoded string or not.
func IsEncoded(argv0 string) bool {
	// @param    string    argv0   String to be checked that it is MIME encoded or not
	// @return   bool              true: Not MIME encoded string
    //                             false: MIME encoded string
	argv0  = strings.ToUpper(argv0)
	match := false
	for {
		// =?UTF-8?B?44OL44Oj44O844Oz?=
		if !strings.Contains(argv0, "=?") { break } // Begins with "=?"
		if !strings.Contains(argv0, "?=") { break } // Ends with "?="
		if len(argv0) < 8                 { break } // String length should be 8 or more
		if strings.Contains(argv0, "?B?") || strings.Contains(argv0, "?Q?") { match = true }
		break
	}
	return match
}

// DecodeH() decodes the value of email header which is a MIME-Encoded string.
func DecodeH(argv0 string) (string, error) {
	// @param    string    argvs  MIME-Encoded text
	// @return   string           MIME-Decoded text
	toreadable := "" // Human readble text (has decoded)
	stringlist := []string{}
	replacingc := []string{".", "[", "]"}
	decodingif := new(mime.WordDecoder); if CharacterSet(argv0) != "UTF-8" {
		// The character set is not UTF-8
		decodingif.CharsetReader = func(c string, v io.Reader) (io.Reader, error) {
			eo, _ := charset.Lookup(c)
			return eo.NewDecoder().Reader(v), nil
		}
	}

	if strings.Contains(argv0, " ") {
		// The argument string include 1 or more space characters
		stringlist = strings.Split(argv0, " ")

	} else {
		// The argument string does not contain any space characters
		stringlist = append(stringlist, argv0)
	}

	for j, e := range stringlist {
		// Check and decode each part of the string
		if IsEncoded(e) == false {
			// Is not MIME-Encoded text part
			if j > 0 { toreadable += " " }; toreadable += e
			continue
		}

		// MIME-Encoded text part such as "=?UTF-8?B?44OL44Oj44O844Oz?="
		if strings.HasPrefix(e, "=?") == false {
			// For example, "[=?UTF-8?B?...]"
			for _, c := range replacingc { e = strings.Replace(e, c + "=?", "=?", -1) }
		}

		if strings.HasSuffix(e, "?=") == false {
			// For example, "=?UTF-8?B?....?=."
			for _, c := range replacingc { e = strings.Replace(e, "?=" + c, "?=", -1) }
		}

		if f, nyaan := decodingif.DecodeHeader(e); nyaan == nil {
			// Successfully decoded
			if j > 0 { toreadable += " " }; toreadable += f

		} else {
			// Failed to decode
			if j > 0 { toreadable += " " }; toreadable += e
		}
	}
	return toreadable, nil
}

// DecodeB() decodes Base64 encoded text.
func DecodeB(argv0 string, argv1 string) (string, error) {
	// @param    string     argv0  Base64 Encoded text
	// @param    string     argv1  Character set name
	// @return   string            MIME-Decoded text
	if len(argv0)  < 8 { return argv0, nil }
	if len(argv1) == 0 { argv1 = "utf-8"   }

	decodingif := new(mime.WordDecoder)
	base64text := strings.TrimSpace(argv0)
	base64text  = strings.Join(strings.Split(base64text, "\n"), "")
	base64text  = fmt.Sprintf("=?%s?B?%s?=", argv1, base64text)
	plainvalue := ""

	if plain, nyaan := decodingif.Decode(base64text); nyaan != nil {
		// Failed to decode the base64-encoded text
		return "", nyaan

	} else {
		// Successfully decoded
		plainvalue = plain
	}
	return plainvalue, nil
}

// DecodeQ() decodes Quoted-Pritable encdoed text
func DecodeQ(argv0 string) (string, error) {
	// @param    string     argv0 Quoted-Printable Encoded text
	// @return   string           MIME-Decoded text
	readstring := strings.NewReader(argv0)
	decodingif := quotedprintable.NewReader(readstring)
	plainvalue := ""

	plain, nyaan := io.ReadAll(decodingif); if nyaan != nil {
		// Failed to decode the quoted-printable text
		plainvalue = argv0
	}
	if len(plain) > 0 { plainvalue = string(plain) }

	return plainvalue, nyaan
}

