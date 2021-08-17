package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type kubeClient struct {
	Config           *rest.Config
	Context          string
	Clientset        *kubernetes.Clientset
	Namespace        string
	enforceNamespace bool
}

func (c *kubeClient) connectClient(namespace string) {
	var err error

	c.Config, err = rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	c.Clientset, err = kubernetes.NewForConfig(c.Config)
	if err != nil {
		panic(err)
	}

	c.Namespace = &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
}

func (c *kubeClient) CreateNamespace(namespace string) {

	c.Clientset.CoreV1().Namespaces().Create(context.Background(), nsName, metav1.CreateOptions{})
	return
}

func (c *kubeClient) DeleteNamespace(ns string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	clientset.CoreV1().Namespaces().Delete(context.Background(), ns, metav1.DeleteOptions{})
	return
}
