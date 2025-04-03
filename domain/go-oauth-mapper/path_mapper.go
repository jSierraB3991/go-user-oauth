package gooauthmapper

import "regexp"

func ConvertPathToRegex(path string) string {
	// Expresión regular para encontrar el número al final de la URL
	re := regexp.MustCompile(`/\d+$`)
	// Reemplazamos el número con `/.+[0-9]+`
	return re.ReplaceAllString(path, `/.*[0-9]+`)
}
