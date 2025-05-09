package service

var CityWaySvc *CityWayService

type CityWayService struct {
	*Service
}

func newCityWayService(svc *Service) {
	CityWaySvc = &CityWayService{
		svc,
	}
}
