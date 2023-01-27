package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func GetParamFromJson(body []byte, path string) (string, error) {
	var x map[string]interface{}
	json.Unmarshal(body, &x)
	childLevel := strings.Count(path, ".")
	children := strings.Split(path, ".")
	ok := true
	for i := 0; i < childLevel; i++ {
		x, ok = x[children[i]].(map[string]interface{})
	}

	if ok {
		val := x[children[childLevel]]
		return fmt.Sprintf("%v", val), nil
	}

	return "", errors.New("can not retrieve the parameter")
}
