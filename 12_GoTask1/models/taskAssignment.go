package models

import "time"

type TaskAssignment struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"creation-time"`
	UpdatedAt time.Time `json:"updation-time"`
}

func (ta TaskAssignment) Save() {

}

func (ta TaskAssignment) Update() {

}

func (ta TaskAssignment) Delete() {

}
