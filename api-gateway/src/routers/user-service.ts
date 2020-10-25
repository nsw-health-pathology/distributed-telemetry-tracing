import express from 'express'
import Axios from 'axios'
import bodyParser from 'body-parser';

const router = express.Router();
const AUTH_SERVICE_BASE_URL = process.env.AUTH_SERVICE_BASE_URL;

router.use(bodyParser)

router.post('/login', async (req, res) => {
    const body = req.body

    const url = `${AUTH_SERVICE_BASE_URL}`
    const response = await Axios.post(url, body)

    for (const h in response.headers) {
        res.header(h, response.headers[h])
    }

    res.status(response.status)
    res.send(response.data)
})

export { router as AuthRouter };
