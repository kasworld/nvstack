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

// slice based stack for fsm
package nvstack

import (
	// "errors"
	"sync"
)

type element struct {
	Name  string
	Value interface{}
}

type NVStack struct {
	stackData []element
	mutex     sync.Mutex
}

func New() *NVStack {
	return &NVStack{
		stackData: make([]element, 0),
	}
}

func (s NVStack) Len() int {
	return len(s.stackData)
}

func (s *NVStack) Push(name string, value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stackData = append(s.stackData, element{name, value})
}

func (s *NVStack) Pop() (string, interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.IsEmpty() {
		v := s.stackData[s.Len()-1]
		s.stackData = s.stackData[:s.Len()-1]
		return v.Name, v.Value
	}
	return "", nil
}

func (s *NVStack) Swap(name string, value interface{}) (string, interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.IsEmpty() {
		v := s.stackData[s.Len()-1]
		s.stackData[s.Len()-1] = element{name, value}
		return v.Name, v.Value
	} else {
		s.stackData = append(s.stackData, element{name, value})
		return "", nil
	}
}

func (s NVStack) Peek() (string, interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.IsEmpty() {
		v := s.stackData[s.Len()-1]
		return v.Name, v.Value
	}
	return "", nil
}

func (s NVStack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *NVStack) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stackData = nil
}
