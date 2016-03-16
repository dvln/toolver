// Copyright © 2015-2016 Erik Brady <brady@dvln.org> and Docker Inc
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

package toolver

import (
	"testing"
)

func assertVersion(t *testing.T, a, b string, result int) {
	if r := Version(a).compareTo(Version(b)); r != result {
		t.Fatalf("Unexpected version comparison result. Found %d, expected %d", r, result)
	}
}

func TestCompareVersion(t *testing.T) {
	assertVersion(t, "1.12", "1.12", 0)
	assertVersion(t, "1.0.0", "1", 0)
	assertVersion(t, "1", "1.0.0", 0)
	assertVersion(t, "1.05.00.0156", "1.0.221.9289", 1)
	assertVersion(t, "1", "1.0.1", -1)
	assertVersion(t, "1.0.1", "1", 1)
	assertVersion(t, "1.0.1", "1.0.2", -1)
	assertVersion(t, "1.0.2", "1.0.3", -1)
	assertVersion(t, "1.0.3", "1.1", -1)
	assertVersion(t, "1.1", "1.1.1", -1)
	assertVersion(t, "1.1.1", "1.1.2", -1)
	assertVersion(t, "1.1.2", "1.2", -1)
}

// Very simple test to see if we get executable info back as expected
func TestExecutableInfo(t *testing.T) {
	name, buildDate, err := ExecutableInfo()
	if name == "" {
		t.Fatal("ExecutableInfo() returned empty tool name, that's not good")
	}
	if buildDate == "" {
		t.Fatal("ExecutableInfo() returned empty build date, that's not good")
	}
	if err != nil {
		t.Fatalf("ExecutableInfo() returned an error and shouldn't have, err:%s\n", err)
	}
}
