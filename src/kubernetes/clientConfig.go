package kubernetes

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
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
	/*
		c.Namespace = &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		}
	*/
	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := c.Clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = c.Clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found example-xxxxx pod in default namespace\n")
		}

		time.Sleep(10 * time.Second)
	}

}

func (c *kubeClient) CreateNamespace(namespace string) {

	//c.Clientset.CoreV1().Namespaces().Create(context.Background(), nsName, metav1.CreateOptions{})
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
