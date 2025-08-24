package stock_repository

import (
	"context"
	"errors"
	"math"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Seems really of to define this here
const ItemsPerPage int64 = 12

func (repo *Repository) HasProduct(ctx context.Context, id string) bool {
	err := repo.collection.FindOne(ctx, id).Err()

	return err == nil || !errors.Is(err, mongo.ErrNoDocuments)
}

func (repo *Repository) TotalItems(ctx context.Context, filter bson.M) (int64, error) {
	total, err := repo.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (repo *Repository) TotalPages(ctx context.Context, filter bson.M) (int64, error) {
	total, err := repo.TotalItems(ctx, filter)
	if err != nil {
		return 0, err
	}

	// Calculate the quantity of pages
	return int64(math.Ceil(float64(total) / float64(ItemsPerPage))), nil
}

func (repo *Repository) GetBrands(ctx context.Context, filter bson.M) ([]string, error) {
	dbFilter := bson.M{}

	// Exclude selected brands if present
	if brandFilter, ok := filter["brand"].(bson.M); ok {
		if inBrands, ok := brandFilter["$in"].([]string); ok && len(inBrands) > 0 {
			dbFilter["brand"] = bson.M{"$nin": inBrands}
		}
	}

	// Apply search filter if present
	if searchVal, ok := filter["search"].(string); ok && searchVal != "" {
		searchRegex := bson.M{"$regex": searchVal, "$options": "i"} // case-insensitive
		if existing, exists := dbFilter["brand"].(bson.M); exists {
			existing["$regex"] = searchVal
			existing["$options"] = "i"
			dbFilter["brand"] = existing
		} else {
			dbFilter["brand"] = searchRegex
		}
	}

	// Get distinct brands
	brands, err := repo.collection.Distinct(ctx, "brand", dbFilter)
	if err != nil {
		return []string{}, err
	}

	// Convert to string and apply optional limit
	var brandsStrings []string
	limit := 0
	if l, ok := filter["limit"].(int); ok {
		limit = l
	}

	for _, brand := range brands {
		str, ok := brand.(string)
		if !ok {
			continue
		}

		// If search is present, filter again in Go (optional safety)
		if searchVal, ok := filter["search"].(string); ok && searchVal != "" {
			if !strings.Contains(strings.ToLower(str), strings.ToLower(searchVal)) {
				continue
			}
		}

		brandsStrings = append(brandsStrings, str)
		if limit > 0 && len(brandsStrings) >= limit {
			break
		}
	}

	return brandsStrings, nil
}
