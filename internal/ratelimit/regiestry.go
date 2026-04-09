package ratelimit

var engines = map[string]Engine{}

func Register(name string, engine Engine) {
	engines[name] = engine
}

func Get(name string) Engine {
	return engines[name]
}

// usage:
// engine := Get("redis")
// result, err := engine.Allow(ctx, policy, req)
