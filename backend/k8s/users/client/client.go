package client

import (
	"github.com/okteto/app/backend/k8s/users/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// UserV1Alpha1Interface TBD
type UserV1Alpha1Interface interface {
	Users(namespace string) UserInterface
}

// UserV1Alpha1Client TBD
type UserV1Alpha1Client struct {
	restClient rest.Interface
}

//Get returns the k8s client
func Get() (*UserV1Alpha1Client, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return NewForConfig(config)
}

// NewForConfig TBD
func NewForConfig(c *rest.Config) (*UserV1Alpha1Client, error) {
	v1alpha1.AddToScheme(scheme.Scheme)

	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &UserV1Alpha1Client{restClient: client}, nil
}

// Users TBD
func (c *UserV1Alpha1Client) Users(namespace string) UserInterface {
	return &userClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}