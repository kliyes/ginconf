package config

var defaultConfig = []byte(
	`[base]
	 listen_address = :8080
	 mode = test
	 
	 [database]
	 host = localhost
	 port = 3306
	 name = project
	 username = root
	 password = root
	 charset = utf8
	 locale = UTC
	`,
)
