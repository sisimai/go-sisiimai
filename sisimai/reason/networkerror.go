// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _      _                      _    _____                     
// | \ | | ___| |___      _____  _ __| | _| ____|_ __ _ __ ___  _ __ 
// |  \| |/ _ \ __\ \ /\ / / _ \| '__| |/ /  _| | '__| '__/ _ \| '__|
// | |\  |  __/ |_ \ V  V / (_) | |  |   <| |___| |  | | | (_) | |   
// |_| \_|\___|\__| \_/\_/ \___/|_|  |_|\_\_____|_|  |_|  \___/|_|   
import "strings"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["NetworkError"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			"could not connect and send the mail to",
			"dns records for the destination computer could not be found",
			"hop count exceeded - possible mail loop",
			"host is unreachable",
			"host name lookup failure",
			"host not found, try again",
			"mail forwarding loop for ",
			"malformed name server reply",
			"malformed or unexpected name server reply",
			"maximum forwarding loop count exceeded",
			"message looping",
			"message probably in a routing loop",
			"no route to host",
			"too many hops",
			"unable to resolve route ",
			"unrouteable mail domain",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "networkerror" or not
	Truth["NetworkError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is networkerror, false: is not networkerror
		return false
	}
}

