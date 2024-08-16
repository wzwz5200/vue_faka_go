package GetPay

import (
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
)

const (
	appID      = "9021000139677594"
	privateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC1wEr4E2M9nYJsRqjNUmZpFf9/uI0qwkuE8WlR84dd2NOHw0MKaE5DTPoZa66g4oEpsEQTvZTiT7q3vatzjVJfL2xtarpTOMys+O8xrc4QT3akZEoXPwKeLZGsujilFGAP1sowPWuawRX+uLP5t9tgJ9MA9fU+McQyyd/I02oH2jHh6/HagsAbrdN9T4V/mDt4mDqq6A7OEkjzs7Ps5qayHHX5EG8E6j5UlgBNfDQzD7bY+cWxTX6lF1o6jdIUFNiGaXHd7Js/0HVvWGevHHn27/NDh7EuYMoIYmnBH/prQvkOti6WqGNKpLa71zVLlOQ5bncyIY3d5MNKwiStqHm3AgMBAAECggEARINtf2DHm1WB8dEdFvFF+704KGxogsLldwMOIb98uureqqH35rd1MTeWW/jUxn88E3wH3a0sr7dAer08IX4XnBwjcoe2H2Mc/OOGP6L2N/3MiqHnsCWnlfs1m/hDHC+3GqFsbIzqwFajNPdpmOwu/WtlxknxnfVisYaaHnlb0gNS7kPH0GgITp5vbV/eolUamwlqI1Fejxp9j9QB+ylowIlGi6QbcrxfmiOgVxNK+n/eEToxycdEiwluCM6GhjSZ/4OalRN5/osCRHq+uFtzcX6bRM+CMH7EbWUgeXyf0fbo/QuQ9dxQ2gBCS+AfQgdcEfcacoxNsXPuNl44ff78cQKBgQDexy0d6apna4/dhSHOVVU5F2y2c9k4U12pAdEtMGAvqPHoySEy2EHg9yUOhkWmvmnjKxGmmtpQiQ/xiXKt64GidKqFmbuEA+dffTJRDul72BMpX9K+/ll+utRtBMNk8z2SynyNPNTHJoAUgckgKOHLYL9UALj+DzaUz0Lk+kTBfwKBgQDQ2t1v7ruYFoULU/w5pQ6zgB80Syp3BGN7NRFlmJsGTiB6gULg38LWOasSZC9Eb1K51bdDXcXbfOf4FSCT+GD5CMlbUnsQ+Zot0MWSdDL+pkv/v44eXKjFMVUZhrRr85LDGkM0eFSM+lctg9x0zGq51NPbasl09hJA50ftFH/zyQKBgQCpTrgkJqSB6sYwdXCGzmVxeTL+yraITxs8SYw+iqhfMEeBQfJAIQiYP1vjWPpwSgBRTDci5kcKs8/xIiOMuBISdBXwaTCQmSLreEuYPPwHSeTuKcwRqV48qSKuI9OX0iC2gbr84AFZxRHMBALltQw9M67U2aO+ObOijo3pVaaRYwKBgCD/4+sSqnteW4ktrPWiuc7s2IFXuw7xF5LZELWfxibZ41HX6KxFTrCwjfIq+Dj06fpI5Vr0jxWeB7zwAaS/ovrWQ5J1VtKYzZ6dlQoN6BzKQ8nWB2uOsm/t2odc/FbuNmszVBkPRjS8PVgItKWTwu03zn2lws0DMGEm6ftwuMLZAoGAZCIAJvE9RyphwI1hFmFbTGrho69RZHulVcDMWC7sOYzVJCVgRwZ+7qKFBCzkphieui+FOUPr1wQLpx48r97XYl632BI0+6cTbvIRxUY8hN4oONlzyuC04jDsTif+pE2D0L9+BdN1/oSP3RRKYQiRfHGBzVy1Q0lQAHvXrNbEw6c="
)

var PAys *alipay.Client

func Getpay() *alipay.Client {
	client, err := alipay.NewClient("9021000139677594", privateKey, false)
	if err != nil {
		fmt.Println(err)

	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).                                                                       // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).                                                                      // 设置签名类型，不设置默认 RSA2
							SetReturnUrl("http://localhost:5173/stock/").                                                  // 设置返回URL 前端订单页面
							SetNotifyUrl("https://2ff1-2409-8a30-ae2e-fde0-f952-19e7-69ed-8978.ngrok-free.app/Pay/notify") // 设置异步通知URL

	client.AutoVerifySign([]byte("./alipayPublicCert.crt bytes"))
	client.SetCertSnByPath("./appPublicCert.crt", "./alipayRootCert.crt", "./alipayPublicCert.crt")

	PAys = client
	return client
}

func Getpayp() *alipay.Client {

	return PAys
}
