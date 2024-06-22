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

// IWkWebContext Root Interface
type IWkWebContext interface {
	IObject
	GetCookieManager() IWkCookieManager                                        // function
	RegisterURIScheme(aScheme string, aDelegate IWkSchemeRequestDelegateEvent) // procedure
	SetCacheModel(cachemodel WebKitCacheModel)                                 // procedure
	DownloadURI(uri string)                                                    // procedure
	SetWetPreferredLanguages(languages IStrings)                               // procedure
}

// TWkWebContext Root Object
type TWkWebContext struct {
	TObject
}

func NewWkWebContext(aWebContext WebKitWebContext) IWkWebContext {
	r1 := wkWebContextImportAPI().SysCallN(0, uintptr(aWebContext))
	return AsWkWebContext(r1)
}

// WkWebContextRef -> IWkWebContext
var WkWebContextRef wkWebContext

// wkWebContext TWkWebContext Ref
type wkWebContext uintptr

func (m *wkWebContext) Default() IWkWebContext {
	var resultWkWebContext uintptr
	wkWebContextImportAPI().SysCallN(1, uintptr(unsafePointer(&resultWkWebContext)))
	return AsWkWebContext(resultWkWebContext)
}

func (m *wkWebContext) New() IWkWebContext {
	var resultWkWebContext uintptr
	wkWebContextImportAPI().SysCallN(4, uintptr(unsafePointer(&resultWkWebContext)))
	return AsWkWebContext(resultWkWebContext)
}

func (m *TWkWebContext) GetCookieManager() IWkCookieManager {
	var resultWkCookieManager uintptr
	wkWebContextImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultWkCookieManager)))
	return AsWkCookieManager(resultWkCookieManager)
}

func (m *TWkWebContext) RegisterURIScheme(aScheme string, aDelegate IWkSchemeRequestDelegateEvent) {
	wkWebContextImportAPI().SysCallN(5, m.Instance(), PascalStr(aScheme), GetObjectUintptr(aDelegate))
}

func (m *TWkWebContext) SetCacheModel(cachemodel WebKitCacheModel) {
	wkWebContextImportAPI().SysCallN(6, m.Instance(), uintptr(cachemodel))
}

func (m *TWkWebContext) DownloadURI(uri string) {
	wkWebContextImportAPI().SysCallN(2, m.Instance(), PascalStr(uri))
}

func (m *TWkWebContext) SetWetPreferredLanguages(languages IStrings) {
	wkWebContextImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(languages))
}

var (
	wkWebContextImport       *imports.Imports = nil
	wkWebContextImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkWebContext_Create", 0),
		/*1*/ imports.NewTable("WkWebContext_Default", 0),
		/*2*/ imports.NewTable("WkWebContext_DownloadURI", 0),
		/*3*/ imports.NewTable("WkWebContext_GetCookieManager", 0),
		/*4*/ imports.NewTable("WkWebContext_New", 0),
		/*5*/ imports.NewTable("WkWebContext_RegisterURIScheme", 0),
		/*6*/ imports.NewTable("WkWebContext_SetCacheModel", 0),
		/*7*/ imports.NewTable("WkWebContext_SetWetPreferredLanguages", 0),
	}
)

func wkWebContextImportAPI() *imports.Imports {
	if wkWebContextImport == nil {
		wkWebContextImport = NewDefaultImports()
		wkWebContextImport.SetImportTable(wkWebContextImportTables)
		wkWebContextImportTables = nil
	}
	return wkWebContextImport
}
