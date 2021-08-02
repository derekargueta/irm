package util

func Percent(part, total int) float32 {
	return (float32(part) / float32(total)) * 100
}
