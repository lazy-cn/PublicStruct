package wechat

import (
	"encoding/xml"
)

//微信支付统一下单请求
type OrderRequestSl struct {
	XMLName     		xml.Name	`xml:"xml"`
	Appid				string		`json:"appid" xml:"appid"`									//公众号appid
	Mch_id				string		`json:"mch_id" xml:"mch_id"` 								//服务商商户号
	Sub_appid			string		`json:"sub_appid,omitempty" xml:"sub_appid,omitempty"`		//子商户公众号
	Sub_mch_id			string		`json:"sub_mch_id" xml:"sub_mch_id"`						//子商户号
	Device_info			string		`json:"device_info,omitempty" xml:"device_info,omitempty"`	//设备号，终端设备号，pc网页或公众号内支付请传“WEB”
	Nonce_str			string		`json:"nonce_str" xml:"nonce_str"`							//随机字符串
	Sign				string		`json:"sign" xml:"sign"`									//签名
	Sign_type			string		`json:"sign_type,omitempty" xml:"sign_type,omitempty"`		//签名类型，目前支持HMAC-SHA256和MD5.默认为MD5
	Body				string		`json:"body" xml:"body"`									//商品描述
	Detail				string		`json:"detail,omitempty" xml:"detail,omitempty"`			//商品详情
	Attach				string		`json:"attach,omitempty" xml:"attach,omitempty"`			//附加数据
	Out_trade_no		string		`json:"out_trade_no" xml:"out_trade_no"`					//商户订单号
	Fee_type			string		`json:"fee_type,omitempty" xml:"fee_type,omitempty"`		//货币类型
	Total_fee			int			`json:"total_feel" xml:"total_feel"`						//总金额
	Spbill_create_ip	string		`json:"spbill_create_ip" xml:"spbill_create_ip"`			//终端IP
	Time_start			string		`json:"time_start,omitempty" xml:"time_start,omitempty"`	//交易起始时间
	Time_expire			string		`json:"time_expire,omitempty" xml:"time_expire,omitempty"`  //交易结束时间
	Goods_tag			string		`json:"goods_tag,omitempty" xml:"goods_tag,omitempty"`		//订单优惠标记
	Notify_url			string		`json:"notify_url" xml:"notify_url"`						//通知地址
	Trade_type			string		`json:"trade_type" xml:"trade_type"`						//交易类型
	Product_id			string		`json:"product_id,omitempty" xml:"product_id,omitempty"`	//商品id，trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义。
	Limit_pay			string 		`json:"limit_pay,omitempty" xml:"limit_pay,omitempty"`		//指定支付方式，no_credit--指定不能使用信用卡支付
	Openid				string		`json:"openid" xml:"openid"`								//用户标识
	Sub_openid			string		`json:"sub_openid,omitempty" xml:"sub_opneid,omitempty"`	//子商户用户标识  trade_type=JSAPI，此参数必传，用户在主商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid。
	Scene_info			string		`json:"scene_info,omitempty" xml:"scene_info,omitempty"`	//场景信息
	RequestXML       	string 		`xml:"-"`      								                //存放最终请求xml串
}


//公众号支付统一下单返回
type OrderResponseSl struct {
	XMLName      		xml.Name 	`xml:"xml"`
	Return_code  		string   	`json:"return_code" xml:"return_code"`          	//返回状态码
	Return_msg   		string   	`json:"return_msg" xml:"return_msg"`           		//返回信息
	
	//以下字段在return_code为SUCCESS的时候有返回
	Appid        		string 		`json:"appid" xml:"appid"`                  		//服务商公众号ID
	Mch_id       		string 		`json:"mch_id" xml:"mch_id"`                 		//服务商商户号
	Sub_appid			string		`json:"sub_appid" xml:"sub_appid,omitempty"`		//jsapi：子商户公众号id  MiniProgram：小程序的appid
	Sub_mch_id			string		`json:"sub_mch_id" xml:"sub_mch_id"`				//子商户商户号
	Device_info  		string 		`json:"device_info" xml:"device_info,omitempty"`  	//设备号
	Nonce_str    		string 		`json:"nonce_str" xml:"nonce_str"`              	//随机字符串
	Sign         		string 		`json:"sign" xml:"sign"`                   			//签名
	Result_code  		string 		`json:"result_code" xml:"result_code"`            	//业务结果
	Err_code     		string 		`json:"err_code" xml:"err_code,omitempty"`     		//错误代码
	Err_code_des 		string 		`json:"err_code_des" xml:"err_code_des,omitempty"` 	//错误代码描述
	
	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	Trade_type   		string 		`json:"trade_type" xml:"trade_type"`   	//交易类型
	Prepay_id    		string 		`json:"prepay_id" xml:"prepay_id"`    	//预支付交易会话标识
	Code_url     		string 		`xml:"code_url,omitempty"`     			//二维码链接
	ReponseXML   		string 		`xml:"-"`                      			//存储返回的xml串
}


