package tcp

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// Server is the tcp server struct
type Server struct {
	Ports []string
}

// NewServer creates a new tcp server
func NewServer(ports []string) *Server {
	return &Server{ports}
}

// Start starts the tcp server
func (t *Server) Start() {
	var wg sync.WaitGroup
	wg.Add(len(t.Ports))
	for _, port := range t.Ports {
		go func(port string, wg *sync.WaitGroup) {
			fmt.Printf("Listening on tcp port: %v\n", port)
			listen, err := net.Listen("tcp", ":"+port)
			if err != nil {
				log.Println("Unable to net.Listen. Received error: ", err)
				wg.Done()
				return
			}
			// addr, err := net.ResolveTCPAddr("tcp", ":"+port)
			// if err != nil {
			// 	log.Println("Unable to net.ResolveTCPAddr. Received error: ", err)
			// 	wg.Done()
			// 	return
			// }
			// listen2, err := net.ListenTCP("tcp", addr)
			// if err != nil {
			// 	log.Println("Unable to net.ListenTCP. Received error: ", err)
			// 	wg.Done()
			// 	return
			// }
			for {
				conn, err := listen.Accept()
				if err != nil {
					log.Fatal(err)
					// handle error
				}

				// conn2, err := listen2.AcceptTCP()
				// if err != nil {
				// 	log.Fatal(err)
				// 	// handle error
				// }
				// Enable Keepalives
				// err = conn2.SetKeepAlive(true)
				// if err != nil {
				// 	fmt.Printf("Unable to set keepalive - %s", err)
				// }
				go handleConnection(conn)
				// go handleConnection(conn2)

			}
		}(port, &wg)
	}
	wg.Wait()
	fmt.Println("TCP Server Stopped")
}

func handleConnection(conn net.Conn) {
	fmt.Println("Handle Conection received: connection")
	remHost, remPort, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		fmt.Printf("Failed to split remote host and port: %v\n", err)
		return
	}
	locHost, locPort, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		fmt.Printf("Failed to split remote host and port: %v\n", err)
		return
	}
	str := fmt.Sprintf(`Date: %v, InIp: %v, InPort: %v, DestIP: %v, DestPort: %v`,
		time.Now().Format("20060102150405"), remHost, remPort, locHost, locPort)
	fmt.Println(str)
	data := make([]byte, 50)
	n, err := conn.Read(data)
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}
	defer conn.Close()
	fmt.Printf("Received data from %v, of length %v data is %s\n", conn.RemoteAddr(), n, data[:n])
	conn.Write(data[:n])
}
