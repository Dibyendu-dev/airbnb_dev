import axios from "axios";
import { serverConfig } from "../config";

export const getAvailableRooms = async (
  roomsCategoryId: number,
  checkInDate: string,
  checkOutDate: string
) => {

        const response = await axios.get(`${serverConfig.HOTEL_SERVICE_URL}rooms/available`, {
            params: {
                roomsCategoryId,
                checkInDate,
                checkOutDate
            }
        })

    return response.data;
};


export const updateBookingIdToRooms = async (bookingId: number,roomIds: number[]) => {
   
    const response = await axios.post(`${serverConfig.HOTEL_SERVICE_URL}rooms/update-booking-id`,{
        bookingId,
        roomIds
    })

    return response.data;
}
