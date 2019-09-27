package testdata

func Parameter(s string) {
	if s == "" {
		return
	}

	if len(s) == 0 { // want "should not check for empty string by length"
		return
	}

	if 0 == len(s) { // want "should not check for empty string by length"
		return
	}

	if len(s) == 42 {
		return
	}

	if 42 == len(s) {
		return
	}

	if s != "" {
		return
	}

	if len(s) != 0 { // want "should not check for empty string by length"
		return
	}

	if 0 != len(s) { // want "should not check for empty string by length"
		return
	}

	if len(s) != 42 {
		return
	}

	if 42 != len(s) {
		return
	}

	if 0 < len(s) { // want "should not check for empty string by length"
		return
	}

	if len(s) < 42 {
		return
	}

	if 42 < len(s) {
		return
	}

	if len(s) <= 0 { // want "should not check for empty string by length"
		return
	}

	if 0 <= len(s) { // want "should not check for empty string by length"
		return
	}

	if len(s) <= 42 {
		return
	}

	if 42 <= len(s) {
		return
	}

	if len(s) > 0 { // want "should not check for empty string by length"
		return
	}

	if len(s) > 42 {
		return
	}

	if 42 > len(s) {
		return
	}

	if len(s) >= 0 { // want "should not check for empty string by length"
		return
	}

	if 0 >= len(s) { // want "should not check for empty string by length"
		return
	}

	if len(s) >= 42 {
		return
	}

	if 42 >= len(s) {
		return
	}
}
