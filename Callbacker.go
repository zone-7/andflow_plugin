package andflow

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

var conn net.Conn
var client *rpc.Client

func instance(HostPort string)*rpc.Client{
	var err error
	conn ,err = net.Dial("tcp", fmt.Sprintf("127.0.0.1:%s",HostPort))
	if err != nil {
		log.Println("连接宿主失败: ", err)

		return  nil
	}
	client = jsonrpc.NewClient(conn)
	return client
}
func ConnectRpcHost(HostPort string)(*rpc.Client,error){
	var err error
	if conn != nil {
		err := conn.SetReadDeadline(time.Now().Add(time.Second * 1))
		if err != nil {
			log.Println("网络端口重连", err)
			client = nil
		}
	}

	if client == nil {
		client = instance(HostPort)
	}


	return client,err
}

type InitCallbackerImpl struct{
	HostPort string
}


func (i *InitCallbackerImpl) initCallback(method string,params []interface{})([]interface{},error) {
	conn ,err:= ConnectRpcHost(i.HostPort)
	if err!=nil{
		return nil, err
	}

	res:=CallbackResponse{}

	req:=CallbackRequest{}
	req.Callbacker="Init"
	req.Method = method
	req.Params = params

	err = conn.Call("RpcExecuter.Execute",req , &res)
	if err!=nil{
		return nil,err
	}
	if res.Results==nil{
		res.Results=[]interface{}{}
	}
	return res.Results,err
}

