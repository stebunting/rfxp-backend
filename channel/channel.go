package channel

type Channel struct {
	Number    int  `json:"number"`
	FreqStart int  `json:"freqStart"`
	FreqEnd   int  `json:"freqEnd"`
	Indoors   bool `json:"indoors"`
	Outdoors  bool `json:"outdoors"`
}
