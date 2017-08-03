#!/usr/bin/env python

import socket
import sys
import time

# Create a TCP/IP socket
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Connect the socket to the port where the server is listening
server_address = ('192.168.0.107', 3333)
print >>sys.stderr, 'connecting to %s port %s' % server_address
sock.connect(server_address)


i=40
inc_dec=10
while True:
	if i >= 160:
		inc_dec = -10
	elif i <= 10:
		inc_dec = 10

	i += inc_dec	

	if i>160 or i<10:
		i=40
	# Send data
	message = "90:15:" + str(i)
	print >>sys.stderr, 'sending "%s"' % message
	sock.sendall(message)

	# Look for the response
	amount_received = 0
	amount_expected = len(message)
	#q=0
	#while amount_received < amount_expected or q < 10:
		#data = sock.recv(16)
		#amount_received += len(data)
		#print >>sys.stderr, 'received "%s"' % data
		#q+=1

	time.sleep(0.25)

#	print >>sys.stderr, 'closing socket'
#	sock.close()