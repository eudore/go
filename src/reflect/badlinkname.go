// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import (
	"internal/abi"
	_ "unsafe"
)

// As of Go 1.22, the symbols below are found to be pulled via
// linkname in the wild. We provide a push linkname here, to
// keep them accessible with pull linknames.
// This may change in the future. Please do not depend on them
// in new code.

//go:linkname add
//go:linkname ifaceIndir
//go:linkname mapassign
//go:linkname rtypeOff
//go:linkname toType
//go:linkname typesByString
//go:linkname valueInterface

// The compiler doesn't allow linknames on methods, for good reasons.
// We use this trick to push linknames of the methods.
// Do not call them in this package.

//go:linkname badlinkname_rtype_Align reflect.(*rtype).Align
func badlinkname_rtype_Align(*rtype) int

//go:linkname badlinkname_rtype_AssignableTo reflect.(*rtype).AssignableTo
func badlinkname_rtype_AssignableTo(*rtype, Type) bool

//go:linkname badlinkname_rtype_Bits reflect.(*rtype).Bits
func badlinkname_rtype_Bits(*rtype) int

//go:linkname badlinkname_rtype_ChanDir reflect.(*rtype).ChanDir
func badlinkname_rtype_ChanDir(*rtype) ChanDir

//go:linkname badlinkname_rtype_Comparable reflect.(*rtype).Comparable
func badlinkname_rtype_Comparable(*rtype) bool

//go:linkname badlinkname_rtype_ConvertibleTo reflect.(*rtype).ConvertibleTo
func badlinkname_rtype_ConvertibleTo(*rtype, Type) bool

//go:linkname badlinkname_rtype_Elem reflect.(*rtype).Elem
func badlinkname_rtype_Elem(*rtype) Type

//go:linkname badlinkname_rtype_Field reflect.(*rtype).Field
func badlinkname_rtype_Field(*rtype, int) StructField

//go:linkname badlinkname_rtype_FieldAlign reflect.(*rtype).FieldAlign
func badlinkname_rtype_FieldAlign(*rtype) int

//go:linkname badlinkname_rtype_FieldByIndex reflect.(*rtype).FieldByIndex
func badlinkname_rtype_FieldByIndex(*rtype, []int) StructField

//go:linkname badlinkname_rtype_FieldByName reflect.(*rtype).FieldByName
func badlinkname_rtype_FieldByName(*rtype, string) (StructField, bool)

//go:linkname badlinkname_rtype_FieldByNameFunc reflect.(*rtype).FieldByNameFunc
func badlinkname_rtype_FieldByNameFunc(*rtype, func(string) bool) (StructField, bool)

//go:linkname badlinkname_rtype_Implements reflect.(*rtype).Implements
func badlinkname_rtype_Implements(*rtype, Type) bool

//go:linkname badlinkname_rtype_In reflect.(*rtype).In
func badlinkname_rtype_In(*rtype, int) Type

//go:linkname badlinkname_rtype_IsVariadic reflect.(*rtype).IsVariadic
func badlinkname_rtype_IsVariadic(*rtype) bool

//go:linkname badlinkname_rtype_Key reflect.(*rtype).Key
func badlinkname_rtype_Key(*rtype) Type

//go:linkname badlinkname_rtype_Kind reflect.(*rtype).Kind
func badlinkname_rtype_Kind(*rtype) Kind

//go:linkname badlinkname_rtype_Len reflect.(*rtype).Len
func badlinkname_rtype_Len(*rtype) int

//go:linkname badlinkname_rtype_Method reflect.(*rtype).Method
func badlinkname_rtype_Method(*rtype, int) Method

//go:linkname badlinkname_rtype_MethodByName reflect.(*rtype).MethodByName
func badlinkname_rtype_MethodByName(*rtype, string) (Method, bool)

//go:linkname badlinkname_rtype_Name reflect.(*rtype).Name
func badlinkname_rtype_Name(*rtype) string

//go:linkname badlinkname_rtype_NumField reflect.(*rtype).NumField
func badlinkname_rtype_NumField(*rtype) int

//go:linkname badlinkname_rtype_NumIn reflect.(*rtype).NumIn
func badlinkname_rtype_NumIn(*rtype) int

//go:linkname badlinkname_rtype_NumMethod reflect.(*rtype).NumMethod
func badlinkname_rtype_NumMethod(*rtype) int

//go:linkname badlinkname_rtype_NumOut reflect.(*rtype).NumOut
func badlinkname_rtype_NumOut(*rtype) int

//go:linkname badlinkname_rtype_Out reflect.(*rtype).Out
func badlinkname_rtype_Out(*rtype, int) Type

//go:linkname badlinkname_rtype_PkgPath reflect.(*rtype).PkgPath
func badlinkname_rtype_PkgPath(*rtype) string

//go:linkname badlinkname_rtype_Size reflect.(*rtype).Size
func badlinkname_rtype_Size(*rtype) uintptr

//go:linkname badlinkname_rtype_String reflect.(*rtype).String
func badlinkname_rtype_String(*rtype) string

//go:linkname badlinkname_rtype_ptrTo reflect.(*rtype).ptrTo
func badlinkname_rtype_ptrTo(*rtype) *abi.Type