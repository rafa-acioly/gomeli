package environments

type site string

const (
	ARGENTINA  site = "MLA"
	BOLIVIA    site = "MBO"
	BRASIL     site = "MLB"
	COLOMBIA   site = "MCO"
	COSTA_RICA site = "MCR"
	EQUADOR    site = "MEC"
	HONDURAS   site = "MHN"
	GUATEMALA  site = "MGT"
	CHILE      site = "MLC"
	MEXICO     site = "MLM"
	NICARAGUA  site = "MNI"
	PARAGUAI   site = "MPY"
	SALVADOR   site = "MSV"
	URUGUAI    site = "MLU"
	VENEZUELA  site = "MLV"
	PANAMA     site = "MPA"
	PERU       site = "MPE"
	PORTUGAL   site = "MPT"
	DOMINICANA site = "MRD"
)

type Storage interface {
	Set(name, value string) error
	Get(name string) interface{}
	Has(name string) bool
	Remove(name string) error
}

type Configuration interface {
	GetStorage() Storage
}

type configuration struct {
	storage Storage
}

func (c configuration) GetStorage() Storage {
	return c.storage
}

func NewConfiguration(storage Storage) Configuration {
	return &configuration{
		storage: storage,
	}
}

type EnvironmentConfig interface {
	GetSite() site
	GetConfiguration() Configuration
}

type environment struct {
	site          site
	configuration Configuration
}

func (e environment) GetSite() site {
	return e.site
}

func (e environment) GetConfiguration() Configuration {
	return e.configuration
}

func NewEnvironmentConfig(site site, configuration Configuration) EnvironmentConfig {
	return &environment{
		site:          site,
		configuration: configuration,
	}
}
