package models

type Account struct {
	ID      uint   `gorm:"primaryKey json:id"`
	Name    string `gorm:"size:100 json:name"`
	Balance int    `gorm:"not null json:balance"`
}
