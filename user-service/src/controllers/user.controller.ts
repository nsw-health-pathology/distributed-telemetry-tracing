import { IError, IHttpResponse, jsonResponse } from "../models/http-response";
import { IUser } from "../models/user";
import { IUserService } from "../services/user.service";
import { StatusCodes } from 'http-status-codes';

export interface IUserController {
    getUser(userName: string): IHttpResponse<IUser | IError>
}

export class UserController implements IUserController {

    constructor(private readonly userSvc: IUserService) { }

    public getUser(userName: string): IHttpResponse<IUser | IError> {

        if (!userName) {
            return jsonResponse({ message: 'Missing username from request' }, StatusCodes.BAD_REQUEST);
        }

        const user = this.userSvc.getUser(userName);
        if (user) {
            return jsonResponse(user)
        }

        return jsonResponse({ message: 'User not found' }, StatusCodes.NOT_FOUND)
    }

}