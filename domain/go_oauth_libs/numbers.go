package gooauthlibs

import (
	"strconv"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func GetUintsFromStrings(strings []string) ([]uint, error) {
	var result []uint
	for _, v := range strings {
		userInt, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, eliotlibs.ConvertIntToUint(userInt))
	}
	return result, nil
}
