/*
Copyright 2020 The Flux authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/fluxcd/flagger/pkg/apis/flagger/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// CanaryLister helps list Canaries.
// All objects returned here must be treated as read-only.
type CanaryLister interface {
	// List lists all Canaries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Canary, err error)
	// Canaries returns an object that can list and get Canaries.
	Canaries(namespace string) CanaryNamespaceLister
	CanaryListerExpansion
}

// canaryLister implements the CanaryLister interface.
type canaryLister struct {
	listers.ResourceIndexer[*v1beta1.Canary]
}

// NewCanaryLister returns a new CanaryLister.
func NewCanaryLister(indexer cache.Indexer) CanaryLister {
	return &canaryLister{listers.New[*v1beta1.Canary](indexer, v1beta1.Resource("canary"))}
}

// Canaries returns an object that can list and get Canaries.
func (s *canaryLister) Canaries(namespace string) CanaryNamespaceLister {
	return canaryNamespaceLister{listers.NewNamespaced[*v1beta1.Canary](s.ResourceIndexer, namespace)}
}

// CanaryNamespaceLister helps list and get Canaries.
// All objects returned here must be treated as read-only.
type CanaryNamespaceLister interface {
	// List lists all Canaries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Canary, err error)
	// Get retrieves the Canary from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Canary, error)
	CanaryNamespaceListerExpansion
}

// canaryNamespaceLister implements the CanaryNamespaceLister
// interface.
type canaryNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.Canary]
}
