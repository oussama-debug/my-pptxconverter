package models

type PPTXInformation struct {
	app         string `xml:"Application"`
	app_version string `xml:"AppVersion"`
	app_company string `xml:"Company"`
	words_count int    `xml:"WordsCount"`
	format      string `xml:"PresentationFormat"`
}

func NewPPTXInformation() *PPTXInformation {
	return &PPTXInformation{}
}

func (p *PPTXInformation) GetAppVersion() string {
	return p.app_version
}

func (p *PPTXInformation) GetCompany() string {
	return p.app_company
}

func (p *PPTXInformation) GetWordsCount() int {
	return p.words_count
}

func (p *PPTXInformation) GetFormat() string {
	return p.format
}
