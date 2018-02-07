package routes

import (
	rstypes "git.containerum.net/ch/json-types/resource-service"
	umtypes "git.containerum.net/ch/json-types/user-manager"
	"git.containerum.net/ch/resource-service/server"
	"git.containerum.net/ch/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var srv server.ResourceService

// SetupRoutes sets up a router
func SetupRoutes(app *gin.Engine, server server.ResourceService) {
	srv = server

	app.Use(utils.SaveHeaders)
	app.Use(utils.PrepareContext)
	app.Use(utils.RequireHeaders(umtypes.UserIDHeader, umtypes.UserRoleHeader))
	app.Use(utils.SubstituteUserMiddleware)

	rstypes.RegisterCustomTagsGin(binding.Validator)

	ns := app.Group("/namespace")
	{
		ns.POST("", createNamespaceHandler)

		ns.GET("", getUserNamespacesHandler)
		ns.GET("/:ns_label", getUserNamespaceHandler)
		ns.GET("/:ns_label/access", getUserNamespaceAccessesHandler)
		ns.GET("/:ns_label/volumes", getVolumesLinkedWithUserNamespaceHandler)

		ns.DELETE("/:ns_label", deleteUserNamespaceHandler)

		ns.PUT("/:ns_label/name", renameUserNamespaceHandler)
		ns.PUT("/:ns_label/access", setUserNamespaceAccessHandler)
		ns.PUT("/:ns_label", resizeUserNamespaceHandler)

		deployment := ns.Group("/:ns_label/deployment")
		{
			deployment.POST("", createDeploymentHandler)

			deployment.GET("", getDeploymentsHandler)
			deployment.GET("/:deploy_label", getDeploymentByLabelHandler)

			deployment.DELETE("/:deploy_label", deleteDeploymentByLabelHandler)

			deployment.PUT("/:deploy_label/image", setContainerImageHandler)
			deployment.PUT("/:deploy_label", replaceDeploymentHandler)
			deployment.PUT("/:deploy_label/replicas", setReplicasHandler)
		}
	}

	nss := app.Group("/namespaces")
	{
		nss.GET("", utils.RequireAdminRole, getAllNamespacesHandler)

		nss.DELETE("", utils.RequireAdminRole, deleteAllUserNamespacesHandler)
	}

	vol := app.Group("/volume")
	{
		vol.POST("", createVolumeHandler)

		vol.GET("", getUserVolumesHandler)
		vol.GET("/:vol_label", getUserVolumeHandler)
		vol.GET("/:vol_label/access", getUserVolumeAccessesHandler)

		vol.DELETE("/:vol_label", deleteUserVolumeHandler)

		vol.PUT("/:vol_label/name", renameUserVolumeHandler)
		vol.PUT("/:vol_label/access", setUserVolumeAccessHandler)
		vol.PUT("/:vol_label", resizeUserVolumeHandler)
	}

	vols := app.Group("/volumes")
	{
		vols.GET("", utils.RequireAdminRole, getAllVolumesHandler)

		vols.DELETE("", utils.RequireAdminRole, deleteAllUserVolumesHandler)
	}

	app.GET("/access", getUserResourceAccessesHandler)

	adm := app.Group("/adm")
	{
		adm.PUT("/access", utils.RequireAdminRole, setUserResourceAccessesHandler)
	}
}
