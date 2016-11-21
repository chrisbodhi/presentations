require 'fileutils'
require 'securerandom'

def clean_up
  uuid = SecureRandom.uuid
  new_filename = "glitch-#{uuid}.gif"

  FileUtils.remove(Dir.glob('artifacts/f*.gif'))
  File.rename('glitch.gif', new_filename)
  FileUtils.mv(new_filename, "glitched/#{new_filename}")
end
