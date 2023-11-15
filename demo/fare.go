package demo

func CalculateFare(kind string, n uint) uint {
	switch kind {
	case "MRT":
		return 14 + ((n - 1) * 2)
	case "BTS":
		return 18 + ((n - 1) * 5)
	case "Airport Link":
		return 16 + ((n - 1) * 3)
	default:
		return 0
	}
}
