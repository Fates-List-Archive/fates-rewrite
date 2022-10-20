package config

type Config struct {
	Secrets  Secrets         `yaml:"secrets"`
	Storage  Storage         `yaml:"storage"`
	Servers  Servers         `yaml:"servers"`
	Perms    map[string]Perm `yaml:"perms"`
	Channels Channels        `yaml:"channels"`
	Misc     Misc            `yaml:"misc"`
}

type Secrets struct {
	Token        string `yaml:"token" comment:"Discord main bot token"`
	ClientID     string `yaml:"client_id" default:"1031535884539023370" comment:"Discord Client ID"`
	ClientSecret string `yaml:"client_secret"`
	JApi         string `yaml:"japi_key" comment:"Get this from https://key.japi.rest"`
}

type Storage struct {
	Postgres Postgres `yaml:"postgres"`
	Redis    Redis    `yaml:"redis"`
}

type Postgres struct {
	Host     string `yaml:"host" comment:"Hostname, defaults to PGHOST" required:"false"`
	Port     uint16 `yaml:"port" default:"5432" comment:"Port, defaults to PGPORT" required:"false"`
	User     string `yaml:"user" comment:"Username, defaults to PGUSER" required:"false"`
	Password string `yaml:"password" comment:"Password, defaults to PGPASSWORD" required:"false"`
	Database string `yaml:"database" comment:"Database name"`
}

type Redis struct {
	Host     string `yaml:"host" default:"127.0.0.1" comment:"Hostname, defaults to REDIS_HOST" required:"false"`
	Port     uint16 `yaml:"port" default:"6379" comment:"Port, defaults to REDIS_PORT" required:"false"`
	Password string `yaml:"password" comment:"Password, defaults to REDIS_PASSWORD" required:"false"`
	Database uint32 `yaml:"database" default:"0" comment:"Database, defaults to REDIS_DB" required:"false"`
}

type Servers struct {
	Main uint64 `yaml:"main" default:"789934742128558080" comment:"Main server ID"`
}

type Perm struct {
	Roles []uint64 `yaml:"roles" default:"976891305336655903 # role A, 836326299223195738 # role B etc."`
	Index uint8    `yaml:"index"`
}

type Channels struct {
	BotLogs uint64 `yaml:"bot_logs"`
}

type Misc struct {
	RestrictedVanity []string `yaml:"restricted_vanity" default:"api,docs,add-bot,admin"`
	Test             string   `yaml:"test" default:"test"`
}

func (c Config) Validate() {
	/*if c.Secrets.Token {
		panic("Token is required")
	}*/
	if c.Secrets.ClientID == "" {
		panic("Client ID is required")
	}
	if c.Secrets.ClientSecret == "" {
		panic("Client secret is required")
	}
	if c.Secrets.JApi == "" {
		panic("JApi key is required")
	}
	if c.Storage.Postgres.Database == "" {
		panic("Postgres database is required")
	}
	if c.Servers.Main == 0 {
		panic("Main server is required")
	}
	if c.Channels.BotLogs == 0 {
		panic("Bot logs channel is required")
	}
	if len(c.Misc.RestrictedVanity) == 0 {
		panic("Restricted vanity is required")
	}
}
