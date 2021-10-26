package andflow_plugin

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

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
		if len(params)>=3{
			var userid int
			switch params[0].(type) {
			case string:
				userid, _ = strconv.Atoi(params[0].(string))
				break
			case int, int8, int32, int64:
				value := fmt.Sprintf("%d", params[0])
				userid, _ = strconv.Atoi(value)

				break
			case uint, uint8, uint32, uint64:
				value := fmt.Sprintf("%d", params[0])
				userid, _ = strconv.Atoi(value)

				break
			case float32, float64:

				value := fmt.Sprintf("%.f", params[0])
				userid, _ = strconv.Atoi(value)

				break
			}

			r:=d.PrepareMetadata(userid,params[1].(string),params[2].(string))
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


