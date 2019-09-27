package testdata

func Expression() {
	var s, t string

	_ = s+t == ""

	for s+t == "" {
		return
	}

	if s+t == "" {
		return
	}

	switch {
	case s+t == "":
		return
	}

	_ = len(s+t) == 0 // want "should not check for empty string by length"

	for len(s+t) == 0 { // want "should not check for empty string by length"
		return
	}

	if len(s+t) == 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s+t) == 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 == len(s+t) // want "should not check for empty string by length"

	for 0 == len(s+t) { // want "should not check for empty string by length"
		return
	}

	if 0 == len(s+t) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 == len(s+t): // want "should not check for empty string by length"
		return
	}

	_ = len(s+t) == 42

	for len(s+t) == 42 {
		return
	}

	if len(s+t) == 42 {
		return
	}

	switch {
	case len(s+t) == 42:
		return
	}

	_ = 42 == len(s+t)

	for 42 == len(s+t) {
		return
	}

	if 42 == len(s+t) {
		return
	}

	switch {
	case 42 == len(s+t):
		return
	}

	_ = s+t != ""

	for s+t != "" {
		return
	}

	if s+t != "" {
		return
	}

	switch {
	case s+t != "":
		return
	}

	_ = len(s+t) != 0 // want "should not check for empty string by length"

	for len(s+t) != 0 { // want "should not check for empty string by length"
		return
	}

	if len(s+t) != 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s+t) != 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 != len(s+t) // want "should not check for empty string by length"

	for 0 != len(s+t) { // want "should not check for empty string by length"
		return
	}

	if 0 != len(s+t) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 != len(s+t): // want "should not check for empty string by length"
		return
	}

	_ = len(s+t) != 42

	for len(s+t) != 42 {
		return
	}

	if len(s+t) != 42 {
		return
	}

	switch {
	case len(s+t) != 42:
		return
	}

	_ = 42 != len(s+t)

	for 42 != len(s+t) {
		return
	}

	if 42 != len(s+t) {
		return
	}

	switch {
	case 42 != len(s+t):
		return
	}

	_ = 0 < len(s+t) // want "should not check for empty string by length"

	for 0 < len(s+t) { // want "should not check for empty string by length"
		return
	}

	if 0 < len(s+t) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 < len(s+t): // want "should not check for empty string by length"
		return
	}

	_ = len(s+t) < 42

	for len(s+t) < 42 {
		return
	}

	if len(s+t) < 42 {
		return
	}

	switch {
	case len(s+t) < 42:
		return
	}

	_ = 42 < len(s+t)

	for 42 < len(s+t) {
		return
	}

	if 42 < len(s+t) {
		return
	}

	switch {
	case 42 < len(s+t):
		return
	}

	_ = len(s+t) <= 0 // want "should not check for empty string by length"

	for len(s+t) <= 0 { // want "should not check for empty string by length"
		return
	}

	if len(s+t) <= 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s+t) <= 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 <= len(s+t) // want "should not check for empty string by length"

	for 0 <= len(s+t) { // want "should not check for empty string by length"
		return
	}

	if 0 <= len(s+t) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 <= len(s+t): // want "should not check for empty string by length"
		return
	}

	_ = len(s+t) <= 42

	for len(s+t) <= 42 {
		return
	}

	if len(s+t) <= 42 {
		return
	}

	switch {
	case len(s+t) <= 42:
		return
	}

	_ = 42 <= len(s+t)

	for 42 <= len(s+t) {
		return
	}

	if 42 <= len(s+t) {
		return
	}

	switch {
	case 42 <= len(s+t):
		return
	}

	_ = len(s+t) > 0 // want "should not check for empty string by length"

	for len(s+t) > 0 { // want "should not check for empty string by length"
		return
	}

	if len(s+t) > 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s+t) > 0: // want "should not check for empty string by length"
		return
	}

	_ = len(s+t) > 42

	for len(s+t) > 42 {
		return
	}

	if len(s+t) > 42 {
		return
	}

	switch {
	case len(s+t) > 42:
		return
	}

	_ = 42 > len(s+t)

	for 42 > len(s+t) {
		return
	}

	if 42 > len(s+t) {
		return
	}

	switch {
	case 42 > len(s+t):
		return
	}

	_ = len(s+t) >= 0 // want "should not check for empty string by length"

	for len(s+t) >= 0 { // want "should not check for empty string by length"
		return
	}

	if len(s+t) >= 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s+t) >= 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 >= len(s+t) // want "should not check for empty string by length"

	for 0 >= len(s+t) { // want "should not check for empty string by length"
		return
	}

	if 0 >= len(s+t) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 >= len(s+t): // want "should not check for empty string by length"
		return
	}

	_ = len(s+t) >= 42

	for len(s+t) >= 42 {
		return
	}

	if len(s+t) >= 42 {
		return
	}

	switch {
	case len(s+t) >= 42:
		return
	}

	_ = 42 >= len(s+t)

	for 42 >= len(s+t) {
		return
	}

	if 42 >= len(s+t) {
		return
	}

	switch {
	case 42 >= len(s+t):
		return
	}
}
