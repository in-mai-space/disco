package models

type User struct {
	BaseModel
	Name         string         `json:"name" gorm:"not null" validate:"required"`
	Password     string         `json:"password" gorm:"not null" validate:"required,min=8"`
	Email        string         `json:"email" gorm:"unique;not null" validate:"required,email"`
	Plans        []Plan         `json:"plans" gorm:"foreignKey:UserID"`
	Goals        []Goal         `json:"goals" gorm:"foreignKey:UserID"`
	Notification []Notification `json:"notification" gorm:"foreignKey:UserID"`
	Buddies      []User         `json:"buddies" gorm:"many2many:buddies;foreignKey:ID;joinForeignKey:UserID;References:ID;JoinReferences:BuddyID"`
}

type UserCreateRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserNameUpdateRequestBody struct {
	Name string `json:"name"`
}
