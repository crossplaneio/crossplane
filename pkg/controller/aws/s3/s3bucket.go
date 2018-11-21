/*
Copyright 2018 The Conductor Authors.

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

package s3

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	bucketv1alpha1 "github.com/upbound/conductor/pkg/apis/aws/storage/v1alpha1"
	awsv1alpha1 "github.com/upbound/conductor/pkg/apis/aws/v1alpha1"
	corev1alpha1 "github.com/upbound/conductor/pkg/apis/core/v1alpha1"
	"github.com/upbound/conductor/pkg/clients/aws"
	"github.com/upbound/conductor/pkg/clients/aws/s3"
	"github.com/upbound/conductor/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	controllerName = "s3bucket.aws.conductor.io"
	finalizer      = "finalizer." + controllerName

	errorResourceClient = "Failed to create s3 client"
	errorCreateResource = "Failed to create resource"
	errorSyncResource   = "Failed to sync resource state"
	errorDeleteResource = "Failed to delete resource"
)

var (
	ctx           = context.Background()
	result        = reconcile.Result{}
	resultRequeue = reconcile.Result{Requeue: true}
)

// Add creates a new Instance Controller and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// Reconciler reconciles a S3Bucket object
type Reconciler struct {
	client.Client
	scheme     *runtime.Scheme
	kubeclient kubernetes.Interface
	recorder   record.EventRecorder

	connect func(*bucketv1alpha1.S3Bucket) (s3.Service, error)
	create  func(*bucketv1alpha1.S3Bucket, s3.Service) (reconcile.Result, error)
	sync    func(*bucketv1alpha1.S3Bucket, s3.Service) (reconcile.Result, error)
	delete  func(*bucketv1alpha1.S3Bucket, s3.Service) (reconcile.Result, error)
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	r := &Reconciler{
		Client:     mgr.GetClient(),
		scheme:     mgr.GetScheme(),
		kubeclient: kubernetes.NewForConfigOrDie(mgr.GetConfig()),
		recorder:   mgr.GetRecorder(controllerName),
	}
	r.connect = r._connect
	r.create = r._create
	r.sync = r._sync
	r.delete = r._delete
	return r
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New(controllerName, mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Instance
	err = c.Watch(&source.Kind{Type: &bucketv1alpha1.S3Bucket{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to Instance Secret
	err = c.Watch(&source.Kind{Type: &corev1.Secret{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &bucketv1alpha1.S3Bucket{},
	})
	if err != nil {
		return err
	}

	return nil
}

// fail - helper function to set fail condition with reason and message
func (r *Reconciler) fail(bucket *bucketv1alpha1.S3Bucket, reason, msg string) (reconcile.Result, error) {
	bucket.Status.SetCondition(corev1alpha1.NewCondition(corev1alpha1.Failed, reason, msg))
	return reconcile.Result{Requeue: true}, r.Update(context.TODO(), bucket)
}

// connectionSecret return secret object for this resource
func connectionSecret(bucket *bucketv1alpha1.S3Bucket, accessKey *iam.AccessKey) *corev1.Secret {
	if bucket.APIVersion == "" {
		bucket.APIVersion = bucketv1alpha1.S3BucketKindAPIVersion
	}
	if bucket.Kind == "" {
		bucket.Kind = bucketv1alpha1.S3BucketKind
	}

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:            bucket.ConnectionSecretName(),
			Namespace:       bucket.Namespace,
			OwnerReferences: []metav1.OwnerReference{bucket.OwnerReference()},
		},

		Data: map[string][]byte{
			corev1alpha1.ResourceCredentialsSecretUserKey:     []byte(*accessKey.AccessKeyId),
			corev1alpha1.ResourceCredentialsSecretPasswordKey: []byte(*accessKey.SecretAccessKey),
			corev1alpha1.ResourceCredentialsSecretEndpointKey: []byte(bucket.Endpoint()),
		},
	}
}

func (r *Reconciler) _connect(instance *bucketv1alpha1.S3Bucket) (s3.Service, error) {
	// Fetch AWS Provider
	p := &awsv1alpha1.Provider{}
	providerNamespacedName := types.NamespacedName{
		Namespace: instance.Namespace,
		Name:      instance.Spec.ProviderRef.Name,
	}
	err := r.Get(ctx, providerNamespacedName, p)
	if err != nil {
		return nil, err
	}

	// Check provider status
	if !p.IsValid() {
		return nil, fmt.Errorf("provider is not ready")
	}

	// Get Provider's AWS Config
	config, err := aws.Config(r.kubeclient, p)
	if err != nil {
		return nil, err
	}

	// Bucket Region and client region must match.
	config.Region = instance.Spec.Region

	// Create new S3 S3Client
	return s3.NewClient(config), nil
}

func (r *Reconciler) _create(bucket *bucketv1alpha1.S3Bucket, client s3.Service) (reconcile.Result, error) {
	bucket.Status.UnsetAllConditions()
	bucket.Status.SetCreating()
	util.AddFinalizer(&bucket.ObjectMeta, finalizer)
	return resultRequeue, r.Update(ctx, bucket)
}

func (r *Reconciler) _sync(bucket *bucketv1alpha1.S3Bucket, client s3.Service) (reconcile.Result, error) {
	if bucket.Status.IsCondition(corev1alpha1.Creating) && !bucket.Status.IsCondition(corev1alpha1.Ready) {
		err := client.CreateBucket(&bucket.Spec)
		if err != nil {
			return r.fail(bucket, errorCreateResource, err.Error())
		}

		if bucket.Status.IAMUsername == nil {
			bucket.Status.IAMUsername = s3.GenerateBucketUsername(&bucket.Spec)
		}
		accessKeys, err := client.CreateUser(bucket.Status.IAMUsername, &bucket.Spec)
		if err != nil {
			return r.fail(bucket, errorCreateResource, err.Error())
		}

		secret := connectionSecret(bucket, accessKeys)
		secret.OwnerReferences = append(secret.OwnerReferences, bucket.OwnerReference())
		bucket.Status.ConnectionSecretRef = corev1.LocalObjectReference{Name: secret.Name}

		_, err = util.ApplySecret(r.kubeclient, secret)
		if err != nil {
			return r.fail(bucket, errorCreateResource, err.Error())
		}

		bucket.Status.UnsetCondition(corev1alpha1.Creating)
		bucket.Status.SetReady()
		return result, r.Update(ctx, bucket)
	}

	return result, nil
}

func (r *Reconciler) _delete(bucket *bucketv1alpha1.S3Bucket, client s3.Service) (reconcile.Result, error) {
	if bucket.Spec.ReclaimPolicy == corev1alpha1.ReclaimDelete {
		if err := client.Delete(bucket); err != nil {
			return r.fail(bucket, errorDeleteResource, err.Error())
		}
	}

	bucket.Status.SetCondition(corev1alpha1.NewCondition(corev1alpha1.Deleting, "", ""))
	util.RemoveFinalizer(&bucket.ObjectMeta, finalizer)
	return result, r.Update(ctx, bucket)
}

// Reconcile reads that state of the bucket for an Instance object and makes changes based on the state read
// and what is in the Instance.Spec
func (r *Reconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the CRD instance
	bucket := &bucketv1alpha1.S3Bucket{}

	err := r.Get(ctx, request.NamespacedName, bucket)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return result, nil
		}
		// Error reading the object - requeue the request.
		log.Printf("failed to get object at start of reconcile loop: %+v", err)
		return result, err
	}

	s3Client, err := r.connect(bucket)
	if err != nil {
		return r.fail(bucket, errorResourceClient, err.Error())
	}

	// Check for deletion
	if bucket.DeletionTimestamp != nil {
		return r.delete(bucket, s3Client)
	}

	// Create s3 bucket
	if !bucket.Status.IsCondition(corev1alpha1.Ready) && !bucket.Status.IsCondition(corev1alpha1.Creating) {
		return r.create(bucket, s3Client)
	}

	// Initialize the bucket, status update is a noop currently.
	return r.sync(bucket, s3Client)
}
