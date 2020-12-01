package repositories_test

import (
	"errors"
	"testing"
	"time"

	"github.com/genki-sano/money-maneger-server/package/domains"
	"github.com/genki-sano/money-maneger-server/package/interfaces/repositories"
	"github.com/genki-sano/money-maneger-server/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPaymentDataAccess_FindAll(t *testing.T) {
	testCases := []struct {
		name     string
		data     domains.Payments
		expected domains.Payments
	}{
		{
			name: "取得する値がある場合、保存されたすべての値を返すべき",
			data: domains.Payments{
				domains.Payment{
					ID:       "hoge",
					Name:     "太郎",
					Date:     "2020/11/01",
					Price:    "1234",
					Category: "食費",
					Memo:     "スーパー",
				},
				domains.Payment{
					ID:       "fuga",
					Name:     "花子",
					Date:     "2020/10/01",
					Price:    "2345",
					Category: "日用品",
					Memo:     "ドラッグストア",
				},
			},
			expected: domains.Payments{
				domains.Payment{
					ID:       "hoge",
					Name:     "太郎",
					Date:     "2020/11/01",
					Price:    "1234",
					Category: "食費",
					Memo:     "スーパー",
				},
				domains.Payment{
					ID:       "fuga",
					Name:     "花子",
					Date:     "2020/10/01",
					Price:    "2345",
					Category: "日用品",
					Memo:     "ドラッグストア",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockPaymentDatastore(ctrl)
			m.EXPECT().GetAll().Return(testCase.data, nil)

			reps := repositories.NewPaymentRepository(m)
			actual, err := reps.FindAll()

			assert.Len(t, actual, len(testCase.expected))
			for i, item := range actual {
				assert.Exactly(t, item.ID, testCase.expected[i].ID)
				assert.Exactly(t, item.Name, testCase.expected[i].Name)
				assert.Exactly(t, item.Date, testCase.expected[i].Date)
				assert.Exactly(t, item.Price, testCase.expected[i].Price)
				assert.Exactly(t, item.Category, testCase.expected[i].Category)
				assert.Exactly(t, item.Memo, testCase.expected[i].Memo)
			}

			assert.NoError(t, err)
		})
	}
}
func TestPaymentDataAccess_GetByDate(t *testing.T) {
	testCases := []struct {
		name     string
		data     domains.Payments
		expected domains.Payments
	}{
		{
			name: "取得する値がある場合、指定した月の値のみを返すべき",
			data: domains.Payments{
				domains.Payment{
					ID:       "hoge",
					Name:     "太郎",
					Date:     "2020/11/01",
					Price:    "1234",
					Category: "食費",
					Memo:     "スーパー",
				},
				domains.Payment{
					ID:       "fuga",
					Name:     "花子",
					Date:     "2020/10/01",
					Price:    "2345",
					Category: "日用品",
					Memo:     "ドラッグストア",
				},
			},
			expected: domains.Payments{
				domains.Payment{
					ID:       "hoge",
					Name:     "太郎",
					Date:     "2020/11/01",
					Price:    "1234",
					Category: "食費",
					Memo:     "スーパー",
				},
			},
		},
		{
			name: "取得する値がない場合、空の値を返すべき",
			data: domains.Payments{
				domains.Payment{
					ID:       "fuga",
					Name:     "花子",
					Date:     "2020/10/01",
					Price:    "2345",
					Category: "日用品",
					Memo:     "ドラッグストア",
				},
			},
			expected: make(domains.Payments, 0),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockPaymentDatastore(ctrl)
			m.EXPECT().GetAll().Return(testCase.data, nil)

			reps := repositories.NewPaymentRepository(m)
			actual, err := reps.GetByDate(time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local))

			assert.Len(t, actual, len(testCase.expected))
			for i, item := range actual {
				assert.Exactly(t, item.ID, testCase.expected[i].ID)
				assert.Exactly(t, item.Name, testCase.expected[i].Name)
				assert.Exactly(t, item.Date, testCase.expected[i].Date)
				assert.Exactly(t, item.Price, testCase.expected[i].Price)
				assert.Exactly(t, item.Category, testCase.expected[i].Category)
				assert.Exactly(t, item.Memo, testCase.expected[i].Memo)
			}

			assert.NoError(t, err)
		})
	}
	t.Run("接続エラーがある場合、エラーを返すべき", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msg := "hoge"

		m := mock.NewMockPaymentDatastore(ctrl)
		m.EXPECT().GetAll().Return(nil, errors.New(msg))

		reps := repositories.NewPaymentRepository(m)
		actual, err := reps.GetByDate(time.Now())

		assert.Empty(t, actual)

		assert.Error(t, err)
		assert.EqualError(t, err, msg)
	})
}
