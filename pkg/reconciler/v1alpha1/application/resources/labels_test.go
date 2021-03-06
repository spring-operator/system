/*
Copyright 2018 The Knative Authors

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
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/projectriff/system/pkg/apis/projectriff"
	projectriffv1alpha1 "github.com/projectriff/system/pkg/apis/projectriff/v1alpha1"
)

func TestMakeLabels(t *testing.T) {
	tests := []struct {
		name string
		app  *projectriffv1alpha1.Application
		want map[string]string
	}{{
		name: "just application name",
		app: &projectriffv1alpha1.Application{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "foo",
				Name:      "bar",
			},
		},
		want: map[string]string{
			projectriff.ApplicationLabelKey: "bar",
		},
	}, {
		name: "pass through labels",
		app: &projectriffv1alpha1.Application{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "baz",
				Name:      "blah",
				Labels: map[string]string{
					"asdf": "bazinga",
					"ooga": "booga",
				},
			},
		},
		want: map[string]string{
			projectriff.ApplicationLabelKey: "blah",
			"asdf":                          "bazinga",
			"ooga":                          "booga",
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := makeLabels(test.app)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("makeLabels (-want, +got) = %v", diff)
			}
		})
	}
}
