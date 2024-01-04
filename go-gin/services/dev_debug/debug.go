package dev_debug

import (
	"fmt"
	"go1/services/logger"
	"reflect"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	
	msg := fmt.Sprintf("Type of variable: %s", t.String())
	logger.DefaultLogger.Info(msg)
}