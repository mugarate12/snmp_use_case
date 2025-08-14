package main

import (
	"testing"
	"time"

	"github.com/gosnmp/gosnmp"
)

func TestWalk(t *testing.T) {
  sys_descr_oid := ".1.3.6.1.2.1.1.1.0"
  params := &gosnmp.GoSNMP{
    Target:    "127.0.0.1",
		Port:      161, 
		Community: "public", 
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

  err := params.Connect()
  if err != nil {
    t.Fatalf("Failed to connect: %v", err)
  }
  defer params.Conn.Close()
  result := walk(params, sys_descr_oid)
  
  if result == nil {
    t.Errorf("Expected result for OID %s, got nil", sys_descr_oid)
  }
}
