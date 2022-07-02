package state

import . "luago/api"

func (self *luaState) CreateTable(nArr,nRec int){
  t:=newLuaTable(nArr,nRec)
  self.state.push(t)
}


