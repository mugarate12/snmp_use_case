FROM php:8.2-cli

# Install SNMP extension and net-snmp tools
RUN apt-get update
RUN apt-get install -y snmp
RUN apt-get install -y --fix-missing libsnmp-dev
RUN apt-get install -y unzip
RUN apt-get install -y git
RUN docker-php-ext-install snmp
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

WORKDIR /app
COPY main.php /app/main.php

RUN composer global require phpunit/phpunit --prefer-dist --no-interaction --no-progress

CMD ["php", "main.php"]
