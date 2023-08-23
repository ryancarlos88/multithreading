package config

import "github.com/spf13/viper"

type Config struct {
	ApiCepPrefix string `mapstructure:"APICEP_PREFIX"`
	ApiCepSuffix string `mapstructure:"APICEP_SUFFIX"`
	ViaCepPrefix string `mapstructure:"VIACEP_PREFIX"`
	ViaCepSufix  string `mapstructure:"VIACEP_SUFFIX"`
}

func LoadConfig(path string) (*Config, error){
	cfg := &Config{}

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}