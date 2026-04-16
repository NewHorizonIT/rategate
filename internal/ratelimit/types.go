package ratelimit

type Request struct {
	Tenant   string
	Key      string
	Cost     int
	User     string
	Endpoint string
}

type Result struct {
	Allowed   bool
	Remaining int
}

type Policy struct {
	Limit    int
	Window   int
	Strategy string
	Burst    int
}
