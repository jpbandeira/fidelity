package repository

import (
	"context"

	"github.com/jp/fidelity/internal/domain"
)

func CreateClient(ctx context.Context, client domain.Client) (domain.Client, error) {
	return domain.Client{}, nil
}
