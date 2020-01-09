package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	// _ "k8s.io/client-go/plugin/pkg/client/auth"

	"gopkg.in/yaml.v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Cluster holds current state of a kubenetes cluster and client
// type Cluster struct {
// 	CS *kubernetes.Clientset
// }

func main() {
	// Create kubernetes client
	cs, err := Client()
	if err != nil {
		log.Fatal(err)
	}

	sg, err := cs.Discovery().ServerGroups()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sg.Groups)
	fmt.Println(sg.SwaggerDoc())
	fmt.Println(sg)

	versionInfo, err := cs.ServerVersion()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(versionInfo.String())
	fmt.Println(versionInfo.Platform)

	output, err := yaml.Marshal(&sg)
	fmt.Println(string(output))
	// csDiscovery, err := cs.Discovery().OpenAPISchema()
	// if err != nil {
	// 	log.Fatalln(err)

	// }

	// fmt.Println(csDiscovery.Swagger)

	// ds, err := cs.AppsV1().DaemonSets("default").List(metav1.ListOptions{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(ds)

}

// Client returns a pointer to a kubernetes Clientset
func Client() (*kubernetes.Clientset, error) {

	// manually use ~/.kube/config for now
	kubeConfigFilePath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	fmt.Println(kubeConfigFilePath)

	kubeConfigFile, err := ioutil.ReadFile(kubeConfigFilePath)
	if err != nil {
		return nil, err
	}

	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeConfigFile)
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}
	return clientSet, err

}
