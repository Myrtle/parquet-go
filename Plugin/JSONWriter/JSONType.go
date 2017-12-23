package JSONWriter

import (
	"fmt"
	"github.com/xitongsys/parquet-go/ParquetType"
	"github.com/xitongsys/parquet-go/parquet"
	"reflect"
)

func JSONTypeToParquetType(val reflect.Value, pT *parquet.Type, cT *parquet.ConvertedType) interface{} {
	if val.Type().Kind() == reflect.Interface && val.IsNil() {
		return nil
	}
	s := fmt.Sprintf("%v", val)
	return ParquetType.StrToParquetType(s, pT, cT)
}