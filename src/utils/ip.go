package utils

import "fmt"

func IntIPToString(ip int) string {
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		byte(ip>>24),
		byte(ip>>16),
		byte(ip>>8),
		byte(ip),
	)
}
