package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/MakiJOJO/douyin-mall-echo/app/pay/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/pay/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/common/utils"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/order"
	"github.com/labstack/echo/v4"
	"github.com/smartwalle/alipay/v3"
)

var aliClient *alipay.Client

// OrderClient orderservice.Client

const (
	appID        = "9021000144656313" // 你的appID
	privateKey   = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCTgI+K//4AuGuqHwTGs/T2oauTv/OPhBVq+BOIpY2gc0badkR0XWd4+F5nT2APO/cj2leUQ+LWoefueDD1In4TvfnV3OltoyMa2ZetRSKYKh7vBM02SZueYyOFtbsJr/4LjPKTJQ8zJN1M0P1CTWkUbOLqgIVy1pAOKpqLN1LrLmZBX97S4RU3UaFO/DlYQ1QlMJeGfQI6yXY5gEp6jL5p6j81fO2oVQkZ6Xy1joUObfmIGC+sVdNhnCpePDwil8aM1r/OnNiVhKMLdohrjzrSx+5PHsoCL/yZDk2HbSUNBrYfNBOccjtXg0R/WB9v9h+LHVaEHyf3s7dFW2F1PW4DAgMBAAECggEAOyJVaeKLUHqfH0rkPU00LhROlyNjX+wSMhpWqnfEuci9ZSP4+bXgn8zi/AQEfNcLk6IbbmNw859hPmeNKRm09fE50hWIt92pW0BU4LBQ2DQ6xpRkORl1fCA+w4JCA/Y8oSDWt0sqNtTWq881WXlzYS7uIhl4ZrvCcQt/fcSmR4ZnTcCklCSViPiNoeqJ6v7HG2Yp7w7kxWaBK5LwAfcfvffia77FiS8JiMfKevn9xDzx1HY0f1U3a67tiWD3D2o0NuODC20IivXeq8dYcZUlR4C3+uv2wuNqOrz7S3G6T8HFdYP65sM6eX2d3h+gdMmRXPBMJuhN8w5A/k5rKqSnIQKBgQD9Dr8v83A+rCs3xQsr6JOGeHtDIFLbeOL2hSwgZQmcuAcGBkRXJCEZ3GI6pGk4W/vKLQJDa5/zLI2cUmNpTjutoIX/nbFXA6CJrj3Svf8BC44ZSZASVnlQhKPWkOiAqKpS7SgZITjMCVNkwbM+D4zxNmAA58KDQwVFlvOoTcjVywKBgQCVN53pqEMMACYYNxnXHHh/m7QMz8adHkaLYFbHm5JGnKuBo/mm8cC9ffB0r8PMlm+1oGt7Dr+rn2nH5mhcMHZUlNxWHVC4vMwYQb6pKIeumJz3nVKGfzQjQqqEmYQ3QCw+LWAZelOh9udyf+AYdlBfZOrEcNfU3UBSpxB3QieBqQKBgQDBSOc879rW89gaw9Uxl3Y+6n1zmxfObLomgzdeEu5Rlro+nsDKMl1aDFu24OdBVfiuxswIxQapzWjocoFd6JRqnMZcpIzUon+XSdAMzHQezz2dEPQLHaORnY4qkAoWYz2dE2liMF7acXER72VBzMzhXJ/dcSe/7Iv1SEQZVDhHGwKBgAMRtcnMkDR6/E2bNIjcKQ5W4Ykx7N/mc4UCYkaQyJ1zM4PjH4tzhYdgQ8Xip6BZp8qQliVd1EtvZ/mYn6TlyklLFo0e5T4ng/srvwQztTa+JNxi/AOQMj5XbLJ1heatBzvwKv3bKkU2kuQkBTP7mwObS8jmmUnjkyMgFJKfZbihAoGBAIuq77SS4eV5bkvWmPwRvxPum5n1BpeFodL0DAIrQgEj52oRpxcdOjPa3m+EbUPQk4l6TBlfNa+nFbz5M6yUFBve2eAxlseoVPhUplexUmNbGbNJ43ax4dKp8mLxaOtEiY0a1ozIaBp4rGBq/Ms2Fgcmf+yQWFiqJRaTMOlfkeVY"
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAghJidiLFG6ew1n04fcutz+5CopeNbSM8rYGvbICbbWE23Fs57+nV/TYFmPMVATLWYyC6mwRx70Ntqr4WSoETvwFKjUkO6DENRhF0qDYkYWAjq/Qt0M1ABRqByUmM1ClgmMT90FBXqdImV91acHAhR2dy1/Bu5ylE+nCNL5EW2NDvtizpQvEbpvqbDVzWrDWA+xFvPKZtNwQ9F0kQaK/GgsDod8fPn3Y0OwJwUB1QjUS9CrGSWzqmZTrnuYD6lJcnA+LfwBJmOLDpDl7EM0idm1CsLh5Rt0ZBnAv1fYoy5vFBuAnBzdZcXEMOfw9XuKVldCNU2MomN77V+/JGyvJBdwIDAQAB"
)

