package services

import (
	"context"
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
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
	return getSvcIp(ctx, svcName)
}

func AllocatePublicIp(ctx context.Context, options *AllocatePublicIpOptions) (*AllocatePublicIpResults, error) {
	svcName := fmt.Sprintf("svc-allocation-id-%s", options.AllocationId)
	svcManifest := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svcName,
			Namespace: K8S_NS,
		},
		Spec: apiv1.ServiceSpec{
			Type:     apiv1.ServiceTypeLoadBalancer,
			Selector: map[string]string{"fakeprovider/ip_allocation": options.AllocationId},
			Ports: []apiv1.ServicePort{
				{
					Name:       "ssh-port",
					Protocol:   "TCP",
					Port:       22,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 22},
				},
				{
					Name:       "kube",
					Protocol:   "TCP",
					Port:       6443,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 6443},
				},
			},
		},
	}
	_, err := client.CoreV1().Services(K8S_NS).Create(ctx, svcManifest, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not create svc for allocation %s, because:%s", options.AllocationId, err)
	}
	time.Sleep(5 * time.Second)
	ip, err := getSvcIp(ctx, svcName)
	if err != nil {
		return nil, err
	}
	return &AllocatePublicIpResults{Ip: ip, AllocationId: options.AllocationId}, nil
}

func getSvcIp(ctx context.Context, svcName string) (string, error) {
	svc, err := client.CoreV1().Services(K8S_NS).Get(ctx, svcName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("SVC not found: %s", err)
	}
	ingresses := svc.Status.LoadBalancer.Ingress
	if len(ingresses) > 0 {
		return ingresses[0].IP, nil
	}
	return "", fmt.Errorf("SVC:%s found with no ingress attached", svc.Spec.ExternalName)
}

type AllocatePublicIpOptions struct {
	AllocationId string
}
type AllocatePublicIpResults struct {
	Ip           string
	AllocationId string
}
