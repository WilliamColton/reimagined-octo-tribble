import socket
import struct

host="127.0.0.1"
port=5678

def close():
    conn.close()
    s.close()
    return

s=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
s.bind((host,port))
s.listen()
conn,_=s.accept()

header=conn.recv(2)
VER, NMETHODS = struct.unpack("!BB", header)
METHODS=conn.recv(NMETHODS)

conn.sendall(struct.pack("!BB",VER,0))

conn.recv(3)
address_type=conn.recv(1)
if address_type==b"\x01":
    to_address=socket.inet_ntoa(conn.recv(4))
elif address_type==b"\x03":
    domain_length = conn.recv(1)[0]   #!!!!!!!!
    to_address=socket.gethostbyname(conn.recv(domain_length))
else:
    close()
to_port=struct.unpack("!H",conn.recv(2))[0]

conn.sendall(b"\x05\x00\x00\x01")
conn.sendall(socket.inet_aton(host))
conn.sendall(struct.pack("!H",port))

data=conn.recv(4096)
to=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
to.connect((to_address,to_port))
to.sendall(data)
data=to.recv(4096)
conn.sendall(data)

    




