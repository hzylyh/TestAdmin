package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model/vo"
	"salotto/service"
	"salotto/utils/qjson"
)

var DashboardSrv = &dashboardService{}

type dashboardService struct {
}

func (ds *dashboardService) GetCaseRunInfo(json *qjson.QJson) (*vo.CaseRunInfoVO, error) {

	var (
		caseRunInfo = &vo.CaseRunInfoVO{}
		total       int64
		success     int64
		fail        int64
		beginTime   string
	)

	rows, err := service.DB.Raw("select t2.begin_time, t2.total, ifnull(t4.success, 0) as success, ifnull(t6.fail, 0) as fail from (select t1.begin_time,count(t1.step_his_id) as 'total', t1.begin_time from t_itf_case_step_run_his t1 group by t1.begin_time) t2 left join (select t3.begin_time,count(t3.step_his_id) as 'success' from t_itf_case_step_run_his t3 where t3.step_status = '成功' group by t3.begin_time) t4 on t2.begin_time = t4.begin_time left join (select t5.begin_time,count(t5.step_his_id) as 'fail' from t_itf_case_step_run_his t5 where t5.step_status = '失败' group by t5.begin_time) t6 on t2.begin_time = t6.begin_time order by t2.begin_time desc limit 7", "").Rows() // (*sql.Rows, error)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//fmt.Println(reflect.Type(caseRunInfo))
	for rows.Next() {
		if err = rows.Scan(&beginTime, &total, &success, &fail); err != nil {
			//if err = rows.Scan(rows, &caseRunInfo); err != nil {
			fmt.Println(err)
			return nil, err
		}

		caseRunInfo.BeginTime = append(caseRunInfo.BeginTime, beginTime)
		caseRunInfo.Total = append(caseRunInfo.Total, total)
		caseRunInfo.Success = append(caseRunInfo.Success, success)
		caseRunInfo.Fail = append(caseRunInfo.Fail, fail)

	}
	return caseRunInfo, nil
}
