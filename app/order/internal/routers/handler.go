package routers

import (
	"github.com/MakiJOJO/douyin-mall-echo/app/order/internal/logic"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

// 辅助函数：将字符串转换为整数，用于处理输入参数的转换
func convertToInt(str string, paramName string, c echo.Context) (int, error) {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0, c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid " + paramName,
		})
	}
	return value, nil
}

// 辅助函数：将字符串转换为 uint64，用于处理输入参数的转换
func convertToUint64(str string, paramName string, c echo.Context) (uint64, error) {
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid " + paramName,
		})
	}
	return value, nil
}

// GetOrderHandlerByUserID 根据用户 ID 获取订单信息
func GetOrderByUserIDHandler(c echo.Context) error {
	userIDStr := c.FormValue("user_id")
	userID, err := convertToInt(userIDStr, "user ID", c)
	if err != nil {
		return err
	}

	orders, err := logic.GetOrderByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get orders by user ID: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, orders)
}

// GetOrderByOrderIdHandler 根据订单 ID 获取订单信息
func GetOrderByOrderIdHandler(c echo.Context) error {
	orderIDStr := c.FormValue("order_id")
	orderID, err := convertToUint64(orderIDStr, "order ID", c)
	if err != nil {
		return err
	}

	order, err := logic.GetOrderByOrderId(orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get order by order ID: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, order)
}

// convertStringToIntSlice 将字符串按分隔符分割并转换为 int 切片
func convertStringToIntSlice(str string, sep string) ([]int, error) {
	strSlice := strings.Split(str, sep)
	intSlice := make([]int, len(strSlice))
	for i, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		intSlice[i] = num
	}
	return intSlice, nil
}

// convertStringToFloat64Slice 将字符串按分隔符分割并转换为 float64 切片
func convertStringToFloat64Slice(str string, sep string) ([]float64, error) {
	strSlice := strings.Split(str, sep)
	floatSlice := make([]float64, len(strSlice))
	for i, s := range strSlice {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		floatSlice[i] = num
	}
	return floatSlice, nil
}

// CreateOrderHandler 创建订单
func CreateOrderHandler(c echo.Context) error {
	userIDStr := c.FormValue("userid")
	userAddr := c.FormValue("useraddr")
	cartIDStr := c.FormValue("cartid")
	productIDsStr := c.FormValue("productids")
	quantitiesStr := c.FormValue("quantities")
	pricesStr := c.FormValue("prices")

	userID, err := convertToInt(userIDStr, "user ID", c)
	if err != nil {
		return err
	}

	cartID, err := convertToInt(cartIDStr, "cart ID", c)
	if err != nil {
		return err
	}

	productIDs, err := convertStringToIntSlice(productIDsStr, ",")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product IDs",
		})
	}

	quantities, err := convertStringToIntSlice(quantitiesStr, ",")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid quantities",
		})
	}

	prices, err := convertStringToFloat64Slice(pricesStr, ",")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid prices",
		})
	}

	orderID, err := logic.CreateOrder(userID, cartID, userAddr, productIDs, quantities, prices)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create order: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]uint64{
		"order_id": orderID,
	})
}

// UpdateOrderAddrHandler 更新订单地址
func UpdateOrderAddrHandler(c echo.Context) error {
	orderIDStr := c.FormValue("orderid")
	userAddr := c.FormValue("useraddr")

	orderID, err := convertToUint64(orderIDStr, "order ID", c)
	if err != nil {
		return err
	}

	if userAddr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user address",
		})
	}

	err = logic.UpdateOrderAddr(orderID, userAddr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update order address: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Order address updated successfully",
	})
}

// CancelOrderHandler 取消订单
func CancelOrderHandler(c echo.Context) error {
	orderIDStr := c.FormValue("orderid")
	orderID, err := convertToUint64(orderIDStr, "order ID", c)
	if err != nil {
		return err
	}

	err = logic.CancelOrder(orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to cancel order: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Order cancelled successfully",
	})
}

// UpdateOrderSuccessHandler 更新订单状态为成功
func UpdateOrderSuccessHandler(c echo.Context) error {
	orderIDStr := c.FormValue("orderid")
	orderID, err := convertToUint64(orderIDStr, "order ID", c)
	if err != nil {
		return err
	}

	err = logic.UpdateOrderSuccess(orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update order status: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Order status updated to success",
	})
}
