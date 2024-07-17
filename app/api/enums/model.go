package enums

type PaymentStatus struct {
	ID   int    `gorm:"primaryKey;not null;autoIncrement" validate:"required"`
	Name string `gorm:"not null;size:50"`
}

type PaymentType struct {
	ID   int    `gorm:"primaryKey;not null;autoIncrement" validate:"required"`
	Name string `gorm:"not null;size:50"`
}
