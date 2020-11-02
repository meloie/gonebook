package services

import "github.com/meloie/gonebook/internal/ent"

type Service struct {
	Database *ent.Client
}
