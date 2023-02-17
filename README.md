# spritz.go
A Spritz cipher implementation in pure Go.

Spritz is a RC4 redesign by *Ron Rivest* and *Jacob Schuldt*
[(PDF)](doc/RS14.pdf).

# Exports

* `Spritz.Encrypt(message, key []byte)`
* `Spritz.Decrypt(message, key []byte)`
* `Spritz.Hash(message, digest []byte)`

# License
Release into the [Public Domain](LICENSE.txt).