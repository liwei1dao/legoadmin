
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type GameBR003ConfData struct {
    Id int32
    Redmultiple float64
    Grennmultiple float64
    Purplemultiple float64
    Numbermultiple float64
}

const TypeId_GameBR003ConfData = 1316695839

func (*GameBR003ConfData) GetTypeId() int32 {
    return 1316695839
}

func NewGameBR003ConfData(_buf map[string]interface{}) (_v *GameBR003ConfData, err error) {
    _v = &GameBR003ConfData{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["id"].(float64); !_ok_ { err = errors.New("id error"); return }; _v.Id = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["redmultiple"].(float64); !_ok_ { err = errors.New("redmultiple error"); return }; _v.Redmultiple = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["grennmultiple"].(float64); !_ok_ { err = errors.New("grennmultiple error"); return }; _v.Grennmultiple = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["purplemultiple"].(float64); !_ok_ { err = errors.New("purplemultiple error"); return }; _v.Purplemultiple = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["numbermultiple"].(float64); !_ok_ { err = errors.New("numbermultiple error"); return }; _v.Numbermultiple = float64(_tempNum_) }
    return
}

