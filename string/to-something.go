// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All
// rights reserved. This software is distributed under The BSD 2-Clause License.
//      _        _             
//  ___| |_ _ __(_)_ __   __ _ 
// / __| __| '__| | '_ \ / _` |
// \__ \ |_| |  | | | | | (_| |
// |___/\__|_|  |_|_| |_|\__, |
//                       |___/ 

package string
import "fmt"
import "strings"
import "golang.org/x/text/encoding"

// ToLF() replace CR and CR/LF to LF.
func ToLF(argv0 *string) *string {
	// @param    *string argv0  Text including CR or CR/LF
	// @return   *string        LF converted text
	if argv0 == nil || *argv0 == "" { return argv0 }

	crreplaced := *argv0; for _, e := range [2]string{"\r\n", "\r"} {
		// Convert CRLF and CR to LF
		if strings.Contains(crreplaced, e) == false { continue }
		crreplaced = strings.ReplaceAll(crreplaced, e, "\n")
	}
	return &crreplaced
}

// ToPlain() converts given HTML text to a plain text.
func ToPlain(argv0 *string) *string {
	// @param    [*string] argv0  HTML text
	// @return   [*string]        Plain text
	if argv0 == nil || *argv0 == "" { return argv0 }

	xhtml := *argv0
	lower := strings.ToLower(*argv0); if strings.Contains(lower, "<body") == false { return argv0 }
	plain := ""
	body0 := -1; for _, e := range []string{">", " ", "\t", "\n"} {
		// Find the position of <body?, and remove the HTML header part
		body0  = strings.Index(lower, "<body" + e); if body0 < 0 { continue }
		body0 += len("<body>") + 1

		if e != ">" { body0 = IndexOnTheWay(lower, ">", body0) + 1 }
		xhtml = xhtml[body0:]
		lower = strings.ToLower(xhtml)

		// Remove string from <style> to </style>
		p0 := strings.Index(lower, "<style");  if p0 < 0 { break }
		p1 := strings.Index(lower, "</style"); if p1 < 0 { break }
		xhtml = xhtml[:p0] + xhtml[p1 + 8:]
	}

	for strings.IndexByte(xhtml, '<') > -1 || strings.IndexByte(xhtml, '>') > -1 {
		// Find "<" from HTML element and remove string until ">"
		p0 := strings.IndexByte(xhtml, '<');     if p0 < 0 { break }
		p1 := IndexOnTheWay(xhtml, ">", p0 + 2); if p1 < 0 { break }

		if p0 >  0 { plain += xhtml[0:p0] + " "      }
		if p0 > p1 { plain += xhtml[p1 + 1:p0] + " " }

		xhtml = xhtml[p1 + 1:]
	}

	// Remove or replace entity references
	plain = strings.ReplaceAll(plain, "&lt;",   "<")
	plain = strings.ReplaceAll(plain, "&gt;",   ">")
	plain = strings.ReplaceAll(plain, "&quot;", `"`)
	plain = strings.ReplaceAll(plain, "&nbsp;", " ")
	plain = strings.ReplaceAll(plain, "&copy;", "(C)")
	plain = strings.ReplaceAll(plain, "&amp;",  "&")
	plain = Sweep(strings.ReplaceAll(plain, "\n", " "))
	return &plain
}

// ToUTF8() converts an encoded text to UTF8 text
func ToUTF8(argv0 []byte, argv1 string) (string, error) {
	// @param    []byte argv0     Some encoded text
	// @param    string argv1     Encoding name of the argv0
	// @return   string, error    Converted string or an error
	if len(argv0) == 0  || argv1 == ""         { return "", nil            }
	if argv1 == "utf-8" || argv1 == "us-ascii" { return string(argv0), nil }

	var encodingif *encoding.Decoder; if encodingif == nil {
		ce := fmt.Errorf("Unsupported encoding: %s, see https://github.com/sisimai/go-sisimai/issues/42", argv1)
		return string(argv0), ce
	}

	utf8string := make([]byte, len(argv0) * 3)
	rightindex, _, nyaan := encodingif.Transform(utf8string, argv0, false)
	if nyaan != nil { return string(argv0), nyaan }
	return string(utf8string[:rightindex]), nil
}

