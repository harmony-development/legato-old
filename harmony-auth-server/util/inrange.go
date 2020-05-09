package util

// InRange checks if a number is between a min and a max
func InRange(num int, min int, max int) bool {
	return num >= min && num <= max
}