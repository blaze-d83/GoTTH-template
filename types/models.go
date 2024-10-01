package types


type Admin struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Pass      string `gorm:"not null"`
}
