package activity

import (
	"time"

	"github.com/nullablenone/life-decision-tracker/domain/decision"
)

type Activity struct {
	ID         int                       `gorm:"primaryKey;autoIncrement"`
	BoardID    string                    `gorm:"type:varchar(50);not null;index"`
	Title      string                    `gorm:"type:varchar(255);not null"`
	DecisionID int                       `gorm:"not null"`
	Decision   decision.DecisionCategory `gorm:"foreignKey:DecisionID"` // Relasi ke tabel Master
	CreatedAt  time.Time                 `gorm:"autoCreateTime"`
}
