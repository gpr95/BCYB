VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.define "n1" do |n1|
	n1.vm.box = "gbarbieru/xenial"
	n1.vm.network "private_network" , ip: "10.9.8.1",
		virtualbox__intnet: "intnet-1"
  end
 
  config.vm.define "n2" do |n2|
	n2.vm.box = "gbarbieru/xenial"
	n2.vm.network "private_network" , ip: "10.9.9.1",
		virtualbox__intnet: "intnet-1"
  end
  
  config.vm.define "n3" do |n3|
	n3.vm.box = "gbarbieru/xenial"
	n3.vm.network "private_network" , ip: "10.9.10.1",
		virtualbox__intnet: "intnet-2"
  end
end
