package repository

import "find-a-friend/fiend-a-friend-api/cmd/http/adpter/storage/postgres"

type Organization struct {
	db *postgres.Db
}

func NewOrganizationRepository(db *postgres.Db) *Organization {
	return &Organization{
		db,
	}
}
