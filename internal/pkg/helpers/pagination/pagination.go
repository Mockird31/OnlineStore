package pagination

import (
	"fmt"
	"net/http"

	"github.com/Mockird31/OnlineStore/config"
	customErrors "github.com/Mockird31/OnlineStore/internal/pkg/helpers/customErrors"
	query "github.com/Mockird31/OnlineStore/internal/pkg/helpers/query"
	model "github.com/Mockird31/OnlineStore/internal/pkg/model"
)

func validatePagination(p *model.Pagination, cfg *config.PaginationConfig) error {
	if p.Offset > cfg.MaxOffset {
		p.Offset = cfg.MaxOffset
	}

	if p.Offset < 0 {
		p.Offset = 0
		return customErrors.ErrInvalidOffset
	}

	if p.Limit > cfg.MaxLimit {
		p.Limit = cfg.MaxLimit
	}

	if p.Limit < 0 {
		p.Limit = 0
		return customErrors.ErrInvalidLimit
	}

	return nil

}

func GetPagination(r *http.Request, cfg *config.PaginationConfig) (*model.Pagination, error) {
	pagination := &model.Pagination{}

	offset, err := query.ReadInt(r.URL.Query(), "offset", cfg.DefaultOffset)
	if err != nil {
		return nil, customErrors.ErrInvalidOffset
	}

	limit, err := query.ReadInt(r.URL.Query(), "limit", cfg.DefaultLimit)
	if err != nil {
		return nil, customErrors.ErrInvalidLimit
	}

	pagination.Offset = offset
	pagination.Limit = limit
	fmt.Println(cfg)
	// err = validatePagination(pagination, cfg)
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println("PAGINATION", pagination.Offset, pagination.Limit)
	return pagination, nil
}
