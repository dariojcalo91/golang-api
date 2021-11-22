package test

import (
	"log"
	"net/http/httptest"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
)

func TestHandleListBills(t *testing.T) {
	ja := jsonassert.New(t)

	w := httptest.NewRecorder()
	log.Println("test code", w.Code)
	if !assert.Equal(t, 200, w.Code) {
		panic("Status not equal")
	}
	ja.Assertf(`
    [
		{
			"id": 1,
			"detail": "Bill #1",
			"cost": "100.50",
			"start_date": "2021-11-21T22:52:41.047154-05:00",
			"end_date": "2021-11-21T22:52:41.047154001-05:00",
			"status": "paid"
		},
		{
			"id": 2,
			"detail": "Bill #2",
			"cost": "50.50",
			"start_date": "2021-11-21T22:52:41.047165-05:00",
			"end_date": "2021-11-21T22:52:41.047165001-05:00",
			"status": "pending"
		},
		{
			"id": 3,
			"detail": "Bill #3",
			"cost": "85.00",
			"start_date": "2021-11-21T22:52:41.047165-05:00",
			"end_date": "2021-11-21T22:52:41.047165001-05:00",
			"status": "paid"
		}
	]`, `
    [
		{
			"id": 1,
			"detail": "Bill #1",
			"cost": "100.50",
			"start_date": "2021-11-21T22:52:41.047154-05:00",
			"end_date": "2021-11-21T22:52:41.047154001-05:00",
			"status": "paid"
		},
		"<<PRESENCE>>",
		"<<PRESENCE>>"
	]`)
}
