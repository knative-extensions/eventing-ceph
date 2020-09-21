/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"testing"
)

func TestCephSource(t *testing.T) {
	testCases := map[string]struct {
		source CephSource
	}{
		"type": {
			source: CephSource{}},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			if kind := tc.source.GetGroupVersionKind(); kind.String() != "sources.knative.dev/v1alpha1, Kind=CephSource" {
				t.Fatalf("Invalid group kind: %s", kind)
			}
			if status := tc.source.GetStatus(); status.GetCondition(CephConditionReady) != nil {
				t.Fatalf("Condition should be nil")
			}
		})
	}
}
