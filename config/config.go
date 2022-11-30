package config

type Conf struct {
	RMQ *RMQ
	DB  *Database
}

func NewConf() *Conf {
	return &Conf{
		RMQ: newRMQ(),
		DB:  newDatabase(),
	}
}
