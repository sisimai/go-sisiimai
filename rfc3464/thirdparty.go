// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _____ _  _    __   _  _     _______         _ ____            _         
// |  _ \|  ___/ ___|___ /| || |  / /_ | || |   / /___ / _ __ __| |  _ \ __ _ _ __| |_ _   _ 
// | |_) | |_ | |     |_ \| || |_| '_ \| || |_ / /  |_ \| '__/ _` | |_) / _` | '__| __| | | |
// |  _ <|  _|| |___ ___) |__   _| (_) |__   _/ /  ___) | | | (_| |  __/ (_| | |  | |_| |_| |
// |_| \_\_|   \____|____/   |_|  \___/   |_|/_/  |____/|_|  \__,_|_|   \__,_|_|   \__|\__, |
//                                                                                     |___/ 

package rfc3464
import "strings"

var ReturnedBy = map[string]func(string) []string {}
var ThirdParty = map[string][]string{
//	"Aol":      []string{"X-Outbound-Mail-Relay-"}, // X-Outbound-Mail-Relay-(Queue-ID|Sender)
	"PowerMTA": []string{"X-PowerMTA-"},            // X-PowerMTA-(VirtualMTA|BounceCategory)
//	"Yandex":   []string{"X-Yandex-"},              // X-Yandex-(Queue-ID|Sender)
}

// is3rdparty() returns true if the argument is a line generated by a MTA that has fields defined
// in RFC3464 inside of a bounce mail the MTA returns
func is3rdparty(argv1 string) bool {
	// @param    string argv1   A line of a bounce mail
	// @return   bool           The line indicates that a bounce mail generated by the 3rd party MTA
	if argv1 == "" { return false }
	if party := returnedby(argv1); party == "" { return false }
	return true
}

// returnedby() returns an MTA name of the 3rd party
func returnedby(argv1 string) string {
	// @param    string argv1   A line of a bounce mail
	// @return   string         An MTA name of the 3rd party
	if argv1 == "" || strings.HasPrefix(argv1, "X-") == false { return "" }
	for e := range ThirdParty {
		// Does the argument include the 3rd party specific field?
		if strings.HasPrefix(argv1, ThirdParty[e][0]) { return e }
	}
	return ""
}

// xfield() returns rfc1894.Field() compatible slice for the specific field of the 3rd party MTA
func xfield(argv1 string) []string {
	// @param    string argv1  A line of the error message
	// @return   []string      rfc1894.Field() compatible slice
	// @see      sisimai/rfc1894/lib.go
	if argv1 == "" { return []string{} }
	if party := returnedby(argv1); party != "" { return ReturnedBy[party](argv1) }
	return []string{}
}

