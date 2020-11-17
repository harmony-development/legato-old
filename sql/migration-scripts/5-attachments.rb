#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{
DROP TABLE Attachments;
ALTER TABLE Messages
    ADD Attachments text[];
}
end