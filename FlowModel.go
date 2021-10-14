package andflow_plugin

import (
	"context"
	"encoding/json"
	"runtime"
	"time"
)

type MetadataOptionLoader runtime.Func

type MetadataAttrsModel map[string]string

type MetadataOptionModel struct {
	Label string `bson:"label" json:"label"`
	Value string `bson:"value" json:"value"`
}

type MetadataPropertiesModel struct {
	Name        string                `bson:"name" json:"name"`
	Title       string                `bson:"title" json:"title"`
	Placeholder string                `bson:"placeholder" json:"placeholder"`
	Element     string                `bson:"element" json:"element"` //表单控件input textarea select
	Default     string                `bson:"default" json:"default"` //默认值
	Attrs       MetadataAttrsModel    `bson:"attrs" json:"attrs"`     //标签属性
	Options     []MetadataOptionModel `bson:"options" json:"options"`
	OptionMode  string				  `bson:"option_mode" json:"option_mode"` //单选single、多选multipy
	OptionsLoader func(userid int ,m *MetadataModel, p *MetadataPropertiesModel)		  	  `bson:"-" json:"-"`
	OptionsLoaderParams []string  `bson:"-" json:"-"`
}

type MetadataModel struct {
	Name     string                    `bson:"name" json:"name"`          //名称
	Title    string                    `bson:"title" json:"title"`        //标题
	Des      string                    `bson:"des" json:"des"`            //描述
	Group    string                    `bson:"group" json:"group"`		   //组
	Template string                    `bson:"template" json:"template"`  //模板
	Css      string                    `bson:"css" json:"css"`            //样式
	Icon     string                    `bson:"icon" json:"icon"`          //图标

	FlowCode string                    `bson:"flow_code" json:"flow_code"`//流程编码
	Params   []MetadataPropertiesModel `bson:"params" json:"params"` 			   //参数定义，与html和js  二选一
	ParamsHtml string                  `bson:"params_html" json:"params_html"`	   //参数配置html界面
	ParamsScript string                `bson:"params_script" json:"params_script"` //参数配置js脚本

}

type ActionModel struct {
	Id string 				 `bson:"id" json:"id"`
	Name string 			 `bson:"name" json:"name"`
	Title string		 	 `bson:"title" json:"title"`
	Icon string              `bson:"icon" json:"icon"`
	Des string 				 `bson:"des" json:"des"`
	Left string			 	 `bson:"left" json:"left"`
	Top string 				 `bson:"top" json:"top"`
	Width string 			 `bson:"width" json:"width"`
	Height string            `bson:"height" json:"height"`
	BodyWidth string		 `bson:"body_width" json:"body_width"`
	BodyHeight string        `bson:"body_height" json:"body_height"`
	Params map[string]string `bson:"params" json:"params"`
	Filter string 			 `bson:"filter" json:"filter"`
	Script string 			 `bson:"script" json:"script"`
}

type LinkModel struct{
	Title string 			`bson:"title" json:"title"`
	SourceId string 		`bson:"source_id" json:"source_id"`
	SourcePosition string 	`bson:"source_position" json:"source_position"`
	TargetId string 		`bson:"target_id" json:"target_id"`
	TargetPosition string 	`bson:"target_position" json:"target_position"`
	Filter string 			`bson:"filter" json:"filter"`
	Active string           `bson:"active" json:"active"`
}
type FlowParamModel struct {
	Name string	   `bson:"name" json:"name"`
	Value string 	`bson:"value" json:"value"`
}
type FlowDictModel struct {
	Name string	   `bson:"name" json:"name"`
	Label string   `bson:"label" json:"label"`
}

type FlowModel struct {
	Code string              `bson:"code" json:"code"`
	Name string              `bson:"name" json:"name"`						//流程名称
	FlowType string		     `bson:"flow_type" json:"flow_type"`			//流程类型
	LogType string           `bson:"log_type" json:"log_type"`				//日志类型
	Timeout string           `bson:"timeout" json:"timeout"`				//执行时效
	CacheTimeout string      `bson:"cache_timeout" json:"cache_timeout"`	//缓存时效
	Params []*FlowParamModel `bson:"params" json:"params"`                   //运行参数列表
	Dict []*FlowDictModel  	 `bson:"dict" json:"dict"`                   	 //运行字典列表
	SaveRuntime string       `bson:"save_runtime" json:"save_runtime"`		 //是否保存运行时内容
	ShowActionState string   `bson:"show_action_state" json:"show_action_state"` //是否显示节点运行状态
	ShowActionContent string `bson:"show_action_content" json:"show_action_content"` //是否显示节点运行消息内容

	Theme string             `bson:"theme" json:"theme"`					//皮肤
	LinkType string          `bson:"link_type" json:"link_type"`			//连接线类型
	Actions []*ActionModel   `bson:"actions" json:"actions"`				//节点列表
	Links []*LinkModel 	     `bson:"links" json:"links"`					//连线列表

	CreateUser     int	 	 `bson:"create_user" json:"create_user"`
	UpdateUser     int		 `bson:"update_user" json:"update_user"`
	CreateTime     time.Time  `bson:"create_time" json:"create_time"`
	UpdateTime     time.Time  `bson:"update_time" json:"update_time"`
}



type Action interface {
	GetName() string
	Init(callback interface{})
	PrepareMetadata(flowCode string, metadata string) string
	Filter(ctx context.Context,runtimeId string,preActionId string,actionId string, callback interface{}) (bool, error)
	Exec(ctx context.Context,runtimeId string,preActionId string,actionId string, callback interface{}) (interface{}, error)
}



