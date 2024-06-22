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
	. "github.com/energye/wk/types"
)

// IWkCookieList Root Interface
type IWkCookieList interface {
	IObject
	Length() int32                   // function
	GetCookie(index int32) IWkCookie // function
}

// TWkCookieList Root Object
type TWkCookieList struct {
	TObject
}

func NewWkCookieList(aList PList) IWkCookieList {
	r1 := wkCookieListImportAPI().SysCallN(0, uintptr(aList))
	return AsWkCookieList(r1)
}

func (m *TWkCookieList) Length() int32 {
	r1 := wkCookieListImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TWkCookieList) GetCookie(index int32) IWkCookie {
	var resultWkCookie uintptr
	wkCookieListImportAPI().SysCallN(1, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultWkCookie)))
	return AsWkCookie(resultWkCookie)
}

var (
	wkCookieListImport       *imports.Imports = nil
	wkCookieListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkCookieList_Create", 0),
		/*1*/ imports.NewTable("WkCookieList_GetCookie", 0),
		/*2*/ imports.NewTable("WkCookieList_Length", 0),
	}
)

func wkCookieListImportAPI() *imports.Imports {
	if wkCookieListImport == nil {
		wkCookieListImport = NewDefaultImports()
		wkCookieListImport.SetImportTable(wkCookieListImportTables)
		wkCookieListImportTables = nil
	}
	return wkCookieListImport
}
