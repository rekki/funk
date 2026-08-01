// Package funk provides small, type-safe helpers for working with slices,
// maps, sets, pointers, and strings.
//
// Slice transforms preserve input order. Functions that range over maps do not
// guarantee output order. Unless a function explicitly says it copies its
// input, returned subslices may share storage with the input slice.
package funk
