<?php
$host = '127.0.0.1';
$community = 'public';

$sys_descr_oid = '.1.3.6.1.2.1.1.1.0';
$sys_object_id_oid = '.1.3.6.1.2.1.1.2.0';
$sys_uptime_oid = '.1.3.6.1.2.1.1.3.0';
$sys_contact_oid = ".1.3.6.1.2.1.1.4.0";
$sys_name_oid = ".1.3.6.1.2.1.1.5.0";
$sys_location_oid = ".1.3.6.1.2.1.1.6.0";

$startTime = microtime(true);
echo "Iniciando SNMP Walk... " . gmdate('Y-m-d\TH:i:s\Z', (int)$startTime) . "\n";
echo "\n";

$sys_descr = walk($host, $community, $sys_descr_oid);
$sys_object_id = walk($host, $community, $sys_object_id_oid);
$sys_uptime = walk($host, $community, $sys_uptime_oid);
$sys_contact = walk($host, $community, $sys_contact_oid);
$sys_name = walk($host, $community, $sys_name_oid);
$sys_location = walk($host, $community, $sys_location_oid);

$result = [];
$result[$sys_descr_oid] = $sys_descr;
$result[$sys_object_id_oid] = $sys_object_id;
$result[$sys_uptime_oid] = $sys_uptime;
$result[$sys_contact_oid] = $sys_contact;
$result[$sys_name_oid] = $sys_name;
$result[$sys_location_oid] = $sys_location;

print_r($result);
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
  $result = snmp2_real_walk($host, $community, $oid);
  if ($result === false) {
    echo "Erro na consulta SNMP para OID: $oid\n";
    exit(1);
  }

  return $result;
}
