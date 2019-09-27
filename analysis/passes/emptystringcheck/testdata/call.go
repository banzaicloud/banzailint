package testdata

func callMe() string {
	return ""
}

func Call() {
	if callMe() == "" {
		return
	}

	if len(callMe()) == 0 { // want "should not check for empty string by length"
		return
	}

	if 0 == len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) == 42 {
		return
	}

	if 42 == len(callMe()) {
		return
	}

	if callMe() != "" {
		return
	}

	if len(callMe()) != 0 { // want "should not check for empty string by length"
		return
	}

	if 0 != len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) != 42 {
		return
	}

	if 42 != len(callMe()) {
		return
	}

	if 0 < len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) < 42 {
		return
	}

	if 42 < len(callMe()) {
		return
	}

	if len(callMe()) <= 0 { // want "should not check for empty string by length"
		return
	}

	if 0 <= len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) <= 42 {
		return
	}

	if 42 <= len(callMe()) {
		return
	}

	if len(callMe()) > 0 { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) > 42 {
		return
	}

	if 42 > len(callMe()) {
		return
	}

	if len(callMe()) >= 0 { // want "should not check for empty string by length"
		return
	}

	if 0 >= len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) >= 42 {
		return
	}

	if 42 >= len(callMe()) {
		return
	}
}
