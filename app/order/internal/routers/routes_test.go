package routers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MakiJOJO/douyin-mall-echo/app/order/internal/dal"
	"github.com/labstack/echo/v4"
)

// 测试 RegisterRoutes 函数
func TestRegisterRoutes(t *testing.T) {
	// 创建一个新的 Echo 实例
	e := echo.New()
	// 调用 RegisterRoutes 函数注册路由
	RegisterRoutes(e)

	// 检查是否注册了 /health 路由
	healthRouteFound := false
	for _, route := range e.Routes() {
		if route.Path == "/health" && route.Method == http.MethodGet {
			healthRouteFound = true
			break
		}
	}
	if !healthRouteFound {
		t.Errorf("Expected /health route to be registered, but it was not")
	}

	// 检查是否注册了 /api/order 组的路由
	apiOrderGroupFound := false
	for _, route := range e.Routes() {
		if route.Path == "/api/order/getOrderbyUserId" && route.Method == http.MethodGet {
			apiOrderGroupFound = true
			break
		}
	}
	if !apiOrderGroupFound {
		t.Errorf("Expected /api/order/getOrderbyUserId route to be registered, but it was not")
	}
}

// 测试 healthHandler 函数，当数据库连接正常时
func TestHealthHandler_DbConnected(t *testing.T) {
	// 创建一个新的 Echo 实例
	e := echo.New()
	// 创建一个新的 HTTP 请求
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	// 创建一个新的 HTTP 响应记录器
	rec := httptest.NewRecorder()
	// 创建一个新的 Echo 上下文
	c := e.NewContext(req, rec)

	// 模拟 dal.DbInstance
	dal.Init()

	// 调用 healthHandler 函数
	if err := healthHandler(c); err != nil {
		t.Errorf("healthHandler returned an error: %v", err)
	}

	// 检查响应状态码
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	// 检查响应体
	var response string
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}
	if response != "Healthy" {
		t.Errorf("Expected response body to be 'Healthy', but got '%s'", response)
	}
}

// 测试 healthHandler 函数，当数据库未连接时
func TestHealthHandler_DbNotConnected(t *testing.T) {
	// 创建一个新的 Echo 实例
	e := echo.New()
	// 创建一个新的 HTTP 请求
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	// 创建一个新的 HTTP 响应记录器
	rec := httptest.NewRecorder()
	// 创建一个新的 Echo 上下文
	c := e.NewContext(req, rec)

	// 模拟 dal.DbInstance 为 nil
	dal.DbInstance = nil

	// 调用 healthHandler 函数
	if err := healthHandler(c); err != nil {
		t.Errorf("healthHandler returned an error: %v", err)
	}

	// 检查响应状态码
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, rec.Code)
	}

	// 检查响应体
	var response string
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}
	if response != "database is not connected" {
		t.Errorf("Expected response body to be 'database is not connected', but got '%s'", response)
	}
}
