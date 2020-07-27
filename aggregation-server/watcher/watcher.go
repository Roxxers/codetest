package watcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	endpoints "thirdlight.com/aggregation-server/lib"
	"thirdlight.com/watcher-node/lib"
)

// Watcher represents one node watcher for one directory
type Watcher struct {
	Instance string
	URL      url.URL
	Port     uint
	List     []lib.FileMetadata
	SeqNo    int
	mux      sync.RWMutex
}

func (w *Watcher) PatchList(patch lib.PatchOperation) error {
	// Code here
	return nil
}

// ReqFiles requests the current file list for the node, and updates the internal list of files.
func (w *Watcher) ReqFiles() error {
	// Get file list
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", w.URL.String(), endpoints.FilesEndpoint), nil)
	if err != nil {
		return fmt.Errorf("Error creating request for URL: %s\n%s", w.URL.String(), err)
	}
	client := &http.Client{Timeout: time.Second * 20}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to request files from node on initialsation: %s\n%s", w.URL.String(), err)
	}
	defer resp.Body.Close()

	var files lib.ListResponse
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return fmt.Errorf("Failed to parse file list of instance: %s @ %s\n%s", w.Instance, w.URL.String(), err)
	}

	defer w.mux.Unlock()
	w.mux.Lock()
	w.List = files.Files
	w.SeqNo = files.Sequence
	return nil
}

// Nodes is a helper class for dealing with the list of registered nodes
type Nodes struct {
	List []*Watcher
	mux  sync.RWMutex
}

func (n *Nodes) Find(instanceID string) (*Watcher, error) {
	defer n.mux.RUnlock()
	n.mux.RLock()
	for _, watcher := range n.List {
		if watcher.Instance == instanceID {
			return watcher, nil
		}
	}
	// No matches
	return nil, fmt.Errorf("no node with instance ID: %s", instanceID)
}

func (n *Nodes) New(instanceID string, address string, port uint) (*Watcher, error) {
	// Formatted like this because the url lib does not like normal ip addresses, but is just fine with domains
	// https://github.com/golang/go/issues/19297
	// A solution that works with both and https would be in prod but this works for now on a local machine
	url, err := url.Parse(fmt.Sprintf("http://%s:%d", address, port))
	if err != nil {
		return nil, err
	}

	w := &Watcher{Instance: instanceID, URL: *url, Port: port}
	w.ReqFiles()
	n.mux.Lock()
	n.List = append(n.List, w)
	n.mux.Unlock()
	return w, nil
}

func (n *Nodes) Remove(instanceID string) error {
	defer n.mux.Unlock()
	n.mux.Lock()
	for x, watcher := range n.List {
		if watcher.Instance == instanceID {
			lenList := len(n.List)
			// Remove instance by replacing it with the last in slice, then removing last element
			lastElm := lenList - 1
			n.List[x] = n.List[lastElm]
			n.List = n.List[:lastElm]
			return nil
		}
	}
	return fmt.Errorf("no node with instance ID: %s", instanceID)
}

func (n *Nodes) FetchAllFiles() map[string][]lib.FileMetadata {
	defer n.mux.RUnlock()
	n.mux.RLock()
	// make is used here due to needing the base map to != nil to reference the files key
	files := make(map[string][]lib.FileMetadata)
	files["files"] = make([]lib.FileMetadata, 0)

	for _, watcher := range n.List {
		files["files"] = append(files["files"], watcher.List...)
	}
	return files
}

// CreateNodesList creates the a new wrapper class for the list of nodes available to this server
func CreateNodesList() *Nodes {
	return &Nodes{List: []*Watcher{}}
}
