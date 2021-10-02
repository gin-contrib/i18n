package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/known/structpb"
)

// EqualSlice --
func EqualSlice(a interface{}, b interface{}) bool {
	return strings.EqualFold(fmt.Sprintf("%b", a), fmt.Sprintf("%b", b))
}

// StructToMap --
func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap)
	return
}

// StructToPbStruct --
func StructToPbStruct(obj interface{}) (newStruct *structpb.Struct, err error) {
	newMap, err := StructToMap(obj)
	if err != nil {
		return
	}
	newStruct, err = structpb.NewStruct(newMap)
	return
}

// InterfaceToStruct Convert an interface to a specify struct
func InterfaceToStruct(source interface{}, result interface{}) error {
	sourceJson, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(sourceJson, result)
}
