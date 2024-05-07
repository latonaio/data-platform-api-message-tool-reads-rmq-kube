package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-message-tool-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-message-tool-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var messageTool *[]dpfm_api_output_formatter.MessageTool
	var text *[]dpfm_api_output_formatter.Text
	for _, fn := range accepter {
		switch fn {
		case "MessageTool":
			func() {
				messageTool = c.MessageTool(mtx, input, output, errs, log)
			}()
		case "MessageTools":
			func() {
				messageTool = c.MessageTools(mtx, input, output, errs, log)
			}()
		case "Text":
			func() {
				text = c.Text(mtx, input, output, errs, log)
			}()
		case "Texts":
			func() {
				text = c.Texts(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		MessageTool: messageTool,
		Text:      text,
	}

	return data
}

func (c *DPFMAPICaller) MessageTool(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.MessageTool {
	where := fmt.Sprintf("WHERE MessageTool = '%s'", input.MessageTool.MessageTool)

	if input.MessageTool.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.MessageTool.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_message_tool_message_tool_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, MessageTool DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToMessageTool(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) MessageTools(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.MessageTool {
	where := fmt.Sprintf("WHERE MessageTool = '%s'", input.MessageTool.MessageTool)

	if input.MessageTool.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.MessageTool.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_message_tool_message_tool_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, MessageTool DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToMessageTool(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Text(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Text {
	var args []interface{}
	messageTool := input.MessageTool.MessageTool
	text := input.MessageTool.Text

	cnt := 0
	for _, v := range text {
		args = append(args, messageTool, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_message_tool_text_data
		WHERE (MessageTool, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Texts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Text {
	var args []interface{}
	text := input.MessageTool.Text

	cnt := 0
	for _, v := range text {
		args = append(args, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?),", cnt-1) + "(?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_message_tool_text_data
		WHERE Language IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	//
	data, err := dpfm_api_output_formatter.ConvertToText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
