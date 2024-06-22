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

// IWkContextMenu Root Interface
type IWkContextMenu interface {
	IObject
	Data() WebKitContextMenu                             // function
	GetItemsLength() int32                               // function
	GetItemAtPosition(position int32) IWkContextMenuItem // function
	Prepend(item IWkContextMenuItem)                     // procedure
	Append(item IWkContextMenuItem)                      // procedure
	Insert(item IWkContextMenuItem, position int32)      // procedure
	Remove(item IWkContextMenuItem)                      // procedure
	RemoveAll()                                          // procedure
}

// TWkContextMenu Root Object
type TWkContextMenu struct {
	TObject
}

func NewWkContextMenu() IWkContextMenu {
	r1 := wkContextMenuImportAPI().SysCallN(1)
	return AsWkContextMenu(r1)
}

func NewWkContextMenu1(aContextMenu WebKitContextMenu) IWkContextMenu {
	r1 := wkContextMenuImportAPI().SysCallN(2, uintptr(aContextMenu))
	return AsWkContextMenu(r1)
}

func (m *TWkContextMenu) Data() WebKitContextMenu {
	r1 := wkContextMenuImportAPI().SysCallN(3, m.Instance())
	return WebKitContextMenu(r1)
}

func (m *TWkContextMenu) GetItemsLength() int32 {
	r1 := wkContextMenuImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TWkContextMenu) GetItemAtPosition(position int32) IWkContextMenuItem {
	var resultWkContextMenuItem uintptr
	wkContextMenuImportAPI().SysCallN(4, m.Instance(), uintptr(position), uintptr(unsafePointer(&resultWkContextMenuItem)))
	return AsWkContextMenuItem(resultWkContextMenuItem)
}

func (m *TWkContextMenu) Prepend(item IWkContextMenuItem) {
	wkContextMenuImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(item))
}

func (m *TWkContextMenu) Append(item IWkContextMenuItem) {
	wkContextMenuImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(item))
}

func (m *TWkContextMenu) Insert(item IWkContextMenuItem, position int32) {
	wkContextMenuImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(item), uintptr(position))
}

func (m *TWkContextMenu) Remove(item IWkContextMenuItem) {
	wkContextMenuImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(item))
}

func (m *TWkContextMenu) RemoveAll() {
	wkContextMenuImportAPI().SysCallN(9, m.Instance())
}

var (
	wkContextMenuImport       *imports.Imports = nil
	wkContextMenuImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkContextMenu_Append", 0),
		/*1*/ imports.NewTable("WkContextMenu_Create", 0),
		/*2*/ imports.NewTable("WkContextMenu_Create1", 0),
		/*3*/ imports.NewTable("WkContextMenu_Data", 0),
		/*4*/ imports.NewTable("WkContextMenu_GetItemAtPosition", 0),
		/*5*/ imports.NewTable("WkContextMenu_GetItemsLength", 0),
		/*6*/ imports.NewTable("WkContextMenu_Insert", 0),
		/*7*/ imports.NewTable("WkContextMenu_Prepend", 0),
		/*8*/ imports.NewTable("WkContextMenu_Remove", 0),
		/*9*/ imports.NewTable("WkContextMenu_RemoveAll", 0),
	}
)

func wkContextMenuImportAPI() *imports.Imports {
	if wkContextMenuImport == nil {
		wkContextMenuImport = NewDefaultImports()
		wkContextMenuImport.SetImportTable(wkContextMenuImportTables)
		wkContextMenuImportTables = nil
	}
	return wkContextMenuImport
}
