
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type vector3 struct {
    X float32
    Y float32
    Z float32
}

const TypeId_vector3 = 337790800

func (*vector3) GetTypeId() int32 {
    return 337790800
}

func Newvector3(_buf map[string]interface{}) (_v *vector3, err error) {
    _v = &vector3{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["x"].(float64); !_ok_ { err = errors.New("x error"); return }; _v.X = float32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["y"].(float64); !_ok_ { err = errors.New("y error"); return }; _v.Y = float32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["z"].(float64); !_ok_ { err = errors.New("z error"); return }; _v.Z = float32(_tempNum_) }
    return
}

