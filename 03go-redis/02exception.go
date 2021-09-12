package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)
var ctxs = context.Background()
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.5.145:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//查询一个不存在的key 对其进行健壮性检查
	result, err := rdb.Get(ctxs, "keys").Result()
	if err ==redis.Nil{
		fmt.Println("key2 does not exist")
	}else if err!=nil {
		panic(err)
	}else {
		fmt.Println("keys",result)
	}

}
