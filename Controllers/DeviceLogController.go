package Controllers

import (
	"air-q/Models"
	db "air-q/config"
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreateDeviceLog(c *fiber.Ctx) error {
	var data Models.Device_Log
	//fmt.Println("before : ", data)
	err := c.BodyParser(&data)
	//fmt.Println("after : ", data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Device_Log нэмэх үйлдэл амжилтгүй.",
			"error":   map[string]interface{}{},
		})
	}
	if data.DeviceID == 0 || data.Value == 0 {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "ID юмуу Value дутуу байна.",
			"error":   map[string]interface{}{},
		})
	}
	devLog := Models.Device_Log{
		DeviceID:    data.DeviceID,
		Value:       data.Value,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	db.DB.Create(&devLog)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Log амжилттай нэмэгдлээ.",
		"data":    devLog,
	})
}

func DeviceLogList(c *fiber.Ctx) error {
	var logs []Models.Device_Log
	db.DB.Select("*").Find(&logs)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Нийт Log-ын мэдээлэл.",
			"data":    logs,
		})
}

func GetDeviceLogDetails(c *fiber.Ctx) error {
	logID := c.Params("logId")
	var device Models.Device_Log
	db.DB.First(&device, logID)

	if device.ID == 0 {
		return c.Status(200).JSON(
			fiber.Map{
				"success": false,
				"message": "Log-ны id олдсонгүй.",
				"error":   map[string]interface{}{},
			})
	}

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Device-ны мэдээлэл.",
			"data":    device,
		})
}

func DeleteDeviceLog(c *fiber.Ctx) error {
	logID := c.Params("logId")
	var logs Models.Device_Log
	db.DB.First(&logs, logID)
	if logs.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "log олдсонгүй.",
			"error":   map[string]interface{}{},
		})
	}
	db.DB.Delete(&logs)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Амжилттай устгагдлаа.",
	})
}

func UpdateDeviceLog(c *fiber.Ctx) error {
	logID := c.Params("logId")
	var logs Models.Device_Log
	db.DB.First(&logs, logID)

	if logs.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Log олдсонгүй.",
			"error":   map[string]interface{}{},
		})
	}

	var data Models.Device_Log
	//var updatedData Models.Organization

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Засах үйлдэл амжилтгүй.",
			"error":   map[string]interface{}{},
		})
	}

	if data.DeviceID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "log id шаардлагатай.",
			"error":   map[string]interface{}{},
		})
	}

	//fmt.Println(data)

	if data.ID != 0 {
		logs.ID = data.ID
	}
	if data.DeviceID != 0 {
		logs.DeviceID = data.DeviceID
	}
	if data.Value != 0 {
		logs.Value = data.Value
	}
	logs.CreatedDate = logs.CreatedDate
	logs.UpdatedDate = time.Now()

	db.DB.Save(&logs)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Амжилттай.",
		"data":    logs,
	})
}
