
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type TestTestExcelBean1 struct {
    X1 int32
    X2 string
    X3 int32
    X4 float32
}

const TypeId_TestTestExcelBean1 = -1738345160

func (*TestTestExcelBean1) GetTypeId() int32 {
    return -1738345160
}

func NewTestTestExcelBean1(_buf map[string]interface{}) (_v *TestTestExcelBean1, err error) {
    _v = &TestTestExcelBean1{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["x1"].(float64); !_ok_ { err = errors.New("x1 error"); return }; _v.X1 = int32(_tempNum_) }
    { var _ok_ bool; if _v.X2, _ok_ = _buf["x2"].(string); !_ok_ { err = errors.New("x2 error"); return } }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["x3"].(float64); !_ok_ { err = errors.New("x3 error"); return }; _v.X3 = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["x4"].(float64); !_ok_ { err = errors.New("x4 error"); return }; _v.X4 = float32(_tempNum_) }
    return
}
