// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package order

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Order) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Order[number], err)
}

func (x *Order) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Order) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CreateOrderRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateOrderRequest[number], err)
}

func (x *CreateOrderRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateOrderRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CartId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateOrderRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserAddr, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateOrderRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.ProductIds = append(x.ProductIds, v)
			return offset, err
		})
	return offset, err
}

func (x *CreateOrderRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.Quantities = append(x.Quantities, v)
			return offset, err
		})
	return offset, err
}

func (x *CreateOrderRequest) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v float32
			v, offset, err = fastpb.ReadFloat(buf, _type)
			if err != nil {
				return offset, err
			}
			x.Prices = append(x.Prices, v)
			return offset, err
		})
	return offset, err
}

func (x *CreateOrderResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateOrderResponse[number], err)
}

func (x *CreateOrderResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *CreateOrderResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderPaySuccessRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderPaySuccessRequest[number], err)
}

func (x *OrderPaySuccessRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *OrderPaySuccessResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderPaySuccessResponse[number], err)
}

func (x *OrderPaySuccessResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Order) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *Order) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Order) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *CreateOrderRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *CreateOrderRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CreateOrderRequest) fastWriteField2(buf []byte) (offset int) {
	if x.CartId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetCartId())
	return offset
}

func (x *CreateOrderRequest) fastWriteField3(buf []byte) (offset int) {
	if x.UserAddr == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUserAddr())
	return offset
}

func (x *CreateOrderRequest) fastWriteField4(buf []byte) (offset int) {
	if len(x.ProductIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 4, len(x.GetProductIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetProductIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *CreateOrderRequest) fastWriteField5(buf []byte) (offset int) {
	if len(x.Quantities) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 5, len(x.GetQuantities()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetQuantities()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *CreateOrderRequest) fastWriteField6(buf []byte) (offset int) {
	if len(x.Prices) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 6, len(x.GetPrices()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteFloat(buf[offset:], numTagOrKey, x.GetPrices()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *CreateOrderResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateOrderResponse) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *CreateOrderResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMessage())
	return offset
}

func (x *OrderPaySuccessRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *OrderPaySuccessRequest) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *OrderPaySuccessResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *OrderPaySuccessResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMessage())
	return offset
}

func (x *Order) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *Order) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetId())
	return n
}

func (x *Order) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetUserId())
	return n
}

func (x *CreateOrderRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *CreateOrderRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *CreateOrderRequest) sizeField2() (n int) {
	if x.CartId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetCartId())
	return n
}

func (x *CreateOrderRequest) sizeField3() (n int) {
	if x.UserAddr == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUserAddr())
	return n
}

func (x *CreateOrderRequest) sizeField4() (n int) {
	if len(x.ProductIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(4, len(x.GetProductIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetProductIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *CreateOrderRequest) sizeField5() (n int) {
	if len(x.Quantities) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(5, len(x.GetQuantities()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetQuantities()[numIdxOrVal])
			return n
		})
	return n
}

func (x *CreateOrderRequest) sizeField6() (n int) {
	if len(x.Prices) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(6, len(x.GetPrices()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeFloat(numTagOrKey, x.GetPrices()[numIdxOrVal])
			return n
		})
	return n
}

func (x *CreateOrderResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateOrderResponse) sizeField1() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetOrderId())
	return n
}

func (x *CreateOrderResponse) sizeField2() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMessage())
	return n
}

func (x *OrderPaySuccessRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *OrderPaySuccessRequest) sizeField1() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetOrderId())
	return n
}

func (x *OrderPaySuccessResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *OrderPaySuccessResponse) sizeField1() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMessage())
	return n
}

var fieldIDToName_Order = map[int32]string{
	1: "Id",
	2: "UserId",
}

var fieldIDToName_CreateOrderRequest = map[int32]string{
	1: "UserId",
	2: "CartId",
	3: "UserAddr",
	4: "ProductIds",
	5: "Quantities",
	6: "Prices",
}

var fieldIDToName_CreateOrderResponse = map[int32]string{
	1: "OrderId",
	2: "Message",
}

var fieldIDToName_OrderPaySuccessRequest = map[int32]string{
	1: "OrderId",
}

var fieldIDToName_OrderPaySuccessResponse = map[int32]string{
	1: "Message",
}
