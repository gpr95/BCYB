VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.network "private_network", ip: "192.168.50.4",
    virtualbox__intnet: "mynetwork"

  config.vm.define "nat_1" do |nat_1|
    nat_1.vm.box = "alpine/alpine64"
    nat_1.vm.synced_folder '.', '/vagrant', disabled: true
  end
end
