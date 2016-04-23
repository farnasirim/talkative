package web

type CountResponse struct {
	Hostname              string `yaml:"hostname"`
	RequestedPath         string `yaml:"requestedPath"`
	TotalNumberOfRequests int    `yaml:totalNumberOfRequests"`
	Count                 int    `yaml:"count"`
	CurrentTimeUnixNano   int64  `yaml:"currentTimeUnixNano"`
	UptimeUnixNano        int64  `yaml:"uptimeUnixTime"`
	UptimeHumanReadable   string `yaml:"uptimeHumanReadable"`
}

type CountResetResponse struct {
	Hostname              string `yaml:"hostname"`
	RequestedPath         string `yaml:"requestedPath"`
	TotalNumberOfRequests int    `yaml:totalNumberOfRequests"`
	Count                 int    `yaml:"count"`
	SetTo                 int    `yaml:"setTo"`
	CurrentTimeUnixNano   int64  `yaml:"currentTimeUnixNano"`
	UptimeUnixNano        int64  `yaml:"uptimeUnixTime"`
	UptimeHumanReadable   string `yaml:"uptimeHumanReadable"`
}

type InfoResponse struct {
	Hostname              string `yaml:"hostname"`
	RequestedPath         string `yaml:"requestPath"`
	RemoteAddress         string `yaml:"remoteAddress"`
	TotalNumberOfRequests int    `yaml:totalNumberOfRequests"`
	CurrentTimeUnixNano   int64  `yaml:"currentTimeUnixNano"`
	UptimeUnixNano        int64  `yaml:"uptimeUnixTime"`
	UptimeHumanReadable   string `yaml:"uptimeHumanReadable"`
}
