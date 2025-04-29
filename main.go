// Copyright 2025 Homin Lee <ff4500@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/suapapa/go_khaiii/pkg/khaiii"
)

func main() {
	khaiii, err := khaiii.New()
	if err != nil {
		log.Fatalf("failed to open khaiii: %v", err)
	}
	defer khaiii.Close()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/analyze", func(c *gin.Context) {
		var request struct {
			Text string `json:"text"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		AnalyzeResult := khaiii.Analyze(request.Text, "")
		c.JSON(200, AnalyzeResult)
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
