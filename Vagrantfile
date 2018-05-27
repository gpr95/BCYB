# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.ssh.forward_agent = true

  config.vm.define "r1" do |r1|
    r1.vm.synced_folder '.', '/vagrant', disabled: true
    r1.vm.hostname = "r1"
    r1.vm.box = "tmatilai/openbsd-5.6"

    # External
    r1.vm.network :private_network, ip: "10.0.10.2",
    :netmask => "255.255.255.0"
    # Internal
    r1.vm.network :private_network, ip: "172.0.1.2",
    :netmask => "255.255.255.0"
    r1.vm.provider "virtualbox" do |v|
      v.memory = 512
      v.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
      v.customize ["modifyvm", :id, "--nicpromisc3", "allow-all"]
    end

    r1.vm.provision "shell",
    inline: "echo 'nameserver 172.0.1.10' > /etc/resolv.conf"

    r1.vm.provision "shell",
    inline: "export PKG_PATH=http://ftp.eu.openbsd.org/pub/OpenBSD/`uname -r`/packages/`arch -s` && pkg_add python-2.7.8"
  end

  config.vm.define "r2" do |r2|
    r2.vm.synced_folder '.', '/vagrant', disabled: true
    r2.vm.hostname = "r2"
    r2.vm.box = "tmatilai/openbsd-5.6"

    # External
    r2.vm.network :private_network, ip: "10.0.20.2",
    :netmask => "255.255.255.0"
    # Internal
    r2.vm.network :private_network, ip: "172.0.2.2",
    :netmask => "255.255.255.0"
    r2.vm.provider "virtualbox" do |v|
      v.memory = 512
      v.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
      v.customize ["modifyvm", :id, "--nicpromisc3", "allow-all"]
    end

    r2.vm.provision "shell",
    inline: "echo 'nameserver 172.0.1.10' > /etc/resolv.conf"

    r2.vm.provision "shell",
    inline: "export PKG_PATH=http://ftp.eu.openbsd.org/pub/OpenBSD/`uname -r`/packages/`arch -s` && pkg_add python-2.7.8"
  end

  config.vm.define "zone0" do |zone0|
    zone0.vm.hostname = "zone0"
    zone0.ssh.insert_key = false
    zone0.vm.box = "ubuntu/trusty64"
    zone0.vm.network :private_network, ip: "172.0.1.10",
    :netmask => "255.255.255.0"
    zone0.vm.network :private_network, ip: "172.0.2.10",
    :netmask => "255.255.255.0"
    zone0.vm.provider "virtualbox" do |v|
      v.memory = 768
      v.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
    end
  end

  config.vm.define "zone1" do |zone1|
    zone1.vm.hostname = "zone1"
    zone1.ssh.insert_key = false
    zone1.vm.box = "ubuntu/trusty64"
    zone1.vm.network :private_network, ip: "172.0.1.11",
    :netmask => "255.255.255.0"
    zone1.vm.network :private_network, ip: "172.0.2.11",
    :netmask => "255.255.255.0"
    zone1.vm.provider "virtualbox" do |v|
      v.memory = 768
      v.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
    end
  end

  config.vm.define "c1" do |c1|
    c1.vm.hostname = "c1"
    c1.ssh.insert_key = false
    c1.vm.box = "ubuntu/trusty64"
    c1.vm.network :private_network, ip: "10.0.10.10",
    :netmask => "255.255.255.0"
    c1.vm.provider "virtualbox" do |v|
      v.memory = 512
    end
  end

  config.vm.define "c2" do |c2|
    c2.vm.hostname = "c2"
    c2.ssh.insert_key = false
    c2.vm.box = "ubuntu/trusty64"
    c2.vm.network :private_network, ip: "10.0.20.10",
    :netmask => "255.255.255.0"
    c2.vm.provider "virtualbox" do |v|
      v.memory = 512
    end
  end

  config.vm.define "c3" do |c3|
    c3.vm.hostname = "c3"
    c3.ssh.insert_key = false
    c3.vm.box = "ubuntu/trusty64"
    c3.vm.network :private_network, ip: "10.0.20.11",
    :netmask => "255.255.255.0"
    c3.vm.provider "virtualbox" do |v|
      v.memory = 512
    end
  end

end
