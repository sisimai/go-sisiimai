// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      ____  __                                _          _         
//  _ __| |__   ___  ___| |_   / /  \/  | ___  ___ ___  __ _  __ _  ___| |    __ _| |__  ___ 
// | '__| '_ \ / _ \/ __| __| / /| |\/| |/ _ \/ __/ __|/ _` |/ _` |/ _ \ |   / _` | '_ \/ __|
// | |  | | | | (_) \__ \ |_ / / | |  | |  __/\__ \__ \ (_| | (_| |  __/ |__| (_| | |_) \__ \
// |_|  |_| |_|\___/|___/\__/_/  |_|  |_|\___||___/___/\__,_|\__, |\___|_____\__,_|_.__/|___/
//                                                           |___/                           
import "libsisimai.org/sisimai/sis"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["MessageLabs"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      https://www.broadcom.com/products/cybersecurity/email
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"securityerror": []string{"Please turn on SMTP Authentication in your mail client"},
			"userunknown":   []string{"542 ", " Rejected", "No such user"},
		}

		for e := range messagesof {
			// Each key is an error reason name
			if sisimoji.ContainsAny(fo.DiagnosticCode, messagesof[e]) { return e }
		}
		return ""
	}
}

