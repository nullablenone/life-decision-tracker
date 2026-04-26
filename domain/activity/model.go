package activity

import (
	"time"

	"github.com/nullablenone/life-decision-tracker/domain/decision"
)

type Board struct {
	ID        string    `gorm:"primaryKey;type:varchar(50)" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Activity struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	BoardID string `gorm:"type:varchar(50);not null;index"`
	Board      Board                     `gorm:"foreignKey:BoardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title      string                    `gorm:"type:varchar(255);not null"`
	DecisionID int                       `gorm:"not null"`
	Decision   decision.DecisionCategory `gorm:"foreignKey:DecisionID"`
	CreatedAt  time.Time                 `gorm:"autoCreateTime"`
}
