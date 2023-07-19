package repository

type Book struct {
	Id     int64  `gorm:"primaryKey;not null"`
	Title  string `gorn:"not null"`
	Author string `gorn:"not null"`
}
