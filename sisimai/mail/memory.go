// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package mail

//                  _ _    __                                         
//  _ __ ___   __ _(_) |  / / __ ___   ___ _ __ ___   ___  _ __ _   _ 
// | '_ ` _ \ / _` | | | / / '_ ` _ \ / _ \ '_ ` _ \ / _ \| '__| | | |
// | | | | | | (_| | | |/ /| | | | | |  __/ | | | | | (_) | |  | |_| |
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|\___|_| |_| |_|\___/|_|   \__, |
//                                                              |___/ 
import "io"

func (this *EmailEntity) readMemory() (*string, error) {
	// @return   *string  Contents of the each email in the this.payload[]
	// @return   error    It has reached to the end of the email
	if this.Size == 0 || this.offset == int64(len(this.payload)) { return nil, io.EOF }

	emailblock := this.payload[this.offset]
	this.offset++
	return &emailblock, nil
}

