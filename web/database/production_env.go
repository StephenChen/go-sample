package database

import "context"

// 互联网公司核心线上业务都会在代码中把SQL放在显眼的位置提供给DBA评审
const (
	getAllByProductIDAndCustomerID = `select * from p_orders where product_id in (:product_id) and customer_id=:customer_id`
)

func GetAllByProductIDAndCustomerID(ctx context.Context, productIDs []uint64, customerID uint64) ([]Order, error) {
	var orderList []Order

	params := map[string]interface{}{
		"product_id":  productIDs,
		"customer_id": customerID,
	}

	// getAllByProductIDAndCustomerID 是 const 类型的 sql 字符串
	sql, args, err := sqlutil_Named(getAllByProductIDAndCustomerID, params)
	if err != nil {
		return nil, err
	}

	var sqldbInstance interface{}
	err = dao_QueryList(ctx, sqldbInstance, sql, args, &orderList)
	if err != nil {
		return nil, err
	}

	return orderList, err
}

type Order struct{}

func sqlutil_Named(sql string, params map[string]interface{}) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func dao_QueryList(ctx context.Context, instance interface{}, sql interface{}, args interface{}, i *[]Order) error {
	return nil
}
