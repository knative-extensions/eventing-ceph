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
	"context"
	"testing"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

func ParseURL(u string, t *testing.T) (url *apis.URL) {
	var err error
	if url, err = apis.ParseURL(u); err == nil {
		return
	}
	t.Fatalf("Failed to parse URL: %s", err.Error())
	return
}

func TestCephSourceValidate(t *testing.T) {
	testCases := map[string]struct {
		source CephSource
	}{
		"validate ok": {
			source: CephSource{Spec: CephSourceSpec{
				ServiceAccountName: "default",
				Port:               "9999",
				SourceSpec: duckv1.SourceSpec{
					Sink: duckv1.Destination{URI: ParseURL("http://hello.world", t)},
				},
			},
			},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			if err := tc.source.Validate(context.TODO()); err != nil {
				t.Fatalf("Source validation should succeed")
			}
		})
	}
}

func TestCephSourceValidateFail(t *testing.T) {
	testCases := map[string]struct {
		source CephSource
	}{
		"missing port": {
			source: CephSource{Spec: CephSourceSpec{
				ServiceAccountName: "default",
				SourceSpec: duckv1.SourceSpec{
					Sink: duckv1.Destination{URI: ParseURL("http://hello.world", t)},
				},
			},
			},
		},
		"invalid port number": {
			source: CephSource{Spec: CephSourceSpec{
				ServiceAccountName: "default",
				Port:               "12345678",
				SourceSpec: duckv1.SourceSpec{
					Sink: duckv1.Destination{URI: ParseURL("http://hello.world", t)},
				},
			},
			},
		},
		"missing sink": {
			source: CephSource{Spec: CephSourceSpec{
				ServiceAccountName: "default",
				Port:               "9999",
				SourceSpec:         duckv1.SourceSpec{},
			},
			},
		},
		"missing service": {
			source: CephSource{Spec: CephSourceSpec{
				Port: "9999",
				SourceSpec: duckv1.SourceSpec{
					Sink: duckv1.Destination{URI: ParseURL("http://hello.world", t)},
				},
			},
			},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			if err := tc.source.Validate(context.TODO()); err == nil {
				t.Fatalf("Source validation should fail")
			}
		})
	}
}