func (i *InitCallbackerImpl)GetFlowPluginPath() string{
	res,_:= i.initCallback( "GetFlowPluginPath",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}


type ActionCallbackerImpl struct {
	RuntimeId string
	HostPort string
}

func (p *ActionCallbackerImpl)actionCallback(runtimeId string,method string,params []interface{})([]interface{},error){
	conn ,err:= ConnectRpcHost(p.HostPort)
	if err!=nil{
		return nil, err
	}




	res:=CallbackResponse{}

	req:=CallbackRequest{}
	req.RuntimeId = runtimeId
	req.Callbacker="Action"
	req.Method = method
	req.Params = params

	err = conn.Call("RpcExecuter.Execute",req , &res)
	if err!=nil{
		return nil,err
	}
	if res.Results==nil{
		res.Results=[]interface{}{}
	}
	return res.Results,err
}


func (p *ActionCallbackerImpl) GetRuntime() string{

	res,_:= p.actionCallback(p.RuntimeId,"GetRuntime",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}

func (p *ActionCallbackerImpl) GetFlow() string{
	res,_:= p.actionCallback(p.RuntimeId,"GetFlow",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl)GetFlowCode() string{
	res,_:= p.actionCallback(p.RuntimeId,"GetFlowCode",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl)GetFlowName()string{
	res,_:= p.actionCallback(p.RuntimeId,"GetFlowName",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl)GetFlowType()string{
	res,_:= p.actionCallback(p.RuntimeId,"GetFlowType",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl)GetFlowLogType()string{
	res,_:= p.actionCallback(p.RuntimeId,"GetFlowLogType",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}

func (p *ActionCallbackerImpl)GetFlowDict(name string)string{
	res,_:= p.actionCallback(p.RuntimeId,"GetFlowDict",[]interface{}{name})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}

func (p *ActionCallbackerImpl)GetFlowTimeout()string{
	res,_:= p.actionCallback(p.RuntimeId,"GetFlowTimeout",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}


func (p *ActionCallbackerImpl) GetAction(aid string) string{
	res,_:= p.actionCallback(p.RuntimeId,"GetAction",[]interface{}{aid})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}


func (p *ActionCallbackerImpl) GetActionName(aid string) string{
	res,_:= p.actionCallback(p.RuntimeId,"GetActionName",[]interface{}{aid})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl) GetActionTitle(aid string) string{
	res,_:= p.actionCallback(p.RuntimeId,"GetActionTitle",[]interface{}{aid})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl) GetActionIcon(aid string) string{
	res,_:= p.actionCallback(p.RuntimeId,"GetActionIcon",[]interface{}{aid})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}

func(p *ActionCallbackerImpl) GetActionScript(aid string) string{
	res,_:= p.actionCallback(p.RuntimeId,"GetActionScript",[]interface{}{aid})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func(p *ActionCallbackerImpl) GetActionFilter(aid string) string{
	res,_:= p.actionCallback(p.RuntimeId,"GetActionFilter",[]interface{}{aid})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}

func (p *ActionCallbackerImpl) GetActionParam(aid,name string) string{
	res,_:= p.actionCallback(p.RuntimeId,"GetActionParam",[]interface{}{aid,name})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl) 	GetActionParams(aid string) map[string]string{
	res,_:= p.actionCallback(p.RuntimeId,"GetActionParams",[]interface{}{aid})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(map[string]string)
		if ok{
			return r
		}
	}
	return nil
}

func (p *ActionCallbackerImpl) GetRuntimeActionData(aid string,name string) interface{}{
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeActionData",[]interface{}{aid,name})

	if len(res)>0{

		return res[0]

	}
	return nil
}

func (p *ActionCallbackerImpl) SetRuntimeActionData(aid string,name string,value interface{}) {
	p.actionCallback(p.RuntimeId,"SetRuntimeActionData",[]interface{}{aid,name,value})
}



func (p *ActionCallbackerImpl) SetRuntimeActionIcon(aid string,icon string){
	p.actionCallback(p.RuntimeId,"SetRuntimeActionIcon",[]interface{}{aid,icon})
}

func (p *ActionCallbackerImpl) GetRuntimeActionDatas(aid string) map[string]interface{}{
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeActionDatas",[]interface{}{aid})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(map[string]interface{})
		if ok{
			return r
		}
	}
	return nil
}

func (p *ActionCallbackerImpl) LogRuntimeActionError(aid string,content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeActionError",[]interface{}{aid,content})

}
func (p *ActionCallbackerImpl) LogRuntimeActionInfo(aid string, content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeActionInfo",[]interface{}{aid,content})

}
func (p *ActionCallbackerImpl) LogRuntimeActionWarn(aid string, content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeActionWarn",[]interface{}{aid,content})


}


func (p *ActionCallbackerImpl) LogRuntimeLinkError(sid string,tid string, content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeLinkError",[]interface{}{sid,tid,content})

}
func (p *ActionCallbackerImpl) LogRuntimeLinkInfo(sid string,tid string, content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeLinkInfo",[]interface{}{sid,tid,content})

}

func (p *ActionCallbackerImpl) LogRuntimeLinkWarn(sid string,tid string, content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeLinkWarn",[]interface{}{sid,tid,content})
}




func (p *ActionCallbackerImpl) LogRuntimeFlowError(content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeFlowError",[]interface{}{content})

}
func (p *ActionCallbackerImpl) LogRuntimeFlowInfo(content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeFlowInfo",[]interface{}{content})
}
func (p *ActionCallbackerImpl) LogRuntimeFlowWarn(content string){
	p.actionCallback(p.RuntimeId,"LogRuntimeFlowWarn",[]interface{}{content})
}

func (p *ActionCallbackerImpl) ShowMessage(aid string, content_type string , content string){
	p.actionCallback(p.RuntimeId,"ShowMessage",[]interface{}{aid,content_type,content})
}

func (p *ActionCallbackerImpl) ShowRuntimeState(){
	p.actionCallback(p.RuntimeId,"ShowRuntimeState",[]interface{}{})
}

func (p *ActionCallbackerImpl) ShowText(aid string,content string){
	p.actionCallback(p.RuntimeId,"ShowText",[]interface{}{aid,content})
}
func (p *ActionCallbackerImpl) ShowGrid(aid string,content string){
	p.actionCallback(p.RuntimeId,"ShowGrid",[]interface{}{aid,content})

}
func (p *ActionCallbackerImpl) ShowChart(aid string,content string){
	p.actionCallback(p.RuntimeId,"ShowChart",[]interface{}{aid,content})
}
func (p *ActionCallbackerImpl) ShowHtml(aid string,content string){
	p.actionCallback(p.RuntimeId,"ShowHtml",[]interface{}{aid,content})
}

func (p *ActionCallbackerImpl) ShowForm(aid string,content string){
	p.actionCallback(p.RuntimeId,"ShowForm",[]interface{}{aid,content})
}

func (p *ActionCallbackerImpl) ShowWeb(aid string,content string){
	p.actionCallback(p.RuntimeId,"ShowWeb",[]interface{}{aid,content})
}

func (p *ActionCallbackerImpl) GetRuntimeId()string{
	return p.RuntimeId
}

func (p *ActionCallbackerImpl) GetRuntimeDes()string{
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeDes",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl) GetRuntimeContextId()string {
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeContextId",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}

func (p *ActionCallbackerImpl) GetRuntimeClientId()string {
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeClientId",[]interface{}{})

	if len(res)>0{

		rt := res[0]
		r,ok:=rt.(string)
		if ok{
			return r
		}
	}
	return ""
}
func (p *ActionCallbackerImpl) SetRuntimeIserror(iserror int){
	p.actionCallback(p.RuntimeId,"SetRuntimeIserror",[]interface{}{iserror})

}
func (p *ActionCallbackerImpl) GetRuntimeIserror() int{
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeIserror",[]interface{}{})

	if len(res)>0{
		rt := res[0]
		r,ok:=rt.(int)
		if ok{
			return r
		}
	}
	return 0
}

func (p *ActionCallbackerImpl) SetRuntimeFlowState(state int){
	p.actionCallback(p.RuntimeId,"SetRuntimeFlowState",[]interface{}{state})
}

func (p *ActionCallbackerImpl) GetRuntimeFlowState()int {
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeFlowState",[]interface{}{})

	if len(res)>0{
		rt := res[0]
		r,ok:=rt.(int)
		if ok{
			return r
		}
	}
	return 0
}
func (p *ActionCallbackerImpl) SetRuntimeFlowBeginTime(t int64){
	p.actionCallback(p.RuntimeId,"SetRuntimeFlowBeginTime",[]interface{}{t})
}
func (p *ActionCallbackerImpl) GetRuntimeFlowBeginTime(  )int64{
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeFlowBeginTime",[]interface{}{})

	if len(res)>0{
		rt := res[0]
		r,ok:=rt.(int64)
		if ok{
			return r
		}
	}
	return 0
}
func (p *ActionCallbackerImpl) SetRuntimeFlowEndTime(t int64){
	p.actionCallback(p.RuntimeId,"SetRuntimeFlowEndTime",[]interface{}{t})
}
func (p *ActionCallbackerImpl) GetRuntimeFlowEndTime()int64{
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeFlowEndTime",[]interface{}{})

	if len(res)>0{
		rt := res[0]
		r,ok:=rt.(int64)
		if ok{
			return r
		}
	}
	return 0
}
func (p *ActionCallbackerImpl) SaveRuntime(){
	p.actionCallback(p.RuntimeId,"SaveRuntime",[]interface{}{})
}



func (p *ActionCallbackerImpl) GetRuntimeData(name string)interface{} {
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeFlowEndTime",[]interface{}{name})

	if len(res)>0{
		return res[0]

	}
	return nil
}
func (p *ActionCallbackerImpl) GetRuntimeDatas() map[string]interface{}{
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeDatas",[]interface{}{})

	if len(res)>0{
		rt := res[0]
		r,ok:=rt.(map[string]interface{})
		if ok{
			return r
		}
	}
	return nil
}
func (p *ActionCallbackerImpl) SetRuntimeData(name string,value interface{}) {
	p.actionCallback(p.RuntimeId,"SetRuntimeData",[]interface{}{name,value})
}

func (p *ActionCallbackerImpl) GetRuntimeParam(name string)interface{} {
	res,_:= p.actionCallback(p.RuntimeId,"GetRuntimeParam",[]interface{}{name})

	if len(res)>0{
		rt := res[0]
		r,ok:=rt.(interface{})
		if ok{
			return r
		}
	}
	return nil
}

func (p *ActionCallbackerImpl) SetRuntimeParam(name string,value interface{}) {

	p.actionCallback(p.RuntimeId,"SetRuntimeParam",[]interface{}{name,value})

}

func (p *ActionCallbackerImpl) ExecuteFlowByCode(clientId string,contextId string, code string,params map[string]interface{})(string,error){
	res,_:= p.actionCallback(p.RuntimeId,"ExecuteFlowByCode",[]interface{}{clientId,contextId,code,params})

	if len(res)==2{
		rt := res[0]
		err := res[1]
		r,ok1:=rt.(string)
		e,ok2:=err.(error)
		if ok1 && ok2{
			return r,e
		}
	}

	return "",nil
}


func (p *ActionCallbackerImpl) ExecuteFlowByModel(clientId string,contextId string, flow string,params map[string]interface{})(string,error){
	res,_:= p.actionCallback(p.RuntimeId,"ExecuteFlowByModel",[]interface{}{clientId,contextId,flow,params})

	if len(res)==2{
		rt := res[0]
		err := res[1]
		r,ok1:=rt.(string)
		e,ok2:=err.(error)
		if ok1 && ok2{
			return r,e
		}
	}

	return "",nil

}

func (p *ActionCallbackerImpl) ExtendRuntimeAction(aid string, rt string){

	p.actionCallback(p.RuntimeId,"ExtendRuntimeAction",[]interface{}{aid,rt})

}
