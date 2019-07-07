import axios from 'axios'
import './global'
axios.default.timeout = 5000
axios.defaults.headers.post['Content-Type'] = 'application/json'

const instance = axios.create()
instance.defaults.headers.post['Content-Type'] = 'application/json'

export default {
  // 获取所有用户
  userIn (data) {
    return instance.get('/api/userin?username=' + data)
  }

}
