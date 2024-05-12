package support

import (
	"fmt"
	"net/http"

	"Adlet/core/backend/container"
)

/*
Members: If there is no value in given from front, we will initialize it with -1 by default
*/
func CheckValidation(writer http.ResponseWriter, PassedValue string) bool {
	var Mapping map[string]map[string][]int = container.MapAll
	var CuttedValue, CuttedKey string
	for i := len(PassedValue) - 1; i >= 0; i-- {
		if PassedValue[i] == '-' {
			CuttedValue = PassedValue[:i]
			CuttedKey = PassedValue[i+1:]
			break
		}
	}
	fmt.Printf("The cutted string key is: %s\n The cutted string value is %s\n ", CuttedKey, CuttedValue)
	var ok bool
	var ID []int
	var flag bool = false
	_, ok = Mapping[CuttedKey]
	if ok {
		ID = Mapping[CuttedKey][CuttedValue]
		flag = true
	} else {
		for key := range Mapping {
			value, ok := Mapping[key][PassedValue]
			if ok {
				flag = true
				ID = value
				break
			}

		}
	}
	if !flag {
		fmt.Println("There is such data bitch")
		return false
	} else {
		container.Array_Available = ID
		fmt.Println("ID:", ID)
		return true
	}
}
