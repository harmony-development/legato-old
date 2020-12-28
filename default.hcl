# Configuration for the software itself; not external dependencies
# Note that these fields are the defaults; you can omit them if you
# do not wish to change them.
Server {
	# The port the server should listen to. Clients default to 2289 if not explicitly
	# specified in URLs, so it should be left at 2289.
	Port = 2289

	# Enables the CORS in REST, useful for browser support
	# you should probably disable this if using a proxy
	UseCORS = true

	# The location of the private and public keys of the server; used to establish
	# trust with other servers when federating.
	PrivateKeyPath = "harmony-key.pem"
	PublicKeyPath = "harmony-key.pub"

	# How Legato should store attachments. There are two backends:
	#
	# - PureFlatfile: This backend exclusively uses a flatfile directory layout
	# - DatabaseFlatfile: This backend stores files in a flatfile directory layout, but stores metadata in the database.
	StorageBackend = "PureFlatfile"

	# The starting point for IDs of messages, channels, guilds, etc.
	SnowflakeStart = 0

	# Assorted opinion-related things for administrators to play with.
	Policies {
		# Avatar-related policies
		Avatar {
			# The geometry avatars should be stored at
			Width = 256
			Height = 256

			# The JPEG quality avatars will be stored with
			Quality = 50

			# Whether or not avatars will be cropped into squares by the server
			Crop = true
		}

		# Requirements for usernames
		Username {
			MinLength = 2
			MaxLength = 20
		}

		# Requirements for passwords. See Unicode for definitions of which
		# characters count as lowercase, uppercase, numbers, and symbols.
		Password {
			MinLength = 5
			MaxLength = 256
			MinLower = 1
			MinUpper = 1
			MinNumbers = 1
			MinSymbols = 0
		}

		# Limitations on attachments
		Attachments {
			# How many attachments a single message can have
			MaximumCount = 10
		}

		# Debug-related policies; useful for developing Legato
		Debug {
			LogErrors = true
			LogRequests = true
			RespondWithErrors = false
			ResponseErrorsIncludeTrace = false

			# This will cause stream-related events to become VERY verbose.
			VerboseStreamHandling = false
		}

		# Session-related policies
		Sessions {
			# How long user sessions last in nanoseconds. The default is 48 hours.
			Duration = 172800000000000
		}

		# Limitations on how many items the in-memory caches of Legato can hold.
		# When Legato runs into the item ceiling, least recently used items are
		# removed from memory and must be loaded from the database when referenced
		# in the future.
		Caches {
			# Cached information on guild owners
			Owner = 5096

			# Cached sessions; sessions not in cache will be loaded from database
			Sessions = 5096
		}

		# Policies for specific endpoints
		APIs {
			# Policies for API endpoints related to messages
			Messages {
				# The maximum amount of messages a client can request at once
				MaximumGetAmount = 50
			}
		}

		# Federation-related policies
		Federation {
			# The size of the nonce used to establish a relationship between
			# the foreign homeserver a user of your homeserver is connecting to
			# and your homeserver. Larger nonces are more secure, but take more
			# time to generate.
			NonceLength = 32

			# How large the queue of guild left notifications can get.
			GuildLeaveNotificationQueueLength = 64
		}
	}
}

# Database-related configuration. This should point to a Postgres database.
Database {
	Host = "127.0.0.1"
	Username = ""
	Password = ""
	Port = 5432
	SSL = false
	Name = "harmony"
}

# Flatfile-related configuration
Flatfile {

	# A relative or absolute directory path to where files will be stored
	MediaPath = "flatfile"
}

# Integration with Sentry for monitoring your Legato instance. See Sentry
# documentation for what value the DSN key should hold.
Sentry {
	Enabled = false
	DSN = ""

	# Whether or not to attach stack traces of errors when reporting them to Sentry.
	AttachStacktraces = true
}
