<?php
$host = '';
$community = 'public';
$serial_filter = "414C434CFC653AE2";

$startTime = microtime(true);
echo "Iniciando SNMP Walk... " . gmdate('Y-m-d\TH:i:s\Z', (int)$startTime) . "\n";
echo "\n";

$Index     = walk($host, $community, "1.3.6.1.4.1.637.61.1.35.10.1.1.24");

foreach ($Index as $key => $val) {
  preg_match('/(\d+)(?!.*\d)/', $key, $matches);
  $Client = $matches[0];
  $F1     = $matches[0];

  $serial = str_replace("Hex-STRING: ", "", get($host, $community, "1.3.6.1.4.1.637.61.1.35.10.1.1.5.$Client"));
  $serial = str_replace("\"", "", $serial);
  $serial = str_replace(" ", "", $serial);

  if ($serial != $serial_filter) {
    continue;
  }

  $rx             = str_replace("INTEGER: ", "", get($host, $community, "1.3.6.1.4.1.637.61.1.35.10.14.1.2.$Client"));
  $voltagem       = str_replace("INTEGER: ", "", get($host, $community, "1.3.6.1.4.1.637.61.1.35.10.14.1.3.$Client"));
  $tx             = str_replace("INTEGER: ", "", get($host, $community, "1.3.6.1.4.1.637.61.1.35.10.14.1.4.$Client"));
  $bias           = str_replace("INTEGER: ", "", get($host, $community, "1.3.6.1.4.1.637.61.1.35.10.14.1.5.$Client"));
  $temperatura    = str_replace("INTEGER: ", "", get($host, $community, "1.3.6.1.4.1.637.61.1.35.10.14.1.6.$Client"));
  $nome           = str_replace("STRING: ", "", get($host, $community, "1.3.6.1.4.1.637.61.1.35.10.1.1.24.$Client"));

  $nome = str_replace("\"", "", $nome);

  echo "ONU: $v\n";
  echo "Serial: $serial\n";
  echo "Nome: $nome\n";
  echo "RX: $rx\n";
  echo "TX: $tx\n";
  echo "Voltagem: $voltagem mV\n";
  echo "Bias: $bias uA\n";
  echo "Temperatura: $temperatura °C\n";
  echo "-------------------------------------\n";
}

echo "\n";

$endTime = microtime(true);
$executionTime = round(($endTime - $startTime) * 1000, 2);

echo "SNMP Walk concluído em " . gmdate('Y-m-d\TH:i:s\Z', (int)$endTime) . "\n";
echo "\033[1mTempo total $executionTime ms\033[0m\n";

/**
 * Função para realizar SNMP Walk
 *
 * @param string $host Endereço do host SNMP
 * @param string $community Comunidade SNMP
 * @param string $oid OID a ser consultado
 * @return array Resultado do SNMP Walk
 */
function walk($host, $community, $oid)
{
  $result = snmp2_real_walk($host, $community, $oid, 10000000);
  if ($result === false) {
    echo "Erro na consulta SNMP para OID: $oid\n";
  }

  return $result;
}

function get($host, $community, $oid)
{
  $result = snmp2_get($host, $community, $oid, 1000000, 5);
  if ($result === false) {
    echo "Erro na consulta SNMP para OID: $oid\n";
    return '';
  }

  return $result;
}
