package service

import (
	"github.com/garyburd/redigo/redis"
	"../config"

)


func RedisQuery(RedisQuery string)(RedisQueryResult string){
	c, err := redis.Dial("tcp", config.RedisAddr)
	checkErr(err)
	defer c.Close()
	result, err := redis.String(c.Do("GET", RedisQuery))
	checkErr(err)
	return result
}

func RedisWrite(WriteKey string,WriteValue string){
	c, err := redis.Dial("tcp", config.RedisAddr)
	checkErr(err)
	defer c.Close()
	_, err = c.Do("SET", WriteKey, WriteValue)
	checkErr(err)
}


func QueryRedisAndCleanSituation(ControlSituation string)(QueryResult string){
	c, err := redis.Dial("tcp", config.RedisAddr)
	checkErr(err)
	defer c.Close()
	result, err := redis.String(c.Do("GET", ControlSituation))
	checkErr(err)
	_, err = c.Do("SET", ControlSituation,"" )
	checkErr(err)
	return result

}