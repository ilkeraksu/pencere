package pencere

func NewProperties() Properties {
	return make(map[string]interface{})
}

type Properties map[string]interface{}

func (this Properties) GetValue(name string) interface{} {
	return this[name]
}

func (this Properties) GetString(key string, def string) string {
	v, ok := this[key]
	if !ok {
		return def
	}
	r, ok := v.(string)
	if !ok {
		return def
	}
	return r
}

func (this Properties) GetInt(key string, def int) int {
	v, ok := this[key]
	if !ok {
		return def
	}
	r, ok := v.(int)
	if !ok {
		return def
	}
	return r
}

func (this Properties) SetValue(name string, value interface{}) {
	this[name] = value
}

func (this Properties) SetInt(name string, value int) {
	this[name] = value
}

func (this Properties) SetString(name string, value string) {

	this[name] = value
}

func (this Properties) GetColor(key string, def Color) Color {
	v, ok := this[key]
	if !ok {
		return def
	}
	r, ok := v.(Color)
	if !ok {
		return def
	}
	return r
}

func (this Properties) SetColor(name string, value Color) {
	this[name] = value
}
