package octopusdeploy

import (
	"context"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandOfflinePackageDropDeploymentTarget(d *schema.ResourceData) *octopusdeploy.DeploymentTarget {
	endpoint := octopusdeploy.NewOfflinePackageDropEndpoint()

	if v, ok := d.GetOk("applications_directory"); ok {
		endpoint.ApplicationsDirectory = v.(string)
	}

	if v, ok := d.GetOk("destination"); ok {
		endpoint.Destination = expandOfflinePackageDropDestination(v)
	}

	if v, ok := d.GetOk("working_directory"); ok {
		endpoint.WorkingDirectory = v.(string)
	}

	deploymentTarget := expandDeploymentTarget(d)
	deploymentTarget.Endpoint = endpoint
	return deploymentTarget
}

func flattenOfflinePackageDropDeploymentTarget(deploymentTarget *octopusdeploy.DeploymentTarget) map[string]interface{} {
	if deploymentTarget == nil {
		return nil
	}

	flattenedDeploymentTarget := flattenDeploymentTarget(deploymentTarget)
	endpointResource, _ := octopusdeploy.ToEndpointResource(deploymentTarget.Endpoint)
	flattenedDeploymentTarget["applications_directory"] = endpointResource.ApplicationsDirectory
	flattenedDeploymentTarget["destination"] = flattenOfflinePackageDropDestination(endpointResource.Destination)
	flattenedDeploymentTarget["working_directory"] = endpointResource.WorkingDirectory
	return flattenedDeploymentTarget
}

func getOfflinePackageDropDeploymentTargetDataSchema() map[string]*schema.Schema {
	dataSchema := getOfflinePackageDropDeploymentTargetSchema()
	setDataSchema(&dataSchema)

	deploymentTargetDataSchema := getDeploymentTargetDataSchema()

	deploymentTargetDataSchema["offline_package_drop_deployment_target"] = &schema.Schema{
		Computed:    true,
		Description: "A list of offline package drop deployment targets that match the filter(s).",
		Elem:        &schema.Resource{Schema: dataSchema},
		Optional:    true,
		Type:        schema.TypeList,
	}

	delete(deploymentTargetDataSchema, "communication_styles")
	delete(deploymentTargetDataSchema, "deployment_targets")
	deploymentTargetDataSchema["id"] = getDataSchemaID()

	return deploymentTargetDataSchema
}

func getOfflinePackageDropDeploymentTargetSchema() map[string]*schema.Schema {
	offlinePackageDropDeploymentTargetSchema := getDeploymentTargetSchema()

	offlinePackageDropDeploymentTargetSchema["applications_directory"] = &schema.Schema{
		Required: true,
		Type:     schema.TypeString,
	}

	offlinePackageDropDeploymentTargetSchema["destination"] = &schema.Schema{
		Computed: true,
		Elem:     &schema.Resource{Schema: getOfflinePackageDropDestinationSchema()},
		Optional: true,
		MaxItems: 1,
		Type:     schema.TypeList,
	}

	offlinePackageDropDeploymentTargetSchema["working_directory"] = &schema.Schema{
		Required: true,
		Type:     schema.TypeString,
	}

	return offlinePackageDropDeploymentTargetSchema
}

func setOfflinePackageDropDeploymentTarget(ctx context.Context, d *schema.ResourceData, deploymentTarget *octopusdeploy.DeploymentTarget) {
	if deploymentTarget == nil {
		return
	}

	endpointResource, err := octopusdeploy.ToEndpointResource(deploymentTarget.Endpoint)
	if err != nil {
		return
	}

	d.Set("applications_directory", endpointResource.ApplicationsDirectory)
	d.Set("destination", flattenOfflinePackageDropDestination(endpointResource.Destination))
	d.Set("working_directory", endpointResource.WorkingDirectory)

	setDeploymentTarget(ctx, d, deploymentTarget)
}
