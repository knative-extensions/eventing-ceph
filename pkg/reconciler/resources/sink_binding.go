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

package resources

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/eventing/pkg/apis/sources/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
	"knative.dev/pkg/tracker"
)

func SinkBindingName(source, subject string) string {
	return kmeta.ChildName(fmt.Sprintf("%s-%s", source, subject), "-sinkbinding")
}

func MakeSinkBinding(owner kmeta.OwnerRefable, source duckv1.SourceSpec, subject tracker.Reference) *v1.SinkBinding {
	sb := &v1.SinkBinding{
		ObjectMeta: metav1.ObjectMeta{
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(owner),
			},
			Name:      SinkBindingName(owner.GetObjectMeta().GetName(), subject.Name),
			Namespace: owner.GetObjectMeta().GetNamespace(),
		},
		Spec: v1.SinkBindingSpec{
			SourceSpec: source,
			BindingSpec: duckv1.BindingSpec{
				Subject: subject,
			},
		},
	}

	sb.SetDefaults(context.Background())
	return sb
}
