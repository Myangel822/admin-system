import  axios  from 'axios'
axios.default.timeout = 100000
const request = axios.create({
	baseURL:"http://localhost:8000",
})
request.interceptors.request.use(
  config => {
    // do something before request is sent

   
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      
      config.headers["Cookies"] = document.cookie
    
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

export default request