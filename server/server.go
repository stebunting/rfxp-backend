package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stebunting/rfxp-backend/router"
)

type Server struct {
	port       int
	httpServer *http.Server
}

type ServerConfig struct {
	Port int
}

func NewServer(config ServerConfig) *Server {
	s := &Server{
		port: config.Port,
	}

	return s
}

func (s *Server) Start() {
	s.httpServer = &http.Server{
		Addr: fmt.Sprintf(":%d", s.port),
	}

	http.HandleFunc("/getdata", router.GetData)

	fmt.Printf("Server listening on port %d...\n", s.port)
	err := s.httpServer.ListenAndServe()

	if err != http.ErrServerClosed {
		panic(err)
	} else {
		fmt.Println("Server Stopped")
	}
}

func (s *Server) Stop() {
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			panic(err)
		}
	}
}
