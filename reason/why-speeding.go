// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                      _ _             
// / ___| _ __   ___  ___  __| (_)_ __   __ _ 
// \___ \| '_ \ / _ \/ _ \/ _` | | '_ \ / _` |
//  ___) | |_) |  __/  __/ (_| | | | | | (_| |
// |____/| .__/ \___|\___|\__,_|_|_| |_|\__, |
//       |_|                            |___/ 

package reason
import "strings"
import "libsisimai.org/sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["Speeding"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"mail sent from your IP address has been temporarily rate limited",
			"please try again slower",
			"receiving mail at a rate that prevents additional messages from being delivered",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "speeding" or not
	ProbesInto["Speeding"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is speeding, false: is not speeding

		// Action: failed
		// Status: 4.7.1
		// Remote-MTA: dns; smtp.example.jp
		// Diagnostic-Code: smtp; 451 4.7.1 <mx.example.org[192.0.2.2]>: Client host rejected: Please try again slower
		if fo        == nil        { return false }
		if fo.Reason == "speeding" { return true  }
		return IncludedIn["Speeding"](strings.ToLower(fo.DiagnosticCode))
	}
}

