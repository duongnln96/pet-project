package config

type otherConfig map[string]interface{}

func (m otherConfig) Get(key string) interface{} {
	return m[key]
}
