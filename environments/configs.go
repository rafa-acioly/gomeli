package environments

type site string

const (
	// Mercado livre does not have a sandbox environment yet,
	// every transaction will be made on the main api.
	// this separation will serve in case they implement a test environment
	wsHostProduction = "https://api.mercadolibre.com"
	wsHostSandbox    = "https://api.mercadolibre.com"

	oauthURIProduction = "/oauth/token"
	oauthURISandbox    = "/oauth/token"

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

// GetWsAuth returns all site endpoint for authentication for each available country,
// since there's no sandbox environment yet, we keep the URLS the same.
// more information: https://api.mercadolibre.com/sites
var (
	wsAuthProduction = map[site]string{
		ARGENTINA:  "https://auth.mercadolibre.com.ar",
		BOLIVIA:    "https://auth.mercadolibre.com.bo",
		BRASIL:     "https://auth.mercadolivre.com.br",
		COLOMBIA:   "https://auth.mercadolibre.com.co",
		COSTA_RICA: "https://auth.mercadolibre.com.cr",
		EQUADOR:    "https://auth.mercadolibre.com.ec",
		HONDURAS:   "https://auth.mercadolibre.com.hn",
		GUATEMALA:  "https://auth.mercadolibre.com.gt",
		CHILE:      "https://auth.mercadolibre.cl",
		MEXICO:     "https://auth.mercadolibre.com.mx",
		NICARAGUA:  "https://auth.mercadolibre.com.ni",
		PARAGUAI:   "https://auth.mercadolibre.com.py",
		SALVADOR:   "https://auth.mercadolibre.com.sv",
		URUGUAI:    "https://auth.mercadolibre.com.uy",
		VENEZUELA:  "https://auth.mercadolibre.com.ve",
		PANAMA:     "https://auth.mercadolibre.com.pa",
		PERU:       "https://auth.mercadolibre.com.pe",
		PORTUGAL:   "https://auth.mercadolibre.com.pt",
		DOMINICANA: "https://auth.mercadolibre.com.do",
	}

	wsAuthSandbox = map[site]string{
		ARGENTINA:  "https://auth.mercadolibre.com.ar",
		BOLIVIA:    "https://auth.mercadolibre.com.bo",
		BRASIL:     "https://auth.mercadolivre.com.br",
		COLOMBIA:   "https://auth.mercadolibre.com.co",
		COSTA_RICA: "https://auth.mercadolibre.com.cr",
		EQUADOR:    "https://auth.mercadolibre.com.ec",
		HONDURAS:   "https://auth.mercadolibre.com.hn",
		GUATEMALA:  "https://auth.mercadolibre.com.gt",
		CHILE:      "https://auth.mercadolibre.cl",
		MEXICO:     "https://auth.mercadolibre.com.mx",
		NICARAGUA:  "https://auth.mercadolibre.com.ni",
		PARAGUAI:   "https://auth.mercadolibre.com.py",
		SALVADOR:   "https://auth.mercadolibre.com.sv",
		URUGUAI:    "https://auth.mercadolibre.com.uy",
		VENEZUELA:  "https://auth.mercadolibre.com.ve",
		PANAMA:     "https://auth.mercadolibre.com.pa",
		PERU:       "https://auth.mercadolibre.com.pe",
		PORTUGAL:   "https://auth.mercadolibre.com.pt",
		DOMINICANA: "https://auth.mercadolibre.com.do",
	}
)
