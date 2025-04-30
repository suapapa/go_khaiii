// Copyright 2025 Homin Lee <ff4500@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suapapa/go_khaiii/pkg/khaiii"

	_ "github.com/suapapa/go_khaiii/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Khaiii API
// @version 1.0
// @description This is a REST API for Khaiii Korean morphological analyzer.
// @host localhost:8080
// @BasePath /v1
func main() {
	var secret string
	secretB, err := os.ReadFile("/token/secret.txt")
	if err != nil {
		fmt.Printf("WARN: failed to read secret: %v\n", err)
	} else {
		secret = strings.TrimSpace(string(secretB))
	}

	khaiii, err := khaiii.New(nil)
	if err != nil {
		log.Fatalf("failed to open khaiii: %v", err)
	}
	defer khaiii.Close()

	root := cmp.Or(os.Getenv("ROOT_PATH"), "/")
	if strings.HasSuffix(root, "/") {
		root = root[:len(root)-1]
	}

	router := gin.Default()
	v1 := router.Group("/v1")

	// @Summary Ping the server
	// @Description Check if the server is running
	// @Tags Health
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router /{root}/ping [get]
	v1.GET(root+"/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// @Summary Analyze Korean text
	// @Description Perform morphological analysis on the input text
	// @Tags Analysis
	// @Accept json
	// @Produce json
	// @Param request body reqInput true "Text to analyze"
	// @Success 200 {object} map[string]interface{} "Analysis result"
	// @Failure 400 {object} map[string]string "Invalid request"
	// @Router /v1/analyze [post]
	v1.POST(root+"/analyze", func(c *gin.Context) {
		var request reqInput
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		AnalyzeResult := khaiii.Analyze(request.Text, "")
		c.JSON(200, gin.H{"data": AnalyzeResult})
	})

	v1.GET(root+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// add middleware to check secret
	router.Use(func(c *gin.Context) {
		if c.Request.Method == "GET" || secret == "" {
			c.Next()
			return
		}

		if c.GetHeader("Authorization") != "Bearer "+secret {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

type reqInput struct {
	Text string `json:"text"`
}
