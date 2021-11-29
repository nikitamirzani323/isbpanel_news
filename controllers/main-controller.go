package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_news/config"
	"github.com/nikitamirzani323/isbpanel_news/helpers"
)

type login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type response_login struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}

const PATH string = config.PATH_API

func Login(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(login)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_login{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username": client.Username,
			"password": client.Password,
		}).
		Post(PATH + "api/login")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_login)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"token":  result.Token,
		"key":    result.Key,
		"time":   time.Since(render_page).String(),
	})
}
