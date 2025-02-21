// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      _______          _                            ____   ___   ___ _____ 
// | | |__   ___  ___| |_   / / ____|_  _____| |__   __ _ _ __   __ _  ___|___ \ / _ \ / _ \___  |
// | | '_ \ / _ \/ __| __| / /|  _| \ \/ / __| '_ \ / _` | '_ \ / _` |/ _ \ __) | | | | | | | / / 
// | | | | | (_) \__ \ |_ / / | |___ >  < (__| | | | (_| | | | | (_| |  __// __/| |_| | |_| |/ /  
// |_|_| |_|\___/|___/\__/_/  |_____/_/\_\___|_| |_|\__,_|_| |_|\__, |\___|_____|\___/ \___//_/   
//                                                              |___/                             

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc1123"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/smtp/reply"
import "libsisimai.org/sisimai/smtp/status"
import sisiaddr "libsisimai.org/sisimai/address"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from Microsoft Exchange Server 2007: https://www.microsoft.com/microsoft-365/exchange/email
	InquireFor["Exchange2007"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		proceedsto := uint8(0)
		mailsender := []string{"postmaster@outlook.com", ".onmicrosoft.com"}
		emailtitle := []string{
			// Subject:            Content-Language:
			"Undeliverable",    // en-US
			"Non_remis_",       // fr-FR
			"Non remis ",       // fr-FR
			"Non recapitabile", // it-CH
			"Olevererbart",     // sv-SE
		}
		boundaries := []string{
			"Original Message Headers",
			"Original message headers:",             // en-US
			"tes de message d'origine :",            // fr-FR/En-têtes de message d'origine
			"Intestazioni originali del messaggio:", // it-CH
			"Ursprungshuvuden:",                     // sv-SE
		}
		startingof := map[string][]string{
			"error":   []string{" RESOLVER.", " QUEUE."},
			"message": []string{
				"Error Details",
				"Diagnostic information for administrators:",           // en-US
				"Informations de diagnostic pour les administrateurs",  // fr-FR
				"Informazioni di diagnostica per gli amministratori",   // it-CH
				"Diagnostisk information f",                            // sv-SE
			},
			"rhost":   []string{
				"DSN generated by:",
				"Generating server",        // en-US
				"Serveur de g",             // fr-FR/Serveur de génération
				"Server di generazione",    // it-CH
				"Genererande server",       // sv-SE
			},
		}
		ndrsubject := map[string]string{
			"SMTPSEND.DNS.NonExistentDomain": "hostunknown",   // 554 5.4.4 SMTPSEND.DNS.NonExistentDomain
			"SMTPSEND.DNS.MxLoopback":        "networkerror",  // 554 5.4.4 SMTPSEND.DNS.MxLoopback
			"RESOLVER.ADR.BadPrimary":        "systemerror",   // 550 5.2.0 RESOLVER.ADR.BadPrimary
			"RESOLVER.ADR.RecipNotFound":     "userunknown",   // 550 5.1.1 RESOLVER.ADR.RecipNotFound
			"RESOLVER.ADR.RecipientNotFound": "userunknown",   // 550 5.1.1 RESOLVER.ADR.RecipientNotFound
			"RESOLVER.ADR.ExRecipNotFound":   "userunknown",   // 550 5.1.1 RESOLVER.ADR.ExRecipNotFound
			"RESOLVER.ADR.RecipLimit":        "toomanyconn",   // 550 5.5.3 RESOLVER.ADR.RecipLimit
			"RESOLVER.ADR.InvalidInSmtp":     "systemerror",   // 550 5.1.0 RESOLVER.ADR.InvalidInSmtp
			"RESOLVER.ADR.Ambiguous":         "systemerror",   // 550 5.1.4 RESOLVER.ADR.Ambiguous, 420 4.2.0 RESOLVER.ADR.Ambiguous
			"RESOLVER.RST.AuthRequired":      "securityerror", // 550 5.7.1 RESOLVER.RST.AuthRequired
			"RESOLVER.RST.NotAuthorized":     "rejected",      // 550 5.7.1 RESOLVER.RST.NotAuthorized
			"RESOLVER.RST.RecipSizeLimit":    "exceedlimit",   // 550 5.2.3 RESOLVER.RST.RecipSizeLimit
			"QUEUE.Expired":                  "expired",       // 550 4.4.7 QUEUE.Expired
		}
		if sisimoji.ContainsAny(bf.Headers["subject"][0], emailtitle) { proceedsto++ }
		if sisimoji.ContainsAny(bf.Headers["from"][0], mailsender)    { proceedsto++ }
		if sisimoji.ContainsAny(bf.Payload, startingof["error"])      { proceedsto++ }
		if sisimoji.ContainsAny(bf.Payload, startingof["message"])    { proceedsto++ }
		if len(bf.Headers["content-language"]) > 0                    { proceedsto++ }
		if proceedsto < 2 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)              // Points the current cursor position
		recipients := uint8(0)              // The number of 'Final-Recipient' header
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if sisimoji.HasPrefixAny(e, startingof["message"]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 || e == "" { continue }

			// Diagnostic information for administrators:
			//
			// Generating server: mta2.neko.example.jp
			//
			// kijitora@example.jp
			// //550 5.1.1 RESOLVER.ADR.RecipNotFound; not found ////
			//
			// Original message headers:
			if strings.IndexByte(e, ' ') < 0 && strings.IndexByte(e, '@') > 1 {
				// This line includes an email address only
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e)
				recipients += 1

			} else {
				// Try to pick the remote hostname and status code, reply code from the error message
				if sisimoji.HasPrefixAny(e, startingof["rhost"]) {
					// Generating server: SG2APC01HT234.mail.protection.outlook.com
					// DSN generated by:       NEKONYAAN0022.apcprd01.prod.exchangelabs.com
					if cv := rfc1123.Find(e); rfc1123.IsInternetHost(cv) { v.Rhost = cv }

				} else {
					// #550 5.1.1 RESOLVER.ADR.RecipNotFound; not found ##
					// #550 5.2.3 RESOLVER.RST.RecipSizeLimit; message too large for this recipient ##
					cr := reply.Find(e, "")
					cs := status.Find(e, "")
					if cr != "" || cs != "" || strings.Contains(e, "Remote Server ") {
						// Remote Server returned '550 5.1.1 RESOLVER.ADR.RecipNotFound; not found'
						// 3/09/2016 8:05:56 PM - Remote Server at mydomain.com (10.1.1.3) returned '550 4.4.7 QUEUE.Expired; message expired'
						v.ReplyCode  = cr
						v.Status     = cs
						v.Diagnosis += e + " "
					}
				}
			}
		}

		for recipients == 0 {
			// Try to pick the recipient address from the following formatted bounce message:
			//   Original Message Details
			//   Created Date:   4/29/2017 11:23:34 PM
			//   Sender Address: neko@example.com
			//   Recipient Address:      kijitora@neko.kyoto.example.jp
			//   Subject:        Nyaan?
			p1 := strings.Index(emailparts[0], "Original Message Details"); if p1 < 0 { break }
			p2 := strings.Index(emailparts[0], "\nRecipient Address: ");    if p2 < 0 { break }
			p3 := sisimoji.IndexOnTheWay(emailparts[0], "\n", p2 + 20);     if p3 < 0 { break }
			cv := sisiaddr.S3S4(emailparts[0][p2 + 20:p3])
			if rfc5322.IsEmailAddress(cv) { dscontents[0].Recipient = cv; recipients++ }
		}
		if recipients == 0  { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			p0 := -1
			p1 := strings.IndexByte(e.Diagnosis, ';')
			for _, r := range startingof["error"] {
				// Find an error message and an error code 
				p0 = strings.Index(e.Diagnosis, r); if p0 > -1 { break }
			}
			if p0 < 0 || p1 < 0 { continue }

			// #550 5.1.1 RESOLVER.ADR.RecipNotFound; not found ##
			if cv := e.Diagnosis[p0 + 1:p1]; len(ndrsubject[cv]) > 0 { e.Reason = ndrsubject[cv] }
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

