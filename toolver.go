// Copyright © 2015 Erik Brady <brady@dvln.org> and Docker Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package toolver is for generic routines to compare tools versions
// for basic greater than, equal, less than class comparisons.
package toolver

import (
	"strconv"
	"strings"
)

// Version provides utility methods for comparing tool versions
type Version string

func (v Version) compareTo(other Version) int {
	var (
		currTab  = strings.Split(string(v), ".")
		otherTab = strings.Split(string(other), ".")
	)

	max := len(currTab)
	if len(otherTab) > max {
		max = len(otherTab)
	}
	for i := 0; i < max; i++ {
		var currInt, otherInt int

		if len(currTab) > i {
			currInt, _ = strconv.Atoi(currTab[i])
		}
		if len(otherTab) > i {
			otherInt, _ = strconv.Atoi(otherTab[i])
		}
		if currInt > otherInt {
			return 1
		}
		if otherInt > currInt {
			return -1
		}
	}
	return 0
}

// LessThan checks if a version is less than another version
func (v Version) LessThan(other Version) bool {
	return v.compareTo(other) == -1
}

// LessThanOrEqualTo checks if a version is less than or equal to another
func (v Version) LessThanOrEqualTo(other Version) bool {
	return v.compareTo(other) <= 0
}

// GreaterThan checks if a version is greater than another one
func (v Version) GreaterThan(other Version) bool {
	return v.compareTo(other) == 1
}

// GreaterThanOrEqualTo checks ia version is greater than or equal to another
func (v Version) GreaterThanOrEqualTo(other Version) bool {
	return v.compareTo(other) >= 0
}

// Equal checks if a version is equal to another
func (v Version) Equal(other Version) bool {
	return v.compareTo(other) == 0
}
