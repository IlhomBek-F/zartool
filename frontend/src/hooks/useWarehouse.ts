import { useEffect, useState } from "react";
import type { ResponseMetaType } from "../core/models/base-model";
import type { WarehouseToolType } from "../core/models/warehouse-tool-model";
import { getRentTools as _getRentTools, deleteTool as _deleteTool, addNewTool as _addNewTool, updateTool as _updateTool} from "../api/warehouse";
import type { CreateRentToolRequestType, UpdateRentToolRequestType } from "../core/models/rent-tool-model";

export function useWarehouse() {
    const [dataSource, setTools] = useState<{meta: ResponseMetaType, data: WarehouseToolType[]}>();
    
    useEffect(() => {
       getRentTools()
    }, [])

    const getRentTools = (page = 1, error?: () => void) => {
           _getRentTools(page)
            .then(({data, meta}) => {
               setTools({meta, data: data.map((t) => ({...t, key: t.id}))})
            }).catch(error)
    }

    const deleteTool = (id: number, error: () => void, success: () => void) => {
         _deleteTool(id).then(success).catch(error)
    }

    const addNewTool = (tool: CreateRentToolRequestType[], error: () => void, success: () => void) => {
        _addNewTool(tool).then(success).catch(error)
    }

    const updateTool = (payload: UpdateRentToolRequestType, error: () => void, success: () => void) => {
            _updateTool(payload).then(success).catch(error)
    }

    return {dataSource, getRentTools, deleteTool, addNewTool, updateTool}
}