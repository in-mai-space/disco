package types

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	errors "github.com/in-mai-space/disco/errors"
)

type BaseController[T any] struct{}

// parsing useing data transfer object
func (b *BaseController[T]) ParseDTO(c *fiber.Ctx, dto interface{}, dtoType string) error {
	if err := c.BodyParser(dto); err != nil {
		return errors.ParseFailedError(dtoType)
	}
	return nil
}

func (b *BaseController[T]) ValidateID(c *fiber.Ctx, modelType string) (uuid.UUID, error) {
	id := c.Params("id")
	idAsUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errors.ValidateIdFailedError(modelType)
	}
	return idAsUUID, nil
}

func (b *BaseController[T]) HandleError(c *fiber.Ctx, err error) error {
	return errors.HandleError(c, err)
}
