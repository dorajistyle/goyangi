package interfaceHelper

import (
    "errors"
    "strconv"
)

// GetInt64 recognize a type of interface and return it's value
func GetInt64(unknown interface{}) (int64, error) {
    switch i := unknown.(type) {
    case float32:
        return int64(i), nil
    case float64:
        return int64(i), nil
    case int:
        return int64(i), nil
    case int8:
        return int64(i), nil
    case int16:
        return int64(i), nil
    case int32:
        return int64(i), nil
    case int64:
        return i, nil
    case uint:
        return int64(i), nil
    case uint8:
        return int64(i), nil
    case uint16:
        return int64(i), nil
    case uint32:
        return int64(i), nil
    case uint64:
        return int64(i), nil
    case string:
         v, err := strconv.ParseInt(i, 10, 64)
         return v, err
    // ...other cases...
    default:
        return -1, errors.New("get int64: unknown value is of incompatible type")
    }
}