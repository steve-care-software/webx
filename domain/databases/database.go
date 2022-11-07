package databases

import "net/url"

type database struct {
	head        Head
	pendings    Entries
	connections []url.URL
}

func createDatabase(
	head Head,
) Database {
	return createDatabaseInternally(head, nil, nil)
}

func createDatabaseWithPendings(
	head Head,
	pendings Entries,
) Database {
	return createDatabaseInternally(head, pendings, nil)
}

func createDatabaseWithConnections(
	head Head,
	connections []url.URL,
) Database {
	return createDatabaseInternally(head, nil, connections)
}

func createDatabaseWithPendingsAndConnections(
	head Head,
	pendings Entries,
	connections []url.URL,
) Database {
	return createDatabaseInternally(head, pendings, connections)
}

func createDatabaseInternally(
	head Head,
	pendings Entries,
	connections []url.URL,
) Database {
	out := database{
		head:        head,
		pendings:    pendings,
		connections: connections,
	}

	return &out
}

// Head returns the head
func (obj *database) Head() Head {
	return obj.head
}

// HasPendings returns true if there is pending entries, false otherwise
func (obj *database) HasPendings() bool {
	return obj.pendings != nil
}

// Pendings returns the pending entries, if any
func (obj *database) Pendings() Entries {
	return obj.pendings
}

// HasConnections returns true if there is connections, false otherwise
func (obj *database) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns the connections, if any
func (obj *database) Connections() []url.URL {
	return obj.connections
}
