package GetPay

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/GetPay"
	GetPayReq "github.com/flipped-aurora/gin-vue-admin/server/model/GetPay/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Key_use"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"golang.org/x/exp/rand"
)

type Get_payService struct{}

//

func generateOrderNumber() string {
	timestamp := time.Now().Unix()
	randNum := 10000 + rand.Intn(10000)
	orderNumber := fmt.Sprintf("%d%d", timestamp, randNum)
	return orderNumber
}

// CreateGet_pay 创建订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (PayService *Get_payService) CreateGet_pay(Pay *GetPay.Get_pay) (err error) {
	err = global.GVA_DB.Create(Pay).Error
	return err
}

// DeleteGet_pay 删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (PayService *Get_payService) DeleteGet_pay(ID string) (err error) {
	err = global.GVA_DB.Delete(&GetPay.Get_pay{}, "id = ?", ID).Error
	return err
}

// DeleteGet_payByIds 批量删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (PayService *Get_payService) DeleteGet_payByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]GetPay.Get_pay{}, "id in ?", IDs).Error
	return err
}

// UpdateGet_pay 更新订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (PayService *Get_payService) UpdateGet_pay(Pay GetPay.Get_pay) (err error) {
	err = global.GVA_DB.Model(&GetPay.Get_pay{}).Where("id = ?", Pay.ID).Updates(&Pay).Error
	return err
}

// GetGet_pay 根据ID获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (PayService *Get_payService) GetGet_pay(ID string) (Pay GetPay.Get_pay, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Pay).Error
	return
}

// GetGet_payInfoList 分页获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (PayService *Get_payService) GetGet_payInfoList(info GetPayReq.Get_paySearch) (list []GetPay.Get_pay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&GetPay.Get_pay{})
	var Pays []GetPay.Get_pay
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Out_trade_no != "" {
		db = db.Where("out_trade_no = ?", info.Out_trade_no)
	}
	if info.Key != "" {
		db = db.Where("key = ?", info.Key)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Pays).Error
	return Pays, total, err
}

// GetOrder 请实现方法
// Author [yourname](https://github.com/yourname)
func (PayService *Get_payService) GetOrder(u_add string) string {
	// Implement your business logic here

	Order := generateOrderNumber()
	newStatus := GetPay.Get_pay{
		Status:       "WAIT_BUYER_PAY",
		Out_trade_no: Order,
		U_add:        u_add,
	}
	db := global.GVA_DB.Model(&GetPay.Get_pay{})
	db.Create(&newStatus)
	return Order
}

// AlilPay 请实现方法
// Author [yourname](https://github.com/yourname)
func (PayService *Get_payService) AlilPay(order_on string, price float64) string {
	// Implement your business logic here
	bm := make(gopay.BodyMap)

	bm.Set("subject", "条码支付").
		Set("scene", "bar_code1").
		Set("auth_code", "1121216117131111").
		Set("out_trade_no", order_on).
		Set("total_amount", price).
		Set("timeout_express", "2m")

	aliRsp, err := PAys.TradePagePay(context.Background(), bm)
	if err != nil {

		fmt.Println(err)

		return err.Error()
	}

	return aliRsp
}

// Notify 请实现方法
// Author [yourname](https://github.com/yourname)
func (PayService *Get_payService) Notify(c *gin.Context) (err error) {
	// 请在这里实现自己的业务逻辑

	dbz := global.GVA_DB
	// dbs := global.GVA_DB.Model(&Key_use.Key_uses{})
	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request)

	if err != nil {

		fmt.Println("err", err.Error())
		return
	}

	fmt.Println("-----AlipayNotify notifyReq: ", notifyReq)

	maps_out_trade_no := notifyReq["out_trade_no"]

	out_trade_Status := notifyReq["trade_status"]

	myString := fmt.Sprintf("%v", out_trade_Status)

	newStatus := GetPay.Get_pay{}

	newUse := Key_use.Key_uses{}

	dbz.Where("out_trade_no = ?", maps_out_trade_no).First(&newStatus) //

	newStatus.Status = myString

	dbz.Where("use = ?", false).Order("id desc").Find(&newUse)

	newStatus.Key = newUse.Keys
	ok, err := alipay.VerifySignWithCert("./alipayPublicCert.crt", notifyReq)
	boolValue := false

	if newUse.Use != nil {
		boolValue = *newUse.Use
	}

	if newUse.Use != nil {

		fmt.Println()
	}

	if boolValue {
		fmt.Println("已经使用")
		return

	} else {
		fmt.Println("未使用")

		boolValue = true
		newUse.Use = &boolValue
	}

	if ok {
		fmt.Println("支付成功", maps_out_trade_no)
		boolValue = true
		dbz.Save(&newStatus)
		dbz.Save(&newUse)

	} else {
		fmt.Println("支付失败:", err)

	}

	return dbz.Error
}

// Stock 请实现方法
// Author [yourname](https://github.com/yourname)
func (PayService *Get_payService) Stock(out_trade_no string) (key string) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&GetPay.Get_pay{})

	newkeys := GetPay.Get_pay{}

	db.Where("out_trade_no = ?", out_trade_no).First(&newkeys)
	if newkeys.ID == 0 {
		fmt.Println("订单不存在与服务器数据库")

		return
	}

	key = newkeys.Key

	return key
}

// Query_order 请实现方法
// Author [yourname](https://github.com/yourname)
func (PayService *Get_payService) Query_order(ctx *gin.Context, U_addr string) (key []string, err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&GetPay.Get_pay{})
	newkeys := []GetPay.Get_pay{}

	db.Where("u_add = ?", U_addr).Limit(10).Find(&newkeys)

	// 在这里处理每批数据并将其赋值给结构体
	for _, user := range newkeys {
		// 假设您有一个自定义结构体 MyUser，您可以将数据赋值给它
		myUser := GetPay.Get_pay{

			Key: user.Key,
		}
		// 处理 myUser 数据，例如存储到其他地方或进行其他操作
		fmt.Println(myUser.Key)

		key = append(key, user.Key)
	}

	return key, db.Error
}
