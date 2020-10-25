import express from 'express'
import { AuthRouter } from './routers/user-service'
import { TodoRouter } from './routers/todo-service'

const app = express()
const PORT = process.env.PORT || 8080;

// API - Auth Service
app.use('/user-service', AuthRouter)

// API - Todo Service
app.use('/todo-service', TodoRouter)

app.listen(PORT, () => {
    console.log(`API Gateway listening on port ${PORT}`)
})