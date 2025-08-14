package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gosnmp/gosnmp"
)

func main() {
	startTime := time.Now()
  fmt.Println("Iniciando SNMP Walk...", startTime.Format(time.RFC3339))
  fmt.Printf("\n")
  
  // Configuration parameters for SNMP
	params := &gosnmp.GoSNMP{
    Target:    "127.0.0.1",
		Port:      161, 
		Community: "public", 
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		// Retries: 3,
	}

	sys_descr_oid := ".1.3.6.1.2.1.1.1.0"
  sys_object_id_oid := ".1.3.6.1.2.1.1.2.0"
  sys_uptime_oid := ".1.3.6.1.2.1.1.3.0"
  sys_contact_oid := ".1.3.6.1.2.1.1.4.0"
  sys_name_oid := ".1.3.6.1.2.1.1.5.0"
  sys_location_oid := ".1.3.6.1.2.1.1.6.0"

	// connect to the SNMP agent
	err := params.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer params.Conn.Close()

	results := make(map[string]interface{})

  sys_desc_result := walk(params, sys_descr_oid)
  results[sys_descr_oid] = sys_desc_result

  sys_object_id_result := walk(params, sys_object_id_oid)
  results[sys_object_id_oid] = sys_object_id_result

  sys_uptime_result := walk(params, sys_uptime_oid)
  results[sys_uptime_oid] = sys_uptime_result

  sys_contact_result := walk(params, sys_contact_oid)
  results[sys_contact_oid] = sys_contact_result

  sys_name_result := walk(params, sys_name_oid)
  results[sys_name_oid] = sys_name_result

  sys_location_result := walk(params, sys_location_oid)
  results[sys_location_oid] = sys_location_result

	// show results
	for oid, value := range results {
		fmt.Printf("%s: %v\n", oid, value)
	}
  fmt.Printf("\n")

  endTime := time.Now()
  fmt.Printf("SNMP Walk concluído em %s\n", endTime.Format(time.RFC3339))
  fmt.Printf("\033[1mTempo total: %s\033[0m\n", endTime.Sub(startTime))
}

func walk(params *gosnmp.GoSNMP, oid string) map[string]interface{} {
  results := make(map[string]interface{}) 

  err := params.Walk(oid, func(pdu gosnmp.SnmpPDU) error {
    // by default, gosnmp returns raw data (in bytes), for example: 
		// if value is "OCTET STRING" it will be returned as []byte (bytes array)
		// if value is "INTEGER" it will be returned as int64
		// is value is "STRING" it will be returned as string
		// if value is "COUNTER" it will be returned as uint64
		// in PHP you can use snmp2_real_walk to get all values in a single call. In this case, PHP parse the data and return it as an associative array.
		if b, ok := pdu.Value.([]byte); ok {
			results[pdu.Name] = string(b)
		} else {
			results[pdu.Name] = pdu.Value
		}

		return nil
	})
  
	if err != nil {
		log.Fatalf("Erro ao fazer SNMP Walk: %v", err)
    os.Exit(1)
	}

  return results
}
