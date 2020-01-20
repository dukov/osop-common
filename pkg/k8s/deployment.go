package k8s

import (
	k8sapps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Deployment abstraction
type Deployment struct {
	Obj *k8sapps.Deployment
}

// NewDeployment creates new Deployment object
func NewDeployment(name, namespace string, replicas *int32, labels map[string]string) Deployment {
	return Deployment{
		Obj: &k8sapps.Deployment{
			TypeMeta: metav1.TypeMeta{APIVersion: k8sapps.SchemeGroupVersion.String(), Kind: "Deployment"},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
				Labels:    labels,
			},
			Spec: k8sapps.DeploymentSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: labels,
				},
				Replicas: replicas,
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: labels,
					},
					Spec: corev1.PodSpec{},
				},
			},
		},
	}
}

// AddContainer extends Deployment pod with a Container
func (d Deployment) AddContainer(c Container) {
	d.Obj.Spec.Template.Spec.Containers = append(d.Obj.Spec.Template.Spec.Containers, *c.Obj)
}

// AddVolume extends Deployment pod with a Volume
func (d Deployment) AddVolume(v Volume) {
	d.Obj.Spec.Template.Spec.Volumes = append(d.Obj.Spec.Template.Spec.Volumes, *v.Obj)
}
