#!/usr/bin/env bash

apt-get update
apt-get install -y apache2 apache2-dev
if ! [ -L /var/www ]; then
  rm -rf /var/www
  ln -fs /vagrant /var/www
fi

printf "\nServerName localhost" | sudo tee --append /etc/apache2/httpd.conf
printf "\nServerName localhost" | sudo tee --append /etc/apache2/apache2.conf
