package decision

type DecisionCategory struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(20);not null"`
	Description string `gorm:"type:varchar(200);not null" `
	Value       int    `gorm:"type:int;not null"`
}
