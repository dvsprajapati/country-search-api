package cache

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, Value interface{})
}
