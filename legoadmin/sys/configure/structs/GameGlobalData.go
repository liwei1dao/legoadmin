
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type GameGlobalData struct {
    MaxChar int32
    Daytaskreward *atn
}

const TypeId_GameGlobalData = 1285622687

func (*GameGlobalData) GetTypeId() int32 {
    return 1285622687
}

func NewGameGlobalData(_buf map[string]interface{}) (_v *GameGlobalData, err error) {
    _v = &GameGlobalData{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["max_char"].(float64); !_ok_ { err = errors.New("max_char error"); return }; _v.MaxChar = int32(_tempNum_) }
    { var _ok_ bool; var _x_ map[string]interface{}; if _x_, _ok_ = _buf["daytaskreward"].(map[string]interface{}); !_ok_ { err = errors.New("daytaskreward error"); return }; if _v.Daytaskreward, err = Newatn(_x_); err != nil { return } }
    return
}

