package k8s

import (
	corev1 "k8s.io/api/core/v1"
)

// Volume abstraction
type Volume struct {
	Obj corev1.Volume
}

// NewVolume creates ConfigMap based volume object
func NewVolume(name, objRefName string) Volume {
	return Volume{
		Obj: corev1.Volume{
			Name: name,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: objRefName,
					},
				},
			},
		},
	}
}
