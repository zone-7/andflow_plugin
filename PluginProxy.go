package andflow

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

var PluginPort string

type RpcExecuter struct {
	action Action

}

func (r *RpcExecuter)Execute(request ActionRequest , response *ActionResponse ) error{
	defer func(res *ActionResponse) {
		if info := recover(); info != nil {

			res.Code=-1
			res.Msg = fmt.Sprintf( "%v",info)

		}
	}(response)

	var err error
	if response.Results==nil{
		response.Results=make([]interface{},0)
	}
	if request.Method=="GetName"{
		name:=r.action.GetName()
		response.Results = append(response.Results,name)
	}
	if request.Method=="Init"{
		callbacker := &InitCallbackerImpl{HostPort:request.HostPort}
		r.action.Init(callbacker)
	}
	if request.Method=="Filter"{
		if len(request.Method)==3{
			panic("Filter 参数不正确")
		}

		runtimeId,ok1:=request.Params[0].(string)
		preActionId,ok2:=request.Params[1].(string)
		actionId,ok3:=request.Params[2].(string)
		if ok1 && ok2 && ok3{

			callbacker := &ActionCallbackerImpl{HostPort:request.HostPort,RuntimeId:runtimeId}

			ctx:=context.Background()

			pass,err:=r.action.Filter(ctx, runtimeId,preActionId,actionId,callbacker)
			if err==nil{
				response.Results = []interface{}{pass,err}
			}else{
				panic(err)
			}


		}else{
			panic("Filter 参数格式不正确")
		}
	}
	if request.Method=="Exec"{
		if len(request.Method)==3{
			panic("Exec 参数不正确")
		}

		runtimeId,ok1:=request.Params[0].(string)
		preActionId,ok2:=request.Params[1].(string)
		actionId,ok3:=request.Params[2].(string)
		if ok1 && ok2 && ok3{
			callbacker := &ActionCallbackerImpl{HostPort:request.HostPort, RuntimeId:runtimeId}

			ctx:=context.Background()

			r,err:=r.action.Exec(ctx, runtimeId,preActionId,actionId,callbacker)
			if err==nil{
				response.Results = []interface{}{r,err}
			}else{
				panic(err)
			}

		}else{
			panic("Exec 参数格式不正确")
		}

	}

	return err
}



func StartPluginServer(action Action) error{
	if len(os.Args)<2{
		log.Println("参数不正确")
		return errors.New("参数不正确")
	}
	PluginPort = os.Args[1]
	log.Println("正在启动插件",PluginPort)

	executer:=new(RpcExecuter)
	executer.action = action

	rpc.Register(executer) // 注册rpc服务

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", PluginPort) )
	if err != nil {
		log.Fatalln("启动插件失败: ", err)
		return err
	}

	log.Println("插件启动完成",PluginPort)

	for {
		log.Println("插件等待服务",PluginPort)
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			continue
		}

		go func(conn net.Conn) { // 并发处理客户端请求
			//fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}




func InitPlugin(d Action) error{
	var err error
	response := &ActionResponse{}
	defer func(res *ActionResponse) {
		if info := recover();info!=nil{
			response.Code=-1
			response.Msg = fmt.Sprintf("%v",info)
		}
		re,_:=json.Marshal(response)

		fmt.Println("andflow_plugin_response="+string(re))
	}(response)

	args := os.Args
	if len(args)<2{
		panic("参数错误")
	}

	reqStr:=args[1]

	reqJson,err := base64.StdEncoding.DecodeString(reqStr)
	if err!=nil{
		panic("参数格式错误")
	}

	fmt.Println("param="+string(reqJson))

	request := ActionRequest{}
	err = json.Unmarshal(reqJson, &request)
	if err!=nil{
		panic(err)
	}

	method:=request.Method
	params:=request.Params
	hostPort:=request.HostPort

	res:=make([]interface{},0)
	switch method {
	case "GetName":
		r := d.GetName()
		res = append(res, r)
	case "PrepareMetadata":
		if len(params)>=2{
			r:=d.PrepareMetadata(params[0].(string),params[1].(string))
			res= append(res, r)
		}else{
			panic("参数错误")
		}
	case "Filter":
		if len(params)>=3 {
			runtimeId := params[0].(string)
			preActionId := params[1].(string)
			actionId := params[2].(string)
			ctx,cancel:=context.WithCancel(context.Background())
			defer cancel()

			var actionCallback ActionCallbacker
			actionCallback = &ActionCallbackerImpl{RuntimeId:runtimeId,HostPort:hostPort}
			pass,err := d.Filter(ctx,runtimeId,preActionId,actionId,actionCallback)
			res = append(res, pass,err)
		}else{
			panic("参数错误")
		}

	case "Exec":
		if len(params)>=3 {
			runtimeId := params[0].(string)
			preActionId := params[1].(string)
			actionId := params[2].(string)
			ctx,cancel:=context.WithCancel(context.Background())
			defer cancel()
			var actionCallback ActionCallbacker
			actionCallback = &ActionCallbackerImpl{RuntimeId:runtimeId,HostPort:hostPort}
			r,err := d.Exec(ctx,runtimeId,preActionId,actionId,actionCallback)
			res = append(res, r, err)
		}else{
			panic("参数错误")
		}

	case "Init":
		initCallback := &InitCallbackerImpl{HostPort:hostPort}
		d.Init(initCallback)
	}

	response.Results = res

	return nil
}


