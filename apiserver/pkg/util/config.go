package util

// ClientOptions contains configuration needed to create a Kubernetes client
type ClientOptions struct {
	QPS   float32
	Burst int
}

// TODO: this needs to be revised.
const (
	// Label keys
	RayClusterLabelKey            = "ray.io/cluster"
	RayClusterNameLabelKey        = "ray.io/cluster-name"
	RayClusterUserLabelKey        = "ray.io/user"
	RayClusterVersionLabelKey     = "ray.io/version"
	RayClusterEnvironmentLabelKey = "ray.io/environment"
	RayNodeTypeLabelKey           = "ray.io/node-type"
	RayNodeGroupLabelKey          = "ray.io/group"

	KubernetesApplicationNameLabelKey = "app.kubernetes.io/name"
	KubernetesManagedByLabelKey       = "app.kubernetes.io/managed-by"

	// Annotation keys
	// Role level
	RayClusterComputeTemplateAnnotationKey = "ray.io/compute-template"
	RayClusterImageAnnotationKey           = "ray.io/compute-image"

	RayClusterDefaultImageRepository = "rayproject/ray"
)

const (
	// The application name
	ApplicationName = "kuberay"

	// The component name for apiserver
	ComponentName = "kuberay-apiserver"
)
