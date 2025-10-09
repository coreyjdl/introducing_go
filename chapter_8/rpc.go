package main

import ("fmt"; "net"; "net/rpc")

type Server struct{}
func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil { panic(err) }

	for {
		conn, err := ln.Accept()
		if err != nil { continue }

		go rpc.ServeConn(conn)
	}
}

func client() {
	conn, err := rpc.Dial("tcp", "localhost:9999")
	if err != nil { panic(err) }

	var result int64

	err = conn.Call("Server.Negate", int64(34), &result)
	if err != nil { panic(err) }

	fmt.Println("Result:", result)
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}