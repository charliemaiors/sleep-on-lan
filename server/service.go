package server

import (
	"github.com/kardianos/service"
)

var logger service.Logger

type sol struct {
}

func (s *sol) Start(srv service.Service) error {
	go StartServer()
	return nil
}

func (s *sol) Stop(srv service.Service) error {
	return nil
}

//InstallService install sleep on lan as a service in current os
func InstallService(port string) {
	srvConf := &service.Config{
		Name:        "sleeponlan",
		DisplayName: "Sleep On Lan Service",
		Description: "Simple sleep on lan service",
		Arguments:   []string{"--port", port},
	}

	sleeponlan := &sol{}

	s, err := service.New(sleeponlan, srvConf)

	if err != nil {
		panic(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		panic(err)
	}

	err = s.Install()
	if err != nil {
		logger.Error(err)
	}
}
