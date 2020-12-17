package typesense

import (
	"github.com/GianOrtiz/typesense-go"
	"github.com/harmony-development/legato/server/config"
	"github.com/ztrue/tracerr"
)

type Dependencies struct {
	Config *config.Config
}

type TypesenseBackend struct {
	Dependencies
	*typesense.Client
}

func New(deps Dependencies) (*TypesenseBackend, error) {
	client := typesense.NewClient(&typesense.Node{
		Host:     deps.Config.Search.Typesense.Host,
		Port:     deps.Config.Search.Typesense.Port,
		Protocol: deps.Config.Search.Typesense.Protocol,
		APIKey:   deps.Config.Search.Typesense.APIKey,
	}, 2)
	if err := client.Ping(); err != nil {
		err = tracerr.Wrap(err)
		return nil, err
	}

	messagesSchema := typesense.CollectionSchema{
		Name:                "messages",
		DefaultSortingField: "created_at",
		Fields: []typesense.CollectionField{
			{
				Name: "content",
				Type: "string",
			},
			{
				Name: "author",
				Type: "int64",
			},
			{
				Name: "created_at",
				Type: "int64",
			},
		},
	}

	if _, err := client.CreateCollection(messagesSchema); err != nil {
		return nil, err
	}

	return &TypesenseBackend{
		Dependencies: deps,
		Client:       client,
	}, nil
}
