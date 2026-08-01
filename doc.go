// Package funk provides a small set of type-safe helpers for slices, maps,
// and sets.
//
// Slice transforms preserve input order. Functions that range over maps do not
// guarantee output order. Nil and empty slices are treated equivalently;
// callers should not rely on whether an empty result is nil. Unless a function
// explicitly says it copies its input, returned subslices may share storage
// with the input slice.
package funk
