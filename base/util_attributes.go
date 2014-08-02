package base

import (
	"fmt"
)

// This file contains utility functions relating to Attributes and Attribute specifications.

// NonClassFloatAttributes returns all FloatAttributes which
// aren't designated as a class Attribute.
func NonClassFloatAttributes(d DataGrid) []Attribute {
	classAttrs := d.AllClassAttributes()
	allAttrs := d.AllAttributes()
	ret := make([]Attribute, 0)
	for _, a := range allAttrs {
		matched := false
		if _, ok := a.(*FloatAttribute); !ok {
			continue
		}
		for _, b := range classAttrs {
			if a.Equals(b) {
				matched = true
				break
			}
		}
		if !matched {
			ret = append(ret, a)
		}
	}
	return ret
}

// NonClassAttrs returns all Attributes which aren't designated as a
// class Attribute.
func NonClassAttributes(d DataGrid) []Attribute {
	classAttrs := d.AllClassAttributes()
	allAttrs := d.AllAttributes()
	return AttributeDifferenceReferences(allAttrs, classAttrs)
}

// ResolveAllAttributes returns AttributeSpecs describing
// all of the Attributes.
func ResolveAllAttributes(d DataGrid, attrs []Attribute) []AttributeSpec {
	ret := make([]AttributeSpec, len(attrs))
	for i, a := range attrs {
		spec, err := d.GetAttribute(a)
		if err != nil {
			panic(fmt.Errorf("Error resolving Attribute %s: %s", a, err))
		}
		ret[i] = spec
	}
	return ret
}

// GetAllAttributeSpecs retrieves every Attribute specification
// from a given DataGrid. Useful in conjunction with MapOverRows.
func GetAllAttributeSpecs(from DataGrid) []AttributeSpec {
	attrs := from.AllAttributes()
	return GetSomeAttributeSpecs(from, attrs)
}

// GetSomeAttributeSpecs returns a subset of Attribute specifications
// from a given DataGrid.
func GetSomeAttributeSpecs(from DataGrid, attrs []Attribute) []AttributeSpec {
	ret := make([]AttributeSpec, len(attrs))
	for i, a := range attrs {
		as, err := from.GetAttribute(a)
		if err != nil {
			panic(err)
		}
		ret[i] = as
	}
	return ret
}

func buildAttrSet(a []Attribute) map[Attribute]bool {
	ret := make(map[Attribute]bool)
	for _, a := range a {
		ret[a] = true
	}
	return ret
}

// AttributeIntersect returns the intersection of two Attribute slices.
//
// IMPORTANT: result is ordered in order of the first []Attribute argument.
//
// IMPORTANT: result contains only Attributes from a1.
func AttributeIntersect(a1, a2 []Attribute) []Attribute {
	ret := make([]Attribute, 0)
	for _, a := range a1 {
		matched := false
		for _, b := range a2 {
			if a.Equals(b) {
				matched = true
				break
			}
		}
		if matched {
			ret = append(ret, a)
		}
	}
	return ret
}

// AttributeIntersectReferences returns the intersection of two Attribute slices.
//
// IMPORTANT: result is not guaranteed to be ordered.
//
// IMPORTANT: done using pointers for speed, use AttributeDifference
// if the Attributes originate from different DataGrids.
func AttributeIntersectReferences(a1, a2 []Attribute) []Attribute {
	a1b := buildAttrSet(a1)
	a2b := buildAttrSet(a2)
	ret := make([]Attribute, 0)
	for a := range a1b {
		if _, ok := a2b[a]; ok {
			ret = append(ret, a)
		}
	}
	return ret
}

// AttributeDifference returns the difference between two Attribute
// slices: i.e. all the values in a1 which do not occur in a2.
//
// IMPORTANT: result is ordered the same as a1.
//
// IMPORTANT: result only contains values from a1.
func AttributeDifference(a1, a2 []Attribute) []Attribute {
	ret := make([]Attribute, 0)
	for _, a := range a1 {
		matched := false
		for _, b := range a2 {
			if a.Equals(b) {
				matched = true
				break
			}
		}
		if !matched {
			ret = append(ret, a)
		}
	}
	return ret
}

// AttributeDifferenceReferences returns the difference between two Attribute
// slices: i.e. all the values in a1 which do not occur in a2.
//
// IMPORTANT: result is not guaranteed to be ordered.
//
// IMPORTANT: done using pointers for speed, use AttributeDifference
// if the Attributes originate from different DataGrids.
func AttributeDifferenceReferences(a1, a2 []Attribute) []Attribute {
	a1b := buildAttrSet(a1)
	a2b := buildAttrSet(a2)
	ret := make([]Attribute, 0)
	for a := range a1b {
		if _, ok := a2b[a]; !ok {
			ret = append(ret, a)
		}
	}
	return ret
}
