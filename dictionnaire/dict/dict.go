package dict
	
import (
	"encoding/json"
	"fmt"
	"os"
)

type dict struct {
	pathFile string
    m map[string]string
}


func New(path string) dict {
	return dict{pathFile: path,m: make(map[string]string)}
}

func (d dict) List() {
	for key, value := range d.m {
		fmt.Println(key, value)
	}
}

func (d dict) Get(key string) {
	fmt.Println(d.m[key])
}

func (d dict) Add(key string, value string) {
	d.m[key] = value
}

func (d dict) Remove(key string) {
	delete(d.m, key)
}

func (d dict) Update(key string, new_value string) {
	d.m[key] = new_value
}

func (d dict) SaveToFile() {
	// Marshal the slice into JSON format
	jsonData, err := json.MarshalIndent(d.m, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Open a file for writing (create if not exists, truncate if exists)
	file, err := os.Create(d.pathFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to dict.json")
}

func (d dict) LoadFromFile() {
	// Open the JSON file for reading
	file, err := os.Open(d.pathFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the JSON data from the file into a slice of Person
	var dicts map[string]string 
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dicts)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the read data
	for key, value := range dicts {
		fmt.Printf("%s: %s\n", key, value)
	}
}