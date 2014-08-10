package cgo_sample

// This file does not use cgo and so is not in go/build's CgoFiles. This is so
// that the package is detected by go/build.
//
// There is a bug in srclib-go where packages containing ONLY CgoFiles and no
// GoFiles are not detected.

func Bar() {}
