package main

import (
	"fmt"
)

func get(dict map[string]string, key string) {
	fmt.Println(dict[key])
}

func add(dict map[string]string, key string, value string) {
	dict[key] = value
}

func remove(dict map[string]string, key string) {
	delete(dict, key)
}

func update(dict map[string]string, key string, new_value string) {
	dict[key] = new_value
}

func list(dict map[string]string) {
	for key, value := range dict {
		fmt.Println(key, value)
	}
}



func main() {
    dict := make(map[string]string)
	add(dict, "a", "1")
	add(dict, "b", "2")
	add(dict, "c", "3")
	list(dict)
	get(dict, "a")
	remove(dict, "a")
	update(dict, "b", "4")
	list(dict)
}