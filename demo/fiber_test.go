package demo_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

type Product struct {
	ID   string
	Name string
}

func TestRoutes(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Taylor"},
		{ID: "2", Name: "Swift"},
		{ID: "3", Name: "1989"},
	}

	app := fiber.New()

	{
		app.Get("/products", func(c *fiber.Ctx) error {
			return c.JSON(products)
		})

		req := httptest.NewRequest("GET", "/products", nil)
		res, err := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, 200, res.StatusCode, "status code")
		assert.NoError(t, err, "response error")
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
		assert.JSONEq(
			t,
			`[{"ID": "1", "Name": "Taylor"}, {"ID": "2", "Name": "Swift"}, {"ID": "3", "Name": "1989"}]`,
			string(body),
			"response body",
		)
	}

	{
		app.Get("/products/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			product, found := lo.Find(products, func(item Product) bool {
				return item.ID == id
			})

			if !found {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return c.JSON(product)
		})

		req := httptest.NewRequest("GET", "/products/1", nil)
		res, err := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, 200, res.StatusCode, "status code")
		assert.NoError(t, err, "response error")
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
		assert.JSONEq(
			t,
			`{"ID": "1", "Name": "Taylor"}`,
			string(body),
			"response body",
		)
	}

	{
		app.Post("/products", func(c *fiber.Ctx) error {
			form := new(Product)

			if err := c.BodyParser(form); err != nil {
				return err
			}

			form.ID = fmt.Sprintf("%d", len(products)+1)
			products = append(products, *form)

			return c.Status(fiber.StatusCreated).JSON(form)
		})

		req := httptest.NewRequest("POST", "/products", strings.NewReader(`{"Name": "Taylor's version"}`))
		req.Header.Add("Content-Type", "application/json")
		res, err := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, 201, res.StatusCode, "status code")
		assert.NoError(t, err, "response error")
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
		assert.JSONEq(
			t,
			`{"ID": "4", "Name": "Taylor's version"}`,
			string(body),
			"response body",
		)
	}
}
