package fetchers

import "github.com/steve-care-software/webx/domain/databases/entities"

type fetcher struct {
	entity  entities.Entity
	content Content
}

func createFetcher(
	entity entities.Entity,
	content Content,
) Fetcher {
	out := fetcher{
		entity:  entity,
		content: content,
	}

	return &out
}

// Entity returns the entity
func (obj *fetcher) Entity() entities.Entity {
	return obj.entity
}

// Content returns the content
func (obj *fetcher) Content() Content {
	return obj.content
}
