package logic

import (
	"errors"

	"github.com/MakiJOJO/douyin-mall-echo/app/pay/utils"
)

func AliPay(order_id string, total_price string) (string, error) {
	payUrl := utils.ZfbPay(order_id, total_price)
	if payUrl == "" {
		return "", errors.New("支付宝支付失败")
	}
	return payUrl, nil
}
