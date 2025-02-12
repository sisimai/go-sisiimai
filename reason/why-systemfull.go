// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____            _                 _____      _ _ 
// / ___| _   _ ___| |_ ___ _ __ ___ |  ___|   _| | |
// \___ \| | | / __| __/ _ \ '_ ` _ \| |_ | | | | | |
//  ___) | |_| \__ \ ||  __/ | | | | |  _|| |_| | | |
// |____/ \__, |___/\__\___|_| |_| |_|_|   \__,_|_|_|
//        |___/                                      

package reason
import "strings"
import "libsisimai.org/sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["SystemFull"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"mail system full",
			"requested mail action aborted: exceeded storage allocation", // MS Exchange
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "systemfull" or not
	ProbesInto["SystemFull"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is systemfull, false: is not systemfull
		return false
	}
}

