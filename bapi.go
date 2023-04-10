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
	x := NewElement(params[1], params[2], &api.Page, nil, params[3], params[4])
	ObjMap[params[1]] = append(ObjMap[params[1]], x)
}

func (api *API) ElementsChangeID(params []string) {
	selected := ObjMap[params[0]]
	for el := range selected {
		selected[el].ChangeID(params[1])
	}
}

// TODO
// func (api *API) ChangeParent(params []string) {}
// func (api *API) AppendChild(params [string]) {}
// func (api *API) ChangeClass(params [string]) {}
// func (api *API) ChangeRef(params [string]) {}
// func (api *API) ChangeTXT(params [string]) {}
// func (api *API) RemoveOneFromEnd(params [string]) {}
// func (api *API) FlushElement(params [string]) {}
// func (api *API) FlushElementKeepChildren(params [string]) {}