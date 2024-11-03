package configs

type AppConfigs struct {
	FPS              int
	InputRefreshRate int
}

func NewAppConfigs() *AppConfigs {
	return &AppConfigs{
		FPS:              15,
		InputRefreshRate: 10,
	}
}
