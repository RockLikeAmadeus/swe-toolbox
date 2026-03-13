# Network Concepts

Networking is just about exchanging arbitrary bytes of data between two or more computers, which means we need three things:

1. A way of identifying the source and destination machines
2. A way of maintaining data integrity
3. A way of routing data between the two computers (out of potentially billions of others)

And ultimately networking hardware and software is just about those three things.

Computer networks are packet-switched, meaning data is split up into individual packets containing a variable number of bytes that are transmitted over shared lines. Because the lines are shared (as opposed to the dedicated circuits in a circuit-switched network), routing is necessary, and that routing is effectively layered, like a piece of mail being sent overseas via planes and then large trucks and then smaller trucks until it reaches its destination.

## LANs

When we say computers are on the same LAN, that can mean either

1. They're connected by physical Ethernet cables
2. They're using the same WiFi access point

Both of these setups use the Ethernet protocol.

## [Network Programming](./network-programming/README.md)