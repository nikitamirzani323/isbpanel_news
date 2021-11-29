package controllers

import (
	"log"
	"strings"
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
type checkpage struct {
	Page string `json:"page" validate:"required"`
}
type news struct {
	News string `json:"news_search" `
}
type newssave struct {
	Page          string `json:"page" validate:"required"`
	Sdata         string `json:"sdata" validate:"required"`
	News_idrecord int    `json:"news_id"`
	News_category int    `json:"news_category" validate:"required"`
	News_title    string `json:"news_title" validate:"required"`
	News_descp    string `json:"news_descp" validate:"required"`
	News_url      string `json:"news_url" validate:"required"`
	News_image    string `json:"news_image" validate:"required"`
}
type newsdelete struct {
	Page    string `json:"page" validate:"required"`
	News_id int    `json:"news_id"`
}
type responseerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type response_login struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}
type response_checkpage struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
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
			"username":  client.Username,
			"password":  client.Password,
			"ipaddress": "22.22.22.22",
			"timezone":  "desktop",
		}).
		Post(PATH + "api/login")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_login)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"token":  result.Token,
		"time":   time.Since(render_page).String(),
	})
}
func Checkpage(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(checkpage)
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

	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_checkpage{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"page": client.Page,
		}).
		Post(PATH + "api/valid")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_checkpage)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": result.Status,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}

func News(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(news)
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

	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_checkpage{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"news_search": client.News,
		}).
		Post(PATH + "api/news")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_checkpage)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": result.Status,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Newssave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(newssave)
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

	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_checkpage{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"page":          client.Page,
			"sdata":         client.Sdata,
			"news_id":       client.News_idrecord,
			"news_category": client.News_category,
			"news_title":    client.News_title,
			"news_descp":    client.News_descp,
			"news_url":      client.News_url,
			"news_image":    client.News_image,
		}).
		Post(PATH + "api/newssave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_checkpage)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Newsdelete(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(newsdelete)
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

	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_checkpage{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"page":    client.Page,
			"news_id": client.News_id,
		}).
		Post(PATH + "api/newsdelete")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_checkpage)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": result.Status,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Categorynews(c *fiber.Ctx) error {

	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_checkpage{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		Post(PATH + "api/categorynews")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_checkpage)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": result.Status,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
