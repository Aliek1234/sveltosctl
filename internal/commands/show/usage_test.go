/*
Copyright 2022. projectsveltos.io. All rights reserved.

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

package show_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2/textlogger"

	configv1beta1 "github.com/projectsveltos/addon-controller/api/v1beta1"
	libsveltosv1beta1 "github.com/projectsveltos/libsveltos/api/v1beta1"
	"github.com/smartbronchit/sveltosctl/internal/commands/show"
	"github.com/smartbronchit/sveltosctl/internal/utils"
)

var _ = Describe("Usage", func() {
	It("showUsage displays per resource, associated list of CAPI clusters", func() {
		configMap := configv1beta1.PolicyRef{
			Namespace: randomString(),
			Name:      randomString(),
			Kind:      string(libsveltosv1beta1.ConfigMapReferencedResourceKind),
		}

		secret := configv1beta1.PolicyRef{
			Namespace: randomString(),
			Name:      randomString(),
			Kind:      string(libsveltosv1beta1.SecretReferencedResourceKind),
		}

		clusterProfile1 := generateClusterProfile()
		clusterProfile1.Spec.PolicyRefs = []configv1beta1.PolicyRef{
			configMap,
		}
		clusterProfile1.Status.MatchingClusterRefs = []corev1.ObjectReference{
			{Namespace: randomString(), Name: randomString()},
		}

		clusterProfile2 := generateClusterProfile()
		clusterProfile2.Spec.PolicyRefs = []configv1beta1.PolicyRef{
			secret,
		}
		clusterProfile2.Status.MatchingClusterRefs = []corev1.ObjectReference{
			{Namespace: randomString(), Name: randomString()},
		}

		old := os.Stdout // keep backup of the real stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		initObjects := []client.Object{clusterProfile1, clusterProfile2}
		scheme, err := utils.GetScheme()
		Expect(err).To(BeNil())
		c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(initObjects...).Build()

		utils.InitalizeManagementClusterAcces(scheme, nil, nil, c)
		err = show.ShowUsage(context.TODO(), "", "", "",
			textlogger.NewLogger(textlogger.NewConfig(textlogger.Verbosity(1))))
		Expect(err).To(BeNil())

		w.Close()
		var buf bytes.Buffer
		_, err = io.Copy(&buf, r)
		Expect(err).To(BeNil())

		/*
			// Example of expected output
			`+----------------+--------------------+---------------+-----------------------+
			| RESOURCE KIND  | RESOURCE NAMESPACE | RESOURCE NAME |       CLUSTERS        |
			+----------------+--------------------+---------------+-----------------------+
			| ClusterProfile |                    | gauuu53n7r    | hme095dqji/yads0fjhoj |
			| ClusterProfile |                    | qa8kxyhq9e    | p1d3rlx2sx/5trz9p06tk |
			| ConfigMap      | gkxc9niba3         | o5fafy6bnn    | hme095dqji/yads0fjhoj |
			| Secret         | 224c2ibzhz         | qkspgp7vp1    | p1d3rlx2sx/5trz9p06tk |
			+----------------+--------------------+---------------+-----------------------+`
		*/

		lines := strings.Split(buf.String(), "\n")
		verifyClusterProfileUsage(lines, clusterProfile1)
		verifyClusterProfileUsage(lines, clusterProfile2)
		verifyUsage(lines, string(libsveltosv1beta1.ConfigMapReferencedResourceKind),
			configMap.Namespace, configMap.Name, &clusterProfile1.Status.MatchingClusterRefs[0])
		verifyUsage(lines, string(libsveltosv1beta1.SecretReferencedResourceKind),
			secret.Namespace, secret.Name, &clusterProfile2.Status.MatchingClusterRefs[0])
		os.Stdout = old
	})
})

func verifyClusterProfileUsage(lines []string, clusterProfile *configv1beta1.ClusterProfile) {
	for i := range clusterProfile.Status.MatchingClusterRefs {
		verifyUsage(lines, configv1beta1.ClusterProfileKind, "", clusterProfile.Name,
			&clusterProfile.Status.MatchingClusterRefs[i])
	}
}

func verifyUsage(lines []string, kind, namespace, name string, matchingCluster *corev1.ObjectReference) {
	clusterInfo := fmt.Sprintf("%s/%s", matchingCluster.Namespace, matchingCluster.Name)
	found := false
	for i := range lines {
		if strings.Contains(lines[i], kind) &&
			strings.Contains(lines[i], namespace) &&
			strings.Contains(lines[i], name) &&
			strings.Contains(lines[i], clusterInfo) {

			found = true
		}
	}
	Expect(found).To(BeTrue())
}

func generateClusterProfile() *configv1beta1.ClusterProfile {
	return &configv1beta1.ClusterProfile{
		ObjectMeta: metav1.ObjectMeta{
			Name: randomString(),
		},
		Spec: configv1beta1.Spec{
			ClusterSelector: libsveltosv1beta1.Selector{
				LabelSelector: metav1.LabelSelector{
					MatchLabels: map[string]string{"zone": "west"},
				},
			},
			SyncMode: configv1beta1.SyncModeContinuous,
		},
	}
}
