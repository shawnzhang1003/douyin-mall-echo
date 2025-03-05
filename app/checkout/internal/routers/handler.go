package routers

import (
	"fmt"
	"net/http"

	"github.com/MakiJOJO/douyin-mall-echo/app/checkout/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/checkout/internal/logic"
	"github.com/labstack/echo/v4"
)

type CheckoutReq struct {
	UserId    uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Firstname string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Address   *Address `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

type Address struct {
	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       string `protobuf:"bytes,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
}

func HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func healthHandler(c echo.Context) error {
	if dal.DbInstance == nil {
		return c.JSON(http.StatusInternalServerError, "database is not connected")
	}
	return c.JSON(http.StatusOK, dal.DbInstance.Health())
}

func CheckoutHandler(c echo.Context) error {
	myReq := &CheckoutReq{}
	if err := c.Bind(myReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	retInfo, err := logic.Checkout(myReq.UserId,
		fmt.Sprintf("%s, %s, %s, %s, %s",
			myReq.Address.StreetAddress,
			myReq.Address.City,
			myReq.Address.State,
			myReq.Address.Country,
			myReq.Address.ZipCode))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resp := map[string]interface{}{
		"message":  "checkout successfully!",
		"order_id": retInfo.OrderId,
		"pay_url":  retInfo.PayUrl,
	}
	return c.JSON(http.StatusOK, resp)
}
