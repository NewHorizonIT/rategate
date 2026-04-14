package policy

// Define strategy types
type Strategy uint8

const (
	TokenBucket = iota
)

// Define the Policy struct
type Policy struct {
	Limit    int32
	Window   int32
	Burst    int32
	Strategy Strategy
	Version  int32
}

// Define the PolicyRepository interface for data access
type IPolicyRepository interface {
	GetPolicy(key Key) (*Policy, error)
	SetPolicy(key Key, policy *Policy) error
	DeletePolicy(key Key) error
}

// Define Key struct for identifying policies
type Key struct {
	TenantID   uint64
	APIKeyID   uint64
	EndpointID uint64
}
