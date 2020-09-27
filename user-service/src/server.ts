import express from 'express';
import * as AppInsights from 'applicationinsights'

import { IUserController, UserController } from './controllers/user.controller';
import { InMemoryDatabase } from './repository/in-memory/in-memory-database';
import { IUserRepository } from './repository/user.repository';
import { IUserService, UserService } from './services/user.service';

const repo: IUserRepository = new InMemoryDatabase();
const svc: IUserService = new UserService(repo);
const userController: IUserController = new UserController(svc);

// Environment Config
const port = process.env.PORT || 8000

// const appInsightsKey = process.env.APPINSIGHTS_INSTRUMENTATIONKEY || ""
AppInsights.setup();

// Seed database
repo.addUser({ username: 'admin@admin.com', password: 'P@55w0rd', firstName: 'Admin', lastName: 'Admin' })

const app = express();

app.get('/user', (req, res) => {
    const userName = req.query['username'] as string;
    const response = userController.getUser(userName);

    res.status(response.statusCode);
    Object.entries(response.headers || {}).map(value => {
        const [headerName, headerValue] = value;
        res.setHeader(headerName, headerValue as string)
    });

    res.send(response.body)
})


app.listen(port, () => {
    console.log(`Server running on port ${port}`)
})