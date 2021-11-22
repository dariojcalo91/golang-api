package adapters

import "github.com/gin-gonic/gin"

func (s *Server) routes() {
	s.router = gin.New()
	core := s.router.Group("/")
	core.Use()
	{
		core.GET("/bills", s.handleListBills)
		core.POST("/bill", s.handleCreateBill)
		core.POST("/bill/:bill_id", s.handleUpdateBill)
		core.DELETE("/bill/:bill_id", s.handleDeleteBill)
	}
}
