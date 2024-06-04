package api

import (
	"fmt"
	"io"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Printf("New connection client: %s \n", ws.RemoteAddr())

	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) HandleWsFeed(ws *websocket.Conn) {
	for {
		payload := fmt.Sprintf("Order -> %d", time.Now().UnixNano())
		if _, err := ws.Write([]byte(payload)); err != nil {
			fmt.Printf("Write message error: %s", err)
		}
		time.Sleep(time.Second * 2)
	}
}

func (s *Server) broadcastMsg(b []byte) {
	for ws := range s.conns {
		if _, err := ws.Write(b); err != nil {
			fmt.Printf("Write message error: %s", err)
		}
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read message error: %s \n", err)
			continue
		}
		msg := buf[:n]

		fmt.Println(string(msg))
		fmt.Println(("Message received from: " + ws.RemoteAddr().String()))

		go s.broadcastMsg(msg)
	}
}
