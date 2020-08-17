package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// ListDeployments : Login to current cluster context and count/list the deployments in there
func ListDeployments(cluster string) (clientset *kubernetes.Clientset, err error) {
	config, _ := buildConfig(cluster)
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deployments, err := clientset.ExtensionsV1beta1().Deployments("").List(metav1.ListOptions{})
	fmt.Printf("There are %d deployments in the cluster %v:\n\n", len(deployments.Items), cluster)
	for _, deployment := range deployments.Items {
		fmt.Printf("%+v\n", deployment.Name)
	}
	return
}

// ListPods : Login to current cluster context and count/list the pods in there
func ListPods(cluster string) (clientset *kubernetes.Clientset, err error) {
	config, err := buildConfig(cluster)
	if err != nil {
		panic(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	fmt.Printf("There are %d pods in the cluster %v:\n\n", len(pods.Items), cluster)
	for _, pod := range pods.Items {
		fmt.Printf("%v \t %v\n", pod.Name, pod.Status.Phase)
	}
	return
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func buildConfig(cluster string) (config *rest.Config, err error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: *kubeconfig},
		&clientcmd.ConfigOverrides{
			CurrentContext: cluster,
		}).ClientConfig()

	return
}

func main() {
	if os.Args[1] == "pods" {
		ListPods(os.Args[2])
	}
	if os.Args[1] == "deployments" {
		ListDeployments(os.Args[2])
	}
}
