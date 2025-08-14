# SNMP Study Use Case

This repository contains a simple use case to study SNMP (Simple Network Management Protocol) using Go and PHP.

# Requirements

- Docker
- Docker Compose
- Make

# Overview

The use case involves performing SNMP operations to retrieve system information from a device. The code examples provided demonstrate how to perform SNMP walks to retrieve various system OIDs (Object Identifiers) such as system description, system object ID, system uptime, system contact, and system location.

### SNMP Simulator

It is recommended to use an SNMP simulator for testing purposes. The [snmpsim](https://hub.docker.com/r/tandrup/snmpsim) Docker image can be used to simulate SNMP responses.

### Go Implementation

This implementation uses the `gosnmp` library to perform SNMP operations. The main function initializes the SNMP parameters, performs a walk operation on the specified OIDs, and prints the results.

### PHP Implementation

This implementation uses the built-in SNMP functions in PHP to perform SNMP operations. It defines a function to perform a walk operation on the specified OIDs and prints the results.

# Usage

## Snmp Simulator

To run the SNMP simulator, use the following command:

```bash
make snmp-simulator
```

## Go Implementation

To build the Go implementation, use the following command:

```bash
make build-go
```

To run the Go implementation, use the following command:

```bash
make run-go
```

## PHP Implementation

To buid the PHP implementation, use the following command:

```bash
make build-php
```

To run the PHP implementation, use the following command:

```bash
make run-php
```
