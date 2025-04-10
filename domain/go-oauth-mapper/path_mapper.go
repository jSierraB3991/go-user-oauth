package gooauthmapper

import "regexp"

func ConvertPathToRegex(path string) string {
	re := regexp.MustCompile(`/\d+$`)
	return re.ReplaceAllString(path, `/.*[0-9]+`)
}
