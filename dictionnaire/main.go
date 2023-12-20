package main

import (
	"dictionnaire/dict"
)


func main() {
    dict := dict.New("dict/dict.json")
	dict.Add("hello", "world")
	dict.Add("foo", "bar")
	dict.Add("foo2", "bar2")
	dict.List()
	dict.Remove("foo2")
	dict.Update("hello", "world2")
	dict.List()
	dict.SaveToFile()
	dict.LoadFromFile()
}