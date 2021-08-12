// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/jenkins-x/lighthouse/pkg/apis/lighthouse/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LighthouseBreakpointLister helps list LighthouseBreakpoints.
// All objects returned here must be treated as read-only.
type LighthouseBreakpointLister interface {
	// List lists all LighthouseBreakpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LighthouseBreakpoint, err error)
	// LighthouseBreakpoints returns an object that can list and get LighthouseBreakpoints.
	LighthouseBreakpoints(namespace string) LighthouseBreakpointNamespaceLister
	LighthouseBreakpointListerExpansion
}

// lighthouseBreakpointLister implements the LighthouseBreakpointLister interface.
type lighthouseBreakpointLister struct {
	indexer cache.Indexer
}

// NewLighthouseBreakpointLister returns a new LighthouseBreakpointLister.
func NewLighthouseBreakpointLister(indexer cache.Indexer) LighthouseBreakpointLister {
	return &lighthouseBreakpointLister{indexer: indexer}
}

// List lists all LighthouseBreakpoints in the indexer.
func (s *lighthouseBreakpointLister) List(selector labels.Selector) (ret []*v1alpha1.LighthouseBreakpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LighthouseBreakpoint))
	})
	return ret, err
}

// LighthouseBreakpoints returns an object that can list and get LighthouseBreakpoints.
func (s *lighthouseBreakpointLister) LighthouseBreakpoints(namespace string) LighthouseBreakpointNamespaceLister {
	return lighthouseBreakpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LighthouseBreakpointNamespaceLister helps list and get LighthouseBreakpoints.
// All objects returned here must be treated as read-only.
type LighthouseBreakpointNamespaceLister interface {
	// List lists all LighthouseBreakpoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LighthouseBreakpoint, err error)
	// Get retrieves the LighthouseBreakpoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LighthouseBreakpoint, error)
	LighthouseBreakpointNamespaceListerExpansion
}

// lighthouseBreakpointNamespaceLister implements the LighthouseBreakpointNamespaceLister
// interface.
type lighthouseBreakpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LighthouseBreakpoints in the indexer for a given namespace.
func (s lighthouseBreakpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LighthouseBreakpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LighthouseBreakpoint))
	})
	return ret, err
}

// Get retrieves the LighthouseBreakpoint from the indexer for a given namespace and name.
func (s lighthouseBreakpointNamespaceLister) Get(name string) (*v1alpha1.LighthouseBreakpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lighthousebreakpoint"), name)
	}
	return obj.(*v1alpha1.LighthouseBreakpoint), nil
}