func ZfbPay(orderID string, totalPrice string) string {
	serverAddr := config.GlobalConfig.HOST

	ethoAddr := utils.MustGetLocalIPv4()
	serverAddr = ethoAddr + serverAddr

	kServerDomain := "http://" + serverAddr

	fmt.Println("kServerDomain", kServerDomain)
	aliClient, err := alipay.New(appID, privateKey, false)
	if err != nil {
		panic(err)
	}
	err = aliClient.LoadAliPayPublicKey(aliPublicKey)
	if err != nil {
		panic(err)
	}
	// var p = alipay.TradeWapPay{}
	p := alipay.TradePagePay{}

	p.NotifyURL = kServerDomain + "/api/v1/pay/notify" // 支付宝回调

	p.ReturnURL = kServerDomain + "/api/v1/pay/callback" // 支付后调转页面
	fmt.Println(" p.ReturnURL", p.ReturnURL)
	fmt.Println("p.NotifyURL", p.NotifyURL)
	p.Subject = "douyin-mall-echo" // 标题
	p.OutTradeNo = orderID         // 传递一个唯一单号
	p.TotalAmount = totalPrice     // 金额
	// p.ProductCode = "QUICK_WAP_WAY"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY" // 网页支付
	var url2 *url.URL
	url2, err = aliClient.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	payURL := url2.String()
	println(payURL)
	return payURL
}

func Callback(c echo.Context) error {
	// 获取查询参数
	charset := c.QueryParam("charset")
	outTradeNo := c.QueryParam("out_trade_no")
	method := c.QueryParam("method")
	totalAmount := c.QueryParam("total_amount")
	sign := c.QueryParam("sign")
	// 其他参数
	// 你可以继续提取其他参数，根据需要

	// 打印参数以进行调试
	log.Printf("charset: %s, out_trade_no: %s, method: %s, total_amount: %s, sign: %s", charset, outTradeNo, method, totalAmount, sign)

	// 在这里可以进行签名验证，查询订单状态等逻辑

	// 返回响应
	if outTradeNo == "" {
		return c.String(http.StatusBadRequest, "参数错误")
	}
	orderId, err := strconv.ParseUint(outTradeNo, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "订单ID转换错误")
	}
	// // 修改订单状态并返回结果
	// r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// cli, err := orderservice.NewClient("order", client.WithResolver(r)) // 指定 Resolver

	paySuccessReq := &order.OrderPaySuccessRequest{
		OrderId: orderId,
	}
	resp, err := client.OrderClient.OrderPaySuccess(context.Background(), paySuccessReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("resp: %v", resp)

	return c.String(http.StatusOK, "回调已成功处理")
}

func Notify(c echo.Context) error {
	// 解析支付宝传递的通知数据
	notifyData, err := c.FormParams()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "解析通知数据失败"})
	}

	// 验证通知签名
	aliClient, err := alipay.New(appID, privateKey, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "创建支付宝客户端失败"})
	}
	err = aliClient.LoadAliPayPublicKey(aliPublicKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "加载支付宝公钥失败"})
	}

	// 验证支付是否成功
	tradeStatus := notifyData.Get("trade_status")
	if tradeStatus != "TRADE_SUCCESS" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "支付未成功"})
	}

	// 获取订单号和支付金额等信息
	orderID := notifyData.Get("out_trade_no")
	totalAmount := notifyData.Get("total_amount")
	fmt.Printf("订单号: %s, 支付金额: %s\n", orderID, totalAmount)

	// 处理业务逻辑，比如更新订单状态
	// TODO: 更新订单状态到数据库等

	// 返回支付宝要求的应答
	return c.String(http.StatusOK, "success")
}
