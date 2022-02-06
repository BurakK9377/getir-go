package record

import "time"

type Record struct {
	Key        string    `json:"key,omitempty" bson:"key,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	TotalCount int       `json:"totalCount,omitempty" bson:"totalCount,omitempty"`
}
