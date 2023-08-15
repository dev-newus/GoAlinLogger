package logging

import "testing"

func TestSetOu(t *testing.T){
	logger := NewLogger(Info)
	logger.SetOutput("./log/test.log")
}

func TestNewLogger(t *testing.T){
	logger := NewLogger(Info)
	logger.Info("test")
}