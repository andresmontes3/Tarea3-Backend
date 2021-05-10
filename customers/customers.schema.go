package customers

type Customer struct {
	CustomerID   string `json:"customer_id"`
	CompanyName  string `json:"company_name"`
	ContactName  string `json:"contact_name"`
	ContactTitle string `json:"contact_title"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Region       string `json:"region"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
}
