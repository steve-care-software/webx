
# Answers
An answer send a reply to an invite request that we received.
```
// List the answers:
retList = network.answer.request list;

// Retrieve a request:
retRequest = network.answer.request retrieve {
    identifier: "myHash",
};

// Create the answer:
retAnswer = network.answer.request create {
    request: retRequest,
    isApproved: true,
    reason: "This is a reason the other party will see", // optional
};

// Brodcast the answer:
network.answer.request brodcast {
    answer: retAnswer,
};

// Retrieve a specific answer:
retAnswer = network.answer.request retrieve {
    identifier: "myHash",
};

// Then we can access the answer properties:
retAnswer.request
retAnswer.isApproved
retInvite.reason
retInvite.signature
```

