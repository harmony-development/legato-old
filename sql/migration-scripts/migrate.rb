require __dir__+'/migration_utils.rb'

vers = get_version()

if vers < CURRENT_VERSION
    ((vers+1)..CURRENT_VERSION).to_a.each do |version|
        puts "Migrating from #{version-1} -> #{version}..."
        system "ruby #{__dir__}/#{version}*.rb", exception: true
        set_version(CURRENT_VERSION)
        puts "Migrated to #{version}!"
    end
else
    puts "Nothing to migrate!"
end
