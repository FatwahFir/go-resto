package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID          string    `json:"id" gorm:"column:id;primaryKey;<-:create"`
	CategoryID  string    `json:"category_id" gorm:"column:category_id;" validate:"required"`
	Name        string    `json:"name" gorm:"column:name;" validate:"required"`
	Price       int       `json:"price" gorm:"column:price;" validate:"required"`
	PriceModal  int       `json:"price_modal" gorm:"column:price_modal;" validate:"required"`
	Description string    `json:"description" gorm:"column:description;" validate:"required"`
	IsAvailable bool      `json:"is_available" gorm:"column:is_available;"`
	Image       string    `json:"image" gorm:"column:image;"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Category    *Category `gorm:"foreignKey:category_id;references:id"`
}

func (m *Menu) TableName() string {
	return "menus"
}

func (m *Menu) BeforeCreate(db *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}

//CREATE MIGRATIONS
// migrate create -ext sql -dir db/migrations create_table_menus

//UP
//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations up

//DOWN
//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations down

//UP SPECIFIC VERSION
//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations UP (1..N)

//DOWN SPECIFIC VERSION
//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations DOWN (1..N)

//CHECK VERSION
//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations version

//FORCE MIGRATION
//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations force 000001(version)