type ActionRequest struct {
	HostPort string	`json:"host_port"`
	Method string	`json:"method"`
	Params []interface{} `json:"params"`
}

type ActionResponse struct{
	Results []interface{}   `json:"results"`
	Code int				`json:"code"`
	Msg string				`json:"msg"`
}

type CallbackRequest struct {
	Callbacker string		`json:"callbacker"`
	Method string			`json:"method"`
	RuntimeId string		`json:"runtime_id"`
	Params []interface{}	`json:"params"`
}

type CallbackResponse struct{
	Results []interface{}	`json:"results"`
	Code int			    `json:"code"`
	Msg string				`json:"msg"`
}
type InitCallbacker interface {
	GetFlowPluginPath() string
	GetFlowActionPath(name string) string

}

type ActionCallbacker interface {
	GetRuntime() string
	GetFlow() string
	GetFlowCode() string
	GetFlowName()string
	GetFlowType()string
	GetFlowLogType()string
	GetFlowDict(name string)string
	GetFlowTimeout()string
	GetAction(aid string) string
	GetActionName(aid string) string
	GetActionTitle(aid string) string
	GetActionIcon(aid string) string
	GetActionScript(aid string) string
	GetActionFilter(aid string) string
	GetActionParams(aid string) map[string]string
	GetActionParam(aid,name string) string
	ShowMessage(aid string,content_type string , content string)
	ShowRuntimeState()
	ShowText(aid string, content string)
	ShowGrid(aid string, content string)
	ShowChart(aid string, content string)
	ShowHtml(aid string, content string)
	ShowForm(aid string, content string)
	ShowWeb(aid string, content string)
	GetRuntimeActionData(aid ,name string) interface{}
	GetRuntimeActionDatas(aid string) map[string]interface{}
	SetRuntimeActionData(aid ,name string,value interface{})
	SetRuntimeActionIcon(aid ,icon string)
	SetRuntimeData(name string,value interface{})
	GetRuntimeData(name string) interface{}
	GetRuntimeDatas() map[string]interface{}
	GetRuntimeParam(name string) interface{}
	SetRuntimeParam(name string,value interface{})
	LogRuntimeFlowError(content string)
	LogRuntimeFlowInfo(content string)
	LogRuntimeFlowWarn(content string)
	LogRuntimeActionError(aid string , content string)
	LogRuntimeActionInfo(aid  string,content string)
	LogRuntimeActionWarn(aid string ,content string)
	LogRuntimeLinkError(sid string, tid string, content string)
	LogRuntimeLinkInfo(sid string, tid string, content string)
	LogRuntimeLinkWarn(sid string, tid string, content string)
	GetRuntimeId()string
	GetRuntimeDes()string
	GetRuntimeContextId()string
	GetRuntimeClientId()string
	SetRuntimeIserror(iserror int)
	GetRuntimeIserror() int
	SetRuntimeFlowState(state int)
	GetRuntimeFlowState()int
	SetRuntimeFlowBeginTime(t int64)
	GetRuntimeFlowBeginTime(  )int64
	SetRuntimeFlowEndTime(t int64)
	GetRuntimeFlowEndTime()int64
	SaveRuntime()
	ExecuteFlowByCode(clientId string,contextId string, code string,params map[string]interface{})(string,error)
	ExecuteFlowByModel(clientId string,contextId string, flow string,params map[string]interface{})(string,error)
	ExtendRuntimeAction(aid string,rt string)
}

func ParseActionCallbacker(obj interface{}) ActionCallbacker {
	p,ok:=obj.(ActionCallbacker)

	if ok == false {
		return nil
	}
	return p
}

func ParseMetadata(metadata string) *MetadataModel{
	if len(metadata)==0{
		return nil
	}
	meta:=MetadataModel{}

	err := json.Unmarshal([]byte(metadata),&meta)
	if err!=nil{
		return nil
	}
	return &meta
}

func ParseInitCallbacker(obj interface{}) InitCallbacker {
	p,ok:=obj.(InitCallbacker)

	if ok == false {
		return nil
	}
	return p
}

func (t *FlowModel) GetDict(name string)*FlowDictModel {
	if t.Dict==nil || len(t.Dict)==0{
		return nil
	}

	for _,d:=range t.Dict{
		if(d.Name==name){
			return d
		}
	}
	return nil
}

func (t *FlowModel) GetActionModel(id string)*ActionModel {
	for _,node:=range t.Actions{
		if node.Id==id{
			return node
		}
	}
	return nil
}
func (t *FlowModel) GetLinkByTargetId(id string) []*LinkModel{
	lks := make([]*LinkModel,0)
	for _,link := range t.Links{
		if link.TargetId==id{
			lks = append(lks, link)
		}
	}
	return lks
}


func (t *FlowModel) GetLinkBySourceId(id string) []*LinkModel{
	lks := make([]*LinkModel,0)

	for _,link := range t.Links{
		if link.SourceId==id{
			lks = append(lks, link)
		}
	}
	return lks
}

func (t *FlowModel) GetLinkBySourceIdAndTargetId(sid string,tid string) *LinkModel{

	for _,link := range t.Links{
		if link.SourceId==sid && link.TargetId==tid{
			return link
		}
	}
	return nil
}

//获取作为起点的节点
func (t *FlowModel) GetStartActionIds()  []string {

	actions_start:=make([]string,0)

	for _,action:=range t.Actions{
		targets := t.GetLinkByTargetId(action.Id)
		if targets==nil || len(targets)==0{
			actions_start = append(actions_start, action.Id)
		}
	}

	return actions_start
}



func (m *MetadataModel) ToJson() string{
	data,_ := json.Marshal(m)
	return string(data)
}