import express from 'express'
import Axios from 'axios'
import bodyParser from 'body-parser';

import { authenticateToken } from '../middleware/authentication';

const router = express.Router();
const TODO_SERVICE_BASE_URL = process.env.TODO_SERVICE_BASE_URL;

router.use(authenticateToken)
router.use(bodyParser)

router.get('/:id', async (req, res) => {
    const id = req.params.id;

    const url = `${TODO_SERVICE_BASE_URL}/${id}`
    const response = await Axios.get(url)

    for (const h in response.headers) {
        res.header(h, response.headers[h])
    }

    res.status(response.status)
    res.send(response.data)
})

router.post('/', async (req, res) => {
    const body = req.body

    const url = `${TODO_SERVICE_BASE_URL}`
    const response = await Axios.post(url, body)

    for (const h in response.headers) {
        res.header(h, response.headers[h])
    }

    res.status(response.status)
    res.send(response.data)
})

export { router as TodoRouter };