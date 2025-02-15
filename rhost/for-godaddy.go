// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ______       ____            _     _       
//  _ __| |__   ___  ___| |_   / / ___| ___ |  _ \  __ _  __| | __| |_   _ 
// | '__| '_ \ / _ \/ __| __| / / |  _ / _ \| | | |/ _` |/ _` |/ _` | | | |
// | |  | | | | (_) \__ \ |_ / /| |_| | (_) | |_| | (_| | (_| | (_| | |_| |
// |_|  |_| |_|\___/|___/\__/_/  \____|\___/|____/ \__,_|\__,_|\__,_|\__, |
//                                                                   |___/ 

package rhost
import "strings"
import "libsisimai.org/sisimai/sis"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["GoDaddy"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      https://ca.godaddy.com/help/fix-rejected-email-with-a-bounce-error-40685
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		errorcodes := map[string]string{
			// Sender bounces
			// -----------------------------------------------------------------------------------------
			// - 535 Authentication not allowed on IBSMTP Servers. IB401
			// - Authentication is not allowed on inbound mail. This happens when you have incorrect
			//   outgoing SMTP server settings set up in your email client, like Outlook or Gmail.
			// - Set up your email client using the SMTP server setting smtpout.secureserver.net.
			"IB401": "securityerror",

			// - 550 jane@coolexample.com Blank From: addresses are not allowed. Provide a valid From.
			//   IB501
			// - The email message "from" field is blank.
			// - Add a valid "from" address and try to send the email again.
			"IB501": "notcompliantrfc",

			// - 550 jane@coolexample.com IP addresses are not allowed as a From: Address. Provide a
			//   valid From. IB502
			// - Email messages can't be sent "from" an IP address.
			// - Add a valid "from" address and try to send the email again.
			"IB502": "notcompliantrfc",

			// - 550 coolexample.com From: Domain is invalid. Provide a valid From: IB506
			// - The domain doesn't have valid MX Records, an IP address, or there were issues during
			//   the DNS lookup when the email was sent.
			// - Verify that you're sending from a valid domain. Then verify that the domain has valid
			//   DNS records by checking your zone file. If the DNS isn't valid, it must be fixed before
			//   you resend the email.
			"IB506": "rejected",

			// - 550 jane@coolexample.com Invalid SPF record. Inspect your SPF settings, and try again.
			//   IB508
			// - The sending email address's domain has an SPF record that does not authorize the sending
			//   email server to send email from the domain.
			// - Modify the SPF record to include the server you're trying to send from or remove the SPF
			//   record from the domain.
			"IB508": "authfailure",

			// - 421 Temporarily rejected. Reverse DNS for this IP failed. IB108
			// - The IP address attempting to send mail does not have reverse DNS setup, or the DNS
			//   lookup failed.
			// - Verify the sending IP address has reverse DNS set up before resending the email.
			//   GoDaddy manages reverse DNS when using our email services. We do not support custom
			//   reverse DNS.
			"IB108": "requireptr",

			// Content bounces
			// -----------------------------------------------------------------------------------------
			// - 552 This message has been rejected due to content judged to be spam by the internet
			//   community. IB212
			// - The email message contains a link, attachment or pattern caught by our filters as spam.	
			"IB212": "spamdetected",

			// - 552 Virus infected message rejected. IB705
			// - The email message containing a link, attachment or pattern caught by our filters as a
			//   possible virus.	
			"IB705": "virusdetected",

			// Rate limiting bounces
			// -----------------------------------------------------------------------------------------
			// - 452 This message contains too many recipients. Reduce the number of recipients and
			//   retry. IB605
			// - The message has attempted to mail too many recipients.
			// - Reduce the number of recipients and try again.
			"IB605": "toomanyconn",

			// - 421 Connection refused, too many sessions from This IP. Lower the number of concurrent
			//   sessions. IB007
			// - This IP address currently has too many sessions open.
			// - Check with your email provider to reduce the number of open sessions on your IP address
			//   and then try again.
			"IB007": "toomanyconn",

			// - 421 Server temporarily unavailable. Try again later. IB101
			// - The email queue is experiencing higher than normal email volume.
			// - Try again later.
			"IB101": "speeding",

			// - 554 This IP has been temporarily blocked for attempting to send too many messages
			//   containing content judged to be spam by the Internet community. IB110
			// - This IP address has attempted to send too many messages containing content judged to be
			//   spam and has been blocked for an hour.
			// - If you're not sending spam, you'll need to contact your Internet Service Provider (ISP)
			//   to see why your IP address is sending so many emails. Something in your system is
			//   causing the issue, and you'll need to troubleshoot.
			"IB110": "blocked",

			// - 554 This IP has been blocked for the day, for attempting to send too many messages
			//   containing content judged to be spam by the Internet community. IB111
			// - This IP address has attempted to send too many messages containing content judged to be
			//   spam and has been blocked for the remainder of the day.
			"IB111": "blocked",

			// - 554 This IP has been temporarily blocked for attempting to mail too many invalid
			//   recipients. IB112
			// - This IP address has attempted to email too many invalid recipients and has been blocked
			//   for an hour.
			"IB112": "blocked",

			// - 554 This IP has been blocked for the day, for attempting to mail too many invalid
			//   recipients. IB113
			// - This IP address has attempted to email too many invalid recipients and has been blocked
			//   for the remainder of the day.
			"IB113": "blocked",

			// - 550 This IP has sent too many messages this hour. IB504
			// - This IP address has reached the maximum allowed messages for that hour.
			"IB504": "speeding",

			// - 550 This message has exceeded the max number of messages per session. Open a new session
			//   and try again. IB510
			// - This IP address has reached the maximum allowed messages for that session.
			"IB510": "speeding",

			// - 550 This IP has sent too many to too many recipients this hour. IB607
			// - This IP address has reached the maximum allowed recipients for that hour.
			"IB607": "speeding",

			// Remote block list (RBL) bounces
			// -----------------------------------------------------------------------------------------
			// - 554 Connection refused. This IP has a poor reputation on Cloudmark Sender Intelligence
			//   (CSI). IB103
			// - This IP address has a poor reputation on Cloudmark Sender Intelligence (CSI).
			"IB103": "badreputation",

			// - 554 Connection refused. This IP address is listed on the Spamhaus Block List (SBL). IB104
			// - This IP address is listed on the Spamhaus Block List.
			"IB104": "blocked",

			// - 554 Connection refused. This IP address is listed on the Exploits Block List (XBL). IB105
			// - This IP address is listed on the Spamhaus Exploits Block List.
			"IB105": "blocked",

			// - 554 Connection refused. This IP address is listed on the Policy Block List (PBL). IB106
			// - This IP address is listed on the Spamhaus Policy Block List.
			"IB106": "blocked",
		}
		messagesof := map[string][]string{
			"blocked": []string{
				// - 554 RBL Reject.
				// - This IP address was blocked from our internal RBL.
				// - Use the link provided in the bounceback to submit a request to remove this IP address.
				"rbl reject",
				"www.spamhaus.org/query/bl?ip=",
			},
			"expired": []string{
				// - 451 Sorry, I wasn't able to establish an SMTP connection. I'm not going to try again;
				//   this message has been in the queue too long.
				// - The recipient's email address has been misspelled or the recipient's email provider
				//   has blocked sending from your address.
				"has been in the queue too long",

				// - Delivery timeout
				// - This could happen for several reasons.
				"delivery timeout",
			},
			"mailboxfull": []string{
				// - Account storage limit
				// - The recipient's account has reached its storage limit and can't receive email right
				//   now.
				// - The recipient needs to delete messages from their inbox or the receiving folder to
				//   make space for more email.
				"account storage limit",
			},
			"spamdetected": []string{
				// - 552 Message rejected for spam or virus content
				// - The email message contains a link, attachment, or pattern caught by our filters as spam.
				"message rejected for spam or virus content",
			},
			"suspend": []string{
				// - Account disabled
				// - The recipient account exists but its ability to receive mail was disabled.
				// - Typically, these accounts remain permanently disabled, but you can try sending the
				//   email again later.
				"account disabled",
			},
			"systemerror": []string{
				// - This message is looping: it already has my Delivered-To line. (//5.4.6)
				// - The recipient account is forwarding the message in a loop.
				// - This is oftentimes because the receiver has two addresses that forward to each
				//   other. They need to correct their forwarding settings.
				"message is looping",
			},
			"userunknown": []string{
				// - 550 Recipient not found
				// - The recipient is not a valid email address.
				// - Remove the invalid recipient from your email.
				"recipient not found",

				// - Account does not exist
				// - The email address that you sent to does not exist.
				// - Verify that the recipient email address was entered correctly.
				"account does not exist",
			},
		}

		issuedcode := fo.DiagnosticCode
		errorlabel := "IB" + sisimoji.Select(issuedcode, " IB", " ", 0)
		if errorcodes[errorlabel] != "" {
			// 192.0.2.22 has sent to too many recipients this hour. IB607 ...
			return errorcodes[errorlabel]

		} else {
			// There is no " IB***" error code in the error message
			issuedcode = strings.ToLower(issuedcode)
			for e := range messagesof {
				// The key is a bounce reason name
				if sisimoji.ContainsAny(issuedcode, messagesof[e]) { return e }
			}
		}
		return ""
	}
}

