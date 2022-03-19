import socket

HOST = '127.0.0.1' 
PORT = 2035

print(f"server on: {HOST}:{PORT}")

def main():
    srv = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    srv.bind((HOST, PORT))
    srv.listen(1)

    conn, addr = srv.accept()
    print('Connected by', addr)

    try:
        while True:
            data = conn.recv(1024)
            if not data: 
                break
            conn.sendall(b"in future i will'be an adnl server:)")
    except KeyboardInterrupt:
        conn.close()

    # conn.close()


if __name__ == "__main__":
    main()