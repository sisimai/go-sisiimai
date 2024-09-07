// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _       _      _                      _   
// | \ | | ___ | |_   / \   ___ ___ ___ _ __ | |_ 
// |  \| |/ _ \| __| / _ \ / __/ __/ _ \ '_ \| __|
// | |\  | (_) | |_ / ___ \ (_| (_|  __/ |_) | |_ 
// |_| \_|\___/ \__/_/   \_\___\___\___| .__/ \__|
//                                     |_|        
import "strings"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["NotAccept"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			"does not accept mail (nullmx)",
			"host/domain does not accept mail", // iCloud
			"host does not accept mail",        // Sendmail
			"mail receiving disabled",
			"name server: .: host not found",   // Sendmail
			"no mx record found for domain=",   // Oath(Yahoo!)
			"no route for current request",
			"smtp protocol returned a permanent error",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "notaccept" or not
	Truth["NotAccept"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is notaccept, false: is not notaccept
		if fo.Reason == "notaccept"                                                { return true  }
		if fo.ReplyCode == "521" || fo.ReplyCode == "554" || fo.ReplyCode == "556" { return true  }
		if fo.SMTPCommand != "MAIL"                                                { return false }
		return Match["NotAccept"](strings.ToLower(fo.DiagnosticCode))
	}
}

