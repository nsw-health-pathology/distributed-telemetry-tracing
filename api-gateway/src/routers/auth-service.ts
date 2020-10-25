import Axios, { AxiosError } from 'axios'
import bodyParser from 'body-parser';
import promiseRouter from 'express-promise-router';

const router = promiseRouter()
const AUTH_SERVICE_BASE_URL = process.env.AUTH_SERVICE_BASE_URL;

router.use(bodyParser.json())

router.post('/login', async (req, res) => {
    const body = req.body

    const url = `${AUTH_SERVICE_BASE_URL}/login`

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

export { router as AuthRouter };
