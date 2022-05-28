package mp

import "github.com/spf13/viper"

var Config = Options{}

func ReadConfig() {
	Config = Options{
		RunMode: viper.GetString("runmode"),
		Mp: MpOption{
			AppID:          viper.GetString("mp.appid"),
			AppSecret:      viper.GetString("mp.appsecret"),
			Token:          viper.GetString("mp.token"),
			EncodingAesKey: viper.GetString("mp.encodingaeskey"),
		},
		MongoDB: MongoDBOption{
			URI: viper.GetString("mongo.uri"),
		},
	}
}

type Options struct {
	RunMode string
	Mp      MpOption
	MongoDB MongoDBOption
}

type MpOption struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAesKey string
}

type MongoDBOption struct {
	URI string
}
