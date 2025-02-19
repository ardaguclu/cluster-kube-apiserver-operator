package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Build configures the behavior of OpenShift builds for the entire cluster.
// This includes default settings that can be overridden in BuildConfig objects, and overrides which are applied to all builds.
//
// The canonical name is "cluster"
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
// +openshift:api-approved.openshift.io=https://github.com/openshift/api/pull/470
// +openshift:file-pattern=cvoRunLevel=0000_10,operatorName=openshift-controller-manager,operatorOrdering=01
// +openshift:capability=Build
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=builds,scope=Cluster
// +kubebuilder:subresource:status
type Build struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user-settable values for the build controller configuration
	// +required
	Spec BuildSpec `json:"spec"`
}

type BuildSpec struct {
	// additionalTrustedCA is a reference to a ConfigMap containing additional CAs that
	// should be trusted for image pushes and pulls during builds.
	// The namespace for this config map is openshift-config.
	//
	// DEPRECATED: Additional CAs for image pull and push should be set on
	// image.config.openshift.io/cluster instead.
	//
	// +optional
	AdditionalTrustedCA ConfigMapNameReference `json:"additionalTrustedCA"`
	// buildDefaults controls the default information for Builds
	// +optional
	BuildDefaults BuildDefaults `json:"buildDefaults"`
	// buildOverrides controls override settings for builds
	// +optional
	BuildOverrides BuildOverrides `json:"buildOverrides"`
}

type BuildDefaults struct {
	// defaultProxy contains the default proxy settings for all build operations, including image pull/push
	// and source download.
	//
	// Values can be overrode by setting the `HTTP_PROXY`, `HTTPS_PROXY`, and `NO_PROXY` environment variables
	// in the build config's strategy.
	// +optional
	DefaultProxy *ProxySpec `json:"defaultProxy,omitempty"`

	// gitProxy contains the proxy settings for git operations only. If set, this will override
	// any Proxy settings for all git commands, such as git clone.
	//
	// Values that are not set here will be inherited from DefaultProxy.
	// +optional
	GitProxy *ProxySpec `json:"gitProxy,omitempty"`

	// env is a set of default environment variables that will be applied to the
	// build if the specified variables do not exist on the build
	// +optional
	Env []corev1.EnvVar `json:"env,omitempty"`

	// imageLabels is a list of docker labels that are applied to the resulting image.
	// User can override a default label by providing a label with the same name in their
	// Build/BuildConfig.
	// +optional
	ImageLabels []ImageLabel `json:"imageLabels,omitempty"`

	// resources defines resource requirements to execute the build.
	// +optional
	Resources corev1.ResourceRequirements `json:"resources"`
}

type ImageLabel struct {
	// name defines the name of the label. It must have non-zero length.
	Name string `json:"name"`

	// value defines the literal value of the label.
	// +optional
	Value string `json:"value,omitempty"`
}

type BuildOverrides struct {
	// imageLabels is a list of docker labels that are applied to the resulting image.
	// If user provided a label in their Build/BuildConfig with the same name as one in this
	// list, the user's label will be overwritten.
	// +optional
	ImageLabels []ImageLabel `json:"imageLabels,omitempty"`

	// nodeSelector is a selector which must be true for the build pod to fit on a node
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// tolerations is a list of Tolerations that will override any existing
	// tolerations set on a build pod.
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// forcePull overrides, if set, the equivalent value in the builds,
	// i.e. false disables force pull for all builds,
	// true enables force pull for all builds,
	// independently of what each build specifies itself
	// +optional
	ForcePull *bool `json:"forcePull,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type BuildList struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard list's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	Items []Build `json:"items"`
}
