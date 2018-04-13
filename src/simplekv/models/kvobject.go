
package models

import (
	"fmt"
)

//
type  KVObject struct{
	Key string `json:key`
	Value string `json:value`
}

func init() {
	fmt.Println("Hello world init KVOBject")
}

var (
	// global key - value cache
	KV = make (map[string]string )
)

