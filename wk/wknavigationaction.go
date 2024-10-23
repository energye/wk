//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wk

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// IWkNavigationAction Root Interface
type IWkNavigationAction interface {
	IObject
	GetRequest() WebKitURIRequest // function
}

// TWkNavigationAction Root Object
type TWkNavigationAction struct {
	TObject
}

func NewWkNavigationAction(aNavigationAction WebKitNavigationAction) IWkNavigationAction {
	r1 := wkNavigationActionImportAPI().SysCallN(0, uintptr(aNavigationAction))
	return AsWkNavigationAction(r1)
}

func (m *TWkNavigationAction) GetRequest() WebKitURIRequest {
	r1 := wkNavigationActionImportAPI().SysCallN(1, m.Instance())
	return WebKitURIRequest(r1)
}

var (
	wkNavigationActionImport       *imports.Imports = nil
	wkNavigationActionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkNavigationAction_Create", 0),
		/*1*/ imports.NewTable("WkNavigationAction_GetRequest", 0),
	}
)

func wkNavigationActionImportAPI() *imports.Imports {
	if wkNavigationActionImport == nil {
		wkNavigationActionImport = NewDefaultImports()
		wkNavigationActionImport.SetImportTable(wkNavigationActionImportTables)
		wkNavigationActionImportTables = nil
	}
	return wkNavigationActionImport
}
