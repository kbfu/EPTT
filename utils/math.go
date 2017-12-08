package utils

func Round(n float64) int {
	fraction := n - float64(int(n))
	if fraction >= 0.5 {
		return int(n) + 1
	} else {
		return int(n)
	}
}
