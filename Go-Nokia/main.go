package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

  index_oid := "1.3.6.1.4.1.637.61.1.35.10.1.1.24"
  index_result := walk(params, index_oid)

  for oid := range index_result {
    // Extract the last sequence of digits from oid (similar to preg_match('/(\d+)(?!.*\d)/', ...))
    var client string
    oidParts := strings.LastIndexByte(oid, '.')

    if (oidParts == -1) {
      continue;
    }
    client = oid[oidParts+1:]

    serial_oid := fmt.Sprintf("1.3.6.1.4.1.637.61.1.35.10.1.1.5.%s", client)
    serial_result, _ := get(params, serial_oid)
    
    var serialHex string
    if b, ok := serial_result.([]byte); ok {
      serialHex = fmt.Sprintf("%X", b)
    } else {
      serialHex = fmt.Sprintf("%v", serial_result)
    }
    
    if serialHex != SERIAL {
      continue
    }

    rx_oid := fmt.Sprintf("1.3.6.1.4.1.637.61.1.35.10.14.1.2.%s", client)
    tx_oid := fmt.Sprintf("1.3.6.1.4.1.637.61.1.35.10.14.1.4.%s", client)
    volt_oid := fmt.Sprintf("1.3.6.1.4.1.637.61.1.35.10.14.1.3.%s", client)
    bias := fmt.Sprintf("1.3.6.1.4.1.637.61.1.35.10.14.1.5.%s", client)
    temp := fmt.Sprintf("1.3.6.1.4.1.637.61.1.35.10.14.1.6.%s", client)
    name := fmt.Sprintf("1.3.6.1.4.1.637.61.1.35.10.1.1.24.%s", client)

    rx, _ := get(params, rx_oid)
    tx, _ := get(params, tx_oid)
    volt, _ := get(params, volt_oid)
    bias_val, _ := get(params, bias)
    temp_val, _ := get(params, temp)
    name_val_raw, _ := get(params, name)
    
    var name_val string
    if b, ok := name_val_raw.([]byte); ok {
      name_val = string(b)
    } else {
      name_val = fmt.Sprintf("%v", name_val_raw)
    }

    fmt.Printf("Rx Power: %v dBm\n", rx)
    fmt.Printf("Tx Power: %v dBm\n", tx)
    fmt.Printf("Voltage: %v V\n", volt)
    fmt.Printf("Bias Current: %v mA\n", bias_val)
    fmt.Printf("Temperature: %v °C\n", temp_val)
    fmt.Printf("Name: %v\n", name_val)
  }

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

func get(params *gosnmp.GoSNMP, oid string) (interface{}, error) {
  result, err := params.Get([]string{oid})
  if err != nil {
    return nil, fmt.Errorf("erro ao fazer SNMP Get: %v", err)
  }

  if len(result.Variables) == 0 {
    return nil, fmt.Errorf("nenhum resultado encontrado para OID %s", oid)
  }

  // by default, gosnmp returns raw data (in bytes), for example: 
  // if value is "OCTET STRING" it will be returned as []byte (bytes array)
  // if value is "INTEGER" it will be returned as int64
  // if value is "STRING" it will be returned as string
  // if value is "COUNTER" it will be returned as uint64
  return result.Variables[0].Value, nil
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
