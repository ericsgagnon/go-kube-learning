package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfigFilePath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	fmt.Println(kubeconfigFilePath)

	kubeconfigFile, err := ioutil.ReadFile(kubeconfigFilePath)
	if err != nil {
		log.Fatal(err)
	}

	kubeconfig, err := clientcmd.RESTConfigFromKubeConfig(kubeconfigFile)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		log.Fatalln(err)
	}

	versionInfo, err := clientset.ServerVersion()
	if err != nil {
		log.Fatalln(err)

	}
	fmt.Println(versionInfo.String())
	fmt.Println(versionInfo.Platform)
	csDiscovery, err := clientset.Discovery().OpenAPISchema()
	if err != nil {
		log.Fatalln(err)

	}

	fmt.Println(csDiscovery.Swagger)

}
