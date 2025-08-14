FROM php:8.2-cli

# Install SNMP extension and net-snmp tools
RUN apt-get update
RUN apt-get install -y snmp
RUN apt-get install -y --fix-missing libsnmp-dev
RUN docker-php-ext-install snmp

WORKDIR /app
COPY main.php /app/main.php

CMD ["php", "main.php"]
