package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"` //omitempty -> omits if value is empty
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
