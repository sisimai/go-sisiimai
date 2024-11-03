// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis
//  ____       _ _                      __  __       _   _            
// |  _ \  ___| (_)_   _____ _ __ _   _|  \/  | __ _| |_| |_ ___ _ __ 
// | | | |/ _ \ | \ \ / / _ \ '__| | | | |\/| |/ _` | __| __/ _ \ '__|
// | |_| |  __/ | |\ V /  __/ |  | |_| | |  | | (_| | |_| ||  __/ |   
// |____/ \___|_|_| \_/ \___|_|   \__, |_|  |_|\__,_|\__|\__\___|_|   
//                                |___/                               
import "strings"
import "sisimai/address"
import "sisimai/rfc1123"
import "sisimai/rfc1894"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

var Fields1894 = rfc1894.FIELDTABLE()
type DeliveryMatter struct {
	Action       string     // The value of Action header
	Agent        string     // MTA name
	Alias        string     // The value of alias entry(RHS)
	Command      string     // SMTP command in the message body
	Date         string     // The value of Last-Attempt-Date header
	Diagnosis    string     // The value of Diagnostic-Code header
	FeedbackType string     // Feedback type
	Lhost        string     // The value of Received-From-MTA header
	Reason       string     // Temporary reason of bounce
	Recipient    string     // The value of Final-Recipient header
	ReplyCode    string     // SMTP Reply Code
	Rhost        string     // The value of Remote-MTA header
	Spec         string     // Protocl specification
	Status       string     // The value of Status header
}

// Select() returns the current value of the sis.DeliveryMatter{}
func(this *DeliveryMatter) Select(argv0 string) string {
	// @param    string argv0  A lower-cased member of sis.DeliveryMatter{}
	// @return   string        The value of the member specified at argv0
	switch argv0 {
		case "action":       return this.Action
		case "agent":        return this.Agent
		case "alias":        return this.Alias
		case "command":      return this.Command
		case "date":         return this.Date
		case "diagnosis":    return this.Diagnosis
		case "feedbacktype": return this.FeedbackType
		case "lhost":        return this.Lhost
		case "reason":       return this.Reason
		case "recipient":    return this.Recipient
		case "replycode":    return this.ReplyCode
		case "rhost":        return this.Rhost
		case "spec":         return this.Spec
		case "status":       return this.Status
		default:             return ""
	}
}

// Update() set the argument into the member of sis.DeliveryMatter
func(this *DeliveryMatter) Update(argv0 string, argv1 string) bool {
	// @param    string argv0  A lower-cased member name of sis.DeliveryMatter{}
	// @param    string argv1  The value to be updated
	// @return   bool          true if it has successfully updated
	if argv0 == "" { return false } // Member name is missing
	if argv1 == "" { return false } // Empty value is not allowd in this function

	actionlist := []string{"delayed", "delivered", "expanded", "failed", "relayed"}
	feedbacklo := []string{"abuse", "dkim", "fraud", "miscategorized", "not-spam", "opt-out", "virus", "other"}

	switch argv0 {
		default: return false
		case "action":
			// Only valid values are accepted
			if this.Action == argv1 || sisimoji.EqualsAny(argv1, actionlist) == false { return false }
			this.Action = argv1

		case "agent":
			// Any value is accepted
			if this.Agent == argv1 { return false }
			this.Agent = argv1

		case "alias":
			// Only valid email addresses are accepted
			if this.Alias == argv1 || sisiaddr.IsEmailAddress(argv1) == false { return false }
			this.Alias = argv1

		case "command":
			// Only valid values are accepted
			if this.Command == argv1 || command.Test(argv1) == false { return false }
			this.Command = argv1

		case "date":
			// Any value is accepted
			if this.Date == argv1 { return false }
			this.Date = argv1

		case "diagnosis":
			// Any value is accepted
			if this.Diagnosis == argv1 { return false }
			this.Diagnosis = argv1

		case "feedbacktype":
			// Only valid values are accepted
			if this.FeedbackType == argv1 || sisimoji.EqualsAny(argv1, feedbacklo) == false { return false }
			this.FeedbackType = argv1

		case "lhost":
			// Only valid hostnames are accepted
			if this.Lhost == argv1 || rfc1123.IsValidHostname(argv1) == false { return false }
			this.Lhost = strings.ToLower(argv1)

		case "reason":
			// Only valid reason names are accepted
			if this.Reason == argv1 { return false }
			this.Reason = strings.ToLower(argv1)

		case "recipient":
			// Only valid email addresses are accepted
			if this.Recipient == argv1 || sisiaddr.IsEmailAddress(argv1) == false { return false }
			this.Recipient = argv1

		case "replycode":
			// Only valid SMTP reply codes are accepted
			if this.ReplyCode == argv1 || reply.Test(argv1) == false { return false }
			this.ReplyCode = argv1

		case "rhost":
			// Only valid hostnames are accepted
			if this.Rhost == argv1 || rfc1123.IsValidHostname(argv1) == false { return false }
			this.Rhost = strings.ToLower(argv1)

		case "spec":
			// Any value is accepted
			if this.Spec == argv1 { return false }
			this.Spec = argv1

		case "status":
			// Only valid SMTP status codes are accepted
			if this.Status == argv1 || status.Test(argv1) == false { return false }
			this.Status = argv1
	}
	return true
}

