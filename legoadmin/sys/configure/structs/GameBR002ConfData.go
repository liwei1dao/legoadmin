
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type GameBR002ConfData struct {
    Id int32
    Bettime int32
    Finishtime int32
}

const TypeId_GameBR002ConfData = -1170816994

func (*GameBR002ConfData) GetTypeId() int32 {
    return -1170816994
}

func NewGameBR002ConfData(_buf map[string]interface{}) (_v *GameBR002ConfData, err error) {
    _v = &GameBR002ConfData{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["id"].(float64); !_ok_ { err = errors.New("id error"); return }; _v.Id = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["bettime"].(float64); !_ok_ { err = errors.New("bettime error"); return }; _v.Bettime = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["finishtime"].(float64); !_ok_ { err = errors.New("finishtime error"); return }; _v.Finishtime = int32(_tempNum_) }
    return
}

