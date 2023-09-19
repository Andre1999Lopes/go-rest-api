package models

import (
	"time"

	"gopkg.in/validator.v2"
)

type Aluno struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Nome      string `json:"nome" validate:"nonzero"`
	Cpf       string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	Rg        string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
}

func ValidaDadosDeAlunos(aluno *Aluno) error {
	if err := validator.Validate((aluno)); err != nil {
		return err
	}
	return nil
}
