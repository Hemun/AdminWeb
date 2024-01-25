package Controllers

import (
	"air-q/Models"
	db "air-q/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func isValidEmail(email string) bool {
	// Regular expression for basic email validation
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

func emailExists(email string) bool {
	var user Models.User
	db.DB.First(&user, email)
	if user.Email == "" {
		return false
	}
	return true
}

func Register(c *fiber.Ctx) error {
	var userData Models.User
	err := c.BodyParser(&userData)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"Message": "Бүргүүлэх хүсэлтэнд алдаа гарлаа.",
		})
	}
	if !isValidEmail(userData.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"Message": "Invalid email format",
		})
	}
	if emailExists(userData.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"Message": "Invalid email format",
		})
	}

	user := Models.User{
		FirstName:   userData.FirstName,
		LastName:    userData.LastName,
		Age:         userData.Age,
		Phone:       userData.Phone,
		Password:    userData.Password,
		Email:       userData.Email,
		Gender:      userData.Gender,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	db.DB.Create(&user)

	return nil
}

func Login(c *fiber.Ctx) error {
	UserID := c.Params("userId")
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Нэвтрэх хүсэлтэнд алдаа гарлаа.",
		})
	}

	//Нууц үг хоосон байгаа эсэх
	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Нууц үг оруулж өгнө үү!!!",
			"error":   map[string]interface{}{},
		})
	}
	var user Models.User
	db.DB.Where("id = ?", UserID).First(&user)

	//Хэрэглэгч байдаг эсэх
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Хэрэглэгч олдсонгүй",
			"error":   map[string]interface{}{},
		})
	}
	//Нууц үг таарч байгаа эсэх.
	//fmt.Println("--------------------------------")
	//fmt.Println("--------------DB Passcode------------------", cashier.Passcode)
	//fmt.Println("--------------DB Passcode typeOf------------------", reflect.TypeOf(cashier.Passcode))
	//fmt.Println("--------------------------------")
	//
	//fmt.Println("--------------------------------")
	//fmt.Println("--------------body passcode------------------", data["passcode"])
	//fmt.Println("--------------DB Passcode typeOf------------------", reflect.TypeOf(data["passcode"]))
	//fmt.Println("--------------------------------")

	if user.Password != data["password"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Нууц үг таарсангүй.",
			"error":   map[string]interface{}{},
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(user.ID)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(), //1 day
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	cookie := fiber.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Token дууссан эсвэл хүчингүй байна.",
		})
	}

	userData := make(map[string]interface{})
	userData["token"] = tokenString
	//fmt.Println(cashierData, " HelloCashierData")

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Амжилттай",
		"data":    userData,
	})

}

func Logout(c *fiber.Ctx) error {
	userID := c.Params("email")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//Нууц үг хэсэг хоосон байгаа эсэх
	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Нууц үг шаардлагатай.",
		})
	}

	var user Models.User
	db.DB.Where("Id = ?", userID).First(&user)

	//Ийм id-тай кассчин байдаг эсэх
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "Хэрэглэгч олдсонгүй.",
		})
	}
	//Нууц үг таарч байгаа эсэх
	if user.Password != data["password"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"Message": "Нууц үг таарсангүй.",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "success",
	})
}
