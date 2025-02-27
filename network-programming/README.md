The network is hardware, and in non-embedded systems, we can only access the hardware _through the OS_. Unix systems expose the [sockets API](https://beej.us/guide/bgnet/) to C programs (there are equivalents for other operating systems). Other languages might provide direct translation to the C API, or might wrap it into a higher abstraction level, for communicating over the network. A _socket_ is simply an abstraction that represents the endpoint within a communication network that can be named and/or addressed, while hiding the details of the underlying transport-layer protocols.

Although there exists the [OSI model](https://en.wikipedia.org/wiki/OSI_model), for most real-world purposes the simplified [internet layer model](https://en.wikipedia.org/wiki/Internet_layer), which contains just four layers, is just as, if not more, useful. When doing network programming, it's important to ask what layer we are concerned with, and in most cases we're working at the application layer (FTP, HTTP). If we're communicating between systems that define their own protocols, we're probably implementing those protocols on top of the transport layer (TCP, UDP, or even the low-latency QUIC protocol), and are unlikely to need to go deeper than that. The sockets API is what wraps up and hides TCP or UDP from us. If we're going to construct our own implementation of an application-layer protocol, like HTTP, or perhaps something more use-case targeted, we would do it on top of the sockets API.

Both TCP and UDP identify senders and receivers of data using _ports_. The sockets API calls TCP sockets _stream sockets_, and it calls UDP sockets _datagram sockets_. 

Python defines the [socket](https://docs.python.org/3/library/socket.html) library which provides "a straightforward transliteration of the Unix system call and library interface for sockets to Python’s object-oriented style", making it a common choice for network programming when we don't want to work at a higher level than raw C demands.

C# provides helper classes like the `TcpClient` and `TcpListener` for simple use cases, but for more fine-grained control and performance you might instead use the `Socket` class.

Other languages provide varying levels of abstraction on top of the sockets API or its equivalent, but the sockets API, which wraps TCP/IP protocols and other messy details, is in most cases the lowest level at which you'll need to write code for non-embedded systems.

---

Connecting to a remote process through the sockets API looks something like [this simple script](./0_application-layer/webclient.py) for clients. The client script can make requests to whatever real-world host or IP address you provide it.

A simple server that can listen for requests from a client using the sockets API looks like [this script](webserver.py). The server can be reached out to via your web browser. 

If you start up the server script in a terminal window (`python webserver.py`), and then execute the client script in a separate terminal (`python webclient.py 127.0.0.1 28333`), you can see these two scripts talk to each other over the socket API.