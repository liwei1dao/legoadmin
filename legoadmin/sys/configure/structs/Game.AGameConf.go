
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


type GameAGameConf struct {
    _dataMap map[int32]*GameAGameConfData
    _dataList []*GameAGameConfData
}

func NewGameAGameConf(_buf []map[string]interface{}) (*GameAGameConf, error) {
    _dataList := make([]*GameAGameConfData, 0, len(_buf))
    dataMap := make(map[int32]*GameAGameConfData)

    for _, _ele_ := range _buf {
        if _v, err2 := NewGameAGameConfData(_ele_); err2 != nil {
            return nil, err2
        } else {
            _dataList = append(_dataList, _v)
            dataMap[_v.Id] = _v
        }
    }
    return &GameAGameConf{_dataList:_dataList, _dataMap:dataMap}, nil
}

func (table *GameAGameConf) GetDataMap() map[int32]*GameAGameConfData {
    return table._dataMap
}

func (table *GameAGameConf) GetDataList() []*GameAGameConfData {
    return table._dataList
}

func (table *GameAGameConf) Get(key int32) *GameAGameConfData {
    return table._dataMap[key]
}


