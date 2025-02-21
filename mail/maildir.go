// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _    __               _ _     _ _      
//  _ __ ___   __ _(_) |  / / __ ___   __ _(_) | __| (_)_ __ 
// | '_ ` _ \ / _` | | | / / '_ ` _ \ / _` | | |/ _` | | '__|
// | | | | | | (_| | | |/ /| | | | | | (_| | | | (_| | | |   
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|\__,_|_|_|\__,_|_|_|   

package mail
import "io"
import "os"
import "fmt"
import "path/filepath"

// listMaildir() is a Maildir/ reader, works as a iterator.
func(this *EmailEntity) listMaildir() (int, error) {
	// @return   int      The number of email files in the Maildir/
	// @return   error    Errors while reading the Maildir/
	if this.handle == nil {
		// Open the Maildir/
		filehandle, nyaan := os.Open(this.Dir);  if nyaan != nil { return 0, nyaan }
		this.handle = filehandle // Successfully opened the Maildir/
	}
	direntries, nyaan := this.handle.Readdir(0); if nyaan != nil { return 0, nyaan }
	for _, e := range direntries {
		// Read each email file in the Maildir/
		if e.IsDir() == true || e.Size() == 0 { continue }
		this.payload = append(this.payload, e.Name())
	}
	this.handle.Close(); this.handle = nil
	return len(this.payload), nil
}

// readMaildir() reads each email file in the Maildir/, do not call this from inside of the goroutine.
func(this *EmailEntity) readMaildir() (*string, error) {
	// @return   *string  Contents of the each file in the Maildir/
	// @return   error    It has reached to the end of the Maildir/
	if this.Size == 0           { return nil, fmt.Errorf("there is no email file in %s", this.Dir) }
	if this.Size <= this.offset { return nil, io.EOF }

	for {
		// Try to read the email file
		this.File = this.payload[this.offset]; this.offset++
		this.Path = filepath.Clean(filepath.FromSlash(this.Dir + "/" + this.File))
		b, nyaan := os.ReadFile(this.Path); if nyaan != nil || len(b) == 0 {
			// Failed to read the email file or the email file is empty
			if this.offset >= this.Size { return nil, io.EOF }
			continue
		}
		cv := string(b); return &cv, nil
	}
}

