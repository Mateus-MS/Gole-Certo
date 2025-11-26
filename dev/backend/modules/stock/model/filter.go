package product

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type StockFilter struct {
	Brand      []string
	PriceRange []string
}

func (sf *StockFilter) Build() bson.M {
	result := bson.M{}

	if len(sf.Brand) != 0 {
		result["brand"] = bson.M{"$in": sf.Brand}
	}

	if len(sf.PriceRange) != 0 {
		minFloat, _ := strconv.ParseFloat(sf.PriceRange[0], 64)
		maxFloat, _ := strconv.ParseFloat(sf.PriceRange[1], 64)
		result["price"] = bson.M{
			"$gte": minFloat,
			"$lte": maxFloat,
		}
	}

	return result
}

func (sf *StockFilter) ReadBrands(r *http.Request) (err error) {
	var raw string

	// Query the brand parameter from the URL
	if raw, err = utils.GetQueryParam(r, "brands", true, ""); err != nil {
		// If the error is that the parameter NOT IS in the URL, return nil
		if errors.Is(err, utils.ParameterNotInURL) {
			return nil
		}

		return err
	}

	// Convert it into an array of strings
	sf.Brand = strings.Split(raw, "-")
	return nil
}

func (sf *StockFilter) ReadPriceRange(r *http.Request) (err error) {
	var raw string

	// Query the priceRange parameter from the URL
	if raw, err = utils.GetQueryParam(r, "price-range", true, ""); err != nil {
		// If the error is that the parameter NOT IS in the URL, return nil
		if errors.Is(err, utils.ParameterNotInURL) {
			return nil
		}
		return err
	}

	// Split into min and max values
	priceMinMax := strings.Split(raw, "-")

	// It should have 2 of length, if not, it's a badRequest
	if len(priceMinMax) != 2 {
		return errors.New("Price range bad format: " + raw)
	}

	sf.PriceRange = priceMinMax
	return nil
}
