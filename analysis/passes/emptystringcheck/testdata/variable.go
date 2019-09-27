package testdata

func Variable() {
	var s string

	_ = s == ""

	for s == "" {
		return
	}

	if s == "" {
		return
	}

	switch {
	case s == "":
		return
	}

	_ = len(s) == 0 // want "should not check for empty string by length"

	for len(s) == 0 { // want "should not check for empty string by length"
		return
	}

	if len(s) == 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s) == 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 == len(s) // want "should not check for empty string by length"

	for 0 == len(s) { // want "should not check for empty string by length"
		return
	}

	if 0 == len(s) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 == len(s): // want "should not check for empty string by length"
		return
	}

	_ = len(s) == 42

	for len(s) == 42 {
		return
	}

	if len(s) == 42 {
		return
	}

	switch {
	case len(s) == 42:
		return
	}

	_ = 42 == len(s)

	for 42 == len(s) {
		return
	}

	if 42 == len(s) {
		return
	}

	switch {
	case 42 == len(s):
		return
	}

	_ = s != ""

	for s != "" {
		return
	}

	if s != "" {
		return
	}

	switch {
	case s != "":
		return
	}

	_ = len(s) != 0 // want "should not check for empty string by length"

	for len(s) != 0 { // want "should not check for empty string by length"
		return
	}

	if len(s) != 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s) != 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 != len(s) // want "should not check for empty string by length"

	for 0 != len(s) { // want "should not check for empty string by length"
		return
	}

	if 0 != len(s) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 != len(s): // want "should not check for empty string by length"
		return
	}

	_ = len(s) != 42

	for len(s) != 42 {
		return
	}

	if len(s) != 42 {
		return
	}

	switch {
	case len(s) != 42:
		return
	}

	_ = 42 != len(s)

	for 42 != len(s) {
		return
	}

	if 42 != len(s) {
		return
	}

	switch {
	case 42 != len(s):
		return
	}

	_ = 0 < len(s) // want "should not check for empty string by length"

	for 0 < len(s) { // want "should not check for empty string by length"
		return
	}

	if 0 < len(s) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 < len(s): // want "should not check for empty string by length"
		return
	}

	_ = len(s) < 42

	for len(s) < 42 {
		return
	}

	if len(s) < 42 {
		return
	}

	switch {
	case len(s) < 42:
		return
	}

	_ = 42 < len(s)

	for 42 < len(s) {
		return
	}

	if 42 < len(s) {
		return
	}

	switch {
	case 42 < len(s):
		return
	}

	_ = len(s) <= 0 // want "should not check for empty string by length"

	for len(s) <= 0 { // want "should not check for empty string by length"
		return
	}

	if len(s) <= 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s) <= 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 <= len(s) // want "should not check for empty string by length"

	for 0 <= len(s) { // want "should not check for empty string by length"
		return
	}

	if 0 <= len(s) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 <= len(s): // want "should not check for empty string by length"
		return
	}

	_ = len(s) <= 42

	for len(s) <= 42 {
		return
	}

	if len(s) <= 42 {
		return
	}

	switch {
	case len(s) <= 42:
		return
	}

	_ = 42 <= len(s)

	for 42 <= len(s) {
		return
	}

	if 42 <= len(s) {
		return
	}

	switch {
	case 42 <= len(s):
		return
	}

	_ = len(s) > 0 // want "should not check for empty string by length"

	for len(s) > 0 { // want "should not check for empty string by length"
		return
	}

	if len(s) > 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s) > 0: // want "should not check for empty string by length"
		return
	}

	_ = len(s) > 42

	for len(s) > 42 {
		return
	}

	if len(s) > 42 {
		return
	}

	switch {
	case len(s) > 42:
		return
	}

	_ = 42 > len(s)

	for 42 > len(s) {
		return
	}

	if 42 > len(s) {
		return
	}

	switch {
	case 42 > len(s):
		return
	}

	_ = len(s) >= 0 // want "should not check for empty string by length"

	for len(s) >= 0 { // want "should not check for empty string by length"
		return
	}

	if len(s) >= 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(s) >= 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 >= len(s) // want "should not check for empty string by length"

	for 0 >= len(s) { // want "should not check for empty string by length"
		return
	}

	if 0 >= len(s) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 >= len(s): // want "should not check for empty string by length"
		return
	}

	_ = len(s) >= 42

	for len(s) >= 42 {
		return
	}

	if len(s) >= 42 {
		return
	}

	switch {
	case len(s) >= 42:
		return
	}

	_ = 42 >= len(s)

	for 42 >= len(s) {
		return
	}

	if 42 >= len(s) {
		return
	}

	switch {
	case 42 >= len(s):
		return
	}
}
