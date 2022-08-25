package main

import (
  "context"
  "fmt"
  "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  //"k8s.io/client-go/tools/clientcmd"
  "k8s.io/client-go/rest"
  "time"

)

func main() {
  // pass kubeconfig
  //config, error := clientcmd.BuildConfigFromFlags("", "/home/ng-voice/.kube/config")
  config, _ := rest.InClusterConfig()
  // creates the clientset
  clientset, _ := kubernetes.NewForConfig(config)
  // API access
  for {
    pods, _ := clientset.CoreV1().Pods("a").List(context.Background(), v1.ListOptions{})
    fmt.Printf("There are %d pods in the namespace a\n", len(pods.Items))
    for _, pod := range pods.Items{
      fmt.Printf("%s \r\n", pod.Name)
    }
    fmt.Printf("  \n ****** \n ")
    time.Sleep(3*time.Second)
  }
}
