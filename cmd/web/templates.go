package main

import "github.com/koller-m/OnlyDupes/internal/models"

type templateData struct {
	Dupe  *models.Dupe
	Dupes []*models.Dupe
}
