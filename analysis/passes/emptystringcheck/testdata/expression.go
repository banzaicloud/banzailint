package testdata

func Expression() {
	var s, t string

	if s+t == "" {
		return
	}

	if len(s+t) == 0 { // want "should not check for empty string by length"
		return
	}

	if 0 == len(s+t) { // want "should not check for empty string by length"
		return
	}

	if len(s+t) == 42 {
		return
	}

	if 42 == len(s+t) {
		return
	}

	if s+t != "" {
		return
	}

	if len(s+t) != 0 { // want "should not check for empty string by length"
		return
	}

	if 0 != len(s+t) { // want "should not check for empty string by length"
		return
	}

	if len(s+t) != 42 {
		return
	}

	if 42 != len(s+t) {
		return
	}

	if 0 < len(s+t) { // want "should not check for empty string by length"
		return
	}

	if len(s+t) < 42 {
		return
	}

	if 42 < len(s+t) {
		return
	}

	if len(s+t) <= 0 { // want "should not check for empty string by length"
		return
	}

	if 0 <= len(s+t) { // want "should not check for empty string by length"
		return
	}

	if len(s+t) <= 42 {
		return
	}

	if 42 <= len(s+t) {
		return
	}

	if len(s+t) > 0 { // want "should not check for empty string by length"
		return
	}

	if len(s+t) > 42 {
		return
	}

	if 42 > len(s+t) {
		return
	}

	if len(s+t) >= 0 { // want "should not check for empty string by length"
		return
	}

	if 0 >= len(s+t) { // want "should not check for empty string by length"
		return
	}

	if len(s+t) >= 42 {
		return
	}

	if 42 >= len(s+t) {
		return
	}
}
