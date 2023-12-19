package dict
	
import (
	"fmt"
)

type dict struct {
    m map[string]string
}

func New() dict {
	return dict{m: make(map[string]string)}
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

