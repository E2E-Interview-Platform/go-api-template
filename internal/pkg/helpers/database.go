package helpers

import (
	"math"

	sq "github.com/Masterminds/squirrel"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
)

func GetPaginationMetaData(page int, limit int, totalRecords int) dto.Pagination {
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	pagination := dto.Pagination{
		RecordPerPage: limit,
		CurrentPage:   page,
		TotalPage:     totalPages,
		TotalRecords:  totalRecords,
	}

	if page > 1 && totalRecords > 0 {
		pre := min(page-1, totalPages)
		pagination.Previous = &pre
	}

	if page < totalRecords {
		next := page + 1
		pagination.Next = &next
	}

	return pagination
}

func SetQueryBuilderOffsetAndLimit(queryBuilder sq.SelectBuilder, page int, limit int) sq.SelectBuilder {
	offset := (page - 1) * limit
	return queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))
}
