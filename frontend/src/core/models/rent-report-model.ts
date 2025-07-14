import type { RentType } from "./renter-model";

export type RentReport = {
   total_completed_rent: number,
   total_created_rent: number,
   rents: RentType[]
}

export type GetReportRentResponseType = RentReport;