import { useEffect, useState } from "react";
import type { Query, ResponseMetaType } from "../core/models/base-model";
import type { CreateRentRequestType, RentType, UpdateRentRequestType } from "../core/models/renter-model";
import { getRenters, createRent as _createRent, updateRent as _updateRent, completeRent as _completeRent, deleteRent as _deleteRent } from "../api";
import { TABLE_PAGE_SIZE } from "../utils/constants";

export function useRenters() {
  const [data, setData] = useState<{
    meta: ResponseMetaType;
    rents: RentType[];
  }>();

   useEffect(() => {
        getData({page: 1, q: '', page_size: TABLE_PAGE_SIZE})
    }, [])

  const getData = (query: Query, errorResponse?: () => void) => {
    getRenters(query)
      .then(({ meta, data }) => {
        setData({ meta: meta, rents: data.map((r) => ({ ...r, key: r.id })) });
      })
      .catch(errorResponse);
  };

  const createRent = (rent: CreateRentRequestType, error: () => void, success: () => void) => {
    _createRent(rent)
      .then(success)
      .catch(error);
  };

  const updateRent = (updateRentPayload: UpdateRentRequestType, error: () => void, success: () => void) => {
       _updateRent(updateRentPayload)
       .then(success)
       .catch(error)
  }

  const completeRent = (id: number, error: () => void, success: () => void) => {
    _completeRent(id)
        .then(success).catch(error)
  }
            

  const deleteRent = (id: number, error: () => void, success: () => void) => {
        _deleteRent(id)
           .then(success)
           .catch(error)
  }

  return { getData, createRent, updateRent, completeRent,deleteRent, data };
}
