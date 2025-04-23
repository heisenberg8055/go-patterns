package mediator

type StationMaster struct {
	IsPlatformFree bool
	TrainQueue     []Train
}

func NewStationMaster() *StationMaster {
	return &StationMaster{
		IsPlatformFree: true,
	}
}

func (s *StationMaster) canArrive(t Train) bool {
	if s.IsPlatformFree {
		s.IsPlatformFree = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, t)
	return false
}

func (s *StationMaster) notifyAboutDeparture() {
	if !s.IsPlatformFree {
		s.IsPlatformFree = true
	}
	if len(s.TrainQueue) > 0 {
		firstTrainQueue := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		firstTrainQueue.PermitArrival()
	}
}
