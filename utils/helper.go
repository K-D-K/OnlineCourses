package utils

import (
	"OnlineCourses/interfaces"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

// GetRequestBody as JSON Array
func GetRequestBody(reader io.ReadCloser) []map[string]interface{} {
	decoder := json.NewDecoder(reader)
	decoder.Token()
	var body []map[string]interface{}
	for decoder.More() {
		var data map[string]interface{}
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		body = append(body, data)
	}
	return body
}

// ExtractKey from array of object
func ExtractKey(array []map[string]interface{}, key string) []interface{} {
	var result []interface{}
	for _, data := range array {
		result = append(result, data[key])
	}
	return result
}

// ConvertToUintArray from normal array
func ConvertToUintArray(array []interface{}) []uint64 {
	result := make([]uint64, len(array))
	for index, element := range array {
		result[index], _ = strconv.ParseUint(fmt.Sprintf("%v", element), 10, 64)
	}
	return result
}

// ConstructPkVsDataMap from given data
func ConstructPkVsDataMap(array []map[string]interface{}) map[interface{}]interface{} {
	pkVsDataMap := make(map[interface{}]interface{})
	for _, data := range array {
		pkVsDataMap[data["id"]] = data
	}
	return pkVsDataMap
}

// ConvertToJSONMap converts struct into map
func ConvertToJSONMap(data interface{}) []map[string]interface{} {
	byteArr, _ := json.Marshal(data)
	var tempMap []map[string]interface{}
	json.Unmarshal(byteArr, &tempMap)
	return tempMap
}

// ConvertMapIntoStruct .
func ConvertMapIntoStruct(data []map[string]interface{}, dest interface{}) {
	byteArr, _ := json.Marshal(data)
	json.Unmarshal(byteArr, &dest)
}

// GetPKIDs for entities
func GetPKIDs(entities []interfaces.Entity) []uint64 {
	pkIDs := []uint64{}
	for _, entity := range entities {
		if entity.GetPKID() != nil {
			pkIDs = append(pkIDs, *entity.GetPKID())
		}
	}
	return pkIDs
}

// GetPKIDVsEntityMap .
func GetPKIDVsEntityMap(entities []interfaces.Entity) map[uint64]interfaces.Entity {
	pkIDVsEntityMap := make(map[uint64]interfaces.Entity)
	for _, entity := range entities {
		if entity.GetPKID() != nil {
			pkIDVsEntityMap[*entity.GetPKID()] = entity
		}
	}
	return pkIDVsEntityMap
}
