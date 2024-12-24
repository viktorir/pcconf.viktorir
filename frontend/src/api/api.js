import axios from 'axios'

const api = axios.create({
    baseURL: 'https://pcconf.viktorir.ru/api/v1', // https://pcconf.viktorir.ru/api/v1
    timeout: 5000,
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
    },
});

export default api