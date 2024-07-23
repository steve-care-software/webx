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

// Brodcast an invite request:
network.contact.request brodcast {
    request: myRequest,
};

// List the invite requests that are still pending:
retList = network.contact.request list;

// Retrieve a specific invite request:
retInvite = network.contact.request retrieve {
    identifier: "myHash",
};

// Then we can access the invite properties:
retInvite.name
retInvite.description
retInvite.local
retInvite.destination
```

## Answers
An answer send a reply to an invite request that we received.
```
// Retrieve a request:
retRequest = network.contact.request retrieve {
    identifier: "myHash",
};

// Create the answer:
retAnswer = network.contact.request.answer create {
    request: retRequest,
    isApproved: true,
    reason: "This is a reason the other party will see", // optional
};

// Brodcast the answer:
network.contact.request.answer brodcast {
    answer: retAnswer,
};

// List the answers:
retList = network.contact.request.answer list;

// Retrieve a specific answer:
retAnswer = network.contact.request retrieve {
    identifier: "myHash",
};

// Then we can access the answer properties:
retAnswer.request
retAnswer.isApproved
retInvite.reason
retInvite.signature
```

## Contacts
### Ranks
