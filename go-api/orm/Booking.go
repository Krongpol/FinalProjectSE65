package orm

import (
	"time"
	
	"gorm.io/gorm" // framwork ต่อกับ database ภาษา GO
)

// สร้าง structure เพื่อรองรับ json
type Booking struct {
	gorm.Model
	UserID string
	CarID  string
	Start  time.Time
	End    time.Time
}
