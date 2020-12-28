#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{
ALTER TABLE Guilds ADD Metadata bytea;

ALTER TABLE Channels DROP COLUMN Kind;

ALTER TABLE Channels ADD Metadata bytea;

ALTER TABLE Messages ADD Metadata bytea;
}
end