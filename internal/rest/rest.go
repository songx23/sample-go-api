package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	e    *gin.Engine
	port int
}

func New(port int) *Server {
	r := gin.Default()
	s := &Server{
		port: port,
		e:    r,
	}
	s.initRoute()
	return s
}

func (s *Server) initRoute() {
	s.e.GET("/greeting", handleGreeting)
	s.e.GET("/go-crazy", handleGoCrazy)
}

func (s *Server) Run() {
	fmt.Println(fmt.Sprintf("Server running on port: %d", s.port))
	s.e.Run(fmt.Sprintf(":%d", s.port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
