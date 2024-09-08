// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____                                 _ 
// / ___| _   _ ___ _ __   ___ _ __   __| |
// \___ \| | | / __| '_ \ / _ \ '_ \ / _` |
//  ___) | |_| \__ \ |_) |  __/ | | | (_| |
// |____/ \__,_|___/ .__/ \___|_| |_|\__,_|
//                 |_|                     
import "strings"
import "sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["Suspend"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		index := []string{
			" is currently suspended",
			" temporary locked",
			"archived recipient",
			"boite du destinataire archivee",
			"email account that you tried to reach is disabled",
			"has been suspended",
			"inactive account",
			"invalid/inactive user",
			"is a deactivated mailbox", // http://service.mail.qq.com/cgi-bin/help?subtype=1&&id=20022&&no=1000742
			"is unavailable: user is terminated",
			"mailbox currently suspended",
			"mailbox disabled",
			"mailbox is frozen",
			"mailbox unavailable or access denied",
			"recipient rejected: temporarily inactive",
			"recipient suspend the service",
			"this account has been disabled or discontinued",
			"this account has been temporarily suspended",
			"this address no longer accepts mail",
			"this mailbox is disabled",
			"user or domain is disabled",
			"user suspended", // http://mail.163.com/help/help_spam_16.htm
			"vdelivermail: account is locked email bounced",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "suspend" or not
	Truth["Suspend"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is suspend, false: is not suspend
		if fo.Reason == "suspend" { return true }
		if fo.ReplyCode == "525"  { return true }
		return IncludedIn["Suspend"](strings.ToLower(fo.DiagnosticCode))
	}
}

