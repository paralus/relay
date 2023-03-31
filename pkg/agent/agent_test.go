package agent

import (
	"context"
	"errors"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestGetNamespace(t *testing.T) {
	tt := []struct {
		name     string
		fakeNS   *v1.Namespace
		nsName   string
		expectNS *v1.Namespace
		err      error
	}{
		{
			name: "valid/namespace_present",
			fakeNS: &v1.Namespace{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Namespace",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "paralus-system",
				},
			},
			nsName: "paralus-system",
			expectNS: &v1.Namespace{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Namespace",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "paralus-system",
				},
			},
			err: nil,
		},
		{
			name: "valid/namespace_absent",
			fakeNS: &v1.Namespace{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Namespace",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "some-namespace",
				},
			},
			nsName:   "paralus-system",
			expectNS: nil,
			err:      errors.New(`namespaces "paralus-system" not found`),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			fclient := fake.NewClientBuilder().
				WithScheme(scheme.Scheme).
				WithObjects(tc.fakeNS).
				Build()

			got, err := getNamespace(ctx, fclient, tc.nsName)
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("got %v, expects %v", err, tc.err)
			}
			if err == nil && got.Name != tc.expectNS.Name {
				t.Errorf("got %v, expects %v", got, tc.expectNS)
			}
		})
	}
}
