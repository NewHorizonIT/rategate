package policy

import (
	"gorm.io/gorm"
)

// Initialize and register the policy management components, including repositories and services, to be used by the application
func Register(db *gorm.DB) error {
	// Initialize repository
	policyRepo := NewPolicyRepository(db)
	println(policyRepo)
	// Initialize service

	return nil
}
