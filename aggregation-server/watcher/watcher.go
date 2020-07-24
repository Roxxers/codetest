package watcher

import (
	"fmt"
	"net/url"
	"sync"

	"thirdlight.com/aggregation-server/types"
	"thirdlight.com/watcher-node/lib"
)

// Watcher represents one node watcher for one directory
type Watcher struct {
	Instance string
	URL      url.URL
	Port     uint
	List     *types.FileList
}

// Nodes is a helper class for dealing with the list of registered nodes
type Nodes struct {
	List []*Watcher
	mux  sync.Mutex
}

// CreateNodesList creates the a new wrapper class for the list of nodes available to this server
func CreateNodesList() *Nodes {
	return &Nodes{List: []*Watcher{}}
}

func (n *Nodes) Find(instanceID string) (*Watcher, error) {
	for _, watcher := range n.List {
		if watcher.Instance == instanceID {
			return watcher, nil
		}
	}
	// No matches
	return nil, fmt.Errorf("no node with instance ID: %s", instanceID)
}

func (n *Nodes) New(instanceID string, address string, port uint) (*Watcher, error) {
	url, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	w := &Watcher{instanceID, *url, port, &types.FileList{}}
	n.mux.Lock()
	n.List = append(n.List, w)
	n.mux.Unlock()
	return w, nil
}

func (w *Watcher) PatchList(patch lib.PatchOperation) error {
	// Code here
	return nil
}
