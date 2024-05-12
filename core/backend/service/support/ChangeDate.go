package support

func ChangeDate(date []byte) string {
	if len(date) == 0 {
		return ""
	}
	newdate := make([]byte, 4)
	copy(newdate, date[:4])

	for i := 2; i < len(date); i++ {
		date[i-2] = date[i]
	}
	date[0] = date[len(date)-2]
	date[1] = date[len(date)-1]
	for i := len(date) - 4; i < len(date); i++ {
		date[i] = newdate[i-(len(date)-4)]
	}
	return string(date)
}
