import { IUser } from "../models/user";

export interface IUserRepository {
    getUser(username: string): IUser;
    addUser(user: IUser): IUser;
}