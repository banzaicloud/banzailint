package testdata

func callMe() string {
	return ""
}

func Call() {
	_ = callMe() == ""

	for callMe() == "" {
		return
	}

	if callMe() == "" {
		return
	}

	switch {
	case callMe() == "":
		return
	}

	_ = len(callMe()) == 0 // want "should not check for empty string by length"

	for len(callMe()) == 0 { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) == 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(callMe()) == 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 == len(callMe()) // want "should not check for empty string by length"

	for 0 == len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if 0 == len(callMe()) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 == len(callMe()): // want "should not check for empty string by length"
		return
	}

	_ = len(callMe()) == 42

	for len(callMe()) == 42 {
		return
	}

	if len(callMe()) == 42 {
		return
	}

	switch {
	case len(callMe()) == 42:
		return
	}

	_ = 42 == len(callMe())

	for 42 == len(callMe()) {
		return
	}

	if 42 == len(callMe()) {
		return
	}

	switch {
	case 42 == len(callMe()):
		return
	}

	_ = callMe() != ""

	for callMe() != "" {
		return
	}

	if callMe() != "" {
		return
	}

	switch {
	case callMe() != "":
		return
	}

	_ = len(callMe()) != 0 // want "should not check for empty string by length"

	for len(callMe()) != 0 { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) != 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(callMe()) != 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 != len(callMe()) // want "should not check for empty string by length"

	for 0 != len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if 0 != len(callMe()) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 != len(callMe()): // want "should not check for empty string by length"
		return
	}

	_ = len(callMe()) != 42

	for len(callMe()) != 42 {
		return
	}

	if len(callMe()) != 42 {
		return
	}

	switch {
	case len(callMe()) != 42:
		return
	}

	_ = 42 != len(callMe())

	for 42 != len(callMe()) {
		return
	}

	if 42 != len(callMe()) {
		return
	}

	switch {
	case 42 != len(callMe()):
		return
	}

	_ = 0 < len(callMe()) // want "should not check for empty string by length"

	for 0 < len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if 0 < len(callMe()) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 < len(callMe()): // want "should not check for empty string by length"
		return
	}

	_ = len(callMe()) < 42

	for len(callMe()) < 42 {
		return
	}

	if len(callMe()) < 42 {
		return
	}

	switch {
	case len(callMe()) < 42:
		return
	}

	_ = 42 < len(callMe())

	for 42 < len(callMe()) {
		return
	}

	if 42 < len(callMe()) {
		return
	}

	switch {
	case 42 < len(callMe()):
		return
	}

	_ = len(callMe()) <= 0 // want "should not check for empty string by length"

	for len(callMe()) <= 0 { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) <= 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(callMe()) <= 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 <= len(callMe()) // want "should not check for empty string by length"

	for 0 <= len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if 0 <= len(callMe()) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 <= len(callMe()): // want "should not check for empty string by length"
		return
	}

	_ = len(callMe()) <= 42

	for len(callMe()) <= 42 {
		return
	}

	if len(callMe()) <= 42 {
		return
	}

	switch {
	case len(callMe()) <= 42:
		return
	}

	_ = 42 <= len(callMe())

	for 42 <= len(callMe()) {
		return
	}

	if 42 <= len(callMe()) {
		return
	}

	switch {
	case 42 <= len(callMe()):
		return
	}

	_ = len(callMe()) > 0 // want "should not check for empty string by length"

	for len(callMe()) > 0 { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) > 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(callMe()) > 0: // want "should not check for empty string by length"
		return
	}

	_ = len(callMe()) > 42

	for len(callMe()) > 42 {
		return
	}

	if len(callMe()) > 42 {
		return
	}

	switch {
	case len(callMe()) > 42:
		return
	}

	_ = 42 > len(callMe())

	for 42 > len(callMe()) {
		return
	}

	if 42 > len(callMe()) {
		return
	}

	switch {
	case 42 > len(callMe()):
		return
	}

	_ = len(callMe()) >= 0 // want "should not check for empty string by length"

	for len(callMe()) >= 0 { // want "should not check for empty string by length"
		return
	}

	if len(callMe()) >= 0 { // want "should not check for empty string by length"
		return
	}

	switch {
	case len(callMe()) >= 0: // want "should not check for empty string by length"
		return
	}

	_ = 0 >= len(callMe()) // want "should not check for empty string by length"

	for 0 >= len(callMe()) { // want "should not check for empty string by length"
		return
	}

	if 0 >= len(callMe()) { // want "should not check for empty string by length"
		return
	}

	switch {
	case 0 >= len(callMe()): // want "should not check for empty string by length"
		return
	}

	_ = len(callMe()) >= 42

	for len(callMe()) >= 42 {
		return
	}

	if len(callMe()) >= 42 {
		return
	}

	switch {
	case len(callMe()) >= 42:
		return
	}

	_ = 42 >= len(callMe())

	for 42 >= len(callMe()) {
		return
	}

	if 42 >= len(callMe()) {
		return
	}

	switch {
	case 42 >= len(callMe()):
		return
	}
}
