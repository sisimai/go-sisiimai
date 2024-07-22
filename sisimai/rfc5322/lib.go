// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

func FIELDINDEX() []string {
	// The following fields are not referred in Sisimai
	// Resent-From Resent-Sender Resent-Cc Cc Bcc Resent-Bcc In-Reply-To References Comments Keywords
	return []string {
		"Resent-Date", "From", "Sender", "Reply-To", "To", "Message-ID", "Subject", "Return-Path",
		"Received", "Date", "X-Mailer", "Content-Type", "Content-Transfer-Encoding", "Content-Description",
		"Content-Disposition",
	}
}

func HEADERTABLE() map[string][]string {
	return map[string][]string {
		"messageid": []string { "message-id" },
		"subject":   []string { "subject" },
		"listid":    []string { "list-id" },
		"date":      []string { "date", "osted-date", "posted", "resent-date" },
		"addresser": []string {
			"from", "return-path", "reply-to", "errors-to", "reverse-path", "x-postfix-sender",
			"envelope-from", "x-envelope-from",
		},
		"recipient": []string {
			"to", "delivered-to", "forward-path", "envelope-to", "x-envelope-to", "resent-to",
			"apparently-to",
		},
	}
}

