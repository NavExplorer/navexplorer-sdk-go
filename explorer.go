package navexplorer

type ExplorerApi struct {
	client *ExplorerClient
}

func NewExplorerApi(host string, network Network) (*ExplorerApi, error) {
	explorerClient, err := NewClient(host, network)
	if err != nil {
		return nil, err
	}

	return &ExplorerApi{explorerClient}, nil
}
