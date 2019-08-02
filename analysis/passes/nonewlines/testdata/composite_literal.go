// Copyright © 2019 Banzai Cloud
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

package testdata

func CompositeLiteral() {
	type a struct {
		a string
	}

	type b struct {
		b string
	}

	_ = a{ // want "composite literals should not start or end with empty lines"

		a: "a",
	}

	// @formatter:off
	_ = a{ // want "composite literals should not start or end with empty lines"
		a: "a",

	}

	_ = a{ // want "composite literals should not start or end with empty lines"

		a: "a",

	}
	// @formatter:on

	_ = a{
		a: "a",
	}

	_ = a{
	}

	_ = a{

	}

	_ = a{a: "a"}

	_ = map[string]string{
		"key": "value",
	}

	_ = map[string]string{"key": "value"}
}
