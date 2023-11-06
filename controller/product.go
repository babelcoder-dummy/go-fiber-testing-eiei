package controller

import (
	"errors"
	"strconv"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/dto"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var validate = validator.New()

type Product struct {
	Service service.Product
}

func (p *Product) Create(c *fiber.Ctx) error {
	form := new(dto.CreateProductForm)
	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := validate.Struct(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := p.Service.Create(form, image)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var serializedProduct dto.ProductResponse
	copier.Copy(&serializedProduct, &product)

	return c.Status(fiber.StatusCreated).JSON(serializedProduct)
}

func (p *Product) FindAll(c *fiber.Ctx) error {
	term := c.Query("term")
	products := p.Service.FindAll(term)

	var serializedProducts []dto.ProductResponse
	copier.Copy(&serializedProducts, &products)

	return c.JSON(serializedProducts)
}

func (p *Product) FindOne(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := p.Service.FindOne(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(fiber.StatusNotFound)
	}

	var serializedProduct dto.ProductResponse
	copier.Copy(&serializedProduct, &product)

	return c.JSON(serializedProduct)
}

func (p *Product) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	form := new(dto.UpdateProductForm)
	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := validate.Struct(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	image, _ := c.FormFile("image")
	product, err := p.Service.Update(uint(id), image, form)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(fiber.StatusNotFound)
	}

	var serializedProduct dto.ProductResponse
	copier.Copy(&serializedProduct, &product)

	return c.JSON(serializedProduct)
}

func (p *Product) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	p.Service.Delete(uint(id))
	return c.SendStatus(fiber.StatusNoContent)
}
