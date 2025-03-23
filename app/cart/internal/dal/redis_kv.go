package dal

// 初始化连接redis

import ()

type Product struct {
	Id          uint32     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
