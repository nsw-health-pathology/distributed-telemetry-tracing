import { StatusCodes } from "http-status-codes";

export interface IHeaders {
    [key: string]: string;
}

export interface IHttpResponse<T> {
    statusCode: number,
    headers?: IHeaders,
    body: T
}

export interface IError {
    message: string;
}

export const jsonResponse = <T>(body: T, statusCode: StatusCodes = StatusCodes.OK): IHttpResponse<T> => {
    return {
        statusCode, body,
        headers: {
            'Content-Type': 'application/json'
        }
    }
}