// Set() substitutes the argv1 as a value into the member related to argv0
func(this *DeliveryMatter) Set(argv0, argv1 string) bool {
	// @param    string argv0  A key name related to the member of DeliveryMatter struct
	// @param    string argv1  The value to be substituted
	// @return   bool          Returns true if it succesufully substituted
	if len(argv0)             == 0 { return false }
	if len(argv1)             == 0 { return false }
	if len(Fields1894[argv0]) == 0 { return false }

	switch argv0 {
		// Available values are the followings:
		// - "action":             Action    (list)
		// - "arrival-date":       Date      (date)
		// - "diagnostic-code":    Diagnosis (code)
		// - "final-recipient":    Recipient (addr)
		// - "last-attempt-date":  Date      (date)
		// - "original-recipient": Alias     (addr)
		// - "received-from-mta":  Lhost     (host)
		// - "remote-mta":         Rhost     (host)
		// - "reporting-mta":      Rhost     (host)
		// - "status":             Status    (stat)
		// - "x-actual-recipient": Alias     (addr)
		default: return false
		case "action":
			// Action: failed
			this.Action = argv1

		case "arrival-date", "last-attempt-date":
			// Arrival-Date: Thu, 29 Apr 2019 23:34:45 +0900 (JST)
			// Last-Attempt-Date: Mon, 21 May 2018 16:10:00 +0900
			this.Date = argv1

		case "diagnostic-code":
			// Diagnostic-Code: smtp; 550 DMARC check failed.
			this.Diagnosis = argv1

		case "final-recipient":
			// Final-Recipient: RFC822; <kijitora@nyaan.jp>
			v := address.S3S4(argv1); if len(v) == 0 { return false }
			this.Recipient = v

		case "original-recipient", "x-actual-recipient":
			// X-Actual-Recipient: RFC822; kijitora@example.co.jp
			// X-Actual-Recipient: X-Unix; |/var/adm/sm.bin/neko
			v := address.S3S4(argv1); if len(v) == 0 { return false }
			this.Alias = v

		case "received-from-mta":
			// Received-From-MTA: DNS; p225-ix4.kyoto.example.ne.jp
			// Received-From-MTA: dns; sironeko@example.jp
			if strings.Contains(argv1, "@") { argv1 = argv1[strings.Index(argv1, "@") + 1:] }
			this.Lhost = argv1

		case "remote-mta", "reporting-mta":
			// Remote-MTA: dns; mx-aol.mail.gm0.yahoodns.net
			// Remote-MTA: dns; 192.0.2.222 (192.0.2.222, the server for the domain.)
			// Reporting-MTA: dsn; d217-29.smtp-out.amazonses.com
			if strings.Contains(argv1, " ") { argv1 = strings.Split(argv1, " ")[0] }
			this.Rhost = argv1

		case "status":
			// Status: 5.1.1
			// Status: 4.2.2 (Over Quota)
			if strings.Contains(argv1, " ") { argv1 = argv1[0:strings.Index(argv1, " ")] }
			if status.Test(argv1) == false  { return false }
			this.Status = argv1
	}
	return true
}

