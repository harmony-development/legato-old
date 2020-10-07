#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{
ALTER TABLE Messages
    ADD COLUMN Reply_To_ID BIGSERIAL;
}
end