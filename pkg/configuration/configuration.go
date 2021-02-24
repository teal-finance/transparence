package configuration

type ConfigValue struct {
	RPC         []RPC         `mapstructure:"rpc"`
	Useraddress []UserAddress `mapstructure:"useraddress"`
	Tokens      []Token       `mapstructure:"tokens"`
}

type RPC struct {
	Blockchain string `mapstructure:"blockchain"`
	Link       string `mapstructure:"link"`
	Userpass   string `mapstructure:"userpass"`
}

type UserAddress struct {
	Blockchain string `mapstructure:"blockchain"`
	Address    string `mapstructure:"address"`
}

type Token struct {
	Blockchain string `mapstructure:"blockchain"`
	Name       string `mapstructure:"name"`
	Symbol     string `mapstructure:"symbol"`
	Address    string `mapstructure:"address"`
}
