/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/joho/godotenv"
	"go1/cmd"
	"go1/services/env"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env.JwtConfig = env.CreateJwtSetting()
}

func main() {
	cmd.Execute()
}
