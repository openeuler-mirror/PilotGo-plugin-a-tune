package service

import (
	"encoding/json"
	"regexp"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func ResultAnalysis(taskId int, machine_uuid string) error {
	result, err := dao.GetResultByTaskIdAndUUID(taskId, machine_uuid, CommandTypeTune)
	if err != nil {
		return err
	}

	var iterResult model.TuningIterResult
	var iterResults []model.TuningIterResult
	var (
		finalOptimizationResult interface{}
		finalEvaluationValue    interface{}
		baselinePerformance     interface{}
	)
	for _, line := range strings.Split(result, "\n") {
		// 检查日志中是否包含调优结果信息
		if strings.Contains(line, "Current Tuning Progress") {
			// 迭代轮次
			reTuningStep := regexp.MustCompile(`\((\d+)/\d+\)`)
			matches := reTuningStep.FindStringSubmatch(line)
			if len(matches) > 0 {
				iterResult.TuningStep = matches[1]
			}
		} else if strings.Contains(line, "recommand parameters") {
			// 推荐参数
			reRecommandParams := regexp.MustCompile(`The \d+th recommand parameters is: (.+)`)
			matches := reRecommandParams.FindStringSubmatch(line)
			if len(matches) > 0 {
				iterResult.RecommandParameters = extractKeyValue(matches[1])
			}
		} else if strings.Contains(line, "th evaluation value") {
			reEvaluationValue := regexp.MustCompile(`evaluation value: ([^\n]+)`)
			matches := reEvaluationValue.FindStringSubmatch(line)
			if len(matches) > 0 {
				re := regexp.MustCompile(`\(([^)]+)\)\(([^)]+)\)`)
				m := re.FindStringSubmatch(matches[1])
				// 本轮评估值
				iterResult.EvaluationValue = extractKeyValue(m[1])
				// 性能提升
				iterResult.PerformanceImprovement = m[2]
				iterResults = append(iterResults, iterResult)
			}
		} else if strings.Contains(line, "The final optimization result") {
			// 最终优化结果
			reFinalOptimization := regexp.MustCompile(`The final optimization result is: ([^\n]+)`)
			matches := reFinalOptimization.FindStringSubmatch(line)
			if len(matches) > 0 {
				finalOptimizationResult = extractKeyValue(matches[1])
			}
		} else if strings.Contains(line, "The final evaluation value") {
			// 最终评估值
			reFinalEvaluationValue := regexp.MustCompile(`The final evaluation value is: (.+)`)
			matches := reFinalEvaluationValue.FindStringSubmatch(line)
			if len(matches) > 0 {
				finalEvaluationValue = extractKeyValue(matches[1])
			}
		} else if strings.Contains(line, "Baseline Performance is:") {
			// 基准性能
			reBaselinePerformance := regexp.MustCompile(`Baseline Performance is: \(([^)]+)\)`)
			matches := reBaselinePerformance.FindStringSubmatch(line)
			if len(matches) > 0 {
				baselinePerformance = extractKeyValue(matches[1])
			}
		}
	}

	tuningResult := &model.TuningResults{
		TaskID:                  taskId,
		MachineUUID:             machine_uuid,
		TuningResult:            jsonToString(iterResults),
		FinalOptimizationResult: jsonToString(finalOptimizationResult),
		FinalEvaluationValue:    jsonToString(finalEvaluationValue),
		BaselinePerformance:     jsonToString(baselinePerformance),
	}
	if err = dao.SaveResultAnalysis(tuningResult); err != nil {
		return err
	}
	return nil
}

func GetResultAnalisis(taskId, machine_uuid string) (interface{}, error) {
	tune_result, err := dao.GetResultAnalisis(taskId, machine_uuid)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	result["id"] = tune_result.ID
	result["task_id"] = tune_result.TaskID
	result["machine_uuid"] = tune_result.MachineUUID
	result["tuning_result"] = jsonStringToObject(tune_result.TuningResult)
	result["final_optimization_result"] = jsonStringToObject(tune_result.FinalOptimizationResult)
	result["final_evaluation_value"] = jsonStringToObject(tune_result.FinalEvaluationValue)
	result["baseline_performance"] = jsonStringToObject(tune_result.BaselinePerformance)
	return result, err
}

func extractKeyValue(input string) map[string]string {
	result := make(map[string]string)
	re := regexp.MustCompile(`(\w+)=(\w+)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		result[match[1]] = match[2]
	}
	return result
}

func jsonStringToObject(input string) interface{} {
	var jsonString interface{}
	if err := json.Unmarshal([]byte(input), &jsonString); err != nil {
		return nil
	}
	return jsonString
}

func jsonToString(input interface{}) string {
	results, err := json.Marshal(input)
	if err != nil {
		logger.Error("序列化失败：%v", err.Error())
		return ""
	}
	return string(results)
}
