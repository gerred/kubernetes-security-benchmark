// Copyright © 2018 Jimmi Dyson <jimmidyson@gmail.com>
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

package node

import (
	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
	. "github.com/mesosphere/kubernetes-security-benchmark/pkg/ginkgo/matchers"
)

const kubeletProcessName = "kubelet"

func Kubelet(index, subIndex int, missingProcessFunc framework.MissingProcessHandlerFunc) {
	f := framework.New(kubeletProcessName, missingProcessFunc)
	BeforeEach(f.BeforeEach)

	It("[2.1.1] Ensure that the --allow-privileged argument is set to false [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--allow-privileged", "false"))
	})

	It("[2.1.2] Ensure that the --anonymous-auth argument is set to false [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--anonymous-auth", "false"))
	})

	It("[2.1.3] Ensure that the --authorization-mode argument is not set to AlwaysAllow [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithDifferentValue("--authorization-mode", "AlwaysAllow"))
	})

	It("[2.1.4] Ensure that the --client-ca-file argument is set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--client-ca-file"))
	})

	It("[2.1.5] Ensure that the --read-only-port argument is set to 0 [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--read-only-port", "0"))
	})

	It("[2.1.6] Ensure that the --streaming-connection-idle-timeout argument is not set to 0 [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithDifferentValue("--streaming-connection-idle-timeout", "0"))
	})

	It("[2.1.7] Ensure that the --protect-kernel-defaults argument is set to true [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithOptionalValue("--protect-kernel-defaults", "true"))
	})

	It("[2.1.8] Ensure that the --make-iptables-util-chains argument is set to true [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--make-iptables-util-chains", "true"))
	})

	It("[2.1.9] Ensure that the --keep-terminated-pod-volumes argument is set to false [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--keep-terminated-pod-volumes", "false"))
	})

	It("[2.1.10] Ensure that the --hostname-override argument is not set [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlag("--hostname-override"))
	})

	It("[2.1.11] Ensure that the --event-qps argument is set to 0 [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--event-qps", "0"))
	})

	It("[2.1.12] Ensure that the --tls-cert-file and --tls-private-key-file arguments are set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--tls-cert-file"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--tls-private-key-file"))
	})

	It("[2.1.13] Ensure that the --cadvisor-port argument is set to 0 [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--cadvisor-port", "0"))
	})

	It("[2.1.14] Ensure that the RotateKubeletClientCertificate argument is not set to false [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrNotContainValue("--feature-gates", "RotateKubeletClientCertificate=false"))
	})

	It("[2.1.16] Ensure that the RotateKubeletServerCertificate argument is set to true [Scored]", func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--feature-gates", "RotateKubeletServerCertificate=true"))
	})
}
