# Request
A request represents an invitation to someone to connect.

```
// List the invite requests that are still pending:
retList = network.contact.answer.request list;

// Create an invite request:
myRequest = network.contact.answer.request create {
    name: "myName",
    description: "This is some details about me",
    local: "http://myHost:8081",
    destination: "http://destination:8081",
};

// Brodcast an invite request:
network.contact.answer.request brodcast {
    request: myRequest,
};

// Retrieve a specific invite request:
retInvite = network.contact.answer.request retrieve {
    identifier: "myHash",
};

// Then we can access the invite properties:
retInvite.name
retInvite.description
retInvite.local
retInvite.destination
```
