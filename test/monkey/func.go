package monkey

import "fmt"

func MyFunc(uid int64) string {
	u, err := GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// some code...

	return fmt.Sprintf("hello %s\n", u.Name)
}
