/*
   Copyright (c) 2012 Kyle Isom <kyle@tyrfingr.is>

   Permission to use, copy, modify, and distribute this software for any
   purpose with or without fee is hereby granted, provided that the 
   above copyright notice and this permission notice appear in all 
   copies.

   THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL 
   WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED 
   WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE 
   AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL
   DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA
   OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER 
   TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR 
   PERFORMANCE OF THIS SOFTWARE.
*/

/*
fswatch provides simple UNIX file system watching in Go. It is based around
the Watcher struct, which should be initialised with either NewWatcher or
NewAutoWatcher. Both functions accept a variable number of string arguments
specfying the paths to be loaded, which may be globbed, and return a pointer
to a Watcher. This value can be started and stopped with the Start() and
Stop() methods. The Watcher will automatically stop if all the files it is
watching have been deleted.

The Start() method returns a read-only channel that receives Notification
values. The Stop() method closes the channel, and no files will be watched
from that point.

*/
package fswatch
