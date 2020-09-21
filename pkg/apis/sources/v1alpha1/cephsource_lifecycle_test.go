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

func TestCephSourceLifecycle(t *testing.T) {
	testCases := map[string]struct {
		source CephSource
	}{
		"init status": {
			source: CephSource{},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			tc.source.Status.InitializeConditions()
			if cond := tc.source.Status.GetCondition(CephConditionReady).Status; cond != "Unknown" {
				t.Fatalf("Unexpected ready condition: %s", cond)
			}
			if cond := tc.source.Status.GetCondition(CephConditionDeployed).Status; cond != "Unknown" {
				t.Fatalf("Unexpected deployed condition: %s", cond)
			}
			if cond := tc.source.Status.GetCondition(CephConditionSinkProvided).Status; cond != "Unknown" {
				t.Fatalf("Unexpected sink condition: %s", cond)
			}
			tc.source.Status.MarkSink(ParseURL("http://hello.world", t))
			if cond := tc.source.Status.GetCondition(CephConditionSinkProvided).Status; cond != "True" {
				t.Fatalf("unexpected sink condition: %s", cond)
			}
			tc.source.Status.MarkNoSink("no good reason", "%s", "just testing")
			if cond := tc.source.Status.GetCondition(CephConditionSinkProvided).Status; cond != "False" {
				t.Fatalf("Unexpected sink condition: %s", cond)
			}
		})
	}
}
