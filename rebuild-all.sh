#!/bin/bash
n=0
until [ $n -ge 5 ]
do
   vagrant destroy -f && break  # substitute your command here
   n=$[$n+1]
   sleep 15
done

n=0
until [ $n -ge 5 ]
do
   vagrant up && break  # substitute your command here
   n=$[$n+1]
   sleep 15
done

n=0
until [ $n -ge 5 ]
do
   ansible -m ping all && break  # substitute your command here
   n=$[$n+1]
   sleep 15
done

n=0
until [ $n -ge 5 ]
do
   ansible-playbook site.yml && break  # substitute your command here
   n=$[$n+1]
   sleep 15
done

notify-send "Internal network have been rebuild.. BCYB more.."
