package main

import "testing"

func TestDidiParser(t *testing.T) {

	log := "[INFO][2021-04-23T15:36:45.056+0800][..alcon/task-taiphon/service/slowpay.TaskInterruptWhenDriverAssign/task_judge_service.go:323]__slowpay__||traceid=ac18498d6082790b6f192f58102e0303||spanid=cc9b6c7d20770827||hintCode=||hintContent={\"locale\":\"es_MX\",\"app_timeout_ms\":15000,\"location_cityid\":\"52159900\",\"lang\":\"es-MX\",\"trip_country\":\"MX\",\"dsim_env\":\"us01-sim100-v\",\"sample\":{\"code\":128},\"dlang\":\"es-MX\",\"driverUtcOffset\":\"-300\",\"location_country\":\"MX\"}||method=POST||host=10.157.155.136:8000||uri=/slow_pay/task_action/mq||params=||from=10.157.185.183:50054||srcMethod=||caller=||reassign order already send to delayQ, 87975639090828,1619163493\n"

	Dparser(log)

	//println(l)
}
