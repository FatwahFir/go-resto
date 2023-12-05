package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        string    `json:"id" gorm:"column:id;primaryKey;<-:create"`
	Name      string    `json:"name" gorm:"column:name;" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Menus     []Menu    `gorm:"foreignKey:category_id;references:id"`
}

func (c *Category) TableName() string {
	return "categories"
}

func (c *Category) BeforeCreate(db *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.NewString()
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
