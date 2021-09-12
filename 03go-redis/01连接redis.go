package main

import (
	"context"
	"github.com/go-redis/redis"
)
var ctx = context.Background()
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.5.145:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//rdb.Set( "lk", "zs").Err()
	//err := rdb.Set(ctx, "lk", "zs", 10).Err()
	//if err != nil {
	//	panic(err)
	//}
	//result, err := rdb.Get(ctx, "lk").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("lk", result)
   rdb.Set("l","na",10)

}
