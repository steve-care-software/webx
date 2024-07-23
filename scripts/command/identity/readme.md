# Identity
The identity contains a name, signature/voting encryption keys and encryption encryption keys.  The identity data is encrypted using the provided password before being saved on disk.

```
// List identities:
command.identity list;

// Create an identity:
command.identity create {
    name: "myName",
    password: "myPassword",
};

// Delete an identity:
command.identity delete {
    name: "myName",
    password: "myPassword",
};

// Update an identity:
command.identity update {
    name: {
        current: "myName",
        updated: "myNewName", // optional
    },
    password: {
        current: "myPassword",
        updated: "myNewPassword", // optional
    },
};

// Authenticate:
myIdentity = command.identity authenticate  {
    name: "myName",
    password: "myPassword",
};

// Then we can access the identity's properties:
myIdentity.name
myIdentity.signature.publickey
myIdentity.signature.ring
myIdentity.encryption.publickey
```

## Signing a request
A signature tells people that our specific identity has created the signature.
```
// First Authenticate:
myIdentity = command.identity authenticate  {
    name: "myName",
    password: "myPassword",
};

// Sign:
retSignature = command.identity sign {
    identity: myIdentity,
    message: "this is some data",
};
```

## Validating a signature
```
// First Authenticate:
myIdentity = command.identity authenticate  {
    name: "myName",
    password: "myPassword",
};

// Verify the signature:
isVerified = command.identity.signature validate {
    identity: myIdentity,
    signature: myReceivedSignature,
    message: "this is the original message",
};

```

## Voting on a request
A vote makes it impossible to know who specifically created the vote, but can be validated against a ring of public keys.  It enables the signature of messages without telling exactly which public key voted on the request.
```
// First Authenticate:
myIdentity = command.identity authenticate  {
    name: "myName",
    password: "myPassword",
};

// Vote:
retVote = command.identity vote {
    identity: myIdentity,
    message: "this is some data",
};
```

## Validating a evote
```
// First Authenticate:
myIdentity = command.identity authenticate  {
    name: "myName",
    password: "myPassword",
};

// Verify the vote:
isVerified = command.identity.vote validate {
    identity: myIdentity,
    signature: myReceivedSignature,
    message: "this is the original message",
};

```