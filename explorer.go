package navexplorer

import (
	"errors"
	"log"
)

type ExplorerApi struct {
	client *ExplorerClient
}

var (
	ErrorExplorerConnectionError = errors.New("Could not connect to the NavExplorer API")
)

func NewExplorerApi(host string, network string) (*ExplorerApi, error) {
	explorerClient, err := NewClient(host, network)
	if err != nil {
		log.Print(err)
		return nil, ErrorExplorerConnectionError
	}

	return &ExplorerApi{explorerClient}, nil
}
