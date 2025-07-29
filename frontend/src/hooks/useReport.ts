import { useEffect, useState } from "react";
import type { Query, ResponseMetaType } from "../core/models/base-model";
import type { RentReport } from "../core/models/rent-report-model";
import { getRentReport } from "../api";
import { TABLE_PAGE_SIZE } from "../utils/constants";

export function useReport() {
    const [report, setData] = useState<{meta: ResponseMetaType, reportData: RentReport}>();
    
    useEffect(() => {
          getData({page: 1, q: '', page_size: TABLE_PAGE_SIZE})
    }, [])

    const getData = (query: Query, error?: () => void) => {
            getRentReport(query)
            .then(({meta, data}) => {
                setData({meta: meta, reportData: data});
            }).catch(error)
    }
    
    return {report, getData}
}