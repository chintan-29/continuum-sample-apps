# Copyright 2015 Apcera Inc. All rights reserved.

require 'json'
require 'fileutils'

class VagrantHelper
    VMWARE_BOX = "continuum_trial_vagrant_vmware_1436979246"
    VMWARE_CHECKSUM256 = "5b9771390865f43a5262d2f348a25a6909b834adb16be789c2b2a75b1d7c1e85"
    VIRTUALBOX_BOX = "continuum_trial_vagrant_virtualbox_1436979246"
    VIRTUALBOX_CHECKSUM256 = "a0908aea1b18d46c7b7972c40cd8ecf549e0cea33c02ad4da85f6fb1ad8ecbf0"

  def initialize
    find_provider
    validate
  end

  def validate
    validate_config
  end

  def validate_config
    config = File.join(File.dirname(__FILE__), 'continuum-setup.conf')
    unless File.file?(config)
      $stderr.puts %^Please copy continuum-setup.conf.template to continuum-setup.conf and update the value of cluster domain prefix name.^
      exit(1)
    end 
  end

  def setup_box(config)
    if virtualbox?
      config.vm.box_url = ENV['VAGRANT_BOX_URL'] ? ENV['VAGRANT_BOX_URL'] : "https://d1g23m5c3hjtbv.cloudfront.net/#{VIRTUALBOX_BOX}.box"
      config.vm.box_download_checksum_type = 'sha256'
      config.vm.box_download_checksum = VIRTUALBOX_CHECKSUM256
    elsif fusion? || workstation?
      config.vm.box_url = ENV['VAGRANT_BOX_URL'] ? ENV['VAGRANT_BOX_URL'] : "https://d1g23m5c3hjtbv.cloudfront.net/#{VMWARE_BOX}.box"
      config.vm.box_download_checksum_type = 'sha256'
      config.vm.box_download_checksum = VMWARE_CHECKSUM256
    else
      puts "Default provider not found."
      exit(1)
    end
  end

  def fusion_options(v)
    vm_options(v)
    v.vmx['ethernet0.virtualDev'] = 'vmxnet3'
    v.vmx['ethernet1.virtualDev'] = 'vmxnet3'
  end

  def workstation_options(v)
    vm_options(v)
    v.vmx['ethernet0.virtualDev'] = 'vmxnet3'
    v.vmx['ethernet1.virtualDev'] = 'vmxnet3'
    v.vmx['mainMem.freeSpaceCheck'] = 'FALSE'
  end

  def virtualbox_options(v)
    v.name = "continuum_trial"
    v.cpus = 2
    v.memory = 4096
    v.gui = ENV['VAGRANT_GUI'] ? true : false
  end

  def vm_options(v)
    v.vmx['displayName'] = 'continuum_trial'
    v.vmx['numvcpus'] = ENV['SYSTEM_TESTS_CPUS'] || '2'
    v.vmx['memsize'] = ENV['SYSTEM_TESTS_MEMORY'] || '4092'
    v.gui = ENV['VAGRANT_GUI'] ? true : false
  end

  def share_folder(config)
    config.vm.synced_folder "..", "/home/vagrant/continuum-sample-apps"
  end

  def scripts(config)
    config.vm.provision "shell", path: File.join('.', 'setup.rb')
  end

  def find_provider
    @provider_type = get_provider
    ENV['VAGRANT_DEFAULT_PROVIDER'] = 'virtualbox' if ENV['VAGRANT_DEFAULT_PROVIDER'].nil?
    @provider_type = @provider_type || ENV['VAGRANT_DEFAULT_PROVIDER']
    # Default to the correct provider if none is given.
    ENV['VAGRANT_DEFAULT_PROVIDER'] = provider_type
  end

  def get_provider
    idx = 0
    ARGV.each do |p|
      next if p.nil?
      if p.index("--provider") == 0 && (idx + 1) <= ARGV.length
        return ARGV[idx + 1]
      elsif p.include? "--provider="
        return p.split('=')[-1]
      end
      idx = idx + 1
    end
    nil
  end

  def provider_type
    @provider_type || 'virtualbox'
  end

  def fusion?
    @provider_type == 'vmware_fusion'
  end

  def workstation?
    @provider_type == 'vmware_workstation'
  end

  def virtualbox?
    @provider_type == 'virtualbox'
  end
end
