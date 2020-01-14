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

	sg, sr, err := cs.Discovery().ServerGroupsAndResources()
	if err != nil {
		log.Fatal(err)
	}

	output, err := yaml.Marshal(&sr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(output))

	for i := range sr {
		sriy, err := yaml.Marshal(&sr[i])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.
			Printf(
				"\n-------------------------------------------\n%d:\n%s",
				i,
				sriy,
			)

	}
	for i := range sg {
		sgiy, err := yaml.Marshal(&sg[i])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.
			Printf(
				"\n-------------------------------------------\n%d:\n%s",
				i,
				sgiy,
			)

	}

	fmt.Printf("Server Groups: %d\n", len(sg))
	fmt.Printf("Server Resources: %d\n", len(sr))
	// fmt.Println(yaml.Marshal(sr))

	// sg, err := cs.Discovery().ServerGroups()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(sg.Groups)
	// fmt.Println(sg.SwaggerDoc())
	// fmt.Println(sg)

	// versionInfo, err := cs.ServerVersion()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(versionInfo.String())
	// fmt.Println(versionInfo.Platform)

	// output, err := yaml.Marshal(&sg)
	// fmt.Println(string(output))
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
	// fmt.Println(kubeConfigFilePath)
	fmt.Printf("\n\nKubeConfig Path: %s\n\n", kubeConfigFilePath)

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
