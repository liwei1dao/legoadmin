
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type mtn struct {
    K int32
    N int32
}

const TypeId_mtn = 108455

func (*mtn) GetTypeId() int32 {
    return 108455
}

func Newmtn(_buf map[string]interface{}) (_v *mtn, err error) {
    _v = &mtn{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["k"].(float64); !_ok_ { err = errors.New("k error"); return }; _v.K = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["n"].(float64); !_ok_ { err = errors.New("n error"); return }; _v.N = int32(_tempNum_) }
    return
}

