package model

type PapagoReq struct {
	SourceLanguage string `form:"source"`
	TargetLanguage string`form:"target"`
	Text string`form:"text"`
}

type PapagoRep struct {
	Message Message `json:"message"`
}
type Message struct{
	AtType string `json:"@result"`
	AtService string `json:"@service"`
	AtVersion string `json:"@version"`
	Result Result `json:"result"`
}
type Result struct {
	SrcLangType string`json:"srcLangType"`
	TarLangType string`json:"tarLangType"`
	TranslatedText string`json:"translatedText"`
	EngineType string`json:"engineType"`
	Pivot string`json:"pivot"`
}