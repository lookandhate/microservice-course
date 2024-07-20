package config

// GRPCConfig config for GRPC server.
type GRPCConfig struct {
	Port int `yaml:"port" env-default:"8080" env:"GRPC_PORT"`
}
