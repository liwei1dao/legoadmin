
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type TestRectangle struct {
    Width float32
    Height float32
}

const TypeId_TestRectangle = -31893773

func (*TestRectangle) GetTypeId() int32 {
    return -31893773
}

func NewTestRectangle(_buf map[string]interface{}) (_v *TestRectangle, err error) {
    _v = &TestRectangle{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["width"].(float64); !_ok_ { err = errors.New("width error"); return }; _v.Width = float32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["height"].(float64); !_ok_ { err = errors.New("height error"); return }; _v.Height = float32(_tempNum_) }
    return
}

