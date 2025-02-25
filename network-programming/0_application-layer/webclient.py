# this script implements https://beej.us/guide/bgnet0/html/split/project-http-client-and-server.html#project-http-client-and-server

# It can be run from the command line, like
# $ python webclient.py example.com
# which defaults to port 80, or
# $ python webclient.py example.com 8088
# to retrieve the full HTTP response and payload

# In order to better understand the sockets API at a lower level, this project may not use any of the following helper functions:
# • The socket.create_connection() function.
# • The socket.create_server() function.
# • Anything in the urllib modules.

import sys
import socket

# Parse command line args
host = sys.argv[1]
port = sys.argv[2]

# The client connection process using the sockets API is described here
# https://beej.us/guide/bgnet0/html/split/introducing-the-sockets-api.html#client-connection-process

# Step 1. Ask the OS for a Socket:
client_socket = socket.socket()

# Step 2. is the DNS lookup, which Python's implementation of connect() takes care of for you.
# But if you really wanted to, you could do it yourself with `socket.getaddrinfo()`
# Step 3. Connect the socket to the destination IP address on the destination port
client_socket.connect(host, port)

# Step 4. Build and send the data.
# In this case, we're sending HTTP, so that's what we'll build, but it could be anything.
