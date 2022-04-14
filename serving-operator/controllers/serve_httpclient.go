package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"

	rayv1alpha1 "github.com/ray-project/kuberay/api/v1alpha1"
)

var (
	DEPLOY_PATH = "/api/serve/deployments/"
	INFO_PATH   = "/api/serve/deployments/"
	STATUS_PATH = "/api/serve/deployments/status"
	DELETE_PATH = "/api/serve/deployments/"
)

type ServingClusterDeployments struct {
	Deployments []rayv1alpha1.ServeConfigSpec `json:"deployments,omitempty"`
}

type RayDashboardClient struct {
	client       http.Client
	dashboardURL string
}

func (r *RayDashboardClient) intClient(url string) {
	r.client = http.Client{}
	r.dashboardURL = "http://" + url
}

func (r *RayDashboardClient) getDeployments() (string, error) {
	req, err := http.NewRequest("GET", r.dashboardURL+DEPLOY_PATH, nil)
	if err != nil {
		return "", err
	}

	resp, err := r.client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return fmt.Sprintf("%s\n", bytes), nil
}

func (r *RayDashboardClient) updateDeployments(specs []rayv1alpha1.ServeConfigSpec) error {
	servingClusterDeployments := ServingClusterDeployments{
		Deployments: specs,
	}

	deploymentJson, err := json.Marshal(servingClusterDeployments)

	fmt.Println("deploymentJson:", deploymentJson)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", r.dashboardURL+DELETE_PATH, bytes.NewBuffer(deploymentJson))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}

func (r *RayDashboardClient) deleteDeployments() error {
	req, err := http.NewRequest("DELETE", r.dashboardURL+DEPLOY_PATH, nil)
	if err != nil {
		return err
	}

	resp, err := r.client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
