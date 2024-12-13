// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__value = 544
)

const (
    _stack__value = 104
)

const (
    _size__value = 13456
)

var (
    _pcsp__value = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {3601, 104},
        {3605, 48},
        {3606, 40},
        {3608, 32},
        {3610, 24},
        {3612, 16},
        {3614, 8},
        {3618, 0},
        {13456, 104},
    }
)

var _cfunc_value = []loader.CFunc{
    {"_value_entry", 0,  _entry__value, 0, nil},
    {"_value", _entry__value, _size__value, _stack__value, _pcsp__value},
}