// Copyright 2026 Anna Veretennykova
//
// SPDX-License-Identifier: MIT

package engine

import "fmt"

type EngineStage string

const (
	StageInput      EngineStage = "input"
	StageParser     EngineStage = "parser"
	StageValidation EngineStage = "validation"
	StageLayout     EngineStage = "layout"
	StagePagination EngineStage = "pagination"
	StageRender     EngineStage = "render"
)

type Error struct {
	Code    string      `json:"code"`
	Stage   EngineStage `json:"stage"`
	Message string      `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%s/%s] %s", e.Stage, e.Code, e.Message)
}

func NewError(code string, stage EngineStage, message string) *Error {
	return &Error{Code: code, Stage: stage, Message: message}
}
