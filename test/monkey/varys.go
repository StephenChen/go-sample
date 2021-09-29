package monkey

type UserInfo struct {
	Name string
}

func GetInfoByUID(uid int64) (*UserInfo, error) {
	return &UserInfo{Name: "default"}, nil
}
