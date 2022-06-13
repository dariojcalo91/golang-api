package adapters

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type bill struct {
	ID        int       `json:"id"`
	Detail    string    `json:"detail"`
	Cost      string    `json:"cost"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
}

var billList = []bill{
	{ID: 1, Detail: "Bill #1", Cost: "100.50", StartDate: time.Now(), EndDate: time.Now().Add(1), Status: "paid"},
	{ID: 2, Detail: "Bill #2", Cost: "50.50", StartDate: time.Now(), EndDate: time.Now().Add(1), Status: "pending"},
	{ID: 3, Detail: "Bill #3", Cost: "85.00", StartDate: time.Now(), EndDate: time.Now().Add(1), Status: "paid"},
}

func (s *Server) handleListBills(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, billList)
}

func (s *Server) handleCreateBill(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, billList)
}

func (s *Server) handleUpdateBill(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, billList)
}

func (s *Server) handleDeleteBill(c *gin.Context) {
	var err error
	billId, err := strconv.ParseInt(c.Param("bill_id"), 10, 32)
	if err != nil || billId < 1 {
		log.Fatalln("bill Id should be a positive integer")
		return
	}

	for _, bill := range billList {
		if bill.ID == int(billId) {
			log.Println("Deleting: ", bill)
			c.Done()
		}
	}
}
