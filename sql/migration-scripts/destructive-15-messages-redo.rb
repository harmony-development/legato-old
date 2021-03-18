#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{

ALTER TABLE Messages DROP COLUMN Content;
ALTER TABLE Messages DROP COLUMN Embeds;
ALTER TABLE Messages DROP COLUMN Actions;
ALTER TABLE Messages DROP COLUMN Overrides;
ALTER TABLE Messages DROP COLUMN Attachments;

ALTER TABLE Messages DROP COLUMN Attachments;

ALTER TABLE Messages ADD Content bytea;

}
end