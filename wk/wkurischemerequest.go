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

// IWkURISchemeRequest Root Interface
type IWkURISchemeRequest interface {
	IObject
	Data() WebKitURISchemeRequest                                         // function
	Scheme() string                                                       // function
	Uri() string                                                          // function
	Path() string                                                         // function
	WebView() WebKitWebView                                               // function
	Method() string                                                       // function
	Headers() IWkHeaders                                                  // function
	Body() IWkInputStream                                                 // function
	Finish(stream IWkInputStream, streamlength int64, contenttype string) // procedure
	FinishWithResponse(response WebKitURISchemeResponse)                  // procedure
	FinishError(code int32, errorMessage string)                          // procedure
}

// TWkURISchemeRequest Root Object
type TWkURISchemeRequest struct {
	TObject
}

func NewWkURISchemeRequest(uRISchemeRequest WebKitURISchemeRequest) IWkURISchemeRequest {
	r1 := wkURISchemeRequestImportAPI().SysCallN(1, uintptr(uRISchemeRequest))
	return AsWkURISchemeRequest(r1)
}

func (m *TWkURISchemeRequest) Data() WebKitURISchemeRequest {
	r1 := wkURISchemeRequestImportAPI().SysCallN(2, m.Instance())
	return WebKitURISchemeRequest(r1)
}

func (m *TWkURISchemeRequest) Scheme() string {
	r1 := wkURISchemeRequestImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TWkURISchemeRequest) Uri() string {
	r1 := wkURISchemeRequestImportAPI().SysCallN(10, m.Instance())
	return GoStr(r1)
}

func (m *TWkURISchemeRequest) Path() string {
	r1 := wkURISchemeRequestImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TWkURISchemeRequest) WebView() WebKitWebView {
	r1 := wkURISchemeRequestImportAPI().SysCallN(11, m.Instance())
	return WebKitWebView(r1)
}

func (m *TWkURISchemeRequest) Method() string {
	r1 := wkURISchemeRequestImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TWkURISchemeRequest) Headers() IWkHeaders {
	var resultWkHeaders uintptr
	wkURISchemeRequestImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultWkHeaders)))
	return AsWkHeaders(resultWkHeaders)
}

func (m *TWkURISchemeRequest) Body() IWkInputStream {
	var resultWkInputStream uintptr
	wkURISchemeRequestImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultWkInputStream)))
	return AsWkInputStream(resultWkInputStream)
}

func (m *TWkURISchemeRequest) Finish(stream IWkInputStream, streamlength int64, contenttype string) {
	wkURISchemeRequestImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(stream), uintptr(unsafePointer(&streamlength)), PascalStr(contenttype))
}

func (m *TWkURISchemeRequest) FinishWithResponse(response WebKitURISchemeResponse) {
	wkURISchemeRequestImportAPI().SysCallN(5, m.Instance(), uintptr(response))
}

func (m *TWkURISchemeRequest) FinishError(code int32, errorMessage string) {
	wkURISchemeRequestImportAPI().SysCallN(4, m.Instance(), uintptr(code), PascalStr(errorMessage))
}

var (
	wkURISchemeRequestImport       *imports.Imports = nil
	wkURISchemeRequestImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkURISchemeRequest_Body", 0),
		/*1*/ imports.NewTable("WkURISchemeRequest_Create", 0),
		/*2*/ imports.NewTable("WkURISchemeRequest_Data", 0),
		/*3*/ imports.NewTable("WkURISchemeRequest_Finish", 0),
		/*4*/ imports.NewTable("WkURISchemeRequest_FinishError", 0),
		/*5*/ imports.NewTable("WkURISchemeRequest_FinishWithResponse", 0),
		/*6*/ imports.NewTable("WkURISchemeRequest_Headers", 0),
		/*7*/ imports.NewTable("WkURISchemeRequest_Method", 0),
		/*8*/ imports.NewTable("WkURISchemeRequest_Path", 0),
		/*9*/ imports.NewTable("WkURISchemeRequest_Scheme", 0),
		/*10*/ imports.NewTable("WkURISchemeRequest_Uri", 0),
		/*11*/ imports.NewTable("WkURISchemeRequest_WebView", 0),
	}
)

func wkURISchemeRequestImportAPI() *imports.Imports {
	if wkURISchemeRequestImport == nil {
		wkURISchemeRequestImport = NewDefaultImports()
		wkURISchemeRequestImport.SetImportTable(wkURISchemeRequestImportTables)
		wkURISchemeRequestImportTables = nil
	}
	return wkURISchemeRequestImport
}
