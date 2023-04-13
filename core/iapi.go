package core

import (
	"fmt"
)

type API struct {
	Page *Page
	Tree *[]*Element
}

// Instance API -> prints POT structure as JSON object
func (api *API) PrintStructure() {
	for el := range *api.Tree {
		fmt.Println("element", el+1, "  :", string(MakeTree_inJSON((*api.Tree)[el])))
	}
}

// Instance API -> creates new element
func (api *API) NewElement(params []string) {
	x := NewElement(params[0], params[1], api.Page, nil, params[2], params[3])
	(*api.Tree) = append((*api.Tree), x)
	ObjMap[params[0]] = append(ObjMap[params[0]], x)
}

// Instance API -> changes ID on each element with given ID
func (api *API) ElementsChangeID(params []string) {
	for el := range ObjMap[params[0]] {
		// maybe should also update ObjMap. Must check the behavior
		ObjMap[params[0]][el].ChangeID(params[1])
	}
}

// Instance API -> changes PARENT on the first element by ID list
func (api *API) ChangeParent(params []string) {
	for el := range ObjMap[params[0]] {
		ObjMap[params[0]][el].ChangeParent(ObjMap[params[1]][0])
	}
}

func (api *API) AppendChild(params []string) {
	ObjMap[params[0]][0].AppendChild(ObjMap[params[1]][0])
}

func (api *API) ChangeClass(params []string) {
	for el := range ObjMap[params[0]] {
		ObjMap[params[0]][el].ChangeClass(params[1])
	}
}

// TODO
// func (api *API) ChangeClass(params [string]) {}
// func (api *API) ChangeRef(params [string]) {}
// func (api *API) ChangeTXT(params [string]) {}
// func (api *API) TXTRemoveOneFromEnd(params [string]) {}
// func (api *API) TXTRemoveOneFromStart(params [string]) {}
// func (api *API) TXTRemoveInBoundaries(params [string]) {}
// until in pot not complete
// func (api *API) FlushElement(params [string]) {}
// func (api *API) FlushElementKeepChildren(params [string]) {}