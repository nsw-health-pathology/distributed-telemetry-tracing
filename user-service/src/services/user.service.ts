import { IUser } from "../models/user";
import { IUserRepository } from "../repository/user.repository";

export interface IUserService {
    registerNewUser(user: IUser): IUser;
    getUser(userName: string): IUser | undefined;
}

export class UserService implements IUserService {

    constructor(private readonly repo: IUserRepository) { }

    registerNewUser(user: IUser): IUser {
        const existingUser = this.repo.getUser(user.username);
        if (existingUser) {
            throw new Error('Username taken');
        }

        return this.repo.addUser(user);
    }
    getUser(userName: string): IUser | undefined {
        return this.repo.getUser(userName);
    }
}