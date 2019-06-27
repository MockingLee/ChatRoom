import axios from 'axios'

axios.default.timeout = 5000
axios.defaults.headers.post['Content-Type'] = 'application/json'

const instance = axios.create()
instance.defaults.headers.post['Content-Type'] = 'application/json'

export default {
  // 获取所有用户
  getTest (data) {
    return instance.post('api/userin', data)
  }

}
