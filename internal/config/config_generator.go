package conf

type ConfigGenerator struct {
	Generator generator
}

type generator struct {
	PathToGeneratedServer string
	PathToHandlers        string
}
