package model

import "github.com/jinzhu/gorm"

type Customer struct {
	CustomerId int `form:"customer_id" json:"customer_id"` // 自增主键

	Lat float64 `form:"lat" json:"lat"` // 纬度

	Lon float64 `form:"lon" json:"lon"` // 经度

	Name string `form:"name" json:"name"` // 名称
}

//根据主键得到单一实体
func GetCustomerByPrimaryKey(customerId int) (*Customer, error) {
	var customer Customer
	err := db.Where("customer_id = ?", customerId).First(&customer).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &customer, nil
}

//根据主键通过条件编辑实体
func EditCustomerByPrimaryKey(customerId int, maps interface{}) error {
	if err := db.Model(&Customer{}).Where("customer_id = ?", customerId).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

//根据主键删除实体
func DeleteCustomerByPrimaryKey(customerId int) error {
	if err := db.Where("customer_id = ?", customerId).Delete(&Customer{}).Error; err != nil {
		return err
	}
	return nil
}

//插入实体
func AddCustomer(customer *Customer) error {
	if err := db.Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

//根据条件获得分页实体集合
func FindCustomers(pageNum int, pageSize int, maps interface{}) ([]*Customer, error) {
	var (
		customers []*Customer
		err       error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&customers).Error
	} else {
		err = db.Where(maps).Find(&customers).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return customers, nil
}
