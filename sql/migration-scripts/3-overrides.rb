#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{
ALTER TABLE Messages
    ADD Overrides jsonb;
}
    con.exec %{
UPDATE Messages
    SET Overrides = 'null'
    WHERE Overrides IS NULL;
}
end