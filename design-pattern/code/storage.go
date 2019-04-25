// 基类接口
type Storage interface {
	Get(id int) string
	MultiGet(id []int) map[string]interface{}
}

// 子类接口
type StorageList interface {
	Storage
	List(id []int) []interface{}
}

