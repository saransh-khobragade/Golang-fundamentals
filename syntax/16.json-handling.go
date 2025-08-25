package main

import (
	"encoding/json"
	"fmt"
)

// JSON handling
func main() {
	x := `{ "name":"John", "age":30, "city":"New York"}`

	// Parse JSON to map
	var y map[string]interface{}
	json.Unmarshal([]byte(x), &y)
	fmt.Println(y["age"]) // 30

	// Parse map to JSON
	z, _ := json.Marshal(y)
	fmt.Println(string(z)) // {"name":"John","age":30,"city":"New York"}

	// Reading JSON from file
	// data, err := ioutil.ReadFile("go/syntax/data.json")
	// if err != nil {
	//     fmt.Println("Error reading file:", err)
	//     return
	// }
	// var fileData map[string]interface{}
	// json.Unmarshal(data, &fileData)
	// fmt.Println(fileData)

	// Writing JSON to file
	// outputData, _ := json.Marshal(fileData)
	// err = ioutil.WriteFile("go/syntax/data.json", outputData, 0644)
	// if err != nil {
	//     fmt.Println("Error writing file:", err)
	// }

	// Pretty print JSON
	// prettyJSON, _ := json.MarshalIndent(fileData, "", "    ")
	// fmt.Println(string(prettyJSON))
} 