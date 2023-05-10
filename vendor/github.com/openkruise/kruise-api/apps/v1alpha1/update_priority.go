/*
Copyright 2020 The Kruise Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UpdatePriorityStrategy is the strategy to define priority for pods update.
// Only one of orderPriority and weightPriority can be set.
type UpdatePriorityStrategy struct {
	// Order priority terms, pods will be sorted by the value of orderedKey.
	// For example:
	// ```
	// orderPriority:
	// - orderedKey: key1
	// - orderedKey: key2
	// ```
	// First, all pods which have key1 in labels will be sorted by the value of key1.
	// Then, the left pods which have no key1 but have key2 in labels will be sorted by
	// the value of key2 and put behind those pods have key1.
	OrderPriority []UpdatePriorityOrderTerm `json:"orderPriority,omitempty"`
	// Weight priority terms, pods will be sorted by the sum of all terms weight.
	WeightPriority []UpdatePriorityWeightTerm `json:"weightPriority,omitempty"`
}

// UpdatePriorityOrder defines order priority.
type UpdatePriorityOrderTerm struct {
	// Calculate priority by value of this key.
	// Values of this key, will be sorted by GetInt(val). GetInt method will find the last int in value,
	// such as getting 5 in value '5', getting 10 in value 'sts-10'.
	OrderedKey string `json:"orderedKey"`
}

// UpdatePriorityWeightTerm defines weight priority.
type UpdatePriorityWeightTerm struct {
	// Weight associated with matching the corresponding matchExpressions, in the range 1-100.
	Weight int32 `json:"weight"`
	// MatchSelector is used to select by pod's labels.
	MatchSelector metav1.LabelSelector `json:"matchSelector"`
}
