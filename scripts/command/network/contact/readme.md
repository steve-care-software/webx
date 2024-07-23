# Contacts
A contact is an answer that have been accepted.  Contacts are ranked locally, so that others know how much we trust each of our contact.  This creates a web of trust that then tells us how much the people in our contact list trusts a new request.

```
// List the contacts:
retList = network.contact list;

// Retrieve a contact:
retRequest = network.contact retrieve {
    identifier: "myHash",
};

// Create the contact:
retAnswer = network.contact create {
    answer: myAnswer,
    rank: 34, // number between 0 and 1000
    notes: "This is some notes about the ranking", // optional
};

// Retrieve a specific contact:
retContact = network.contact retrieve {
    identifier: "myHash",
};

// Then we can access the contact properties:
retContact.answer
retContact.rank
retContact.notes
retContact.signature
```
