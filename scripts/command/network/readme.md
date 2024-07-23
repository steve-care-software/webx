# Network
The network enables to create decentralized handshake, build contact lists, enable common database sync proposals and message people from our contact list.

## Request
A request represents an invitation to someone to connect.
```
// Create an invite request:
myRequest = network.contact.request create {
    name: "myName",
    description: "This is some details about me",
    local: "http://myHost:8081",
    destination: "http://destination:8081",
};

// List the invite requests that are still pending:
retList = network.contact.request list {
    name: "myName",
    description: "This is some details about me",
    local: "http://myHost:8081",
    destination: "http://destination:8081",
};

// Retrieve a specific invite request:
retInvite = network.contact.request retrieve {
    identifier: "myHash",
};

// Then we can access the invite properties
retInvite.name
retInvite.description
retInvite.local
retInvite.destination
```

## Answers
An answer send a reply to an invite request that we received.
```
```

## Contacts
### Ranks
