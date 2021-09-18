package repository

import "github.com/dena-autumn-hackathon-2021-team-d/dena-autumn-backend/domain/entity"

type Comment interface {
	Post(comment *entity.Comment) error
}
