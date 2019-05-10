package mapdata

import (
	"fmt"
)

//GetStringFromMap get string from map
func GetStringFromMap(value map[string]interface{}, key string) (string, error) {
	tmp, ok := value[key]
	if !ok {
		return "", fmt.Errorf("no key %v in %v", key, value)
	}
	tmpstring, ok := tmp.(string)
	if ok {
		return tmpstring, nil
	}
	return "", fmt.Errorf("%v is not string type", tmp)
}

//GetFloat64FromMap get float64 from map
func GetFloat64FromMap(value map[string]interface{}, key string) (float64, error) {
	tmp, ok := value[key]
	if !ok {
		return 0.0, fmt.Errorf("no key %v in %v", key, value)
	}
	tmpfloat64, ok := tmp.(float64)
	if ok {
		return tmpfloat64, nil
	}
	return 0.0, fmt.Errorf("%v is not float64 type", tmp)
}

//GetInt64FromMap get int64 from map
func GetInt64FromMap(value map[string]interface{}, key string) (int64, error) {
	tmp, ok := value[key]
	if !ok {
		return 0, fmt.Errorf("no key %v in %v", key, value)
	}
	tmpint64, ok := tmp.(int64)
	if ok {
		return tmpint64, nil
	}
	return 0, fmt.Errorf("%v is not int64 type", tmp)
}

//GetIntFromMap get int from map
func GetIntFromMap(value map[string]interface{}, key string) (int, error) {
	tmp, ok := value[key]
	if !ok {
		return 0, fmt.Errorf("no key %v in %v", key, value)
	}
	tmpint, ok := tmp.(int)
	if ok {
		return tmpint, nil
	}
	return 0, fmt.Errorf("%v is not int type", tmp)
}

//GetUint32FromMap get uint32 from map
func GetUint32FromMap(value map[string]interface{}, key string) (uint32, error) {
	tmp, ok := value[key]
	if !ok {
		return 0, fmt.Errorf("no key %v in %v", key, value)
	}
	tmpuint32, ok := tmp.(uint32)
	if ok {
		return tmpuint32, nil
	}
	return 0, fmt.Errorf("%v is not uint32 type", tmp)
}

//GetUint64FromMap get uint64 from map
func GetUint64FromMap(value map[string]interface{}, key string) (uint64, error) {
	tmp, ok := value[key]
	if !ok {
		return 0, fmt.Errorf("no key %v in %v", key, value)
	}
	tmpuint64, ok := tmp.(uint64)
	if ok {
		return tmpuint64, nil
	}
	return 0, fmt.Errorf("%v is not uint64 type", tmp)
}

//GetMapFromMap get map[string]interface{} from map
func GetMapFromMap(value map[string]interface{}, key string) (map[string]interface{}, error) {
	tmp, ok := value[key]
	if !ok {
		return map[string]interface{}{}, fmt.Errorf("no key %v in %v", key, value)
	}
	tmpmap, ok := tmp.(map[string]interface{})
	if ok {
		return tmpmap, nil
	}
	return map[string]interface{}{}, fmt.Errorf("%v is not map[string]interface{} type", tmp)
}

//GetSliceInterfaceFromMap get []interface{} from map
func GetSliceInterfaceFromMap(value map[string]interface{}, key string) ([]interface{}, error) {
	tmp, ok := value[key]
	if !ok {
		return []interface{}{}, fmt.Errorf("no key %v in %v", key, value)
	}
	tmpslice, ok := tmp.([]interface{})
	if ok {
		return tmpslice, nil
	}
	return []interface{}{}, fmt.Errorf("%v is not []interface{} type", tmp)
}
