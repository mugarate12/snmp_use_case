package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/joho/godotenv"
)

var (
  HOST_IP string
  COMMUNITY string
  PORT uint16

  SERIAL string
)

func main() {
  LoadConfig()

	startTime := time.Now()
  fmt.Println("Iniciando SNMP Walk...", startTime.Format(time.RFC3339))
  fmt.Printf("\n")
  
  // Configuration parameters for SNMP
	params := &gosnmp.GoSNMP{
    Target:    HOST_IP,
		Port:      PORT, 
		Community: COMMUNITY, 
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		// Retries: 3,
	}

	// connect to the SNMP agent
	err := params.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer params.Conn.Close()

  // info_oid := "iso.3.6.1.2.1.1.1.0"
  index_oid := "1.3.6.1.4.1.637.61.1.35.10.1.1.24"
  index_result := walk(params, index_oid)

  for oid, value := range index_result {
    fmt.Printf("%s: %v\n", oid, value)
  }

	// show results
	// for oid, value := range results {
	// 	fmt.Printf("%s: %v\n", oid, value)
	// }
  // fmt.Printf("\n")

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

func LoadConfig(path ...string) {
  envPath := ".env"
  if len(path) > 0 && path[0] != "" {
    envPath = path[0]
  }
  
  err := godotenv.Load(envPath)
  if err != nil {
    fmt.Println("Erro ao carregar o arquivo .env")
    os.Exit(1)
  }

  HOST_IP = os.Getenv("HOST_IP")
  COMMUNITY = os.Getenv("COMMUNITY")
  
  portStr := os.Getenv("PORT")
  portInt, err := strconv.Atoi(portStr)

  if err != nil {
    fmt.Printf("Erro ao converter PORT para inteiro: %v\n", err)
    os.Exit(1)
  }
 
  PORT = uint16(portInt)

  SERIAL = os.Getenv("SERIAL")
}
