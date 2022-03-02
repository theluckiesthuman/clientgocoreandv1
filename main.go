package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/kiran/.kube/config", "location to kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pods.Items {
		fmt.Println(p.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}
}
