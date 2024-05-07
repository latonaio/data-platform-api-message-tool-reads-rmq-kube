package dpfm_api_output_formatter

import (
	"data-platform-api-message-tool-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToMessageTool(rows *sql.Rows) (*[]MessageTool, error) {
	defer rows.Close()
	messageTool := make([]MessageTool, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.MessageTool{}

		err := rows.Scan(
			&pm.MessageTool,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &messageTool, nil
		}

		data := pm
		messageTool = append(messageTool, MessageTool{
			MessageTool:			data.MessageTool,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &messageTool, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.MessageTool,
			&pm.Language,
			&pm.MessageToolName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			MessageTool:			data.MessageTool,
			Language:          		data.Language,
			MessageToolName:		data.MessageToolName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
