package server

func is_Valid_Text(text string) bool {
	if text == "" { // empty text handling
		return false
	}
	for _, symbol := range text { // ascii table handling
		if symbol < 32 || symbol > 127 {
			return false
		}
	}
	return true
}
