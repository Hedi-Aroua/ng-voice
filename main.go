package main

import (
  "context"
  "fmt"
  "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  //"k8s.io/client-go/tools/clientcmd"
  "k8s.io/client-go/rest"

)

func main() {
  // pass kubeconfig
  //config, error := clientcmd.BuildConfigFromFlags("", "/home/ng-voice/.kube/config")
  config, _ := rest.InClusterConfig()

  //if error != nil {
    //config, _ := rest.InClusterConfig()
    //fmt.Printf("working inside the cluster\n")
  //}
  // creates the clientset
  clientset, _ := kubernetes.NewForConfig(config)
  // API access
  pods, _ := clientset.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
  fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
  for _, pod := range pods.Items{
    fmt.Printf("%s \n", pod.Name)
  }
}