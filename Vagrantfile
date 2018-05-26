VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.provider "virtualbox" do |v| 
	v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
  end

  config.vm.define "inetgateway" do |ig|
	ig.vm.box = "gbarbieru/xenial"
	ig.vm.hostname = "internet-gateway"
	ig.vm.network "private_network", ip: "10.9.8.8", virtualbox__intnet: "internet-1"
	ig.vm.network "private_network", ip: "10.9.9.8", virtualbox__intnet: "internet-2"
  end
  config.vm.define "router1" do |router1|
	router1.vm.box = "gbarbieru/xenial"
	router1.vm.hostname = "router1"
	router1.vm.network "private_network", ip: "10.9.8.7", virtualbox__intnet: "internet-1"
	router1.vm.network "private_network", ip: "10.9.7.7", virtualbox__intnet: "internet-3"
	router1.vm.network "private_network", ip: "172.16.1.1", virtualbox__intnet: "intnet-1"
	router1.vm.network "private_network", ip: "172.16.3.1", virtualbox__intnet: "intnet-3"
  end
  config.vm.define "router2" do |router2|
	router2.vm.box = "gbarbieru/xenial"
	router2.vm.hostname = "router2"
	router2.vm.network "private_network", ip: "10.9.9.9", virtualbox__intnet: "internet-2"
	router2.vm.network "private_network", ip: "10.9.7.9", virtualbox__intnet: "internet-3"
	router2.vm.network "private_network", ip: "172.16.2.1", virtualbox__intnet: "intnet-2"
  end

  config.vm.define "n1" do |n1|
	n1.vm.box = "gbarbieru/xenial"
	n1.vm.hostname = "n1"
	n1.vm.network "private_network" , ip: "172.16.1.2",
		virtualbox__intnet: "intnet-1"
  end
  config.vm.define "n2" do |n2|
	n2.vm.box = "gbarbieru/xenial"
	n2.vm.hostname = "n2"
	n2.vm.network "private_network" , ip: "172.16.3.3",
		virtualbox__intnet: "intnet-3"
  end
  config.vm.define "n3" do |n3|
	n3.vm.box = "gbarbieru/xenial"
	n3.vm.hostname = "n3"
	n3.vm.network "private_network" , ip: "172.16.2.2",
		virtualbox__intnet: "intnet-2"
  end
end
