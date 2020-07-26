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

func (w *Watcher) PatchList(patch lib.PatchOperation) error {
	// Code here
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
	url, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	// TODO: Add actually getting the full file list here
	w := &Watcher{instanceID, *url, port, &types.FileList{}}
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

// CreateNodesList creates the a new wrapper class for the list of nodes available to this server
func CreateNodesList() *Nodes {
	return &Nodes{List: []*Watcher{}}
}
