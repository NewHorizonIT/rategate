package policy

import (
	"gorm.io/gorm"
)

// Define repository layer for policy management, including interfaces and implementations for data access
type PolicyRepository struct {
	db *gorm.DB
}

func NewPolicyRepository(db *gorm.DB) *PolicyRepository {
	return &PolicyRepository{db: db}
}
