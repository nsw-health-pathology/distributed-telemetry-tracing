import { Request, Response, NextFunction } from 'express'
import jwt from 'jsonwebtoken'

const ACCESS_TOKEN_SECRET = process.env.ACCESS_TOKEN_SECRET;

// https://www.digitalocean.com/community/tutorials/nodejs-jwt-expressjs
export const authenticateToken = (req: Request, res: Response, next: NextFunction) => {

    // Gather the jwt access token from the request header
    const authHeader = req.headers['authorization']
    const token = authHeader && authHeader.split(' ')[1]
    if (token == null) return res.sendStatus(401) // if there isn't any token

    jwt.verify(token, ACCESS_TOKEN_SECRET as string, (err, user: any) => {
        if (err) {
            console.log(err)
            return res.sendStatus(403)
        }

        console.log('valid jwt', user)

        // req.user = user
        next() // pass the execution off to whatever request the client intended
    })
}