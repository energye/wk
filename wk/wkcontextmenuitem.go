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

// IWkContextMenuItem Root Interface
type IWkContextMenuItem interface {
	IObject
	Data() WebKitContextMenuItem             // function
	GetSubmenu() WebKitContextMenu           // function
	GetAction() WkAction                     // function
	GetStockAction() WebKitContextMenuAction // function
	IsSeparator() bool                       // function
	SetSubmenu(submenu IWkContextMenu)       // procedure
}

// TWkContextMenuItem Root Object
type TWkContextMenuItem struct {
	TObject
}

func NewWkContextMenuItem(aContextMenu WebKitContextMenuItem) IWkContextMenuItem {
	r1 := wkContextMenuItemImportAPI().SysCallN(0, uintptr(aContextMenu))
	return AsWkContextMenuItem(r1)
}

// WkContextMenuItemRef -> IWkContextMenuItem
var WkContextMenuItemRef wkContextMenuItem

// wkContextMenuItem TWkContextMenuItem Ref
type wkContextMenuItem uintptr

func (m *wkContextMenuItem) NewFromAction(action WkAction, aLabel string, commandId int32) IWkContextMenuItem {
	var resultWkContextMenuItem uintptr
	wkContextMenuItemImportAPI().SysCallN(6, uintptr(action), PascalStr(aLabel), uintptr(commandId), uintptr(unsafePointer(&resultWkContextMenuItem)))
	return AsWkContextMenuItem(resultWkContextMenuItem)
}

func (m *wkContextMenuItem) NewFromStockAction(action WebKitContextMenuAction) IWkContextMenuItem {
	var resultWkContextMenuItem uintptr
	wkContextMenuItemImportAPI().SysCallN(7, uintptr(action), uintptr(unsafePointer(&resultWkContextMenuItem)))
	return AsWkContextMenuItem(resultWkContextMenuItem)
}

func (m *wkContextMenuItem) NewFromStockActionWithLabel(action WebKitContextMenuAction, aLabel string) IWkContextMenuItem {
	var resultWkContextMenuItem uintptr
	wkContextMenuItemImportAPI().SysCallN(8, uintptr(action), PascalStr(aLabel), uintptr(unsafePointer(&resultWkContextMenuItem)))
	return AsWkContextMenuItem(resultWkContextMenuItem)
}

func (m *wkContextMenuItem) NewWithSubmenu(aLabel string, submenu WebKitContextMenu) IWkContextMenuItem {
	var resultWkContextMenuItem uintptr
	wkContextMenuItemImportAPI().SysCallN(10, PascalStr(aLabel), uintptr(submenu), uintptr(unsafePointer(&resultWkContextMenuItem)))
	return AsWkContextMenuItem(resultWkContextMenuItem)
}

func (m *wkContextMenuItem) NewSeparator() IWkContextMenuItem {
	var resultWkContextMenuItem uintptr
	wkContextMenuItemImportAPI().SysCallN(9, uintptr(unsafePointer(&resultWkContextMenuItem)))
	return AsWkContextMenuItem(resultWkContextMenuItem)
}

func (m *TWkContextMenuItem) Data() WebKitContextMenuItem {
	r1 := wkContextMenuItemImportAPI().SysCallN(1, m.Instance())
	return WebKitContextMenuItem(r1)
}

func (m *TWkContextMenuItem) GetSubmenu() WebKitContextMenu {
	r1 := wkContextMenuItemImportAPI().SysCallN(4, m.Instance())
	return WebKitContextMenu(r1)
}

func (m *TWkContextMenuItem) GetAction() WkAction {
	r1 := wkContextMenuItemImportAPI().SysCallN(2, m.Instance())
	return WkAction(r1)
}

func (m *TWkContextMenuItem) GetStockAction() WebKitContextMenuAction {
	r1 := wkContextMenuItemImportAPI().SysCallN(3, m.Instance())
	return WebKitContextMenuAction(r1)
}

func (m *TWkContextMenuItem) IsSeparator() bool {
	r1 := wkContextMenuItemImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TWkContextMenuItem) SetSubmenu(submenu IWkContextMenu) {
	wkContextMenuItemImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(submenu))
}

var (
	wkContextMenuItemImport       *imports.Imports = nil
	wkContextMenuItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkContextMenuItem_Create", 0),
		/*1*/ imports.NewTable("WkContextMenuItem_Data", 0),
		/*2*/ imports.NewTable("WkContextMenuItem_GetAction", 0),
		/*3*/ imports.NewTable("WkContextMenuItem_GetStockAction", 0),
		/*4*/ imports.NewTable("WkContextMenuItem_GetSubmenu", 0),
		/*5*/ imports.NewTable("WkContextMenuItem_IsSeparator", 0),
		/*6*/ imports.NewTable("WkContextMenuItem_NewFromAction", 0),
		/*7*/ imports.NewTable("WkContextMenuItem_NewFromStockAction", 0),
		/*8*/ imports.NewTable("WkContextMenuItem_NewFromStockActionWithLabel", 0),
		/*9*/ imports.NewTable("WkContextMenuItem_NewSeparator", 0),
		/*10*/ imports.NewTable("WkContextMenuItem_NewWithSubmenu", 0),
		/*11*/ imports.NewTable("WkContextMenuItem_SetSubmenu", 0),
	}
)

func wkContextMenuItemImportAPI() *imports.Imports {
	if wkContextMenuItemImport == nil {
		wkContextMenuItemImport = NewDefaultImports()
		wkContextMenuItemImport.SetImportTable(wkContextMenuItemImportTables)
		wkContextMenuItemImportTables = nil
	}
	return wkContextMenuItemImport
}
