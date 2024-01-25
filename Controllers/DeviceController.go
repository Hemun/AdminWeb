package Controllers

import (
	"air-q/Models"
	db "air-q/config"
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreateDevice(c *fiber.Ctx) error {
	var data Models.Device
	//fmt.Println("before : ", data)
	err := c.BodyParser(&data)
	//fmt.Println("after : ", data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Device нэмэх үйлдэл амжилтгүй.",
			"error":   map[string]interface{}{},
		})
	}
	if data.DeviceName == "" || data.DeviceNo == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Нэр юмуу дугаар дутуу байна.",
			"error":   map[string]interface{}{},
		})
	}
	dev := Models.Device{
		DeviceName:  data.DeviceName,
		DeviceNo:    data.DeviceNo,
		DeviceTspe:  data.DeviceTspe,
		DevicePrice: data.DevicePrice,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	db.DB.Create(&dev)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Device амжилттай нэмэгдлээ.",
		"data":    dev,
	})
}

func DeviceList(c *fiber.Ctx) error {
	var device []Models.Device
	db.DB.Select("*").Find(&device)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Нийт Device-ын мэдээлэл.",
			"data":    device,
		})
}

func GetdeviceDetails(c *fiber.Ctx) error {
	devID := c.Params("deviceId")
	var device Models.Device
	db.DB.First(&device, devID)

	if device.ID == 0 {
		return c.Status(200).JSON(
			fiber.Map{
				"success": false,
				"message": "Device-ны id олдсонгүй.",
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

func DeleteDevice(c *fiber.Ctx) error {
	devID := c.Params("deviceId")
	var device Models.Device
	db.DB.First(&device, devID)
	if device.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "device олдсонгүй.",
			"error":   map[string]interface{}{},
		})
	}
	db.DB.Delete(&device)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Амжилттай устгагдлаа.",
	})
}

func UpdateDevice(c *fiber.Ctx) error {
	devID := c.Params("deviceId")
	var device Models.Device
	db.DB.First(&device, devID)

	if device.DeviceName == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Device олдсонгүй.",
			"error":   map[string]interface{}{},
		})
	}

	var data Models.Device
	//var updatedData Models.Organization

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Засах үйлдэл амжилтгүй.",
			"error":   map[string]interface{}{},
		})
	}

	if data.DeviceName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Organization нэр шаардлагатай.",
			"error":   map[string]interface{}{},
		})
	}

	//fmt.Println(data)

	if data.ID != 0 {
		device.ID = data.ID
	}
	if data.DeviceName != "" {
		device.DeviceName = data.DeviceName
	}
	if data.DeviceNo != "" {
		device.DeviceNo = data.DeviceNo
	}
	if data.DeviceTspe != "" {
		device.DeviceTspe = data.DeviceTspe
	}
	if data.DevicePrice != 0 {
		device.DevicePrice = data.DevicePrice
	}
	device.CreatedDate = device.CreatedDate
	device.UpdatedDate = time.Now()

	db.DB.Save(&device)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Амжилттай.",
		"data":    device,
	})
}
