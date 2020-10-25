import express from 'express'
import { AuthRouter } from './routers/auth-service'
import { TodoRouter } from './routers/todo-service'

const app = express()
const PORT = process.env.PORT || 8080;

// API - Auth Service
app.use('/auth-service', AuthRouter)

// API - Todo Service
app.use('/todo-service', TodoRouter)

app.listen(PORT, () => {
    console.log(`API Gateway listening on port ${PORT}`)
})