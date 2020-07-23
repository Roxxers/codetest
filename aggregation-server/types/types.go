package types

import "thirdlight.com/watcher-node/lib"

// FilesResponse represents JSON object response for the /files endpoint
type FilesResponse struct {
	Files []lib.FileMetadata `json:"files"`
}
