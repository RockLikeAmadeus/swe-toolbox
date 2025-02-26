# this script implements https://beej.us/guide/bgnet0/html/split/project-http-client-and-server.html#project-http-client-and-server

# It can be run from the command line, like
# $ python webserver.py
# which defaults to listening on port 28333, or
# $ python webserver.py 12399
# to handle requests and send back a simple static response message

# In order to better understand the sockets API at a lower level, this project may not use any of the following helper functions:
# • The socket.create_connection() function.
# • The socket.create_server() function.
# • Anything in the urllib modules.

import sys
import socket

# Parse command line args
host = sys.argv[1]
port = 80 if len(sys.argv) < 3 else int(sys.argv[2]) 

# The client connection process using the sockets API is described here
# https://beej.us/guide/bgnet0/html/split/introducing-the-sockets-api.html#client-connection-process

# Step 1. Ask the OS for a Socket:
client_socket = socket.socket()

# Step 2. is the DNS lookup, which Python's implementation of connect() takes care of for you.
# But if you really wanted to, you could do it yourself with `socket.getaddrinfo()`
# Step 3. Connect the socket to the destination IP address on the destination port
client_socket.connect((host, port))

# Step 4. Build and send the data.
# In this case, we're sending a so-called "HTTP request", which is why it has this specific form,
# but it could be anything.
str_msg = """GET / HTTP/1.1\r
Host: example.com\r
Connection: close\r

"""

#  Over the web, the standard string encoding is ISO-8859-1, so convert from a string to an array
# of raw bytes
msg = str_msg.encode("ISO-8859-1")

# Now send it. We do it all at once, which the Python wrapper over the sockets API makes easy for us
client_socket.sendall(msg)

# Now we wait for the response, and we keep reading the response until `recv` is empty, which tells
# us that there's no more data to read
while True:
    reply = client_socket.recv(4096)
    print(reply)
    if (len(reply) == 0):
        break

# Step 5. Close up the connection
client_socket.close()