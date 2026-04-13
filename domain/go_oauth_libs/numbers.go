package gooauthlibs

import "strconv"

func GetUintsFromStrings(strings []string) ([]uint, error) {
	var result []uint
	for _, v := range strings {
		userInt, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, uint(userInt))
	}
	return result, nil
}
