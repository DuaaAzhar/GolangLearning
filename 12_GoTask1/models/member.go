package models

import "time"

type Member struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"creation-time"`
	UpdatedAt time.Time `json:"updation-time"`
}

func (m Member) Save() {

}

func (m Member) Update() {

}

func (m Member) Delete() {

}

func GetMemberById() {

}

func GetAllMembers() {

}
