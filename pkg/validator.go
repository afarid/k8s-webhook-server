/*
Copyright 2018 The Kubernetes Authors.
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

package hook

import (
	"context"
	"encoding/json"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"log"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/validate-v1-pod,mutating=false,failurePolicy=fail,groups="",resources=pods,verbs=create;update,versions=v1,name=vpod.kb.io

// PodValidator validates Pods
type PodValidator struct {
	Client  client.Client
	decoder *admission.Decoder
	Debug   bool
}

// PodValidator admits a pod if a specific annotation exists.
func (v *PodValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	fmt.Println(v.Debug)
	var err error
	if v.Debug {
		r, err := json.Marshal(req)
		log.Printf("req: %s", r)
		if err != nil {
			log.Printf("err: %s", err)
		}
	}
	pod := &corev1.Pod{}

	err = v.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	key := "example-mutating-admission-webhook"
	anno, found := pod.Annotations[key]
	if !found {
		return admission.Denied(fmt.Sprintf("missing annotation %s", key))
	}
	if anno != "foo" {
		return admission.Denied(fmt.Sprintf("annotation %s did not have value %q", key, "foo"))
	}

	resp := admission.Allowed("")
	if v.Debug {
		r, err := json.Marshal(resp)
		log.Printf("res: %s", r)
		if err != nil {
			log.Printf("err: %s", err)
		}
	}
	return resp
}

// PodValidator implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (v *PodValidator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
