package helper

import "fmt"

func BuildKey(tenant, user, endpoint string) string {
	return fmt.Sprintf("rl:%s:%s:%s", tenant, user, endpoint)
}
