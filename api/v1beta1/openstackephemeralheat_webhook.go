/*


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

// Generated by:
//
// operator-sdk create webhook --group osp-director --version v1beta1 --kind OpenStackEphemeralHeat --defaulting
//

package v1beta1

import (
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var openstackephemeralheatlog = logf.Log.WithName("openstackephemeralheat-resource")

// SetupWebhookWithManager -
func (r *OpenStackEphemeralHeat) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-osp-director-openstack-org-v1beta1-openstackephemeralheat,mutating=true,failurePolicy=fail,sideEffects=None,groups=osp-director.openstack.org,resources=openstackephemeralheats,verbs=create;update,versions=v1beta1,name=mopenstackephemeralheat.kb.io

var _ webhook.Defaulter = &OpenStackEphemeralHeat{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *OpenStackEphemeralHeat) Default() {
	openstackephemeralheatlog.Info("default", "name", r.Name)

	//FIXME provide configurable defaults for the values below
	if r.Spec.HeatAPIImageURL == "" {
		r.Spec.HeatAPIImageURL = "quay.io/tripleomaster/openstack-heat-api:current-tripleo"
	}
	if r.Spec.HeatEngineImageURL == "" {
		r.Spec.HeatEngineImageURL = "quay.io/tripleomaster/openstack-heat-engine:current-tripleo"
	}
	if r.Spec.MariadbImageURL == "" {
		r.Spec.MariadbImageURL = "quay.io/tripleomaster/openstack-mariadb:current-tripleo"
	}
	if r.Spec.RabbitImageURL == "" {
		r.Spec.RabbitImageURL = "quay.io/tripleomaster/openstack-rabbitmq:current-tripleo"
	}
	if r.Spec.HeatEngineReplicas == 0 {
		r.Spec.HeatEngineReplicas = 3
	}

}
