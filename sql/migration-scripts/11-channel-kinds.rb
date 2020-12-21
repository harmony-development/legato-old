#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{
ALTER TABLE Channels
    ADD Kind TEXT;

UPDATE Channels SET Kind = 'h.voice' WHERE IsVoice = true;

ALTER TABLE Channels
    DROP COLUMN IsVoice;
}
end