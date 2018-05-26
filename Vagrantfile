# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.ssh.forward_agent = true

  config.vm.define "router" do |router|
    router.vm.synced_folder '.', '/vagrant', disabled: true
    router.vm.hostname = "router"
    router.vm.box = "tmatilai/openbsd-5.6"

    # External
    router.vm.network :private_network, ip: "10.0.10.2",
    :netmask => "255.255.255.0"
    # Internal
    router.vm.network :private_network, ip: "172.0.1.2",
    :netmask => "255.255.255.0"
    router.vm.provider "virtualbox" do |v|
      v.memory = 512
      v.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
      v.customize ["modifyvm", :id, "--nicpromisc3", "allow-all"]
    end

    router.vm.provision "shell",
    inline: "echo 'nameserver 8.8.8.8' > /etc/resolv.conf"

    router.vm.provision "shell",
    inline: "export PKG_PATH=http://ftp.eu.openbsd.org/pub/OpenBSD/`uname -r`/packages/`arch -s` && pkg_add python-2.7.8"
  end

  config.vm.define "router2" do |router2|
    router2.vm.synced_folder '.', '/vagrant', disabled: true
    router2.vm.hostname = "router2"
    router2.vm.box = "tmatilai/openbsd-5.6"

    # External
    router2.vm.network :private_network, ip: "10.0.20.2",
    :netmask => "255.255.255.0"
    # Internal
    router2.vm.network :private_network, ip: "172.0.2.2",
    :netmask => "255.255.255.0"
    router2.vm.provider "virtualbox" do |v|
      v.memory = 512
      v.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
      v.customize ["modifyvm", :id, "--nicpromisc3", "allow-all"]
    end

    router2.vm.provision "shell",
    inline: "echo 'nameserver 8.8.8.8' > /etc/resolv.conf"

    router2.vm.provision "shell",
    inline: "export PKG_PATH=http://ftp.eu.openbsd.org/pub/OpenBSD/`uname -r`/packages/`arch -s` && pkg_add python-2.7.8"
  end

  config.vm.define "zone0" do |zone0|
    zone0.vm.hostname = "zone0"
    zone0.ssh.insert_key = false
    zone0.vm.box = "ubuntu/trusty64"
    zone0.vm.network :private_network, ip: "172.0.1.10",
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
    zone1.vm.network :private_network, ip: "172.0.2.10",
    :netmask => "255.255.255.0"
    zone1.vm.provider "virtualbox" do |v|
      v.memory = 768
      v.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
    end
  end

  config.vm.define "client" do |client|
    client.vm.hostname = "client"
    client.ssh.insert_key = false
    client.vm.box = "ubuntu/trusty64"
    client.vm.network :private_network, ip: "10.0.10.10",
    :netmask => "255.255.255.0"
    client.vm.provider "virtualbox" do |v|
      v.memory = 512
    end
  end

  config.vm.define "client2" do |client2|
    client2.vm.hostname = "client2"
    client2.ssh.insert_key = false
    client2.vm.box = "ubuntu/trusty64"
    client2.vm.network :private_network, ip: "10.0.20.10",
    :netmask => "255.255.255.0"
    client2.vm.provider "virtualbox" do |v|
      v.memory = 512
    end
  end

  config.vm.define "client3" do |client3|
    client3.vm.hostname = "client3"
    client3.ssh.insert_key = false
    client3.vm.box = "ubuntu/trusty64"
    client3.vm.network :private_network, ip: "10.0.20.11",
    :netmask => "255.255.255.0"
    client3.vm.provider "virtualbox" do |v|
      v.memory = 512
    end
  end

end
