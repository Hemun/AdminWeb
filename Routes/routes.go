package Routes

import (
	"air-q/Controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//auth routes
	//app.Post("/user/:userId/login", Controllers.Login)
	//app.Get("/user/:userId/logout", Controllers.Logout)
	//app.Post("/user/register", Controllers.Register)
	//app.Put("/user/userId/update", Controllers.Update)

	//Organization routes
	app.Post("/organizations", Controllers.CreateOrg)
	app.Get("/organizations/", Controllers.OrgList)
	app.Get("/organizations/:orgId", Controllers.GetOrgDetails)
	app.Delete("/organizations/:orgId", Controllers.DeleteOrg)
	app.Put("/organizations/:orgId", Controllers.UpdateOrg)

	//device routes
	app.Get("/devices", Controllers.DeviceList)
	app.Get("/devices/:deviceId", Controllers.GetdeviceDetails)
	app.Post("/devices", Controllers.CreateDevice)
	app.Delete("/devices/:deviceId", Controllers.DeleteDevice)
	app.Put("/devices/:deviceId", Controllers.UpdateDevice)

	//DeviceLog routes
	app.Get("/deviceLogs", Controllers.DeviceLogList)
	app.Get("/deviceLogs/:logId", Controllers.GetDeviceLogDetails)
	app.Post("/deviceLogs", Controllers.CreateDeviceLog)
	app.Delete("/deviceLogs/:logId", Controllers.DeleteDeviceLog)
	app.Put("/deviceLogs/:logId", Controllers.UpdateDeviceLog)

}
