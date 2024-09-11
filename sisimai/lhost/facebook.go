// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

import "sisimai/sis"

func init() {
	// Decode bounce messages from
	InquireFor["Facebook"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Body)            == 0 { return sis.RisingUnderway{} }

        return sis.RisingUnderway{}
    }
}

//       _               _      _______              _                 _    
//  _ __| |__   ___  ___| |_   / /  ___|_ _  ___ ___| |__   ___   ___ | | __
// | '__| '_ \ / _ \/ __| __| / /| |_ / _` |/ __/ _ \ '_ \ / _ \ / _ \| |/ /
// | |  | | | | (_) \__ \ |_ / / |  _| (_| | (_|  __/ |_) | (_) | (_) |   < 
// |_|  |_| |_|\___/|___/\__/_/  |_|  \__,_|\___\___|_.__/ \___/ \___/|_|\_\
//                                                                          
