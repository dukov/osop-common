package k8s

import (
	corev1 "k8s.io/api/core/v1"
)

// Container abstraction
type Container struct {
	Obj *corev1.Container
}

// NewContainer create new Container object
func NewContainer(name, image string, command []string) Container {
	return Container{
		Obj: &corev1.Container{
			Name:    name,
			Image:   image,
			Command: command,
		},
	}

}

// AddVolume extends Container volume mouns with a Volume
func (c Container) AddVolume(vol corev1.VolumeMount) {
	c.Obj.VolumeMounts = append(c.Obj.VolumeMounts, vol)
}
