package models

type Netflix struct {
	// ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string `json:"movie"`
	Watched bool   `json:"watched"`
}
