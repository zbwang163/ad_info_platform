package convert

import "strconv"

func StringToInt64(in string) (int64, error) {
	tem, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, err
	}
	return tem, nil
}
