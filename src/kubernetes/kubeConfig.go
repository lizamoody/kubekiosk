package main

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

//access cluster from Kubernetes API
func main() {
	//loading rules
	//running from container
	//pulls deployment and runs container
	//
	/*
		rules := clientcmd.NewDefaultClientConfigLoadingRules()
		kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
		kconfig, err := kubeconfig.ClientConfig()
	*/

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
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
	/*
		clientset := kubernetes.NewForConfigOrDie(config)
		nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}

		for _, n := range nodeList.Items {
			fmt.Println(n.Name)
		}

		for {
			//access API to get pods
			pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		}
		//uses current context
		//need to figure out pathpods... aren't local
		//config, _ := clientcmd.BuildConfigFromFlags("", "path_here")
	*/
	//create client set
	/*clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		//access API to get pods
		pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	}*/

}

func CreateNamespace(namespace string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	nsName := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-new-namespace",
		},
	}
	clientset.CoreV1().Namespaces().Create(context.Background(), nsName, metav1.CreateOptions{})
	return
}

/*
func (c *Client) DeleteNamespace(namespace string) error {
	return c.Clientset.v1().Namespaces().Delete(namespace, &metav1.DeleteOptions{})
}

func (c *Client) getUserNamespace(user string) error {
	return
}
*/
