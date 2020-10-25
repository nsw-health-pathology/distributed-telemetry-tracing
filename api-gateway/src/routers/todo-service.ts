import Axios, { AxiosError } from 'axios'
import bodyParser from 'body-parser';
import promiseRouter from 'express-promise-router';

import { authenticateToken } from '../middleware/authentication';

const router = promiseRouter()

const TODO_SERVICE_BASE_URL = process.env.TODO_SERVICE_BASE_URL;

router.use(authenticateToken)
router.use(bodyParser.json())

router.get('/:id', async (req, res) => {
    const id = req.params.id;
    console.log('calling todo-service GET /id', id)

    const url = `${TODO_SERVICE_BASE_URL}/${id}`

    try {
        const response = await Axios.get(url)
        console.log(response.data)

        res.status(response.status)
        res.send(response.data)
    } catch (error) {
        const e: AxiosError = error as AxiosError;

        const status = e.response?.status || 500
        res.status(status)
        res.send(e.response?.data || e.message || "request failed")

    }
})

router.post('/', async (req, res) => {
    console.log('calling todo-service POST')
    const body = req.body

    const url = `${TODO_SERVICE_BASE_URL}`

    try {
        const response = await Axios.post(url, body)

        res.status(response.status)
        res.send(response.data)
    } catch (error) {
        const e: AxiosError = error as AxiosError;

        const status = e.response?.status || 500
        res.status(status)
        res.send(e.response?.data || e.message || "request failed")

    }
})

export { router as TodoRouter };