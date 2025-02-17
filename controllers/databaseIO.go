/*
 * Author: Yang Aobo
 * Telegram: @AnAsianGangster
 * Created At: March 17, 2021
 * Updated At: March 18, 2021
 * Last Modified By: Yang Aobo
 */

/**
 * This package contains HTTP handler functions
 *
 *
 * This file contains handler functions that handle I/O operations on key value pairs
 *
 * All functions destructure HTTP requests, call database operations, build response
 * and reply with response
 */
package controllers

import (
	"go-http-redis/databases"
	"go-http-redis/tools"
	"log"

	"github.com/gin-gonic/gin"
)

type KeyValuePair struct {
	Node  string `json:"node"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func FindOneKeyValuePair(context *gin.Context) {
	redisClient := tools.GetRedisClient()
	node := context.Query("node")
	key := context.Query("key")
	value := databases.FindOneKeyValuePair(node, redisClient, key)
	context.JSON(200, gin.H{
		"key":   key,
		"value": value,
	})
}

func CreatOneKeyValuePair(context *gin.Context) {
	redisClient := tools.GetRedisClient()
	var keyValuePair KeyValuePair
	err := context.BindJSON(&keyValuePair)
	if err != nil {
		log.Fatal(err)
	}

	value := databases.AddKeyValuePair(keyValuePair.Node, redisClient, keyValuePair.Key, keyValuePair.Value)
	if value == true {
		context.JSON(201, gin.H{
			"node":  keyValuePair.Node,
			"key":   keyValuePair.Key,
			"value": keyValuePair.Value,
		})
	} else {
		context.JSON(400, gin.H{
			"success": false,
		})
	}
}

func UpdateOneKeyValuePair(context *gin.Context) {
	redisClient := tools.GetRedisClient()
	var keyValuePair KeyValuePair
	err := context.BindJSON(&keyValuePair)
	if err != nil {
		log.Fatal(err)
	}

	value := databases.UpdateKeyValuePair(keyValuePair.Node, redisClient, keyValuePair.Key, keyValuePair.Value)
	if value == true {
		context.JSON(201, gin.H{
			"node":  keyValuePair.Node,
			"key":   keyValuePair.Key,
			"value": keyValuePair.Value,
		})
	} else {
		context.JSON(400, gin.H{
			"success": false,
		})
	}
}

func DeleteOneKeyValuePair(context *gin.Context) {
	redisClient := tools.GetRedisClient()
	node := context.Query("node")
	key := context.Query("key")
	value := databases.DeleteOneKeyValuePair(node, redisClient, key)

	if value == true {
		context.JSON(200, gin.H{
			"message": "delete successfully",
			"key":     key,
			"value":   value,
		})
	} else {
		context.JSON(400, gin.H{
			"success": false,
		})
	}
}
