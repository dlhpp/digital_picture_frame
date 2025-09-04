package yaml

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/utils"
	"gopkg.in/yaml.v3"
)

func GetString(obj interface{}, path string, defaultValue string) string {
	value := Get(obj, path, defaultValue)
	str, ok := value.(string)
	if ok && len(strings.TrimSpace(str)) > 0 {
		return str
	}
	return defaultValue
}

func GetStringArray(obj interface{}, path string, defaultValue []string) []string {
	value := Get(obj, path, defaultValue)
	// logging.Log("GetStringArray", 9, utils.DescribeVariable("value", value))
	// strArray, ok := value.([]string)
	// if ok {
	// 	return strArray
	// }
	// return defaultValue
	val, ok := value.([]interface{})
	if !ok {
		return defaultValue
	}
	return utils.ConvertToSliceStrings(val)
}

func GetInt(obj interface{}, path string, defaultValue int) int {
	value := Get(obj, path, defaultValue)
	val, ok := value.(int)
	if ok {
		return val
	}
	return defaultValue
}

func GetIntArray(obj interface{}, path string, defaultValue []int) []int {
	value := Get(obj, path, defaultValue)
	// logging.Log("GetIntArray", 1, utils.DescribeVariable("value", value))
	// val, ok := value.([]int)
	// logging.Log("GetIntArray", 1, "ok", ok, utils.DescribeVariable("val", val))
	// if ok {
	// 	return val
	// }
	// return defaultValue

	val, ok := value.([]interface{})
	if !ok {
		return defaultValue
	}
	return utils.ConvertToSliceInt(val)
}

func GetFloat(obj interface{}, path string, defaultValue float64) float64 {
	value := Get(obj, path, defaultValue)
	val, ok := value.(float64)
	if ok {
		return val
	}
	return defaultValue
}

func GetFloatArray(obj interface{}, path string, defaultValue []float64) []float64 {
	value := Get(obj, path, defaultValue)
	// val, ok := value.([]float64)
	// if ok {
	// 	return val
	// }
	// return defaultValue
	val, ok := value.([]interface{})
	if !ok {
		return defaultValue
	}
	return utils.ConvertToSliceFloat(val)
}

func Get(obj interface{}, path string, defaultValue interface{}) interface{} {
	if obj == nil || path == "" {
		return defaultValue
	}

	// Split the path by dots
	keys := strings.Split(path, ".")
	current := obj

	for _, key := range keys {
		current = getValueByKey(current, key)
		if current == nil {
			return defaultValue
		}
	}

	return current
}

// getValueByKey retrieves a value from the current object using a single key
func getValueByKey(obj interface{}, key string) interface{} {
	if obj == nil {
		return nil
	}

	v := reflect.ValueOf(obj)

	// Handle pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Map:
		// Try to get the value from the map
		mapValue := v.MapIndex(reflect.ValueOf(key))
		if !mapValue.IsValid() {
			return nil
		}
		return mapValue.Interface()

	case reflect.Slice, reflect.Array:
		// Convert key to integer for slice/array access
		index, err := strconv.Atoi(key)
		if err != nil || index < 0 || index >= v.Len() {
			return nil
		}
		return v.Index(index).Interface()

	case reflect.Struct:
		// Try to get field by name (case-sensitive)
		field := v.FieldByName(key)
		if !field.IsValid() {
			return nil
		}
		return field.Interface()

	default:
		return nil
	}
}

func OpenYamlFile(filePath string) *map[string]any {

	logging.Log("OpenYamlFile", 5, fmt.Sprintf("file: %s", filePath))

	yamlFile, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("OpenYamlFile: error reading YAML file: %w", err))
	}
	defer yamlFile.Close()

	var cfg map[string]interface{}

	decoder := yaml.NewDecoder(yamlFile)
	if err = decoder.Decode(&cfg); err != nil {
		panic(fmt.Errorf("OpenYamlFile: error unmarshaling YAML: %w", err))
	}

	return &cfg
}
