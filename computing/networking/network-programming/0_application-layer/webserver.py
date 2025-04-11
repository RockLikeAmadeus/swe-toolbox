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

BUFFER_SIZE_BYTES = 4096

# Parse command line args
port = 28333 if len(sys.argv) < 2 else int(sys.argv[1]) 

# The server listening process using the sockets API is described here
# https://beej.us/guide/bgnet0/html/split/introducing-the-sockets-api.html#server-listening-process

# Step 1. Ask the OS for a Socket:
server_socket = socket.socket()

# Do that thing beej says here

# Step 2. Bind the socket to a port
# This just tells the OS what port you'll be listening on
server_socket.bind(('', port)) # '' here means choose any local address

# Step 3. Listen for incoming connections
server_socket.listen()

# Step 4. Accept incoming connections
# The accept() function blocks until a client tries to connect
# Then, it returns a brand new socket specifically for the connection to that client
while True:
    new_socket, return_addr = server_socket.accept()
    # Step 5. Send and receive data
    # A real-world server would probably spin off a separate thread process to maintain this 
    # connection and send data back and forth to the client, but in this simple case we'll just send 
    # a static response and close the connection before checking for the next client connection
    request = ""
    while True:
        request += new_socket.recv(BUFFER_SIZE_BYTES).decode("ISO-8859-1")
        # This should be an HTTP request, so we're looking for a blank line to know we're done
        if ("\r\n\r\n" in request):
            break
    print(request)

    # response here

    new_socket.close()

# Step 6. Go back and accept another connection