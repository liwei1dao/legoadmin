
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type GameDJ001ConfData struct {
    Id int32
    Multiple float64
    Nweight int32
    Winweight int32
    Loseweight int32
}

const TypeId_GameDJ001ConfData = -982830297

func (*GameDJ001ConfData) GetTypeId() int32 {
    return -982830297
}

func NewGameDJ001ConfData(_buf map[string]interface{}) (_v *GameDJ001ConfData, err error) {
    _v = &GameDJ001ConfData{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["id"].(float64); !_ok_ { err = errors.New("id error"); return }; _v.Id = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["multiple"].(float64); !_ok_ { err = errors.New("multiple error"); return }; _v.Multiple = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["nweight"].(float64); !_ok_ { err = errors.New("nweight error"); return }; _v.Nweight = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["winweight"].(float64); !_ok_ { err = errors.New("winweight error"); return }; _v.Winweight = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["loseweight"].(float64); !_ok_ { err = errors.New("loseweight error"); return }; _v.Loseweight = int32(_tempNum_) }
    return
}
