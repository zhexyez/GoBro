package main

import (
	"fmt"
)

type API struct {
	Page page
	Tree []*element
}

func (api *API) PrintStructure() {
	for el := range api.Tree {
		fmt.Println("element", el+1, "  :", string(MakeTree_inJSON(api.Tree[el])))
	}
}

func (api *API) NewElement(params []string) {
	// fmt.Println("parent id:", params[0])
	// fmt.Println("id:", params[1])
	// fmt.Println("class:", params[2])
	// fmt.Println("ref:", params[3])
	// fmt.Println("txt:", params[4])
}