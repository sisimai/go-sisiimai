// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ____     ______                     _                 _ _ 
// | | |__   ___  ___| |_   / /\ \   / / ___| ___  ___ _ __   __| |_ __ ___   __ _(_) |
// | | '_ \ / _ \/ __| __| / /  \ \ / /|___ \/ __|/ _ \ '_ \ / _` | '_ ` _ \ / _` | | |
// | | | | | (_) \__ \ |_ / /    \ V /  ___) \__ \  __/ | | | (_| | | | | | | (_| | | |
// |_|_| |_|\___/|___/\__/_/      \_/  |____/|___/\___|_| |_|\__,_|_| |_| |_|\__,_|_|_|
import "strings"
import "sisimai/sis"
import "sisimai/rfc1123"
import "sisimai/rfc5322"
import "sisimai/smtp/reply"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Sendmail version 5
	InquireFor["V5sendmail"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }
		if strings.HasPrefix(bf.Head["subject"][0], "Returned mail: ") == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"   ----- Unsent message follows -----", "  ----- No message was collected -----"}
		startingof := map[string][]string{
			// Error text regular expressions which defined in src/savemail.c
			//   savemail.c:485| (void) fflush(stdout);
			//   savemail.c:486| p = queuename(e->e_parent, 'x');
			//   savemail.c:487| if ((xfile = fopen(p, "r")) == NULL)
			//   savemail.c:488| {
			//   savemail.c:489|   syserr("Cannot open %s", p);
			//   savemail.c:490|   fprintf(fp, "  ----- Transcript of session is unavailable -----\n");
			//   savemail.c:491| }
			//   savemail.c:492| else
			//   savemail.c:493| {
			//   savemail.c:494|   fprintf(fp, "   ----- Transcript of session follows -----\n");
			//   savemail.c:495|   if (e->e_xfp != NULL)
			//   savemail.c:496|       (void) fflush(e->e_xfp);
			//   savemail.c:497|   while (fgets(buf, sizeof buf, xfile) != NULL)
			//   savemail.c:498|       putline(buf, fp, m);
			//   savemail.c:499|   (void) fclose(xfile);
			"error":   []string{"While talking to "},
			"message": []string{"----- Transcript of session follows -----"},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false); if emailparts[1] == "" { return sis.RisingUnderway{} }
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		remotehost := ""                  // The last remote hostname
		curcommand := ""                  // THe last SMTP command
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.Contains(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			//    ----- Transcript of session follows -----
			// While talking to smtp.example.com:
			// >>> RCPT To:<kijitora@example.org>
			// <<< 550 <kijitora@example.org>, User Unknown
			// 550 <kijitora@example.org>... User unknown
			// 421 example.org (smtp)... Deferred: Connection timed out during user open with example.org
			if strings.HasPrefix(e, ">>> ") { curcommand = command.Find(e[4:]) }
			if sisimoji.Aligned(e, []string{" <", "@", ">..."}) || strings.Contains(strings.ToUpper(e), ">>> RCPT TO:") {
				// 550 <kijitora@example.org>... User unknown
				// >>> RCPT To:<kijitora@example.org>
				p0 := strings.Index(e, " ")
				p1 := sisimoji.IndexOnTheWay(e, "<", p0)
				p2 := sisimoji.IndexOnTheWay(e, ">", p1)
				cv := sisiaddr.S3S4(e[p1:p2 + 1])

				if cv == v.Recipient || (curcommand == "MAIL" && strings.HasPrefix(e, "<<< ")) {
					// The recipient address is the same address with the last appeared address
					// like "550 <mikeneko@example.co.jp>... User unknown"
					// Append this line to the string which is keeping error messages
					v.Diagnosis += " " + e
					curcommand   = ""

				} else {
					// The recipient address in this line differs from the last appeared address
					// or is the first recipient address in this bounce message
					if len(v.Recipient) > 0 {
						// There are multiple recipient addresses in the message body.
						dscontents = append(dscontents, sis.DeliveryMatter{})
						v = &(dscontents[len(dscontents) - 1])
					}
					recipients++
					v.Recipient = cv
					v.Rhost     = remotehost
					v.Diagnosis += " " + e
				}
			} else {
				// This line does not include a recipient address
				if strings.Contains(e, startingof["error"][0]) {
					// While talking to mail.example.co.jp:
					cv := rfc1123.Find(e); if rfc1123.IsInternetHost(cv) { remotehost = cv }

				} else {
					// Append this line into the error message string
					v.Diagnosis += " " + e
				}
			}
		}

		if recipients == 0 {
			// There is no recipient address in the error message
			p1 := strings.Index(emailparts[1], "\nTo: ")
			p2 := sisimoji.IndexOnTheWay(emailparts[1], "\n", p1 + 6)

			// Get the recipient address from "To:" header at the original message
			if p1 > 0 { dscontents[0].Recipient = sisiaddr.S3S4(emailparts[1][p1 + 5:p2]) }
			if sisiaddr.IsEmailAddress(dscontents[0].Recipient) == false { return sis.RisingUnderway{} }
			recipients++
		}

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis
			e := &(dscontents[j])

			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			e.ReplyCode = reply.Find(e.Diagnosis, "")
			e.Command   = command.Find(e.Diagnosis)

			// There is no local part in the recipient email address like "@example.jp"
			// Get an email address from the value of Diagnostic-Code: field
			if sisiaddr.IsEmailAddress(e.Recipient) == true  { continue }
			p1 := strings.Index(e.Diagnosis, "<"); if p1 < 0 { continue }
			p2 := strings.Index(e.Diagnosis, ">"); if p2 < 0 { continue }
			e.Recipient = sisiaddr.S3S4(e.Diagnosis[p1:p2 + 1])
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

