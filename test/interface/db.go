package _interface

// DB 数据接口
type DB interface {
	Get(key string) (int, error)
	Add(key string, value int) error
}

// GetFromDB 根据 key 从 DB 查询数据
func GetFromDB(db DB, key string) int {
	v, err := db.Get(key)
	if err != nil {
		return -1
	}
	return v
}

// mockgen -source='db.go' -destination='mocks/db_mock.go' -package=mocks
