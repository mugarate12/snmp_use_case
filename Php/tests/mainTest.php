<?php

use PHPUnit\Framework\TestCase;

require_once __DIR__ . '/../main.php';

class MainTest extends TestCase
{
  public function testWalk()
  {
    $host = '127.0.0.1';
    $community = 'public';
    $sys_descr_oid = '.1.3.6.1.2.1.1.1.0';
    $result = walk($host, $community, $sys_descr_oid);

    $this->assertIsArray($result, "O resultado deve ser um array");
    $this->assertNotEmpty($result, "O resultado não deve estar vazio");
  }
}
