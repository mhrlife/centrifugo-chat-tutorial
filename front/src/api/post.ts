import axiosInstance from "./axios.ts";
import {AxiosResponse} from "axios";


export interface ApiResponse<Response> {
    ok: true;
    data: Response;
}

export interface ApiError {
    ok: false;
    message: string;
    code: number;
}

export async function post<Request, Response>(url: string, args: Request): Promise<Response> {
    const axiosResponse: AxiosResponse<ApiResponse<Response> | ApiError> = await axiosInstance.post(url, args)

    if (!axiosResponse.data.ok) {
        throw axiosResponse.data as ApiError;
    }

    return axiosResponse.data.data as Response;
}