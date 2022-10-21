package databases

import "net/url"

type database struct {
	content     Content
	connections []url.URL
}

func createDatabase(
	content Content,
) Database {
	return createDatabaseInternally(content, nil)
}

func createDatabaseWithConnections(
	content Content,
	connections []url.URL,
) Database {
	return createDatabaseInternally(content, connections)
}

func createDatabaseInternally(
	content Content,
	connections []url.URL,
) Database {
	out := database{
		content:     content,
		connections: connections,
	}

	return &out
}

// Content returns the content
func (obj *database) Content() Content {
	return obj.content
}

// HasConnections returns true if there is connections, false otherwise
func (obj *database) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns the connections, if any
func (obj *database) Connections() []url.URL {
	return obj.connections
}
