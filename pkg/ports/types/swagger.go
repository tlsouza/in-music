package types

type SwaggerConfig struct {
	Schemes  []string       `json:"schemes,omitempty"`
	Swagger  string         `json:"swagger,omitempty"`
	Info     SwaggerInfo    `json:"info,omitempty"`
	Host     string         `json:"host,omitempty"`
	BasePath string         `json:"basePath,omitempty"`
	Paths    map[string]any `json:"paths,omitempty"`
}

type SwaggerInfo struct {
	Version     string `json:"version,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type SwaggerSchema struct {
	Type       any             `json:"type,omitempty"`
	Required   *[]string       `json:"required,omitempty"`
	Item       any             `json:"item,omitempty"`
	Properties *map[string]any `json:"properties,omitempty"`
	Format     string          `json:"format,omitempty"`
}

type DocResponse struct {
	Status      int    `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	Body        any    `json:"schema,omitempty"`
}

type SwaggerResponse struct {
	Description string         `json:"description,omitempty"`
	Schema      *SwaggerSchema `json:"schema,omitempty"`
}

type SwaggerParameters struct {
	In          string         `json:"in,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Example     string         `json:"example,omitempty"`
	Type        any            `json:"type,omitempty"`
	Item        any            `json:"item,omitempty"`
	Format      string         `json:"format,omitempty"`
	Schema      *SwaggerSchema `json:"schema,omitempty"`
}

type SwaggerRouteObject struct {
	Description string            `json:"description,omitempty"`
	Request     any               `json:"request,omitempty"`
	Responses   []DocResponse     `json:"responses,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	Path        map[string]string `json:"path,omitempty"`
	Query       map[string]string `json:"query,omitempty"`
}

type SwaggerDescription struct {
	Description string                      `json:"description,omitempty"`
	Responses   *map[string]SwaggerResponse `json:"responses,omitempty"`
	Parameters  *[]SwaggerParameters        `json:"parameters,omitempty"`
}
