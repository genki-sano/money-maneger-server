package domains_test

import (
	"strconv"
	"testing"

	"github.com/genki-sano/money-maneger-server/package/domains"
	"github.com/stretchr/testify/assert"
)

func TestPayment_NewPayment(t *testing.T) {
	t.Run("指定した形式で支払情報が作成されるべき", func(t *testing.T) {
		id := "hoge"
		name := "太郎"
		date := "2020/11/01"
		price := "1234"
		category := "食費"
		memo := "スーパー"

		actual := domains.NewPayment(id, name, date, price, category, memo)

		expected, _ := strconv.Atoi(price)
		assert.Exactly(t, actual.ID, id)
		assert.Exactly(t, actual.Name, name)
		assert.Exactly(t, actual.Date, date)
		assert.Exactly(t, actual.Price, uint32(expected))
		assert.Exactly(t, actual.Category, category)
		assert.Exactly(t, actual.Memo, memo)
	})
}
