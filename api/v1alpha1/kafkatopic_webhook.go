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

package v1alpha1

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var kafkatopiclog = logf.Log.WithName("kafkatopic-resource")

func (r *KafkaTopic) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-kafka-haase-de-v1alpha1-kafkatopic,mutating=false,failurePolicy=fail,groups=kafka.haase.de,resources=kafkatopics,versions=v1alpha1,name=vkafkatopic.kb.io

var _ webhook.Validator = &KafkaTopic{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *KafkaTopic) ValidateCreate() error {
	kafkatopiclog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *KafkaTopic) ValidateUpdate(old runtime.Object) error {
	kafkatopiclog.Info("validate update", "name", r.Name)

	oldTopic, ok := old.(*KafkaTopic)
	if !ok {
		return fmt.Errorf("unexpected type for old object")
	}

	if oldTopic.Spec.Replicas != r.Spec.Replicas {
		return fmt.Errorf("replicas is read only")
	}

	if oldTopic.Spec.Partitions > r.Spec.Partitions {
		return fmt.Errorf("can't decrease partition count")
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *KafkaTopic) ValidateDelete() error {
	kafkatopiclog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
