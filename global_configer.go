package config

var (
	configer Configer
)

// InitConfig initialize global config
func InitConfig(adapterName, fileName string) error {
	c, err := NewConfig(adapterName, fileName)
	if err != nil {
		return err
	}

	configer = c
	return nil
}

// String returns the string value for a giving key
func String(key string) string {
	return configer.String(key)
}

// DefaultString returns the string value for a given key.
// if err != nil return defaltval
func DefaultString(key string, defaultVal string) string {
	return configer.DefaultString(key, defaultVal)
}

// Int returns the integer value for a given key.
func Int(key string) (int, error) {
	return configer.Int(key)
}

// DefaultInt returns the integer value for a given key.
// if err != nil return defaltval
func DefaultInt(key string, defaultVal int) int {
	return configer.DefaultInt(key, defaultVal)
}

// Set writes a new value for key.
// if write to one section, the key need be "section::key".
// if the section is not existed, it panics.
func Set(key, value string) error {
	return configer.Set(key, value)
}
