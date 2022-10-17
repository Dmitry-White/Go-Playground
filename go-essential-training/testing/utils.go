package main

func getVerb(key string) string {
	switch key {
	case "A_1":
		return PATTERN.A_1
	case "A_2":
		return PATTERN.A_2
	case "A_3":
		return PATTERN.A_3
	default:
		return ""
	}
}
