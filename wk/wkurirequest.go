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

// IWkURIRequest Root Interface
type IWkURIRequest interface {
	IObject
	URI() string          // property
	SetURI(AValue string) // property
	Method() string       // function
	Headers() IWkHeaders  // function
}

// TWkURIRequest Root Object
type TWkURIRequest struct {
	TObject
}

func NewWkURIRequest(aURIRequest WebKitURIRequest) IWkURIRequest {
	r1 := wkURIRequestImportAPI().SysCallN(0, uintptr(aURIRequest))
	return AsWkURIRequest(r1)
}

// WkURIRequestRef -> IWkURIRequest
var WkURIRequestRef wkURIRequest

// wkURIRequest TWkURIRequest Ref
type wkURIRequest uintptr

func (m *wkURIRequest) NewURIRequest(uri string) IWkURIRequest {
	var resultWkURIRequest uintptr
	wkURIRequestImportAPI().SysCallN(3, PascalStr(uri), uintptr(unsafePointer(&resultWkURIRequest)))
	return AsWkURIRequest(resultWkURIRequest)
}

func (m *TWkURIRequest) URI() string {
	r1 := wkURIRequestImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWkURIRequest) SetURI(AValue string) {
	wkURIRequestImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWkURIRequest) Method() string {
	r1 := wkURIRequestImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TWkURIRequest) Headers() IWkHeaders {
	var resultWkHeaders uintptr
	wkURIRequestImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultWkHeaders)))
	return AsWkHeaders(resultWkHeaders)
}

var (
	wkURIRequestImport       *imports.Imports = nil
	wkURIRequestImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkURIRequest_Create", 0),
		/*1*/ imports.NewTable("WkURIRequest_Headers", 0),
		/*2*/ imports.NewTable("WkURIRequest_Method", 0),
		/*3*/ imports.NewTable("WkURIRequest_NewURIRequest", 0),
		/*4*/ imports.NewTable("WkURIRequest_URI", 0),
	}
)

func wkURIRequestImportAPI() *imports.Imports {
	if wkURIRequestImport == nil {
		wkURIRequestImport = NewDefaultImports()
		wkURIRequestImport.SetImportTable(wkURIRequestImportTables)
		wkURIRequestImportTables = nil
	}
	return wkURIRequestImport
}
