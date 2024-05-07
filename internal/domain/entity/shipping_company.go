package entity

type ShippingCompany struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
