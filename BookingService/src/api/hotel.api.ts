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
