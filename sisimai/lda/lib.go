// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lda

//  _     ____    _    
// | |   |  _ \  / \   
// | |   | | | |/ _ \  
// | |___| |_| / ___ \ 
// |_____|____/_/   \_\
import "strings"
import "sisimai/sis"
import sisimoji "sisimai/string"

var LocalAgent = map[string][]string{
	// dovecot/src/deliver/deliver.c
	// 11: #define DEFAULT_MAIL_REJECTION_HUMAN_REASON \
	// 12: "Your message to <%t> was automatically rejected:%n%r"
	"dovecot":    []string{"Your message to ", " was automatically rejected:"},
	"mail.local": []string{"mail.local: "},
	"procmail":   []string{"procmail: "},
	"maildrop":   []string{"maildrop: "},
	"vpopmail":   []string{"vdelivermail: "},
	"vmailmgr":   []string{"vdeliver: "},
}
var MessagesOf = map[string]map[string][]string{
	"dovecot": map[string][]string{
		"mailboxfull": []string{
			"quota exceeded", // Dovecot 1.2 dovecot/src/plugins/quota/quota.c
			"quota exceeded (mailbox for user is full)", // dovecot/src/plugins/quota/quota.c
			"not enough disk space",
		},
		"userunknown": []string{"mailbox doesn't exist: "},
	},
	"mail.local": map[string][]string{
		"mailboxfull": []string{
			"disc quota exceeded",
			"mailbox full or quota exceeded",
		},
		"systemerror": []string{"temporary file write error"},
		"userunknown": []string{
			": unknown user:",
			": user unknown",
			": invalid mailbox path",
			": user missing home directory",
		},
	},
	"procmail": map[string][]string{
		"mailboxfull": []string{"quota exceeded while writing"},
		"systemfull":  []string{"no space left to finish writing"},
	},
	"maildrop": map[string][]string{
		"mailboxfull": []string{"maildir over quota."},
		"userunknown": []string{
			"invalid user specified.",
			"cannot find system user",
		},
	},
	"vpopmail": map[string][]string{
		"filtered":    []string{"user does not exist, but will deliver to "},
		"mailboxfull": []string{"domain is over quota", "user is over quota"},
		"suspend":     []string{"account is locked email bounced"},
		"userunknown": []string{"sorry, no mailbox here by that name."},
	},
	"vmailmgr": map[string][]string{
		"mailboxfull": []string{"delivery failed due to system quota violation"},
		"userunknown": []string{
			"invalid or unknown base user or domain",
			"invalid or unknown virtual user",
			"user name does not refer to a virtual user",
		},
	},
}

// Find() detects the bounce reason from the error message generated by Local Delivery Agent
func Find(fo *sis.Fact) string {
	// @param    *sis.Fact fo    Struct to be detected the reason
	// @return   string          Bounce reason name or an empty string
	if fo.DiagnosticCode == ""                  { return "" }
	if fo.Command != "" && fo.Command != "DATA" { return "" }

	deliversby := "" // LDA; Local Delivery Agent name
	reasontext := "" // Bounce reason
	issuedcode := strings.ToLower(fo.DiagnosticCode)

	for e := range LocalAgent {
		// Find a local delivery agent name from the lower-cased error message
		if sisimoji.ContainsAny(issuedcode, LocalAgent[e]) == false { continue }
		deliversby = e; break
	}
	if deliversby == "" { return "" }

	for e := range MessagesOf[deliversby] {
		// The key nane is a Local Delivery Agent name
		if sisimoji.ContainsAny(issuedcode, MessagesOf[deliversby][e]) == false { continue }
		reasontext = e; break
	}
	return reasontext
}

