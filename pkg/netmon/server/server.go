package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	ListenAddr	string
}

func (s *Server) Run(stopCh <- chan struct{}) {
	log.Printf("Starting netmon, listen: %s", s.ListenAddr)
	app := gin.New()
	RegisterRouter(app)

	server := http.Server{
		Addr: 		s.ListenAddr,
		Handler: 	app,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Fatal: %s", err)
		}
	}()

	<- stopCh
	log.Printf("Stopping netmon ...")
	log.Fatal("Server exited")
}