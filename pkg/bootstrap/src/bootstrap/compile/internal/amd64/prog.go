// Do not edit. Bootstrap copy of /Volumes/Android/buildbot/src/android/build-tools/out/obj/go/src/cmd/compile/internal/amd64/prog.go

//line /Volumes/Android/buildbot/src/android/build-tools/out/obj/go/src/cmd/compile/internal/amd64/prog.go:1
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package amd64

import (
	"bootstrap/compile/internal/gc"
	"bootstrap/internal/obj"
	"bootstrap/internal/obj/x86"
)

const (
	LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
	RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
var progtable = [x86.ALAST & obj.AMask]obj.ProgInfo{
	obj.ATYPE:     {Flags: gc.Pseudo | gc.Skip},
	obj.ATEXT:     {Flags: gc.Pseudo},
	obj.AFUNCDATA: {Flags: gc.Pseudo},
	obj.APCDATA:   {Flags: gc.Pseudo},
	obj.AUNDEF:    {Flags: gc.Break},
	obj.AUSEFIELD: {Flags: gc.OK},
	obj.ACHECKNIL: {Flags: gc.LeftRead},
	obj.AVARDEF:   {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARKILL:  {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARLIVE:  {Flags: gc.Pseudo | gc.LeftRead},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the Intel opcode.
	obj.ANOP:               {Flags: gc.LeftRead | gc.RightWrite},
	x86.AADCL & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry},
	x86.AADCQ & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry},
	x86.AADCW & obj.AMask:  {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry},
	x86.AADDB & obj.AMask:  {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AADDL & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AADDW & obj.AMask:  {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AADDQ & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AADDSD & obj.AMask: {Flags: gc.SizeD | gc.LeftRead | RightRdwr},
	x86.AADDSS & obj.AMask: {Flags: gc.SizeF | gc.LeftRead | RightRdwr},
	x86.AANDB & obj.AMask:  {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AANDL & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AANDQ & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AANDW & obj.AMask:  {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry},

	x86.ABSFL & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.SetCarry},
	x86.ABSFQ & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.SetCarry},
	x86.ABSFW & obj.AMask:   {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.SetCarry},
	x86.ABSRL & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.SetCarry},
	x86.ABSRQ & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.SetCarry},
	x86.ABSRW & obj.AMask:   {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.SetCarry},
	x86.ABSWAPL & obj.AMask: {Flags: gc.SizeL | RightRdwr},
	x86.ABSWAPQ & obj.AMask: {Flags: gc.SizeQ | RightRdwr},

	obj.ACALL & obj.AMask: {Flags: gc.RightAddr | gc.Call | gc.KillCarry},
	x86.ACDQ & obj.AMask:  {Flags: gc.OK, Reguse: AX, Regset: AX | DX},
	x86.ACQO & obj.AMask:  {Flags: gc.OK, Reguse: AX, Regset: AX | DX},
	x86.ACWD & obj.AMask:  {Flags: gc.OK, Reguse: AX, Regset: AX | DX},
	x86.ACLD & obj.AMask:  {Flags: gc.OK},
	x86.ASTD & obj.AMask:  {Flags: gc.OK},

	x86.ACMOVLEQ & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.UseCarry},
	x86.ACMOVLNE & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.UseCarry},
	x86.ACMOVQEQ & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.UseCarry},
	x86.ACMOVQNE & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.UseCarry},
	x86.ACMOVWEQ & obj.AMask: {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.UseCarry},
	x86.ACMOVWNE & obj.AMask: {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.UseCarry},

	x86.ACMPB & obj.AMask:      {Flags: gc.SizeB | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ACMPL & obj.AMask:      {Flags: gc.SizeL | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ACMPQ & obj.AMask:      {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ACMPW & obj.AMask:      {Flags: gc.SizeW | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ACOMISD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ACOMISS & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ACVTSD2SL & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSD2SQ & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSD2SS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSL2SD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSL2SS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSQ2SD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSQ2SS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSS2SD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSS2SL & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTSS2SQ & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTTSD2SL & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTTSD2SQ & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTTSS2SL & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ACVTTSS2SQ & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.ADECB & obj.AMask:      {Flags: gc.SizeB | RightRdwr},
	x86.ADECL & obj.AMask:      {Flags: gc.SizeL | RightRdwr},
	x86.ADECQ & obj.AMask:      {Flags: gc.SizeQ | RightRdwr},
	x86.ADECW & obj.AMask:      {Flags: gc.SizeW | RightRdwr},
	x86.ADIVB & obj.AMask:      {Flags: gc.SizeB | gc.LeftRead | gc.SetCarry, Reguse: AX, Regset: AX},
	x86.ADIVL & obj.AMask:      {Flags: gc.SizeL | gc.LeftRead | gc.SetCarry, Reguse: AX | DX, Regset: AX | DX},
	x86.ADIVQ & obj.AMask:      {Flags: gc.SizeQ | gc.LeftRead | gc.SetCarry, Reguse: AX | DX, Regset: AX | DX},
	x86.ADIVW & obj.AMask:      {Flags: gc.SizeW | gc.LeftRead | gc.SetCarry, Reguse: AX | DX, Regset: AX | DX},
	x86.ADIVSD & obj.AMask:     {Flags: gc.SizeD | gc.LeftRead | RightRdwr},
	x86.ADIVSS & obj.AMask:     {Flags: gc.SizeF | gc.LeftRead | RightRdwr},
	x86.AIDIVB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | gc.SetCarry, Reguse: AX, Regset: AX},
	x86.AIDIVL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | gc.SetCarry, Reguse: AX | DX, Regset: AX | DX},
	x86.AIDIVQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | gc.SetCarry, Reguse: AX | DX, Regset: AX | DX},
	x86.AIDIVW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | gc.SetCarry, Reguse: AX | DX, Regset: AX | DX},
	x86.AIMULB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | gc.SetCarry, Reguse: AX, Regset: AX},
	x86.AIMULL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | gc.ImulAXDX | gc.SetCarry},
	x86.AIMULQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | gc.ImulAXDX | gc.SetCarry},
	x86.AIMULW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | gc.ImulAXDX | gc.SetCarry},
	x86.AINCB & obj.AMask:      {Flags: gc.SizeB | RightRdwr},
	x86.AINCL & obj.AMask:      {Flags: gc.SizeL | RightRdwr},
	x86.AINCQ & obj.AMask:      {Flags: gc.SizeQ | RightRdwr},
	x86.AINCW & obj.AMask:      {Flags: gc.SizeW | RightRdwr},
	x86.AJCC & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJCS & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJEQ & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJGE & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJGT & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJHI & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJLE & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJLS & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJLT & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJMI & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJNE & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJOC & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJOS & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJPC & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJPL & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	x86.AJPS & obj.AMask:       {Flags: gc.Cjmp | gc.UseCarry},
	obj.AJMP & obj.AMask:       {Flags: gc.Jump | gc.Break | gc.KillCarry},
	x86.ALEAW & obj.AMask:      {Flags: gc.LeftAddr | gc.RightWrite},
	x86.ALEAL & obj.AMask:      {Flags: gc.LeftAddr | gc.RightWrite},
	x86.ALEAQ & obj.AMask:      {Flags: gc.LeftAddr | gc.RightWrite},
	x86.AMOVBLSX & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVBLZX & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVBQSX & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVBQZX & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVBWSX & obj.AMask:   {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVBWZX & obj.AMask:   {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVLQSX & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVLQZX & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVWLSX & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVWLZX & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVWQSX & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVWQZX & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVQL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	x86.AMOVB & obj.AMask:      {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move},
	x86.AMOVL & obj.AMask:      {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move},
	x86.AMOVQ & obj.AMask:      {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Move},
	x86.AMOVW & obj.AMask:      {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move},
	x86.AMOVUPS & obj.AMask:    {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	x86.AMOVSB & obj.AMask:     {Flags: gc.OK, Reguse: DI | SI, Regset: DI | SI},
	x86.AMOVSL & obj.AMask:     {Flags: gc.OK, Reguse: DI | SI, Regset: DI | SI},
	x86.AMOVSQ & obj.AMask:     {Flags: gc.OK, Reguse: DI | SI, Regset: DI | SI},
	x86.AMOVSW & obj.AMask:     {Flags: gc.OK, Reguse: DI | SI, Regset: DI | SI},
	obj.ADUFFCOPY & obj.AMask:  {Flags: gc.OK, Reguse: DI | SI, Regset: DI | SI | X0},
	x86.AMOVSD & obj.AMask:     {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move},
	x86.AMOVSS & obj.AMask:     {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move},

	// We use&obj.AMask MOVAPD as a faster synonym for MOVSD.
	x86.AMOVAPD & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move},
	x86.AMULB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | gc.SetCarry, Reguse: AX, Regset: AX},
	x86.AMULL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | gc.SetCarry, Reguse: AX, Regset: AX | DX},
	x86.AMULQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | gc.SetCarry, Reguse: AX, Regset: AX | DX},
	x86.AMULW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | gc.SetCarry, Reguse: AX, Regset: AX | DX},
	x86.AMULSD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | RightRdwr},
	x86.AMULSS & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | RightRdwr},
	x86.ANEGB & obj.AMask:     {Flags: gc.SizeB | RightRdwr | gc.SetCarry},
	x86.ANEGL & obj.AMask:     {Flags: gc.SizeL | RightRdwr | gc.SetCarry},
	x86.ANEGQ & obj.AMask:     {Flags: gc.SizeQ | RightRdwr | gc.SetCarry},
	x86.ANEGW & obj.AMask:     {Flags: gc.SizeW | RightRdwr | gc.SetCarry},
	x86.ANOTB & obj.AMask:     {Flags: gc.SizeB | RightRdwr},
	x86.ANOTL & obj.AMask:     {Flags: gc.SizeL | RightRdwr},
	x86.ANOTQ & obj.AMask:     {Flags: gc.SizeQ | RightRdwr},
	x86.ANOTW & obj.AMask:     {Flags: gc.SizeW | RightRdwr},
	x86.AORB & obj.AMask:      {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AORL & obj.AMask:      {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AORQ & obj.AMask:      {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AORW & obj.AMask:      {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.APOPQ & obj.AMask:     {Flags: gc.SizeQ | gc.RightWrite},
	x86.APUSHQ & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead},
	x86.APXOR & obj.AMask:     {Flags: gc.SizeD | gc.LeftRead | RightRdwr},
	x86.ARCLB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry},
	x86.ARCLL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry},
	x86.ARCLQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry},
	x86.ARCLW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry},
	x86.ARCRB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry},
	x86.ARCRL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry},
	x86.ARCRQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry},
	x86.ARCRW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry},
	x86.AREP & obj.AMask:      {Flags: gc.OK, Reguse: CX, Regset: CX},
	x86.AREPN & obj.AMask:     {Flags: gc.OK, Reguse: CX, Regset: CX},
	obj.ARET & obj.AMask:      {Flags: gc.Break | gc.KillCarry},
	x86.AROLB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.AROLL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.AROLQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.AROLW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ARORB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ARORL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ARORQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ARORW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASALB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASALL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASALQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASALW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASARB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASARL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASARQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASARW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASBBB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry},
	x86.ASBBL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry},
	x86.ASBBQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry},
	x86.ASBBW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry},
	x86.ASETCC & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETCS & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETEQ & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETGE & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETGT & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETHI & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETLE & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETLS & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETLT & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETMI & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETNE & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETOC & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETOS & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETPC & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETPL & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASETPS & obj.AMask:    {Flags: gc.SizeB | gc.RightWrite | gc.UseCarry},
	x86.ASHLB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASHLL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASHLQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASHLW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASHRB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASHRL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASHRQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASHRW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry},
	x86.ASQRTSD & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | RightRdwr},
	x86.ASTOSB & obj.AMask:    {Flags: gc.OK, Reguse: AX | DI, Regset: DI},
	x86.ASTOSL & obj.AMask:    {Flags: gc.OK, Reguse: AX | DI, Regset: DI},
	x86.ASTOSQ & obj.AMask:    {Flags: gc.OK, Reguse: AX | DI, Regset: DI},
	x86.ASTOSW & obj.AMask:    {Flags: gc.OK, Reguse: AX | DI, Regset: DI},
	obj.ADUFFZERO & obj.AMask: {Flags: gc.OK, Reguse: X0 | DI, Regset: DI},
	x86.ASUBB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.ASUBL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.ASUBQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.ASUBW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.ASUBSD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | RightRdwr},
	x86.ASUBSS & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | RightRdwr},
	x86.ATESTB & obj.AMask:    {Flags: gc.SizeB | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ATESTL & obj.AMask:    {Flags: gc.SizeL | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ATESTQ & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.ATESTW & obj.AMask:    {Flags: gc.SizeW | gc.LeftRead | gc.RightRead | gc.SetCarry},
	x86.AUCOMISD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightRead},
	x86.AUCOMISS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RightRead},
	x86.AXCHGB & obj.AMask:    {Flags: gc.SizeB | LeftRdwr | RightRdwr},
	x86.AXCHGL & obj.AMask:    {Flags: gc.SizeL | LeftRdwr | RightRdwr},
	x86.AXCHGQ & obj.AMask:    {Flags: gc.SizeQ | LeftRdwr | RightRdwr},
	x86.AXCHGW & obj.AMask:    {Flags: gc.SizeW | LeftRdwr | RightRdwr},
	x86.AXORB & obj.AMask:     {Flags: gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AXORL & obj.AMask:     {Flags: gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AXORQ & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AXORW & obj.AMask:     {Flags: gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry},
	x86.AXORPS & obj.AMask:    {Flags: gc.LeftRead | RightRdwr},
}

func progflags(p *obj.Prog) uint32 {
	flags := progtable[p.As&obj.AMask].Flags
	if flags&gc.ImulAXDX != 0 && p.To.Type != obj.TYPE_NONE {
		flags |= RightRdwr
	}
	return flags
}

func progcarryflags(p *obj.Prog) uint32 {
	return progtable[p.As&obj.AMask].Flags
}

func proginfo(p *obj.Prog) {
	info := &p.Info
	*info = progtable[p.As&obj.AMask]
	if info.Flags == 0 {
		gc.Fatalf("unknown instruction %v", p)
	}

	if (info.Flags&gc.ShiftCX != 0) && p.From.Type != obj.TYPE_CONST {
		info.Reguse |= CX
	}

	if info.Flags&gc.ImulAXDX != 0 {
		if p.To.Type == obj.TYPE_NONE {
			info.Reguse |= AX
			info.Regset |= AX | DX
		} else {
			info.Flags |= RightRdwr
		}
	}

	// Addressing makes some registers used.
	if p.From.Type == obj.TYPE_MEM && p.From.Name == obj.NAME_NONE {
		info.Regindex |= RtoB(int(p.From.Reg))
	}
	if p.From.Index != x86.REG_NONE {
		info.Regindex |= RtoB(int(p.From.Index))
	}
	if p.To.Type == obj.TYPE_MEM && p.To.Name == obj.NAME_NONE {
		info.Regindex |= RtoB(int(p.To.Reg))
	}
	if p.To.Index != x86.REG_NONE {
		info.Regindex |= RtoB(int(p.To.Index))
	}
	if gc.Ctxt.Flag_dynlink {
		// When -dynlink is passed, many operations on external names (and
		// also calling duffzero/duffcopy) use R15 as a scratch register.
		if p.As == x86.ALEAQ || info.Flags == gc.Pseudo || p.As == obj.ACALL || p.As == obj.ARET || p.As == obj.AJMP {
			return
		}
		if p.As == obj.ADUFFZERO || p.As == obj.ADUFFCOPY || (p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local) || (p.To.Name == obj.NAME_EXTERN && !p.To.Sym.Local) {
			info.Reguse |= R15
			info.Regset |= R15
			return
		}
	}
}
