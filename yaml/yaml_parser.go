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

/* =====================================================================================

DLH 2025-09-04

This file is a golang implementation of Lodash's
awesome _.get() function: https://lodash.com/docs/4.17.15#get
to generically enable loading and using YAML config files with zero setup.

With this approach, for any new project or yaml file I can immediately start
using it without have to setup structs etc.

Another huge advantage is that we don't need to know the structure of the
YAML file in advance.  The paths to get the desired values could be read
dynamically from a database or environment variables or command line flags.

In addition, this same file could be used on generic JSON structures as well.

===================================================================================== */

func GetString(obj any, path string, defaultValue string) string {
	value := Get(obj, path, defaultValue)
	str, ok := value.(string)

	if !ok {
		logging.Log("GetString", 99, "NOT OK", "path", path, "value", value, utils.DescribeVariable("value", value))
		return defaultValue
	}

	if len(strings.TrimSpace(str)) < 1 {
		logging.Log("GetString", 99, "EMPTY RESULT STRING", "path", path, "value", value, utils.DescribeVariable("value", value))
		return defaultValue
	}

	return str
}

func GetStringArray(obj any, path string, defaultValue []string) []string {
	value := Get(obj, path, defaultValue)
	val, ok := value.([]any)
	if !ok {
		logging.Log("GetStringArray", 99, "NOT OK", "path", path, "value", value, utils.DescribeVariable("value", value))
		return defaultValue
	}
	result := utils.ConvertToSliceStrings(val)
	return result
}

func GetInt(obj any, path string, defaultValue int) int {
	value := Get(obj, path, defaultValue)
	val, ok := value.(int)
	if !ok {
		logging.Log("GetInt", 99, "NOT OK", "path", path, "value", value, utils.DescribeVariable("value", value))
		return defaultValue
	}
	return val
}

func GetIntArray(obj any, path string, defaultValue []int) []int {
	value := Get(obj, path, defaultValue)
	val, ok := value.([]any)
	if !ok {
		logging.Log("GetIntArray", 99, "NOT OK", "path", path, "value", value, utils.DescribeVariable("value", value))
		return defaultValue
	}
	return utils.ConvertToSliceInt(val)
}

func GetFloat(obj any, path string, defaultValue float64) float64 {
	value := Get(obj, path, defaultValue)
	val, ok := value.(float64)
	if !ok {
		logging.Log("GetFloat", 99, "NOT OK", "path", path, "value", value, utils.DescribeVariable("value", value))
		return defaultValue
	}
	return val
}

func GetFloatArray(obj any, path string, defaultValue []float64) []float64 {
	value := Get(obj, path, defaultValue)
	val, ok := value.([]any)
	if !ok {
		logging.Log("GetFloatArray", 99, "NOT OK", "path", path, "value", value, utils.DescribeVariable("value", value))
		return defaultValue
	}
	return utils.ConvertToSliceFloat(val)
}

// This implements Lodash's awesome _.get() function: https://lodash.com/docs/4.17.15#get
func Get(obj any, path string, defaultValue any) any {
	if obj == nil || path == "" {
		logging.Log("Get", 99, "INPUT IS EMPTY", "path", path, utils.DescribeVariable("obj", obj))
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
func getValueByKey(obj any, key string) any {
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
		// What does .Interface() do here?
		// Converts back to a standard Go type:
		// Reflection allows you to inspect types and values generically.
		// Interface() reverses this process, returning the actual
		// underlying Go value (e.g., a string, int, or custom struct)
		// that the reflect.Value is holding.
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

	logging.Log("OpenYamlFile", 9, fmt.Sprintf("file: %s", filePath))

	yamlFile, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("OpenYamlFile: error reading YAML file: %w", err))
	}
	defer yamlFile.Close()

	var cfg map[string]any

	decoder := yaml.NewDecoder(yamlFile)
	if err = decoder.Decode(&cfg); err != nil {
		panic(fmt.Errorf("OpenYamlFile: error unmarshaling YAML: %w", err))
	}

	logging.Log("OpenYamlFile", 2, utils.DescribeVariable("cfg", cfg))

	return &cfg
}
