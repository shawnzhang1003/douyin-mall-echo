package routers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/MakiJOJO/douyin-mall-echo/app/product/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/product/internal/logic"
	"github.com/MakiJOJO/douyin-mall-echo/app/product/model"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/labstack/echo/v4"
)

func HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}
	c.Logger().Info("Hello World")
	mtl.Logger.Info("Hello World")

	return c.JSON(http.StatusOK, resp)
}

func DeleteProductHandler(c echo.Context) error {
	productIDStr := c.FormValue("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}
	err = logic.DeleteProduct(uint32(productID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "success")
}

func UpdateProductHandler(c echo.Context) error {
	product_id_str := c.FormValue("product_id")
	product_id, err := strconv.Atoi(product_id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}
	product_name := c.FormValue("name")
	product_description := c.FormValue("description")
	product_price_str := c.FormValue("price")
	product_picture := c.FormValue("picture")
	product_category_name := c.FormValue("category_name")
	product_price, err := strconv.ParseFloat(product_price_str, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid price")
	}
	retProductID, err := logic.UpdateProduct(uint32(product_id), product_name, product_description, product_picture, float32(product_price), product_category_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, retProductID)
	return nil
}

func CreateProductHandler(c echo.Context) error {
	product_name := c.FormValue("name")
	product_description := c.FormValue("description")
	product_price_str := c.FormValue("price")
	product_picture := c.FormValue("picture")
	productCategoryNamesStr := c.FormValue("category_names")
	product_price, err := strconv.ParseFloat(product_price_str, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid price")
	}
	categoryNames := strings.Split(productCategoryNamesStr, ",")
	if len(categoryNames) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Category names are required")
	}
	retProduct, err := logic.CreateProduct(product_name, product_description, product_picture, float32(product_price), categoryNames)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, retProduct)
	return nil
}

func healthHandler(c echo.Context) error {
	if dal.DbInstance == nil {
		return c.JSON(http.StatusInternalServerError, "database is not connected")
	}
	return c.JSON(http.StatusOK, dal.DbInstance.Health())
}

func GetProductsByCategoryIDHandler(c echo.Context) error {
	categoryIDStr := c.FormValue("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}
	products, err := logic.GetProductsByCategoryID(uint32(categoryID))
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, products)
	return nil
}

func GetProductByIDHandler(c echo.Context) error {
	productIDStr := c.FormValue("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}
	// 将读取redis那部分也放在model中吧，这里好像没有复用logic
	product, err := model.GetById(uint32(productID))
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, product)
	return nil
}

func SearchProductsHandler(c echo.Context) error {
	name := c.FormValue("name")
	products, err := logic.SearchProducts(name)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, products)
	return nil
}

func GetAllCategoriesHandler(c echo.Context) error {
	categories, err := logic.GetAllCategories()
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, categories)
	return nil
}

func ListProductsHandler(c echo.Context) error {
	page, err := strconv.Atoi(c.FormValue("page"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid page number")
	}
	pageSize, err := strconv.Atoi(c.FormValue("page_size"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid page size")
	}
	categoryName := c.FormValue("category_name")

	products, err := logic.ListProducts(page, pageSize, categoryName)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, products)
	return nil
}
