package main

import "encoding/json"

type Pet struct {
	ID    int     `json:"id,omitempty"`
	Type  string  `json:"type"`
	Price float32 `json:"price"`
}

func JSONStringToPet(jsonString string) Pet {
	pet := Pet{}

	err := json.Unmarshal([]byte(jsonString), &pet)
	if err != nil {
		return Pet{}
	}

	return pet
}

func PetToJSONString(pet Pet) string {
	data, err := json.Marshal(pet)
	if err != nil {
		return ""
	}

	return string(data)
}
