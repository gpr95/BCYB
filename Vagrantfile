VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.provider "virtualbox"
  default_router = "172.20.10.1"


  config.vm.define "nat_n1" do |nat_n1|
    nat_n1.vm.box = "alpine/alpine64"
    nat_n1.vm.synced_folder '.', '/vagrant', disabled: true
    config.vm.network "public_network", ip: "192.168.3.1", auto_config: false
    config.vm.provision :shell, :inline => "ip route delete default 2>&1 >/dev/null || true; ip route add default via #{default_router}"
  end

  config.vm.define "nat_n2" do |nat_n2|
    nat_n2.vm.box = "alpine/alpine64"
    nat_n2.vm.synced_folder '.', '/vagrant', disabled: true
    config.vm.network "public_network", ip: "192.168.4.1", auto_config: false
    config.vm.provision :shell, :inline => "ip route delete default 2>&1 >/dev/null || true; ip route add default via #{default_router}"
  end

end
