#!/usr/bin/env ruby

require 'pg'
require 'json'

CURRENT_VERSION = 11

def get_config
    return JSON.parse(File.read("config.json"))
end

def get_version
    begin
        return JSON.parse(File.read(".version.json"))["version"]
    rescue => exception
        return CURRENT_VERSION
    end
end

def set_version(vers)
    File.open(".version.json", "w") do |f|
        f.write(JSON.pretty_generate({ :version => vers }))
    end
end

def connect
    config = get_config()["instanceserver"]["DB"]
    con = PG.connect :dbname => "harmony",
                     :user => config["User"],
                     :password => config["Password"],
                     :port => config["Port"],
                     :host => config["Host"]

    return con
end
