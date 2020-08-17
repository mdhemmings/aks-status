package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// ListDeployments : Login to current cluster context and count/list the deployments in there
func ListDeployments() (clientset *kubernetes.Clientset) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deployments, err := clientset.ExtensionsV1beta1().Deployments("").List(metav1.ListOptions{})
	fmt.Printf("There are %d deployments in the cluster:\n", len(deployments.Items))
	//	fmt.Printf("They are %v\n", deployments.Items)
	for _, deployment := range deployments.Items {
		fmt.Printf("%+v\n", deployment.Name)
	}
	return
}

// ListPods : Login to current cluster context and count/list the pods in there
func ListPods() (clientset *kubernetes.Clientset) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	fmt.Printf("There are %d pods in the cluster:\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("%+v\n", pod.Name)
	}
	return
}

func main() {
	return
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
