package code

// GenDiff returns a textual representation of the difference between
// two configuration files in the requested format.
func GenDiff(filepath1, filepath2, format string) (string, error) {
	if _, err := Parse(filepath1); err != nil {
		return "", err
	}
	if _, err := Parse(filepath2); err != nil {
		return "", err
	}
	return "", nil
}
