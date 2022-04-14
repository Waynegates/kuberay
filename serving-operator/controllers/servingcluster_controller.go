/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	rayv1alpha1 "github.com/ray-project/kuberay/api/v1alpha1"
	rayutil "github.com/ray-project/kuberay/apiserver/pkg/util"
	rayiov1alpha1 "github.com/ray-project/kuberay/ray-operator/api/raycluster/v1alpha1"
	"github.com/ray-project/kuberay/ray-operator/controllers/common"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	log              = logf.Log.WithName("servingcluster-controller")
	ctrlLog          = ctrl.Log.WithName("ctrl-servingcluster-controller")
	rayclusterSuffix = "-raycluster"
)

// NewReconciler returns a new reconcile.Reconciler
func NewReconciler(mgr manager.Manager) *ServingClusterReconciler {
	return &ServingClusterReconciler{
		Client:   mgr.GetClient(),
		Scheme:   mgr.GetScheme(),
		Log:      ctrl.Log.WithName("controllers").WithName("RayCluster"),
		Recorder: mgr.GetEventRecorderFor("raycluster-controller"),
	}
}

var _ reconcile.Reconciler = &ServingClusterReconciler{}

// ServingClusterReconciler reconciles a ServingCluster object
type ServingClusterReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=ray.io,resources=servingclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ray.io,resources=servingclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ray.io,resources=servingclusters/finalizers,verbs=update
// +kubebuilder:rbac:groups=ray.io,resources=rayclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ray.io,resources=rayclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ray.io,resources=rayclusters/finalizer,verbs=update
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;create;update
// +kubebuilder:rbac:groups=core,resources=serviceaccount,verbs=get;list;watch;create;delete
// +kubebuilder:rbac:groups="rbac.authorization.k8s.io",resources=roles,verbs=get;list;watch;create;delete;update
// +kubebuilder:rbac:groups="rbac.authorization.k8s.io",resources=rolebindings,verbs=get;list;watch;create;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ServingCluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *ServingClusterReconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	tmplog := r.Log.WithValues("servingcluster", request.NamespacedName)
	log.Info("reconciling ServingCluster", "cluster name", request.Name)
	ctrlLog.Info("ctrl reconciling ServingCluster", "cluster name", request.Name)
	tmplog.Info("tmp reconciling ServingCluster", "cluster name", request.Name)
	// TODO(user): your logic here
	// Get serving cluster instance
	servingClusterInstance := &rayv1alpha1.ServingCluster{}
	if err := r.Get(context.TODO(), request.NamespacedName, servingClusterInstance); err != nil {
		if errors.IsNotFound(err) {
			log.Info("Read request instance not found error!")
		} else {
			log.Error(err, "Read request instance error!")
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Update ray cluster
	rayclusterNamespacedName := request.NamespacedName
	rayclusterNamespacedName.Name = rayclusterNamespacedName.Name + rayclusterSuffix
	rayClusterInstance := &rayiov1alpha1.RayCluster{}

	err := r.Get(context.TODO(), rayclusterNamespacedName, rayClusterInstance)

	if err == nil {
		rayClusterInstance.Spec = servingClusterInstance.Spec.RayClusterSpec

		log.Info("Update ray cluster spec")
		if err := r.Update(context.TODO(), rayClusterInstance); err != nil {
			log.Error(err, "Fail to update ray cluster instance!")
			// Error reading the object - requeue the request.
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
	} else {
		if errors.IsNotFound(err) {
			log.Info("Not found raycluster, creating raycluster!")
			rayClusterInstance, err = r.constructRayClusterForServingCluster(servingClusterInstance)
			if err != nil {
				log.Error(err, "unable to construct raycluster from spec")
				// don't bother requeuing until we get a change to the spec
				return ctrl.Result{}, nil
			}
			if err := r.Create(ctx, rayClusterInstance); err != nil {
				log.Error(err, "unable to create rayCluster for servingCluster", "rayCluster", rayClusterInstance)
				return ctrl.Result{}, err
			}
			log.V(1).Info("created rayCluster for servingCluster run", "rayCluster", rayClusterInstance)
		} else {
			log.Error(err, "Read request instance error!")
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
	}

	clientURL, err := r.fetchDashboardURL(rayClusterInstance)

	if err != nil {
		log.Error(err, "fail to find dashboard url")
		return ctrl.Result{}, err
	}

	if clientURL == "" {
		log.Info("dashboard url is empty")
		return ctrl.Result{}, nil
	}

	rayDashboardClient := RayDashboardClient{}
	rayDashboardClient.intClient(clientURL)
	deploymentString, err := rayDashboardClient.getDeployments()

	if err != nil {
		log.Error(err, "fail to get deployment info")
		return ctrl.Result{}, err
	}

	log.V(1).Info("deployment details", "deployment", deploymentString)

	if len(servingClusterInstance.Spec.ServeConfigSpecs) == 0 {
		if err = rayDashboardClient.deleteDeployments(); err != nil {
			log.Error(err, "fail to delete deployment")
			return ctrl.Result{}, err
		}
	} else {
		if err = rayDashboardClient.updateDeployments(servingClusterInstance.Spec.ServeConfigSpecs); err != nil {
			log.Error(err, "fail to update deployment")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *ServingClusterReconciler) constructRayClusterForServingCluster(servingCluster *rayv1alpha1.ServingCluster) (*rayiov1alpha1.RayCluster, error) {
	name := fmt.Sprintf("%s%s", servingCluster.Name, rayclusterSuffix)

	rayCluster := &rayiov1alpha1.RayCluster{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      servingCluster.Labels,
			Annotations: servingCluster.Annotations,
			Name:        name,
			Namespace:   servingCluster.Namespace,
		},
		Spec: *servingCluster.Spec.RayClusterSpec.DeepCopy(),
	}

	if rayCluster.Annotations == nil {
		rayCluster.Annotations = make(map[string]string)
	}

	rayCluster.Annotations["servingclusterName"] = servingCluster.Name

	if err := ctrl.SetControllerReference(servingCluster, rayCluster, r.Scheme); err != nil {
		return nil, err
	}

	return rayCluster, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServingClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rayv1alpha1.ServingCluster{}).
		Named("servingcluster-controller").
		Owns(&rayiov1alpha1.RayCluster{}).
		Complete(r)
}

func (r *ServingClusterReconciler) fetchDashboardURL(instance *rayiov1alpha1.RayCluster) (string, error) {
	headServices := corev1.ServiceList{}
	filterLabels := client.MatchingLabels{common.RayClusterLabelKey: instance.Name}
	if err := r.List(context.TODO(), &headServices, client.InNamespace(instance.Namespace), filterLabels); err != nil {
		return "", err
	}

	dashboardURL := ""

	if headServices.Items != nil {
		if len(headServices.Items) == 1 {
			r.Log.Info("reconcileServices ", "head service found", headServices.Items[0].Name)
			// TODO: compare diff and reconcile the object
			// For example. ServiceType might be changed or port might be modified
			servicePorts := headServices.Items[0].Spec.Ports

			dashboardPort := int32(-1)

			for _, servicePort := range servicePorts {
				if servicePort.Name == "dashboard" {
					dashboardPort = servicePort.Port
					break
				}
			}

			if dashboardPort == int32(-1) {
				return "", rayutil.NewCustomErrorf(rayutil.CUSTOM_CODE_NOT_FOUND, "dashboard port not found")
			}

			// TODO: Replace localhost with a discover hostname
			dashboardURL = fmt.Sprintf("localhost:%v",
				dashboardPort)
			return dashboardURL, nil
		}

		// This should never happen.
		// We add the protection here just in case controller has race issue or user manually create service with same label.
		if len(headServices.Items) > 1 {
			r.Log.Info("reconcileServices ", "Duplicates head service found", len(headServices.Items))
			return "", rayutil.NewCustomErrorf(rayutil.CUSTOM_CODE_GENERIC, "Duplicates head service found %v", len(headServices.Items))
		}
	}

	// Create head service if there's no existing one in the cluster.
	if headServices.Items == nil || len(headServices.Items) == 0 {
		return "", rayutil.NewCustomErrorf(rayutil.CUSTOM_CODE_GENERIC, "No head service found")
	}

	return "", nil
}
