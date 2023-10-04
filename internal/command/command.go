package command

import (
	"github.com/NooFreeNames/Cripto/internal/service"
)

type CommandManager struct {
	service *service.Service
}

func NewCommandManger(service *service.Service) *CommandManager {
	return &CommandManager{service: service}
}
