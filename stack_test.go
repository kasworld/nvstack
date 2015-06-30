// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// slice based stack
package nvstack

import (
	"testing"
)

func TestPeek(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Error("not empty")
	}
	v := "hello"
	s.Push(v, nil)
	if s.Len() != 1 {
		t.Error("push fail")
	}
	if n, _ := s.Peek(); n != v {
		t.Error("peek fali")
	}

	if n, _ := s.Pop(); n != v {
		t.Error("pop fali")
	}

	if !s.IsEmpty() {
		t.Error("pop fail")
	}

	s.Push(v, nil)
	s.Reset()
	if !s.IsEmpty() {
		t.Error("reset fail")
	}
}
