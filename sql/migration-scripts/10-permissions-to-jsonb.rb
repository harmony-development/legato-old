#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{

ALTER TABLE Permissions
    DROP COLUMN Nodes;

ALTER TABLE Permissions
    ADD COLUMN Nodes jsonb NOT NULL DEFAULT '[]';

}
end
