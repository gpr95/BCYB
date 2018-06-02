#!/usr/bin/env bash

cd /vagrant/mod_antiloris-0.4
apxs2 -a -i -c mod_antiloris.c
service apache2 restart
