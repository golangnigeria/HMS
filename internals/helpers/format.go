package helpers

import "fmt"

func FormatNairaShort(value float32) string {
	v := float64(value)

	switch {
	case v >= 1_000_000:
		return fmt.Sprintf("₦%.1fM", v/1_000_000)
	case v >= 1_000:
		return fmt.Sprintf("₦%.1fK", v/1_000)
	default:
		return fmt.Sprintf("₦%.0f", v)
	}
}
