
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type GameBR004ConfData struct {
    Id int32
    Redmultiple float64
    Grennmultiple float64
    Purplemultiple int32
    Numbermultiple int32
    Step []int32
    Off int32
    Rmin int32
    Rmax int32
    Spacekill float64
    Spacetrun int32
}

const TypeId_GameBR004ConfData = -490758624

func (*GameBR004ConfData) GetTypeId() int32 {
    return -490758624
}

func NewGameBR004ConfData(_buf map[string]interface{}) (_v *GameBR004ConfData, err error) {
    _v = &GameBR004ConfData{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["id"].(float64); !_ok_ { err = errors.New("id error"); return }; _v.Id = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["redmultiple"].(float64); !_ok_ { err = errors.New("redmultiple error"); return }; _v.Redmultiple = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["grennmultiple"].(float64); !_ok_ { err = errors.New("grennmultiple error"); return }; _v.Grennmultiple = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["purplemultiple"].(float64); !_ok_ { err = errors.New("purplemultiple error"); return }; _v.Purplemultiple = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["numbermultiple"].(float64); !_ok_ { err = errors.New("numbermultiple error"); return }; _v.Numbermultiple = int32(_tempNum_) }
     {
                    var _arr_ []interface{}
                    var _ok_ bool
                    if _arr_, _ok_ = _buf["step"].([]interface{}); !_ok_ { err = errors.New("step error"); return }
    
                    _v.Step = make([]int32, 0, len(_arr_))
                    
                    for _, _e_ := range _arr_ {
                        var _list_v_ int32
                        { var _ok_ bool; var _x_ float64; if _x_, _ok_ = _e_.(float64); !_ok_ { err = errors.New("_list_v_ error"); return }; _list_v_ = int32(_x_) }
                        _v.Step = append(_v.Step, _list_v_)
                    }
                }

    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["off"].(float64); !_ok_ { err = errors.New("off error"); return }; _v.Off = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["rmin"].(float64); !_ok_ { err = errors.New("rmin error"); return }; _v.Rmin = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["rmax"].(float64); !_ok_ { err = errors.New("rmax error"); return }; _v.Rmax = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["spacekill"].(float64); !_ok_ { err = errors.New("spacekill error"); return }; _v.Spacekill = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["spacetrun"].(float64); !_ok_ { err = errors.New("spacetrun error"); return }; _v.Spacetrun = int32(_tempNum_) }
    return
}
