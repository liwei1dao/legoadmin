
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type GameAGameConfData struct {
    Id int32
    Control bool
    Poolbetmin float64
    Poolbetmax float64
    Poolcontrolwin bool
    Poolcontrollose bool
    Poolcontroltypes int32
    Poolkilllv int32
    Poolbasemoney float64
    Poolbasemlv int32
    Poolmoneymin float64
    Poolmoneytrun float64
    Poolmoneymax float64
    Poolcontrollosemodulus int32
    Poolcontrolloselvmin int32
    Poolcontrolloselvmax int32
    Poolcontrolwinmodulus int32
    Poolcontrolwinlvmin int32
    Poolcontrolwinlvmax int32
    Luckypooltax int32
    Luckypoolmultiple float64
}

const TypeId_GameAGameConfData = -1241913585

func (*GameAGameConfData) GetTypeId() int32 {
    return -1241913585
}

func NewGameAGameConfData(_buf map[string]interface{}) (_v *GameAGameConfData, err error) {
    _v = &GameAGameConfData{}
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["id"].(float64); !_ok_ { err = errors.New("id error"); return }; _v.Id = int32(_tempNum_) }
    { var _ok_ bool; if _v.Control, _ok_ = _buf["Control"].(bool); !_ok_ { err = errors.New("Control error"); return } }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolbetmin"].(float64); !_ok_ { err = errors.New("Poolbetmin error"); return }; _v.Poolbetmin = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolbetmax"].(float64); !_ok_ { err = errors.New("Poolbetmax error"); return }; _v.Poolbetmax = float64(_tempNum_) }
    { var _ok_ bool; if _v.Poolcontrolwin, _ok_ = _buf["Poolcontrolwin"].(bool); !_ok_ { err = errors.New("Poolcontrolwin error"); return } }
    { var _ok_ bool; if _v.Poolcontrollose, _ok_ = _buf["Poolcontrollose"].(bool); !_ok_ { err = errors.New("Poolcontrollose error"); return } }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolcontroltypes"].(float64); !_ok_ { err = errors.New("Poolcontroltypes error"); return }; _v.Poolcontroltypes = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolkilllv"].(float64); !_ok_ { err = errors.New("Poolkilllv error"); return }; _v.Poolkilllv = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolbasemoney"].(float64); !_ok_ { err = errors.New("Poolbasemoney error"); return }; _v.Poolbasemoney = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolbasemlv"].(float64); !_ok_ { err = errors.New("Poolbasemlv error"); return }; _v.Poolbasemlv = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolmoneymin"].(float64); !_ok_ { err = errors.New("Poolmoneymin error"); return }; _v.Poolmoneymin = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolmoneytrun"].(float64); !_ok_ { err = errors.New("Poolmoneytrun error"); return }; _v.Poolmoneytrun = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolmoneymax"].(float64); !_ok_ { err = errors.New("Poolmoneymax error"); return }; _v.Poolmoneymax = float64(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolcontrollosemodulus"].(float64); !_ok_ { err = errors.New("Poolcontrollosemodulus error"); return }; _v.Poolcontrollosemodulus = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolcontrolloselvmin"].(float64); !_ok_ { err = errors.New("Poolcontrolloselvmin error"); return }; _v.Poolcontrolloselvmin = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolcontrolloselvmax"].(float64); !_ok_ { err = errors.New("Poolcontrolloselvmax error"); return }; _v.Poolcontrolloselvmax = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolcontrolwinmodulus"].(float64); !_ok_ { err = errors.New("Poolcontrolwinmodulus error"); return }; _v.Poolcontrolwinmodulus = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolcontrolwinlvmin"].(float64); !_ok_ { err = errors.New("Poolcontrolwinlvmin error"); return }; _v.Poolcontrolwinlvmin = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Poolcontrolwinlvmax"].(float64); !_ok_ { err = errors.New("Poolcontrolwinlvmax error"); return }; _v.Poolcontrolwinlvmax = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Luckypooltax"].(float64); !_ok_ { err = errors.New("Luckypooltax error"); return }; _v.Luckypooltax = int32(_tempNum_) }
    { var _ok_ bool; var _tempNum_ float64; if _tempNum_, _ok_ = _buf["Luckypoolmultiple"].(float64); !_ok_ { err = errors.New("Luckypoolmultiple error"); return }; _v.Luckypoolmultiple = float64(_tempNum_) }
    return
}

