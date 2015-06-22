golang-chat
===========

This is a simple tcp chat client and server with a simple binary protocol.

This example has a few rough corners where better error handling could be done. It is
only meant as an example to demonstrate how such a simple protocol could be implemented over tcp.

##The Protocol
A 4 byte int in big endian format is sent first. This int represents the length of the message to be sent.
The message is then sent.

This allows one to avoid the messiness of having to choose a delimiter and having arbitrary buffer sizes.
By using such a protocol we can determine the size of the buffer to be used for the message by reading
the 4 byte predecessor and allocating a buffer large enough to accomodate the message.

If a message of more than 4GB needs to be sent, the predecessor can be increased to 8 bytes. This has to be determined
carefully when creating the protocol.