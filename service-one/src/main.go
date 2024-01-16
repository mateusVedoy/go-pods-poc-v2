package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	RenderChi "github.com/go-chi/render"
	RenderPkg "github.com/unrolled/render"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var USER = os.Getenv("USER")
var render *RenderPkg.Render
var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(1)
	go start()
	waitGroup.Wait()
}

func start() {
	contentType := middleware.AllowContentType("application/json")
	render = RenderPkg.New()
	route := chi.NewRouter()
	route.Use(middleware.RequestID)
	route.Use(middleware.RealIP)
	route.Use(middleware.Recoverer)
	route.Use(contentType)
	route.Use(RenderChi.SetContentType(RenderChi.ContentTypeJSON))
	route.Use(middleware.Timeout(60 * time.Second))

	route.Post("/service-two/pods/{kind}/amount/update/{amount}", updatePodsAmount)
	route.Post("/service-two/pods/delete/name/{identifier}", deletePodByName)

	panic(http.ListenAndServe(":8081", route))
}

// delete pods controller
func deletePodByName(w http.ResponseWriter, r *http.Request) {

	podName := chi.URLParam(r, "identifier")

	clientset, clientsetErr := kubeconfig()

	if clientsetErr != nil {
		fmt.Fprintf(w, clientsetErr.Error())
		return
	}

	podsClient := clientset.CoreV1().Pods("default")

	deletePropagation := metav1.DeletePropagationForeground
	deleteOptions := &metav1.DeleteOptions{
		PropagationPolicy: &deletePropagation,
	}

	deleteErr := podsClient.Delete(context.TODO(), podName, *deleteOptions)

	if deleteErr != nil {
		fmt.Fprintf(w, deleteErr.Error())
		return
	}

	scallingErr := decreasePods("default", "service-two")

	if scallingErr != nil {
		fmt.Fprintf(w, scallingErr.Error())
		return
	}

	fmt.Fprintf(w, "Pod deleted successfully")
}

// update pod amount controller
func updatePodsAmount(w http.ResponseWriter, r *http.Request) {

	amount := chi.URLParam(r, "amount")
	amountInt, _ := strconv.Atoi(amount)
	amountInt32 := int32(amountInt)
	kind := chi.URLParam(r, "kind")

	var scallingErr error

	if kind == "deployment" {
		scallingErr = scaleDeploymentPods("default", "service-two", amountInt32)
	}

	if kind == "hpa" {
		scallingErr = scaleHpaPods("default", "service-two", amountInt32)
	}

	if scallingErr != nil {
		fmt.Fprintf(w, scallingErr.Error())
	} else {
		fmt.Fprintf(w, "Pods updated: %s", amount)
	}
}

func decreasePods(namespace, deploymentName string) error {
	clientset, clientsetErr := kubeconfig()

	if clientsetErr != nil {
		return clientsetErr
	}

	deploymentClient := clientset.AppsV1().Deployments(namespace)

	deployment, deplErr := deploymentClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})

	if deplErr != nil {
		return errors.New(fmt.Sprintf("Error getting deployment. Reason: %v", deplErr.Error()))
	}

	*deployment.Spec.Replicas -= 1

	_, updateErr := deploymentClient.Update(context.TODO(), deployment, metav1.UpdateOptions{})

	if updateErr != nil {
		return errors.New(fmt.Sprintf("Error updating deployment. Reason: %v", updateErr.Error()))
	}

	return nil
}

// update pod amount asset
func scaleDeploymentPods(namespace, deploymentName string, replicas int32) error {
	clientset, clientsetErr := kubeconfig()

	if clientsetErr != nil {
		return clientsetErr
	}

	deploymentClient := clientset.AppsV1().Deployments(namespace)

	deployment, deplErr := deploymentClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})

	if deplErr != nil {
		return errors.New(fmt.Sprintf("Error getting deployment. Reason: %v", deplErr.Error()))
	}

	deployment.Spec.Replicas = &replicas

	_, updateErr := deploymentClient.Update(context.TODO(), deployment, metav1.UpdateOptions{})

	if updateErr != nil {
		return errors.New(fmt.Sprintf("Error updating deployment. Reason: %v", updateErr.Error()))
	}

	return nil
}

// update pod amount asset
func scaleHpaPods(namespace, deploymentName string, minReplicas int32) error {

	clientset, clientsetErr := kubeconfig()

	if clientsetErr != nil {
		return clientsetErr
	}

	hpaClient := clientset.AutoscalingV1().HorizontalPodAutoscalers(namespace)

	hpa, hpaErr := hpaClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})

	if hpaErr != nil {
		return errors.New(fmt.Sprintf("Error getting HPA. Reason: %v", hpaErr.Error()))
	}

	_, updateErr := hpaClient.Update(context.TODO(), hpa, metav1.UpdateOptions{})

	if updateErr != nil {
		return errors.New(fmt.Sprintf("Error updating HPA. Reason: %v", updateErr.Error()))
	}

	return nil
}

// common kubeconfig
func kubeconfig() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error config in cluster config. Reason: %v", err.Error()))
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error config kubernetes new config. Reason: %v", err.Error()))
	}

	return clientset, nil
}
