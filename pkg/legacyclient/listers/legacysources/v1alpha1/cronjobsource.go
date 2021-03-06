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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "knative.dev/eventing/pkg/apis/legacysources/v1alpha1"
)

// CronJobSourceLister helps list CronJobSources.
type CronJobSourceLister interface {
	// List lists all CronJobSources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CronJobSource, err error)
	// CronJobSources returns an object that can list and get CronJobSources.
	CronJobSources(namespace string) CronJobSourceNamespaceLister
	CronJobSourceListerExpansion
}

// cronJobSourceLister implements the CronJobSourceLister interface.
type cronJobSourceLister struct {
	indexer cache.Indexer
}

// NewCronJobSourceLister returns a new CronJobSourceLister.
func NewCronJobSourceLister(indexer cache.Indexer) CronJobSourceLister {
	return &cronJobSourceLister{indexer: indexer}
}

// List lists all CronJobSources in the indexer.
func (s *cronJobSourceLister) List(selector labels.Selector) (ret []*v1alpha1.CronJobSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CronJobSource))
	})
	return ret, err
}

// CronJobSources returns an object that can list and get CronJobSources.
func (s *cronJobSourceLister) CronJobSources(namespace string) CronJobSourceNamespaceLister {
	return cronJobSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CronJobSourceNamespaceLister helps list and get CronJobSources.
type CronJobSourceNamespaceLister interface {
	// List lists all CronJobSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CronJobSource, err error)
	// Get retrieves the CronJobSource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CronJobSource, error)
	CronJobSourceNamespaceListerExpansion
}

// cronJobSourceNamespaceLister implements the CronJobSourceNamespaceLister
// interface.
type cronJobSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CronJobSources in the indexer for a given namespace.
func (s cronJobSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CronJobSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CronJobSource))
	})
	return ret, err
}

// Get retrieves the CronJobSource from the indexer for a given namespace and name.
func (s cronJobSourceNamespaceLister) Get(name string) (*v1alpha1.CronJobSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cronjobsource"), name)
	}
	return obj.(*v1alpha1.CronJobSource), nil
}
