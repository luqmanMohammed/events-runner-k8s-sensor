package utils

import (
	"reflect"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

// GetKubeAPIConfig returns a Kubernetes API config.
// Common abstraction to get config in both incluster and out cluster
// scenarios.
func GetKubeAPIConfig(kubeConfigPath string) (*rest.Config, error) {
	if kubeConfigPath == "" {
		klog.V(3).Info("Provided KubeConfig path is empty. Getting config from home")
		if home := homedir.HomeDir(); home != "" {
			kubeConfigPath = home + "/.kube/config"
		}
	}
	return clientcmd.BuildConfigFromFlags("", kubeConfigPath)
}

// GetKubeAPIConfigOrDie wrapper around GetKubeAPIConfig.
// Panics if unable to load config.
func GetKubeAPIConfigOrDie(kubeConfigPath string) *rest.Config {
	config, err := GetKubeAPIConfig(kubeConfigPath)
	if err != nil {
		panic(err)
	}
	return config
}

// StringInSlice returns true if the string is in the slice.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Check if provided value is its zero value
func IsZero(value interface{}) bool {
	return value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

// Find 1st zero value in map of values and return its key else return empty string
func FindZeroValue(values map[string]interface{}) string {
	for k, v := range values {
		if IsZero(v) {
			return k
		}
	}
	return ""
}
