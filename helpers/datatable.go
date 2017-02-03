package helpers

import (
	"reflect"
	"app/passporte/models"
)


func displaytable(map[string]User result)([]models.User)  {

	dataValue := reflect.ValueOf(result)
	var valueSlice []models.User
	for _, key := range dataValue.MapKeys() {
	valueSlice = append(valueSlice, result[key.String()])
	return valueSlice
}