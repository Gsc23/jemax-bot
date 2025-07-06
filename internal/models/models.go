package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ================== BASE MODEL =====================

type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// ================== CUSTOMER =====================

type Customer struct {
	Model
	Phone string   `gorm:"uniqueIndex;not null"` 
	Order []*Order 
}

// ================== ORDER =====================

type Order struct {
	Model
	CustomerID uuid.UUID    `gorm:"not null"`
	Customer   *Customer    
	Status     string       `gorm:"not null;default:'pendente'"`
	Items      []*OrderItem `gorm:"foreignKey:OrderID"` 
	TotalPrice int32        `gorm:"not null;default:0"` 
}

// ================== ORDER ITEM =====================

type OrderItem struct {
	OrderID   uuid.UUID `gorm:"primaryKey"`
	ProductID uuid.UUID `gorm:"primaryKey"`
	Product   *Product
	Quantity  int32 `gorm:"not null"`
	UnitPrice int32 `gorm:"not null"`
}

// ================== INGREDIENT =====================

type Ingredient struct {
	Model
	Name            string `gorm:"uniqueIndex;not null"` 
	QuantityInStock int32  `gorm:"not null;default:0"`   
	Unit            string `gorm:"not null"`             
}

// ================== PRODUCT =====================

type Product struct {
	Model
	SKU         string `gorm:"uniqueIndex;not null"`
	Name        string `gorm:"not null"`
	Description string
	Price       int32 `gorm:"not null"` // Preço em centavos
	IsAvailable bool  `gorm:"default:true"`
	Ingredients []*ProductIngredient
}

// ================== PRODUCT INGREDIENT (RECEITA) =====================

type ProductIngredient struct {
	ProductID        uuid.UUID   `gorm:"primaryKey"`
	IngredientID     uuid.UUID   `gorm:"primaryKey"`
	Ingredient       *Ingredient `gorm:"foreignKey:IngredientID"`
	QuantityRequired int32       `gorm:"not null"`
}

// ================= PROMOTION ====================

type Promotion struct {
	Model
	Name        string `gorm:"not null"`
	Description string
	Value       int32               `gorm:"not null"`
	StartDate   time.Time           `gorm:"not null"`
	EndDate     time.Time           `gorm:"not null"`
	IsActive    bool                `gorm:"default:true"`
	Items       []*PromotionProduct `gorm:"foreignKey:PromotionID"` 
}

// ============== PROMOTION_PRODUCT =================

type PromotionProduct struct {
	PromotionID uuid.UUID `gorm:"primaryKey"`
	ProductID   uuid.UUID `gorm:"primaryKey"`
	Product     *Product
	Quantity    int32 `gorm:"not null"`
}
