package services

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var client *kubernetes.Clientset

const K8S_NS = "runspace"

func InitiateK8SClient() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	return nil
}

func GetMachineIp(ctx context.Context, Machine_id string) (string, error) {
	if client == nil {
		return "", fmt.Errorf("K8S client was not initiated")
	}
	svcName := fmt.Sprintf("svc-%s", Machine_id)
	svc, err := client.CoreV1().Services("runspace").Get(ctx, svcName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("SVC not found: %s", err)
	}
	ingresses := svc.Status.LoadBalancer.Ingress
	if len(ingresses) > 0 {
		return ingresses[0].IP, nil
	}
	return "", fmt.Errorf("SVC:%s found with no ingress attached", svc.Spec.ExternalName)
}
