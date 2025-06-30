package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ================== USER =====================

type User struct {
	Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Phonenumber string    `gorm:"unique;not null"`
	Status      string
	IsAdmin     bool
}

// ================== ORDER =====================

type Order struct {
	Model
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey"`
	Phonenumber string        `gorm:"not null"`
	Items       []*OrderItem   `gorm:"foreignKey:OrderID"`
}

// ================== ITEM =====================

type Item struct {
	Model
	ID    uuid.UUID   `gorm:"type:uuid;primaryKey"`
	Name  string
	Stock int         // estoque total
	Orders []*OrderItem `gorm:"foreignKey:ItemID"`
}

// =============== ORDER_ITEM ==================

type OrderItem struct {
	Model
	OrderID    uuid.UUID
	Order      Order
	ItemID     uuid.UUID
	Item       Item
	Quantity   int64
}

// ================== STORAGE =====================

type Storage struct {
	Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string
	Quantity int
}
