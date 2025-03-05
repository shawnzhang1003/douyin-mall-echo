package main

import (
    "flag"
    "fmt"
    "net/http"
    "sync"
    "time"
	// "io/ioutil"
)

// 发送单个请求并记录响应时间
func sendRequest(method, urlStr string, wg *sync.WaitGroup, results chan int64) {
    defer wg.Done()

    var resp *http.Response
    var err error

    start := time.Now()
    switch method {
    case "GET":
        resp, err = http.Get(urlStr)
    case "POST":
        // 如果是 POST 请求，这里暂时不处理请求体，仅展示如何添加 POST 请求逻辑
        // 若需要发送请求体，可参考之前的 JSON 示例添加相应代码
        resp, err = http.Post(urlStr, "", nil)
    default:
        fmt.Printf("不支持的请求方法: %s\n", method)
        return
    }

    if err != nil {
        fmt.Printf("请求出错: %v\n", err)
        return
    }
    defer resp.Body.Close()

	// // 读取响应体
    // body, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     fmt.Printf("读取响应体出错: %v\n", err)
    //     return
    // }
	// fmt.Println(string(body))
    elapsed := time.Since(start).Milliseconds()
    results <- elapsed
}

// 压测函数
func performLoadTest(method, url string, concurrency int, requests int) {
    var wg sync.WaitGroup
    results := make(chan int64, requests)

    // 记录压测开始时间
    startTime := time.Now()

    // 启动并发请求
    for i := 0; i < concurrency; i++ {
        wg.Add(requests / concurrency)
        for j := 0; j < requests/concurrency; j++ {
            go sendRequest(method, url, &wg, results)
        }
    }

    // 等待所有请求完成
    go func() {
        wg.Wait()
        close(results)
    }()

    // 收集响应时间
    var totalTime int64
    var minTime int64 = 999999
    var maxTime int64 = 0
    count := 0
    for elapsed := range results {
        totalTime += elapsed
        if elapsed < minTime {
            minTime = elapsed
        }
        if elapsed > maxTime {
            maxTime = elapsed
        }
        count++
    }

    // 记录压测结束时间
    endTime := time.Now()
    // 计算总耗时（秒）
    totalDuration := endTime.Sub(startTime).Seconds()

    // 计算平均响应时间
    if count > 0 {
        avgTime := totalTime / int64(count)
        fmt.Printf("请求总数: %d\n", requests)
        fmt.Printf("并发数: %d\n", concurrency)
        fmt.Printf("最小响应时间: %d ms\n", minTime)
        fmt.Printf("最大响应时间: %d ms\n", maxTime)
        fmt.Printf("平均响应时间: %d ms\n", avgTime)

        // 计算 QPS
        qps := float64(requests) / totalDuration
        fmt.Printf("QPS: %.2f\n", qps)
    }
}

func main() {
    // 定义命令行参数
    method := flag.String("method", "GET", "请求方法，如 GET、POST")
    url := flag.String("url", "https://example.com/api/search?userid=123", "要压测的 URL，可直接包含查询参数")
    concurrency := flag.Int("concurrency", 10, "并发数")
    requests := flag.Int("requests", 100, "总请求数")
    flag.Parse()

    // 执行压测
    performLoadTest(*method, *url, *concurrency, *requests)
}