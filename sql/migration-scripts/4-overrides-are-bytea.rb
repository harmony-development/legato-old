#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{
ALTER TABLE Messages
    ALTER COLUMN Overrides TYPE bytea USING NULL;
}
end