package types

import "thirdlight.com/watcher-node/lib"

// FileList represents JSON object response for the /files endpoint
// It is also used to represent parts of this response in worker classes
type FileList struct {
	Files []lib.FileMetadata `json:"files"`
}
