import socket
import struct
import sys

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

ver,nmethods = conn.recv(1),conn.recv(1)
methods = conn.recv(ord(nmethods))
conn.sendall(b'\x05\x00')
ver,cmd,rsv,atype = conn.recv(1),conn.recv(1),conn.recv(1),conn.recv(1)
if ord(cmd)!=1:
    close()
    sys.exit(0)
if ord(atype) == 1:
    # IPv4
    to_addr = socket.inet_ntoa(conn.recv(4))
    to_port = struct.unpack(">H", conn.recv(2))[0]
elif ord(atype) == 3:
    # 域名
    addr_len = ord(conn.recv(1))
    to_addr = conn.recv(addr_len)
    to_port = struct.unpack(">H", conn.recv(2))[0]
    print(to_addr,to_port)
else:
    # 不支持的地址类型(such as IPv6)
    close()
    sys.exit()

to=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
to.connect((to_addr,to_port))
print("to is successful")

reply = b"\x05\x00\x00\x01" + socket.inet_aton(host) + struct.pack(">H", port)
conn.sendall(reply)
print("reply send successful")

data=conn.recv(4096)
print("data recv successful")
to.sendall(data)
data=to.recv(4096)
conn.sendall(data)
print("data change successful")
