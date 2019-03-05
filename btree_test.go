// Copyright 2014 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package btree

import (
	"flag"
	"fmt"
	"math/rand"
//	"os"
//	"reflect"
//	"sort"
//	"strings"
//	"sync"
//	"testing"
	"time"
)

func init() {
	seed := time.Now().Unix()
	fmt.Println(seed)
	rand.Seed(seed)
}

// perm returns a random permutation of n Int items in the range [0, n).
func perm(n int) (out []Item) {
	for _, v := range rand.Perm(n) {
		out = append(out, Int(v))
	}
	return
}

// rang returns an ordered list of Int items in the range [0, n).
func rang(n int) (out []Item) {
	for i := 0; i < n; i++ {
		out = append(out, Int(i))
	}
	return
}

// all extracts all items from a tree in order as a slice.
func all(t *BTree) (out []Item) {
	t.Ascend(func(a Item) bool {
		out = append(out, a)
		return true
	})
	return
}

// rangerev returns a reversed ordered list of Int items in the range [0, n).
func rangrev(n int) (out []Item) {
	for i := n - 1; i >= 0; i-- {
		out = append(out, Int(i))
	}
	return
}

// allrev extracts all items from a tree in reverse order as a slice.
func allrev(t *BTree) (out []Item) {
	t.Descend(func(a Item) bool {
		out = append(out, a)
		return true
	})
	return
}

var btreeDegree = flag.Int("degree", 3, "B-Tree degree")

// Test goes here !





type byInts []Item

func (a byInts) Len() int {
	return len(a)
}

func (a byInts) Less(i, j int) bool {
	return a[i].(Int) < a[j].(Int)
}

func (a byInts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
