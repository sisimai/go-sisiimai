// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____                       _ _     _           _ _   
// | ____|_  _____ ___  ___  __| | |   (_)_ __ ___ (_) |_ 
// |  _| \ \/ / __/ _ \/ _ \/ _` | |   | | '_ ` _ \| | __|
// | |___ >  < (_|  __/  __/ (_| | |___| | | | | | | | |_ 
// |_____/_/\_\___\___|\___|\__,_|_____|_|_| |_| |_|_|\__|
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"
import sisimoji "sisimai/string"

func init() {
	// Try to match that the given text and message patterns
	Match["ExceedLimit"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{"message header size exceeds limit", "message too large"}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "exceedlimit" or not
	Truth["ExceedLimit"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is exceedlimit, false: is not exceedlimit
		if fo.Reason == "exceedlimit"                      { return true }
		if status.Name(fo.DeliveryStatus) == "exceedlimit" { return true }
		return Match["ExceedLimit"](strings.ToLower(fo.DiagnosticCode))
	}
}

