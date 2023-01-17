package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	filename string
	fileType string
	filepath []string
}

func NewDefaultConfig() *Config {
	return &Config{
		filename: "conf",
		fileType: "yaml",
		filepath: []string{"conf"},
	}
}

func (c *Config) LoadFile(addPath string) error {
	viper.SetConfigName(c.filename)
	viper.SetConfigType(c.fileType)
	c.filepath = append(c.filepath, addPath)
	for _, path := range c.filepath {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found err ", err.Error())
			return err
		}
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return nil
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func UnmarshalKey(key string, val interface{}) error {
	return viper.UnmarshalKey(key, val)
}

// Config 是配置组件的配置Struct
//type Config struct {
//	delimiter    string
//	k            *koanf.Koanf
//	tag          string
//	ignoreUnused bool
//}

//const defaultDelimiter = "."
//
//var defaultConfig = NewConfig()
//
//// map[string]interface{} 到 Struct 的映射配置
//func (c *Config) newUnmarshalConf(o interface{}) koanf.UnmarshalConf {
//	return koanf.UnmarshalConf{
//		Tag: c.tag,
//		DecoderConfig: &mapstructure.DecoderConfig{
//			DecodeHook: mapstructure.ComposeDecodeHookFunc(
//				mapstructure.StringToTimeDurationHookFunc(),
//				mapstructure.StringToByteSizeHookFunc(),
//				StringToTimeHookFunc,
//			),
//			Metadata:                 nil,
//			Result:                   o,
//			ZeroFields:               true,
//			WeaklyTypedInput:         false,
//			ErrorUnused:              !c.ignoreUnused,
//			TagName:                  c.tag,
//			FieldNameTransFormMethod: "snake",
//		},
//	}
//}
//
//// unmarshalWithConf 自定义解包配置
//func (c *Config) unmarshal(path string, o interface{}) error {
//	unmarshalConf := c.newUnmarshalConf(o)
//	return c.k.UnmarshalWithConf(path, o, unmarshalConf)
//}
//
//func (c *Config) unmarshalPartially(path string, o interface{}) error {
//	unmarshalConf := c.newUnmarshalConf(o)
//
//	d, err := mapstructure.NewDecoder(unmarshalConf.DecoderConfig)
//	if err != nil {
//		return err
//	}
//
//	rawConfMap := c.k.Get(path)
//	newConfMap := map[string]interface{}{}
//	if confMap, ok := rawConfMap.(map[string]interface{}); ok {
//		for k, v := range confMap {
//			switch v.(type) {
//			case map[string]interface{}:
//			default:
//				newConfMap[k] = v
//			}
//		}
//	}
//	return d.Decode(newConfMap)
//}
//
//func (c *Config) getSubMapKeys(path string) []string {
//	keys := make([]string, 0)
//	rawConfMap := c.k.Get(path)
//	if confMap, ok := rawConfMap.(map[string]interface{}); ok {
//		for k, v := range confMap {
//			switch v.(type) {
//			case map[string]interface{}:
//				keys = append(keys, k)
//			default:
//			}
//		}
//	}
//	return keys
//}
