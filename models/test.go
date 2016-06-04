package models

// TestPingResponse returns API heartbeat status code
type TestPingResponse struct {
	Status int
}

// TestResponse returns test JSON results from jsonplaceholder.typicode.com
type TestResponse struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
