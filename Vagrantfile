VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.provider "libvirt"

  config.vm.define "rhel7_minion" do |rhel7_minion|
    rhel7_minion.vm.box = "alpine/alpine64"
    rhel7_minion.vm.synced_folder '.', '/vagrant', disabled: true
    config.vm.network "private_network",
        :libvirt__network_name => 'primary_vagrant_private_net',
        ip: '192.168.1.2', auto_config: false,
        :libvirt__netmask => '255.255.255.0',
        :libvirt__forward_mode => 'route',
        :libvirt__dhcp_enabled => true
  end

  config.vm.define "rhel5_minion" do |rhel5_minion|
    rhel5_minion.vm.box = "alpine/alpine64"
    rhel5_minion.vm.synced_folder '.', '/vagrant', disabled: true
    rhel5_minion.vm.network 'private_network',
        :libvirt__network_name => 'secondary_vagrant_private_net',
        ip: '192.168.2.3', auto_config: false,
        :libvirt__netmask => '255.255.255.0',
        :libvirt__forward_mode => 'route',
        :libvirt__dhcp_enabled => true
  end

end
