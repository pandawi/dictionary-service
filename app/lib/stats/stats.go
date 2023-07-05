package stats

type StatsImpl struct {
	totalRequests    int
	totalRequestTime int
}

type Stats interface {
	CountWords(wordList []string) int
	CalculateAvgReqProcessingTime() int
	AddRequest(reqTime int)
	GetTotalRequests() int
}

func NewStats() Stats {
	return &StatsImpl{totalRequests: 0, totalRequestTime: 0}
}

func (s *StatsImpl) CountWords(wordList []string) int {
	length := len(wordList)
	return length
}

func (s *StatsImpl) CalculateAvgReqProcessingTime() int {
	if s.totalRequests == 0 {
		return 0
	}
	return s.totalRequestTime / s.totalRequests
}

func (s *StatsImpl) AddRequest(reqTime int) {
	s.totalRequests++
	s.totalRequestTime += reqTime
}

func (s *StatsImpl) GetTotalRequests() int {
	return s.totalRequests
}
