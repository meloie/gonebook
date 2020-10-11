package services

import "gonebook/internal/ent"

type Service struct {
	Database *ent.Client
}
