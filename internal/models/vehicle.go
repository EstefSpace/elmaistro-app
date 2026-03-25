package models

type Vehicle struct {
	ID         int64  `json:"id"`
	CustomerID int64  `json:"customer_id"` // Relation with customer
	Make       string `json:"make"`        // Example: Nissan
	Model      string `json:"model"`       // Example: Sentra
	Year       int    `json:"year"`
	Plate      string `json:"plate"` // Plates (Unique)
	VIN        string `json:"vin"`   // Serial number
	Color      string `json:"color"`
}
