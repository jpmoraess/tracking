package entity

type CustomerIntegration struct {
	ID                    string `json:"id,omitempty" bson:"_id,omitempty"`
	CustomerId            string `json:"customerId" bson:"customerId"`
	CustomerSystemId      string `json:"customerSystemId" bson:"customerSystemId"`
	ShippingCompanyId     string `json:"shippingCompanyId" bson:"shippingCompanyId"`
	CustomerSystemApiKey  string `json:"customerSystemApiKey" bson:"customerSystemApiKey"`
	ShippingCompanyApiKey string `json:"shippingCompanyApiKey" bson:"shippingCompanyApiKey"`
	Valid                 bool   `json:"valid" bson:"valid"`
}
