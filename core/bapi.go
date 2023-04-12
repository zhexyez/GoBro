package core

import (
	"fmt"
)

type API struct {
	Page *Page
	Tree *[]*Element
}

func (api *API) PrintStructure() {
	for el := range *api.Tree {
		fmt.Println("element", el+1, "  :", string(MakeTree_inJSON((*api.Tree)[el])))
	}
}

func (api *API) NewElement(params []string) {
	x := NewElement(params[0], params[1], api.Page, nil, params[2], params[3])
	(*api.Tree) = append((*api.Tree), x)
	ObjMap[params[0]] = append(ObjMap[params[0]], x)
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