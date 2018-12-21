package test_model

import (
	"time"
)

type User struct {
	Serializer
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	Salary     float64   `json:"salary"`
	CreatedAt  time.Time `json:"created_at"`
}

func (u User) Validate() (bool, error) {
	return true, nil
}

func (u User) ToRepresentation(f func(Block)(func(b Block)Block,[]string))([]byte, error){
	return u.Serialize(u,f)
}
