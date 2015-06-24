golang-chat
===========

This is a simple tcp chat client and server with a simple binary protocol.

This example has a few rough corners where better error handling could be done. It is
only meant as an example to demonstrate how such a simple protocol could be implemented over tcp.

##The Protocol
A 4 byte int in big endian format is sent first. This int represents the length of the message to be sent
and is considered the header. The message is then sent.

This allows one to avoid the messiness of having to choose a delimiter and having arbitrary buffer sizes.
By using such a protocol we can determine the size of the buffer to be used for the message by reading
the 4 byte header and allocating a buffer large enough to accomodate the message.

If a message of more than 4GB needs to be sent, the header can be increased to 8 bytes. This has to be determined
carefully when creating the protocol since once defined the protocol cannot be extended without breaking backwards
compatibility.