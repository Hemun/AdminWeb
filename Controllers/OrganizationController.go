package Controllers

import (
	"air-q/Models"
	db "air-q/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreateOrg(c *fiber.Ctx) error {
	var data Models.Organization
	//fmt.Println("before : ", data)
	err := c.BodyParser(&data)
	//fmt.Println("after : ", data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Org нэмэх үйлдэл амжилтгүй.",
			"error":   map[string]interface{}{},
		})
	}
	if data.OrgName == "" || data.OrgRegister == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Нэр юмуу регистер дутуу байна.",
			"error":   map[string]interface{}{},
		})
	}
	org := Models.Organization{
		OrgName:     data.OrgName,
		OrgRegister: data.OrgRegister,
		Phone:       data.Phone,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	db.DB.Create(&org)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Organization амжилттай нэмэгдлээ.",
		"data":    org,
	})
}

func OrgList(c *fiber.Ctx) error {
	var organization []Models.Organization
	db.DB.Select("*").Find(&organization)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Нийт Organization-ны мэдээлэл.",
			"data":    organization,
		})
}

func DeleteOrg(c *fiber.Ctx) error {
	orgID := c.Params("orgId")
	var organization Models.Organization
	db.DB.First(&organization, orgID)
	if organization.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Organization олдсонгүй.",
			"error":   map[string]interface{}{},
		})
	}
	db.DB.Delete(&organization)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Амжилттай устгагдлаа.",
	})
}

func GetOrgDetails(c *fiber.Ctx) error {
	orgID := c.Params("orgId")
	var organization Models.Organization
	db.DB.First(&organization, orgID)

	if organization.ID == 0 {
		return c.Status(200).JSON(
			fiber.Map{
				"success": false,
				"message": "Organization-ны id олдсонгүй.",
				"error":   map[string]interface{}{},
			})
	}

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Organization-ны мэдээлэл.",
			"data":    organization,
		})
}

func UpdateOrg(c *fiber.Ctx) error {
	orgID := c.Params("orgId")
	var organization Models.Organization
	db.DB.First(&organization, orgID)

	if organization.OrgName == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Organization олдсонгүй.",
			"error":   map[string]interface{}{},
		})
	}

	var data Models.Organization
	//var updatedData Models.Organization

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Засах үйлдэл амжилтгүй.",
			"error":   map[string]interface{}{},
		})
	}

	if data.OrgName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Organization нэр шаардлагатай.",
			"error":   map[string]interface{}{},
		})
	}

	fmt.Println(data)

	if data.ID != 0 {
		organization.ID = data.ID
	}
	if data.OrgName != "" {
		organization.OrgName = data.OrgName
	}
	if data.OrgRegister != "" {
		organization.OrgRegister = data.OrgRegister
	}
	if data.Phone != 0 {
		organization.Phone = data.Phone
	}
	organization.CreatedDate = organization.CreatedDate
	organization.UpdatedDate = time.Now()

	db.DB.Save(&organization)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Амжилттай.",
		"data":    organization,
	})
}
