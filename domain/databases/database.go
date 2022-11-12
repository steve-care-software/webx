package databases

import "net/url"

type database struct {
	head        Head
	connections []url.URL
}

func createDatabase(
	head Head,
) Database {
	return createDatabaseInternally(head, nil)
}

func createDatabaseWithConnections(
	head Head,
	connections []url.URL,
) Database {
	return createDatabaseInternally(head, connections)
}

func createDatabaseInternally(
	head Head,
	connections []url.URL,
) Database {
	out := database{
		head:        head,
		connections: connections,
	}

	return &out
}

// Head returns the head
func (obj *database) Head() Head {
	return obj.head
}

// HasConnections returns true if there is connections, false otherwise
func (obj *database) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns the connections, if any
func (obj *database) Connections() []url.URL {
	return obj.connections
}
