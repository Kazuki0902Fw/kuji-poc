package util

import (
	"golang.org/x/exp/constraints"
	goSlices "golang.org/x/exp/slices"
)

// Contains はスライスに指定した要素が含まれているかどうかを返す。
func Contains[T comparable](s []T, e T) bool {
	return goSlices.Contains(s, e)
}

// Map はスライスの各要素に対して関数を適用した結果を返す。
func Map[E, V any](elms []E, fn func(E) V) []V {
	// 空スライスと nil を厳密に区別
	if elms == nil {
		return nil
	}

	outputs := make([]V, len(elms), cap(elms))
	for i, elm := range elms {
		outputs[i] = fn(elm)
	}
	return outputs
}

// MapWithErr はスライスの各要素に対して関数を適用した結果を返す。
func MapWithErr[E, V any](elms []E, fn func(E) (V, error)) ([]V, error) {
	// 空スライスと nil を厳密に区別
	if elms == nil {
		return nil, nil
	}

	outputs := make([]V, len(elms), cap(elms))
	for i, elm := range elms {
		conv, err := fn(elm)
		if err != nil {
			return nil, err
		}
		outputs[i] = conv
	}
	return outputs, nil
}

// FindFunc はスライスの各要素に対して関数を適用した結果が true となる最初の要素を返す。
func FindFunc[E any](elems []E, fn func(E) bool) (e E, found bool) {
	i := goSlices.IndexFunc(elems, fn)
	if i < 0 {
		return
	}
	e = elems[i]
	found = true
	return
}

// FindFuncIndex はスライスの各要素に対して関数を適用した結果が true となる最初の要素のインデックスを返す。
func FindFuncIndex[E any](elems []E, fn func(E) bool) (index int, found bool) {
	i := goSlices.IndexFunc(elems, fn)
	if i < 0 {
		return
	}
	index = i
	found = true
	return
}

// CountFunc はスライスの各要素に対して関数を適用した結果が true となる要素の数を返す。
func CountFunc[E any](elems []E, fn func(E) bool) int {
	return len(FilterFunc(elems, fn))
}

// FilterFunc はスライスの各要素に対して関数を適用した結果が true となる要素のみを返す。
func FilterFunc[E any](elems []E, fn func(E) bool) []E {
	// 空スライスと nil を厳密に区別
	if elems == nil {
		return nil
	}

	n := make([]E, 0, len(elems))
	for _, elm := range elems {
		if fn(elm) {
			n = append(n, elm)
		}
	}
	return n
}

// Uniq はスライスの要素を重複なく並べ替える。
func Uniq[T constraints.Ordered](s []T) []T {
	goSlices.Sort(s)
	return goSlices.Clip(goSlices.Compact(s))
}

// Any はスライスの要素のうち、条件を満たすものがあるかどうかを返す。
func AnyFunc[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}
