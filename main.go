// Copyright 2025 Homin Lee <ff4500@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

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
	khaiii, err := khaiii.New()
	if err != nil {
		log.Fatalf("failed to open khaiii: %v", err)
	}
	defer khaiii.Close()

	router := gin.Default()
	v1 := router.Group("/v1")
	// @Summary Ping the server
	// @Description Check if the server is running
	// @Tags Health
	// @Produce json
	// @Success 200 {object} map[string]string
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// @Summary Analyze Korean text
	// @Description Perform morphological analysis on the input text
	// @Tags Analysis
	// @Accept json
	// @Produce json
	// @Param request body struct{Text string} true "Text to analyze"
	// @Success 200 {object} interface{}
	// @Failure 400 {object} khaiii.AnalyzeResult
	v1.POST("/analyze", func(c *gin.Context) {
		var request reqInput
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		AnalyzeResult := khaiii.Analyze(request.Text, "")
		c.JSON(200, AnalyzeResult)
	})

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run() // listen and serve on 0.0.0.0:8080
}

type reqInput struct {
	Text string `json:"text"`
}
