begin
  require 'simplecov'
  SimpleCov.start do
    add_filter '/spec/'
  end
rescue LoadError
end

require 'rspec'
require 'vcr'
require 'webmock/rspec'
require 'capsulecd'

Dir['spec/support/*.rb'].each { |f| require File.expand_path(f) }


python_configuration = CapsuleCD::Configuration.new(config_file: 'spec/fixtures/live_python_configuration.yml')
chef_configuration = CapsuleCD::Configuration.new(config_file: 'spec/fixtures/live_chef_configuration.yml')
node_configuration = CapsuleCD::Configuration.new(config_file: 'spec/fixtures/live_node_configuration.yml')
ruby_configuration = CapsuleCD::Configuration.new(config_file: 'spec/fixtures/live_ruby_configuration.yml')

def configure_vcr(vcr_config, type, capsulecd_config)
  vcr_config.filter_sensitive_data('<GITHUB_ACCESS_TOKEN>', type) do
    capsulecd_config.source_github_access_token
  end
  vcr_config.filter_sensitive_data('<SUPERMARKET_USERNAME>', type) do
    capsulecd_config.chef_supermarket_username
  end
  vcr_config.filter_sensitive_data('<NPM_AUTH_TOKEN>', type) do
    capsulecd_config.npm_auth_token
  end
  vcr_config.filter_sensitive_data('<PYPI_USERNAME>', type) do
    capsulecd_config.pypi_username
  end
  vcr_config.filter_sensitive_data('<PYPI_PASSWORD>', type) do
    capsulecd_config.pypi_password
  end

end

VCR.configure do |c|
  c.debug_logger = File.open('spec/vcr.log', 'w+')
  c.cassette_library_dir = 'spec/fixtures/vcr_cassettes'
  c.hook_into :webmock
  c.configure_rspec_metadata!
  c.preserve_exact_body_bytes { true }

  c.default_cassette_options = {
      record: ENV['TRAVIS'] ? :none : :once
  }

  configure_vcr(c, :python, python_configuration)
  configure_vcr(c, :ruby, ruby_configuration)
  configure_vcr(c, :node, node_configuration)
  configure_vcr(c, :chef, chef_configuration)
end

# This file was generated by the `rspec --init` command. Conventionally, all
# specs live under a `spec` directory, which RSpec adds to the `$LOAD_PATH`.
# The generated `.rspec` file contains `--require spec_helper` which will cause
# this file to always be loaded, without a need to explicitly require it in any
# files.
#
# Given that it is always loaded, you are encouraged to keep this file as
# light-weight as possible. Requiring heavyweight dependencies from this file
# will add to the boot time of your test suite on EVERY test run, even for an
# individual file that may not need all of that loaded. Instead, consider making
# a separate helper file that requires the additional dependencies and performs
# the additional setup, and require it from the spec files that actually need
# it.
#
# The `.rspec` file also contains a few flags that are not defaults but that
# users commonly want.
#
# See http://rubydoc.info/gems/rspec-core/RSpec/Core/Configuration
RSpec.configure do |config|
  # rspec-expectations config goes here. You can use an alternate
  # assertion/expectation library such as wrong or the stdlib/minitest
  # assertions if you prefer.
  config.expect_with :rspec do |expectations|
    # This option will default to `true` in RSpec 4. It makes the `description`
    # and `failure_message` of custom matchers include text for helper methods
    # defined using `chain`, e.g.:
    #     be_bigger_than(2).and_smaller_than(4).description
    #     # => "be bigger than 2 and smaller than 4"
    # ...rather than:
    #     # => "be bigger than 2"
    expectations.include_chain_clauses_in_custom_matcher_descriptions = true
  end

  # rspec-mocks config goes here. You can use an alternate test double
  # library (such as bogus or mocha) by changing the `mock_with` option here.
  config.mock_with :rspec do |mocks|
    # Prevents you from mocking or stubbing a method that does not exist on
    # a real object. This is generally recommended, and will default to
    # `true` in RSpec 4.
    mocks.verify_partial_doubles = true
  end

end
