package product

type Product interface {
	GetProductID() string
	SetAmmount(int64)
	GetAmmount() int64
}

// WARNING: This function mutates what is passing
// TODO: solve this :P
func MergeLists[T Product](p1, p2 []T) []T {
	merged := append([]T{}, p1...)

	for _, newProd := range p2 {
		found := false
		for i, existing := range merged {
			if newProd.GetProductID() == existing.GetProductID() {
				merged[i].SetAmmount(merged[i].GetAmmount() + newProd.GetAmmount())
				found = true
				break
			}
		}
		if !found {
			merged = append(merged, newProd)
		}
	}

	return merged
}
