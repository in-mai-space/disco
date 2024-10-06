package models

type User struct {
	BaseModel
	Name         string         `json:"name" gorm:"not null"`
	Password     string         `json:"password" gorm:"not null"`
	Email        string         `json:"email" gorm:"unique;not null"`
	Plans        []Plan         `json:"plans" gorm:"foreignKey:UserID"`
	Goals        []Goal         `json:"goals" gorm:"foreignKey:UserID"`
	Notification []Notification `json:"notification" gorm:"foreignKey:UserID"`
	Buddies      []User         `json:"buddies" gorm:"many2many:buddies;foreignKey:ID;joinForeignKey:UserID;References:ID;JoinReferences:BuddyID"`
}
