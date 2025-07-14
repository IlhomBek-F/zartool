import type { RentType } from "./renter-model";

export type RentReport = {
   reports: RentType[],
   total_completed_rent: number,
   total_active_rent: number
}

export type GetReportRentResponseType = RentReport;