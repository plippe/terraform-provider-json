package main

import "fmt"
import "os"
import "encoding/json"

import "./models"
import "../terraform-provider-aws/aws"

func createProviderJson(filePath string, terraformResources []aws.TerraformResource) {
	resources := []models.Resource{}
	for _, terraformResource := range terraformResources {
		resources = append(
			resources,
			models.ResourceBuilder(terraformResource),
		)
	}

	err := os.Remove(filePath)
	if err != nil { panic(err) }

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil { panic(err) }
    defer f.Close()

	for _, resource := range resources {
		bytes, err := json.Marshal(resource)
		if err == nil { _, err = f.Write(bytes) }
		if err == nil { _, err = f.Write([]byte("\n")) }

		if err != nil { panic(err) }
	}
}

func main() {
	createProviderJson("docs/terraform-provider-aws.json", append(aws.DataSources(), aws.Resources()...))
	fmt.Println("done")
}
