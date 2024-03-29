package client_go

import (
	"fmt"
	"github.com/ZhangSetSail/GoStudy/gin-demo/model"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	"sigs.k8s.io/gateway-api/apis/v1beta1"
	"strings"
)

func (m *ManagerClientGo) GetResourcesNameByNamespace(kind, namespace string) (*model.ResourceNamesResponse, error) {
	var names []string
	switch strings.ToLower(kind) {
	case "deployment":
		deploys, err := m.clientSet.AppsV1().Deployments(namespace).List(m.ctx, v1.ListOptions{})
		if err != nil {
			logrus.Errorf("get deployment list failure by namespace: %v", namespace)
			return nil, err
		}
		for _, deploy := range deploys.Items {
			names = append(names, deploy.GetName())
		}
	case "service":
		fmt.Println()
	case "statefulSet":
		fmt.Println()
	case "pod":
		pods, err := m.clientSet.CoreV1().Pods(namespace).List(m.ctx, v1.ListOptions{})
		if err != nil {
			logrus.Errorf("get pod list failure by namespace: %v", namespace)
			return nil, err
		}
		for _, pod := range pods.Items {
			names = append(names, pod.GetName())
		}
	case "configMap":
		fmt.Println()
	case "":
		fmt.Println()
	}
	return &model.ResourceNamesResponse{
		Names:     names,
		Namespace: namespace,
		Kind:      kind,
	}, nil
}

func (m *ManagerClientGo) CreateGateway(namespace string) {
	v1alpha2.TCPRoute{}
	v1alpha2.UDPRoute{}
	v1alpha2.GRPCRoute{}
	v1alpha2.TLSRoute{}

	v1alpha2.HTTPRouteRule{}
	v1beta1.HTTPRoute{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1beta1.HTTPRouteSpec{},
		Status:     v1beta1.HTTPRouteStatus{},
	}
	v1beta1.Gateway{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1beta1.GatewaySpec{},
		Status:     v1beta1.GatewayStatus{},
	}
	v1beta1.GatewayClass{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1beta1.GatewayClassSpec{},
		Status:     v1beta1.GatewayClassStatus{},
	}
	gateway := &v1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-gateway",
			Namespace: "default",
		},
		Spec: v1beta1.GatewaySpec{
			GatewayClassName: "test",
			Listeners: []v1beta1.Listener{
				{
					Name:          "",
					Hostname:      nil,
					Port:          0,
					Protocol:      "",
					TLS:           nil,
					AllowedRoutes: nil,
				},
			},
			Addresses:
		},
	}
	m.gatewayClient.GatewayClasses().Create()
	m.gatewayClient.Gateways(namespace).Create(m.ctx, gateway, metav1.CreateOptions{})
}

func  (m *ManagerClientGo) CreateHttpRoute(namespace string)  {
	m.gatewayClient.GatewayClasses().Create(v1beta1.GatewayClass{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1beta1.GatewayClassSpec{},
		Status:     v1beta1.GatewayClassStatus{},
	})
	m.gatewayClient.HTTPRoutes(namespace).Create(m.ctx,&v1beta1.HTTPRoute{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1beta1.HTTPRouteSpec{
			CommonRouteSpec: v1beta1.CommonRouteSpec{
				ParentRefs: []v1beta1.ParentReference{
					{
						Group:
					},
				},
			},
			Hostnames:       nil,
			Rules:           nil,
		},
	}, metav1.CreateOptions{})
}