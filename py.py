#开始于2021/7/13 
#第一天写了54行代码(包括注释)
import socket
import struct

HOST='127.0.0.1'
PORT=5678

def close():
    conn.close()
    s.close() 
    return

s= socket.socket(socket.AF_INET,socket.SOCK_STREAM)   #定义socket类型，网络通信协议为TCP
s.bind((HOST,PORT))   #套接字绑定的IP与端口
s.listen()
conn,_=s.accept()

#第一次确认
VER=conn.recv(1)
NMETHODS=conn.recv(1)
METHODS=conn.recv(1)
if VER!=b"\x05":
    if NMETHODS!=b"\x01":
        if METHODS!=b"\x00":
            print("连接失败")
            close()

#第二次确认
conn.sendall(b"\x05")
conn.sendall(b"\x00")

#由于返回值为\x00,所以无需第三次验证
#直接开始发送数据
VER=conn.recv(1)
CMD=conn.recv(1)
RSV=conn.recv(1)
ATYP=conn.recv(1)
DST_ADDR=conn.recv(4)
DST_PORT=conn.recv(2)

DST_ADDR=socket.inet_ntop(socket.AF_INET,DST_ADDR)
DST_PORT=struct.unpack(">H",DST_PORT)[0]#如果不加[0]的话返回值是tuple,加了之后该元组的[0]为端口,是一个int

conn.sendall(VER)
conn.sendall(b"\x00")
conn.sendall(RSV)
conn.sendall(b"\x01")
conn.sendall(socket.inet_pton(socket.AF_INET,HOST))
conn.sendall(struct.pack(">H",PORT))

data=conn.recv(4096)

close()






