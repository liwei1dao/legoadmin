
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


type GamePlayerBuff struct {
    _dataMap map[int32]*GamePlayerBuffData
    _dataList []*GamePlayerBuffData
}

func NewGamePlayerBuff(_buf []map[string]interface{}) (*GamePlayerBuff, error) {
    _dataList := make([]*GamePlayerBuffData, 0, len(_buf))
    dataMap := make(map[int32]*GamePlayerBuffData)

    for _, _ele_ := range _buf {
        if _v, err2 := NewGamePlayerBuffData(_ele_); err2 != nil {
            return nil, err2
        } else {
            _dataList = append(_dataList, _v)
            dataMap[_v.Id] = _v
        }
    }
    return &GamePlayerBuff{_dataList:_dataList, _dataMap:dataMap}, nil
}

func (table *GamePlayerBuff) GetDataMap() map[int32]*GamePlayerBuffData {
    return table._dataMap
}

func (table *GamePlayerBuff) GetDataList() []*GamePlayerBuffData {
    return table._dataList
}

func (table *GamePlayerBuff) Get(key int32) *GamePlayerBuffData {
    return table._dataMap[key]
}


